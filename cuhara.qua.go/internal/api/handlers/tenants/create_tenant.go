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
		log := util.LogFromEchoContext(c).With().Str("function", "createTenantHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("createTenantHandler started")

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

		log.Debug().Msg("createTenantHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
