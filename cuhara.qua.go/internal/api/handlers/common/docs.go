package common

import (
	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/middleware"
	"github.com/labstack/echo/v4"
)

func DocsRouter(s *api.Server) *echo.Route {
	return s.Router.Root.GET("/docs", docsHandler(s))
}

func docsHandler(_ *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(middleware.AuthModeKey, middleware.Anonymous)
		return c.File("docs/swagger.yml")
	}
}
