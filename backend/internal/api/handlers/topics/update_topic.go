package topics

import (
	"net/http"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/types"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
)

func UpdateTopicRouter(s *api.Server) *echo.Route {
	return s.Router.APIV1Topics.PATCH("/:id", updateTopicHandler(s))
}

func updateTopicHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := util.LogFromEchoContext(c).With().Str("function", "updateTopicHandler").Logger()
		ctx := c.Request().Context()

		log.Debug().Msg("updateTopicHandler started")

		param := c.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil || id <= 0 {
			return httperrors.ErrInvalidID
		}

		var body types.UpdateTopicRequest
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		res, err := s.Topic.Update(ctx, dto.UpdateTopicRequest{
			ID:   id,
			Name: body.Name,
		})
		if err != nil {
			return err
		}

		log.Debug().Msg("updateTopicHandler successfully executed")

		return c.JSON(http.StatusOK, res.ToTypes())
	}
}
