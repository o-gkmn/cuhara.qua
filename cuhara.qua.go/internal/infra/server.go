package infra

import (
	"net/http"

	"cuhara.qua.go/internal/infra/config"
)

func NewServer(cfg config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr: cfg.Addr,
		ReadTimeout: cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout: cfg.IdleTimeout,
		Handler: handler,
	}
}