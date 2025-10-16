package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func DeleteTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Topics.DELETE("/:id", deleteTopicHandler(s))
}

func deleteTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "deleteTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("deleteTopicHandler started")

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil || id <= 0 {
			return httperrors.ErrInvalidID
		}

		res, err := s.Topic.Delete(ctx, dto.DeleteTopicRequest{
			ID: id,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("deleteTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
