package roles

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteRoleRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Roles.DELETE("/:id", deleteRoleHandler(s))
}

func deleteRoleHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "deleteRoleHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("deleteRoleHandler started")

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil || id <= 0 {
			return httperrors.ErrInvalidID
		}

		res, err := s.Role.Delete(ctx, dto.DeleteRoleRequest{
			ID: id,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("deleteRoleHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
