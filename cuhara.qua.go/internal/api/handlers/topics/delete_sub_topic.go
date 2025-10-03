package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.DELETE("/:subTopicID", deleteSubTopicHandler(s))
}

func deleteSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "deleteSubTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("deleteSubTopicHandler started")

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

		log.Debug().Msg("deleteSubTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
