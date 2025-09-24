package users

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func UpdateUserRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Users.PUT("", updateUserHandler(s))
}

func updateUserHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.UpdateUserRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.User.Update(ctx, dto.UpdateUserRequest{
			ID:         util.PtrToInt64(body.ID),
			Name:       util.StringToPtr(body.Name),
			Email:      util.StringToPtr(body.Email),
			VscAccount: util.StringToPtr(body.VscAccount),
			RoleID:     util.SafeChainInt64(func() *int64 { return body.Role.ID }),
		})
		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}
