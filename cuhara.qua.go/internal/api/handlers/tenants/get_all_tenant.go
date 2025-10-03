package tenants

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetAllRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Tennants.GET("", getAll(s))
}

func getAll(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getAll").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getAll started")

		tenants, err := s.Tennant.GetAll(ctx)
		if err != nil {
			return err
		}

		tenantResponse := make([]types.TenantResponse, len(tenants))
		for i, tenant := range tenants {
			tenantResponse[i] = *tenant.ToTypes()
		}

		log.Debug().Msg("getAll successfully executed")

		return c.JSON(http.StatusOK, tenantResponse)
	}
}
