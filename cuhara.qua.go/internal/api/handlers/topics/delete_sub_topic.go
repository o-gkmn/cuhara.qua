package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeleteSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.DELETE("/:subTopicID", deleteSubTopicHandler(s))
}

func deleteSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()


		var topicIDStr = c.Param("id")
		topicID, err := strconv.ParseInt(topicIDStr, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse topic id")
			return err
		}

		var subTopicIDStr = c.Param("subTopicID")
		subTopicID, err := strconv.ParseInt(subTopicIDStr, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse sub topic id")
			return err
		}
		
		res, err := s.Topic.DeleteSubTopic(ctx, dto.DeleteSubTopicRequest{
			ID:      subTopicID,
			TopicID: topicID,
		})
		if err != nil {
			return err
		}
		
		
		return c.JSON(http.StatusOK, res.ToTypes())
	}
}