package auth

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/go-openapi/swag"
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
			Email:      swag.StringValue(body.Email),
			Password:   swag.StringValue(body.Password),
			VscAccount: swag.StringValue(body.VscAccount),
			RoleID:     swag.Int64Value(body.RoleID),
			TenantID:   swag.Int64Value(body.TenantID),
		})
		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}
