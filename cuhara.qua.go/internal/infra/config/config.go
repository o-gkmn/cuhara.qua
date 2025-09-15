package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr             string
	WriteDatabaseURL string
	ReadDatabaseURL  string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
	JWTSecret        string
	JWTIssuer        string
	JWTTTLMinute     int
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("WARNING: .env file not found: %v", err)
	}
	return Config{
		Addr:             getEnv("ADDR", ":8080"),
		WriteDatabaseURL: mustEnv("WRITE_DATABASE_URL"),
		ReadDatabaseURL:  mustEnv("READ_DATABASE_URL"),
		ReadTimeout:      15 * time.Second,
		WriteTimeout:     15 * time.Second,
		IdleTimeout:      60 * time.Second,
		JWTSecret:        mustEnv("JWT_SECRET"),
		JWTIssuer:        getEnv("JWT_ISSUER", "cuhara.qua"),
		JWTTTLMinute:     envInt("JWT_TTL_MINUTES", 60),
	}, nil
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	if i, err := strconv.Atoi(v); err == nil && i > 0 {
		return i
	}
	return def
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}

func mustEnv(key string) string {
	return os.Getenv(key)
}
