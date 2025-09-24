package middleware

import (
	"context"
	"net/http"
	"strings"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrUnauthorized         = httperrors.NewHTTPError(http.StatusUnauthorized, "unauthorized", "unauthorized")
	ErrForbidden            = httperrors.NewHTTPError(http.StatusForbidden, "forbidden", "forbidden")
	ErrNotFound             = httperrors.NewHTTPError(http.StatusNotFound, "not_found", "not_found")
	ErrInvalidToken         = httperrors.NewHTTPError(http.StatusUnauthorized, "invalid_token", "invalid_token")
	ErrInvalidSubjcet       = httperrors.NewHTTPError(http.StatusUnauthorized, "invalid_subject", "invalid_subject")
	ErrInvalidIssuer        = httperrors.NewHTTPError(http.StatusUnauthorized, "invalid_issuer", "invalid_issuer")
	ErrMissingHeader        = httperrors.NewHTTPError(http.StatusUnauthorized, "missing_header", "missing_header")
	ErrInvalidHeader        = httperrors.NewHTTPError(http.StatusUnauthorized, "invalid_header", "invalid_header")
	ErrInvalidSigningMethod = httperrors.NewHTTPError(http.StatusUnauthorized, "invalid_signing_method", "invalid_signing_method")
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
		Skipper: middleware.DefaultSkipper,
	}
)

const authHeader = "Authorization"

func JWTAuthWithConfig(cfg JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cfg.Skipper != nil && cfg.Skipper(c) {
				return next(c)
			}

			tokenStr, err := bearerTokenFromHeader(c.Request().Header.Get(authHeader))
			if err != nil {
				return err
			}

			claims, err := validateJWT(tokenStr, cfg.S.Config.Auth)
			if err != nil {
				return ErrInvalidToken
			}
			if claims.Subject == "" {
				return ErrInvalidSubjcet
			}
			if claims.Issuer != cfg.S.Config.Auth.JWTIssuer {
				return ErrInvalidIssuer
			}

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, util.CTXKeyUser, claims.Subject)
			ctx = context.WithValue(ctx, util.CTXKeyAuthToken, tokenStr)
			c.SetRequest(c.Request().WithContext(ctx))

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
		return "", ErrMissingHeader
	}
	parts := strings.SplitN(authz, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], authScheme) {
		return "", ErrInvalidHeader
	}
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", ErrInvalidHeader
	}
	return token, nil
}

func validateJWT(tokenStr string, cfg config.AuthServer) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidSigningMethod
			}
			return []byte(cfg.JWTSecret), nil
		},
		jwt.WithIssuer(cfg.JWTIssuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
