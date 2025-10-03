package handlers

import (
	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/handlers/auth"
	"cuhara.qua.go/internal/api/handlers/claims"
	"cuhara.qua.go/internal/api/handlers/common"
	"cuhara.qua.go/internal/api/handlers/roles"
	"cuhara.qua.go/internal/api/handlers/tenants"
	"cuhara.qua.go/internal/api/handlers/topics"
	"cuhara.qua.go/internal/api/handlers/users"
	"github.com/labstack/echo/v4"
)

func AttachAllRoutes(s *api.Server) {
	s.Router.Routes = []*echo.Route{
		common.SwaggerRouter(s),
		common.DocsRouter(s),
		auth.LoginRouter(s),
		auth.RegisterRouter(s),
		roles.GetAllRouter(s),
		roles.CreateRoleRouter(s),
		roles.UpdateRoleRouter(s),
		roles.DeleteRoleRouter(s),
		users.GetUsersRouter(s),
		users.UpdateUserRoute(s),
		users.DeleteUserRoute(s),
		tenants.GetAllRouter(s),
		tenants.CreateTenantRouter(s),
		tenants.UpdateTenantRouter(s),
		tenants.DeleteTenantRouter(s),
		topics.GetAllTopicRouter(s),
		topics.CreateTopicRouter(s),
		topics.UpdateTopicRouter(s),
		topics.DeleteTopicRouter(s),
		topics.GetAllSubTopicRouter(s),
		topics.CreateSubTopicRouter(s),
		topics.DeleteSubTopicRouter(s),
		topics.UpdateSubTopicRouter(s),
		claims.GetAllRouter(s),
		claims.CreateClaimRouter(s),
		claims.UpdateClaimRouter(s),
		claims.DeleteClaimRouter(s),
	}
}
