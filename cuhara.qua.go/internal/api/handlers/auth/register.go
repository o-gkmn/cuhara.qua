package auth

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Auth.POST("/register", registerHandler(s))
}

func registerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.RegisterRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Auth.Register(ctx, dto.RegisterRequest{
			Email:      body.Email,
			Password:   body.Password,
			VscAccount: body.VscAccount,
			RoleID:     *body.Role.Id,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
