package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func CreateSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1SubTopics.POST("", createSubTopicHandler(s))
}

func createSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "createSubTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("createSubTopicHandler started")

		var topicIDStr = c.Param("id")
		topicID, err := strconv.ParseInt(topicIDStr, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse topic id")
			return err
		}

		var body dto.CreateSubTopicRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Topic.CreateSubTopic(ctx, dto.CreateSubTopicRequest{
			TopicID: topicID,
			Name:    body.Name,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("createSubTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
