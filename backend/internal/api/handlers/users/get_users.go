package users

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetUsersRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Users.GET("", getUsersHandler(s))
}

func getUsersHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getUsersHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getUsersHandler started")

		users, err := s.User.GetAll(ctx)
		if err != nil {
			return err
		}

		userResponses := make([]*types.UserResponse, len(users))
		for i, user := range users {
			userResponses[i] = user.ToTypes()
		}

		log.Debug().Msg("getUsersHandler successfully executed")

		return c.JSON(http.StatusOK, userResponses)
	}
}
