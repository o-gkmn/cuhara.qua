package claims

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"github.com/labstack/echo/v4"
)

func DeleteClaimRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Claims.DELETE("/:id", deleteClaimHandler(s))
}

func deleteClaimHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return httperrors.ErrInvalidID
		}

		res, err := s.Claim.Delete(ctx, dto.DeleteClaimRequest{
			ID: id,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}