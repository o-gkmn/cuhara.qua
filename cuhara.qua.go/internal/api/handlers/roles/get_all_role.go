package roles

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"github.com/labstack/echo/v4"
)

func GetAllRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Roles.GET("", getAll(s))
}

func getAll(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		roles, err := s.Role.GetRoles(ctx)
		if err != nil {
			return err
		}

		roleResponse := make([]types.RoleResponse, len(roles))
		for i, role := range roles {
			roleResponse[i] = *role.ToTypes()
		}

		return c.JSON(http.StatusOK, roleResponse)
	}
}
