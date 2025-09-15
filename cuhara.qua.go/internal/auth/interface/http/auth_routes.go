package http

import "github.com/go-chi/chi"

func RegisterAuthRoutes(r chi.Router, ac *AuthController) {
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", ac.Login)
	})
}
