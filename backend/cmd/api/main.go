package main

import (
	"os"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/router"
	"cuhara.qua.go/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func main() {
	_ = gotenv.Load(".env")

	cfg := config.DefaultServiceConfigFromEnv()

	initLogger(cfg.Logger)

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

func initLogger(loggerConfig config.LoggerServer) {
	// Set the global log level
	zerolog.SetGlobalLevel(loggerConfig.Level)

	// Configure output
	if loggerConfig.PrettyPrintConsole {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = log.Output(os.Stderr)
	}

	// Add caller information if configured
	if loggerConfig.LogCaller {
		log.Logger = log.Logger.With().Caller().Logger()
	}
}
