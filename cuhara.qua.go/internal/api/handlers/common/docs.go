package common

import (
	"cuhara.qua.go/internal/api"
	"github.com/labstack/echo/v4"
)

func DocsRouter(s *api.Server) *echo.Route {
	return s.Router.Root.GET("/docs", docsHandler(s))
}

func docsHandler(_ *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.File("docs/swagger.yml")
	}
}
