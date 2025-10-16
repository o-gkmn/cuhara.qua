package common

import (
	"net/http"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/middleware"
	"github.com/labstack/echo/v4"
)

func SwaggerRouter(s *api.Server) *echo.Route {
	return s.Router.Root.GET("/swagger", swaggerHandler(s))
}

func swaggerHandler(_ *api.Server) echo.HandlerFunc {
	const html = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
    <style>
      body { margin: 0; }
      #swagger-ui { position: absolute; top:0; left:0; right:0; bottom:0; }
    </style>
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.onload = () => {
        SwaggerUIBundle({
          url: '/docs',
          dom_id: '#swagger-ui',
          presets: [SwaggerUIBundle.presets.apis],
          layout: "BaseLayout"
        });
      };
    </script>
  </body>
</html>`
	return func(c echo.Context) error {
		c.Set(middleware.AuthModeKey, middleware.Anonymous)
		return c.HTML(http.StatusOK, html)
	}
}
