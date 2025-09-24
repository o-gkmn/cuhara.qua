package handlers

import (
	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/handlers/auth"
	"cuhara.qua.go/internal/api/handlers/common"
	"cuhara.qua.go/internal/api/handlers/users"
	"github.com/labstack/echo/v4"
)

func AttachAllRoutes(s *api.Server) {
	s.Router.Routes = []*echo.Route{
		common.SwaggerRouter(s),
		common.DocsRouter(s),
		auth.LoginRouter(s),
		auth.RegisterRouter(s),
		users.GetUsersRouter(s),
		users.UpdateUserRoute(s),
	}
}
