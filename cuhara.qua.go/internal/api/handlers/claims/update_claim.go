package claims

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

func UpdateClaimRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Claims.PATCH("/:id", updateClaimHandler(s))
}

func updateClaimHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return httperrors.ErrInvalidID
		}

		var body types.UpdateClaimRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Claim.Update(ctx, dto.UpdateClaimRequest{
			ID:          id,
			Name:        body.Name,
			Description: body.Description,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
