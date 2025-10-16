package claims

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteClaimRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Claims.DELETE("/:id", deleteClaimHandler(s))
}

func deleteClaimHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "deleteClaimHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("deleteClaimHandler started")

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

		log.Debug().Msg("deleteClaimHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
