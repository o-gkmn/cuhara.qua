package tenants

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteTenantRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Tennants.DELETE("/:id", deleteTenantHandler(s))
}

func deleteTenantHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "deleteTenantHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("deleteTenantHandler started")

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil || id <= 0 {
			return httperrors.ErrInvalidID
		}

		res, err := s.Tennant.Delete(ctx, dto.DeleteTenantRequest{
			ID: id,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("deleteTenantHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
