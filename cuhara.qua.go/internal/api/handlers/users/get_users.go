package users

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"github.com/labstack/echo/v4"
)

func GetUsersRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Users.GET("", getUsersHandler(s))
}

func getUsersHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		users, err := s.User.GetUsers(ctx)
		if err != nil {
			return err
		}

		userResponses := make([]*types.UserResponse, len(users))
		for i, user := range users {
			userResponses[i] = user.ToTypes()
		}

		return c.JSON(http.StatusOK, userResponses)
	}
}
