package topics

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func CreateSubTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Topics.POST("", createSubTopicHandler(s))
}

func createSubTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var topicID int64
		err := util.BindValidatePathParams(c, &topicID)
		if err != nil {
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

		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}