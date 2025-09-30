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
			ID:         body.Id,
			Name:       body.Name,
			Email:      util.EmailPtrToStringPtr(body.Email),
			VscAccount: body.VscAccount,
			RoleID:     util.SafeChainInt64(func() *int64 { return &body.Role.Id }),
		})
		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}
