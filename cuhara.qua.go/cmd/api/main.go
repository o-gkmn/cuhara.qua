package main

import (
	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/router"
	"cuhara.qua.go/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load(".env")

	cfg := config.DefaultServiceConfigFromEnv()

	s := api.NewServer(cfg)
	err := router.Init(s)
	if err != nil {
		log.Error().Err(err).Msg("router cannot be loaded")
	}

	err = s.InitCmd().Start()
	if err != nil {
		log.Error().Err(err).Msg("server cannot be loaded")
	}
}
