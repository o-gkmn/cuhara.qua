package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func UpdateSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.PATCH("/:subTopicID", updateSubTopicHandler(s))
}

func updateSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "updateSubTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("updateSubTopicHandler started")

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

		var body dto.UpdateSubTopicRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Topic.UpdateSubTopic(ctx, dto.UpdateSubTopicRequest{
			ID:      subTopicID,
			TopicID: topicID,
			Name:    body.Name,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("updateSubTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
