package middleware

import (
	"strings"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

var (
	skipJWTAuthPaths = []string{"/api/v1/auth/login", "/", "/swagger", "/docs"}
)

const AuthModeKey = "auth_mode"

type AuthMode string

const (
	Anonymous AuthMode = "anonymous"
	Forbidden AuthMode = "forbidden"
)

type JWTConfig struct {
	S       *api.Server
	Skipper func(c echo.Context) bool
}

var (
	DefaultJWTConfig = JWTConfig{
		Skipper: skipJWTAuth,
	}
)

const authHeader = "Authorization"

func JWTAuthWithConfig(cfg JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := util.LogFromEchoContext(c).With().Str("middleware", "jwt").Logger()

			if cfg.Skipper != nil && cfg.Skipper(c) {
				return next(c)
			}

			tokenStr, err := bearerTokenFromHeader(c.Request().Header.Get(authHeader))
			if err != nil {
				log.Info().Err(err).Str("token", tokenStr).Msg("invalid bearer token")
				return err
			}

			claims, err := util.Verify(tokenStr)
			if err != nil {
				log.Info().Err(err).Any("claims", claims).Msg("invalid jwt token")
				return httperrors.ErrInvalidToken
			}
			if claims["sub"] == nil {
				log.Info().Any("claims", claims).Msg("invalid subject")
				return httperrors.ErrInvalidSubjcet
			}
			if claims["iss"] != cfg.S.Config.Auth.JWTIssuer {
				log.Info().Any("claims", claims).Msg("invalid issuer")
				return httperrors.ErrInvalidIssuer
			}

			ctx := c.Request().Context()
			ctx = util.SaveContextValue(ctx, util.CTXKeyUser, claims["sub"])
			ctx = util.SaveContextValue(ctx, util.CTXKeyAuthToken, tokenStr)
			c.SetRequest(c.Request().WithContext(ctx))

			log.Debug().Msg("token validation successful")

			return next(c)
		}
	}
}

func JWTAuth(s *api.Server) echo.MiddlewareFunc {
	c := DefaultJWTConfig
	c.S = s
	return JWTAuthWithConfig(c)
}

const authScheme = "Bearer"

func bearerTokenFromHeader(authz string) (string, error) {
	authz = strings.TrimSpace(authz)
	if authz == "" {
		return "", httperrors.ErrMissingHeader
	}
	parts := strings.SplitN(authz, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], authScheme) {
		return "", httperrors.ErrInvalidHeader
	}
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", httperrors.ErrInvalidHeader
	}
	return token, nil
}

func skipJWTAuth(c echo.Context) bool {
	for _, path := range skipJWTAuthPaths {
		if c.Request().URL.Path == path {
			return true
		}
	}
	return false
}
