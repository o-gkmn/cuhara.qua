package roles

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func UpdateRoleRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Roles.PATCH("/:id", updateRoleHandler(s))
}

func updateRoleHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil || id <= 0 {
			return httperrors.ErrInvalidID
		}

		var body types.UpdateRoleRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Role.Update(ctx, dto.UpdateRoleRequest{
			ID:   id,
			Name: body.Name,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
