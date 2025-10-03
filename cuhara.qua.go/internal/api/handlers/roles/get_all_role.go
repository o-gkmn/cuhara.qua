package roles

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetAllRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Roles.GET("", getAll(s))
}

func getAll(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getAll").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getAll started")

		roles, err := s.Role.GetAll(ctx)
		if err != nil {
			return err
		}

		roleResponse := make([]types.RoleResponse, len(roles))
		for i, role := range roles {
			roleResponse[i] = *role.ToTypes()
		}

		log.Debug().Msg("getAll successfully executed")

		return c.JSON(http.StatusOK, roleResponse)
	}
}
