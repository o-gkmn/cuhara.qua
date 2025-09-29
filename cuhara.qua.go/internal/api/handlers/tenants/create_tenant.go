package tenants

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func CreateTenantRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Tennants.POST("", createTenantHandler(s))
}

func createTenantHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.CreateTenantRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Tennant.Create(ctx, dto.CreateTenantRequest{
			Name: body.Name,
		})

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}