package claims

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func CreateClaimRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Claims.POST("", createClaimHandler(s))
}

func createClaimHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var body types.CreateClaimRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Claim.Create(ctx, dto.CreateClaimRequest{
			Name:        body.Name,
			Description: &body.Description,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
