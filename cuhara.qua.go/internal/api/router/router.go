package router

import (
	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/handlers"
	"cuhara.qua.go/internal/api/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Init(s *api.Server) error {
	s.Echo = echo.New()

	s.Echo.Debug = s.Config.Echo.Debug
	s.Echo.HideBanner = true
	s.Echo.Logger.SetOutput(&echoLogger{level: s.Config.Logger.RequestLevel, log: log.With().Str("component", "echo").Logger()})
	echo.NotFoundHandler = NotFoundHandler(s.Config)

	s.Echo.HTTPErrorHandler = HTTPErrorHandlerWithConfig(HTTPErrorHandlerConfig{
		HideInternalServerErrorDetails: s.Config.Echo.HideInternalServerErrorDetails,
	})

	// -
	// General middleware
	if s.Config.Echo.EnableTrailingSlashMiddleware {
		s.Echo.Pre(echoMiddleware.RemoveTrailingSlash())
	} else {
		log.Warn().Msg("Disabling trailing slash middleware due to environment config")
	}

	if s.Config.Echo.EnableRecoverMiddleware {
		s.Echo.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
			LogErrorFunc: middleware.LogErrorFuncWithRequestInfo,
		}))
	} else {
		log.Warn().Msg("Disabling recover middleware due to environment config")
	}

	if s.Config.Echo.EnableSecureMiddleware {
		s.Echo.Use(echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
			Skipper:               echoMiddleware.DefaultSecureConfig.Skipper,
			XSSProtection:         s.Config.Echo.SecureMiddleware.XSSProtection,
			ContentTypeNosniff:    s.Config.Echo.SecureMiddleware.ContentTypeNosniff,
			XFrameOptions:         s.Config.Echo.SecureMiddleware.XFrameOptions,
			HSTSPreloadEnabled:    s.Config.Echo.SecureMiddleware.HSTSPreloadEnabled,
			HSTSExcludeSubdomains: s.Config.Echo.SecureMiddleware.HSTSExcludeSubdomains,
			HSTSMaxAge:            s.Config.Echo.SecureMiddleware.HSTSMaxAge,
			CSPReportOnly:         s.Config.Echo.SecureMiddleware.CSPReportOnly,
			ContentSecurityPolicy: s.Config.Echo.SecureMiddleware.ContentSecurityPolicy,
			ReferrerPolicy:        s.Config.Echo.SecureMiddleware.ReferrerPolicy,
		}))
	} else {
		log.Warn().Msg("Disabling secure middleware due to environment config")
	}

	if s.Config.Echo.EnableRequestIDMiddleware {
		s.Echo.Use(echoMiddleware.RequestID())
	} else {
		log.Warn().Msg("Disabling request id middleware due to environment config")
	}

	if s.Config.Echo.EnableCORSMiddleware {
		s.Echo.Use(echoMiddleware.CORS())
	} else {
		log.Warn().Msg("Disabling cors middleware due to environment config")
	}

	if s.Config.Echo.EnableValidationMiddleware {
		s.Echo.Use(middleware.OpenAPIValidationMiddleware())
	} else {
		log.Warn().Msg("Disabling validation middleware due to environment config")
	}

	s.Router = &api.Router{
		Routes:        nil,
		Root:          s.Echo.Group(""),
		APIV1Auth:     s.Echo.Group("/api/v1/auth"),
		APIV1Users:    s.Echo.Group("/api/v1/users"),
		APIV1Roles:    s.Echo.Group("/api/v1/roles"),
		APIV1Tennants: s.Echo.Group("/api/v1/tenants"),
		APIV1Topics:   s.Echo.Group("/api/v1/topics"),
		APIV1Claims:   s.Echo.Group("/api/v1/claims"),
	}

	handlers.AttachAllRoutes(s)

	return nil
}
