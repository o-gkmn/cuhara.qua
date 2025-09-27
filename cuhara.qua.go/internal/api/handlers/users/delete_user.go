package users

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteUserRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Users.DELETE("/:id", deleteUserHandler(s))
}

func deleteUserHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("id", c.Param("id")).Logger()
		ctx := c.Request().Context()

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return err
		}

		res, err := s.User.Delete(ctx, dto.DeleteUserRequest{
			ID: id,
		})
		if err != nil {
			log.Err(err).Msg("Failed to delete user")
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}
