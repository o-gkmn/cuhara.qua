package claims

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetAllRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Claims.GET("", getAllHandler(s))
}

func getAllHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getAllHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getAllHandler started")

		claims, err := s.Claim.GetAll(ctx)
		if err != nil {
			return err
		}

		claimResponses := make([]types.ClaimResponse, len(claims))
		for i, claim := range claims {
			claimResponses[i] = *claim.ToTypes()
		}

		log.Debug().Msg("getAllHandler successfully executed")

		return c.JSON(http.StatusOK, claimResponses)
	}
}
