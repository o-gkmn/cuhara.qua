package http

import "github.com/go-chi/chi"

func RegisterTennantRoutes(r chi.Router, tc *TennantsController) {
	r.Route("/api/v1/tennants", func(r chi.Router) {
		r.Post("/", tc.Create)
	})
}
