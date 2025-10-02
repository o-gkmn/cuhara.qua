package topics

import (
	"net/http"

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
		ctx := c.Request().Context()

		var topicID int64
		err := util.BindValidatePathParams(c, &topicID)
		if err != nil {
			return err
		}
		
		var subTopicID int64
		err = util.BindValidatePathParams(c, &subTopicID)
		if err != nil {
			return err
		}
		
		res, err := s.Topic.DeleteSubTopic(ctx, dto.DeleteSubTopicRequest{
			ID:      subTopicID,
			TopicID: topicID,
		})
		if err != nil {
			return err
		}
		
		
		return util.ValidateAndReturn(c, http.StatusOK, res.ToTypes())
	}
}