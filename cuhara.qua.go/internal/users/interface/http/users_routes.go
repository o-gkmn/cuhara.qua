package http

import "github.com/go-chi/chi"

func RegisterUserRoutes(r chi.Router, uc *UsersController) {
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", uc.CreateUser)
	})
}
