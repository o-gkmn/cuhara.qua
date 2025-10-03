package topics

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func CreateTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Topics.POST("", createTopicHandler(s))
}

func createTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "createTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("createTopicHandler started")

		var body types.CreateTopicRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Topic.Create(ctx, dto.CreateTopicRequest{
			Name: body.Name,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("createTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
