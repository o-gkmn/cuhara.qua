package topics

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetAllTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Topics.GET("", getAllTopicHandler(s))
}

func getAllTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getAllTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getAllTopicHandler started")

		res, err := s.Topic.GetAll(ctx)
		if err != nil {
			return err
		}

		topicResponse := make([]types.TopicResponse, len(res))
		for i, topic := range res {
			topicResponse[i] = *topic.ToTypes()
		}

		log.Debug().Msg("getAllTopicHandler successfully executed")

		return c.JSON(http.StatusOK, topicResponse)
	}
}
