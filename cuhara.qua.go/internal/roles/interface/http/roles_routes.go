package http

import "github.com/go-chi/chi"

func RegisterRoleRoutes(r chi.Router, rc *RoleController) {
	r.Route("/api/v1/roles", func(r chi.Router) {
		r.Post("/", rc.CreateRole)
	})
}