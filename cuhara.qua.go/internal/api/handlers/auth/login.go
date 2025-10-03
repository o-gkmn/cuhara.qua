package auth

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"

	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func LoginRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Auth.POST("/login", loginHandler(s))
}

func loginHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "loginHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("loginHandler started")

		var body types.LoginRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Auth.Login(ctx, dto.LoginRequest{
			Email:    string(body.Email),
			Password: body.Password,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("loginHandler successfully executed")
		
		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
