package roles

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/go-openapi/swag/conv"
	"github.com/labstack/echo/v4"
)

func CreateRoleRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Roles.POST("", createRoleHandler(s))
}

func createRoleHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.CreateRoleRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Role.Create(ctx, dto.CreateRoleRequest{
			Name:     conv.Value(body.Name),
			TenantID: *util.SafeChainInt64(func() *int64 { return body.Tenant.ID }),
		})
		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}
