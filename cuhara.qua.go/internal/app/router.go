package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"cuhara.qua.go/internal/common/cqrs"
	roleshttp "cuhara.qua.go/internal/roles/interface/http"
	tennanthttp "cuhara.qua.go/internal/tennants/interface/http"
	usershttp "cuhara.qua.go/internal/users/interface/http"
)

func NewRouter(cb *cqrs.CommandBus) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	uc := usershttp.NewUsersController(cb)
	usershttp.RegisterUserRoutes(r, uc)

	rc := roleshttp.NewRoleController(cb)
	roleshttp.RegisterRoleRoutes(r, rc)

	tc := tennanthttp.NewTennantsController(cb)
	tennanthttp.RegisterTennantRoutes(r, tc)

	return r
}
