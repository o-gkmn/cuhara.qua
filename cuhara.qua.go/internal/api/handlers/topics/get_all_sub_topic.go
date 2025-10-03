package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GetAllSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.GET("", getAllSubTopicHandler(s))
}

func getAllSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

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

		return c.JSON(http.StatusOK, res)
	}
}
