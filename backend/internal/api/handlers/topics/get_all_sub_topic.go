package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func GetAllSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.GET("", getAllSubTopicHandler(s))
}

func getAllSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "getAllSubTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("getAllSubTopicHandler started")

		var topicIDStr = c.Param("id")
		topicID, err := strconv.ParseInt(topicIDStr, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse topic id")
			return err
		}

		res, err := s.Topic.GetSubTopics(ctx, dto.GetSubTopicsRequest{
			TopicID: topicID,
		})

		if err != nil {
			return err
		}

		log.Debug().Msg("getAllSubTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res)
	}
}
