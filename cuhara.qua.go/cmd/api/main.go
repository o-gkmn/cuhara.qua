package main

import (
	"log"

	"cuhara.qua.go/internal/app"
	"cuhara.qua.go/internal/infra"
	"cuhara.qua.go/internal/infra/config"
	"cuhara.qua.go/internal/infra/db"

	_ "cuhara.qua.go/docs"
)

// @title Cuhara QUA API
// @version 1.0
// // @description Cuhara QUA API Documentation

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("ERROR: Config not found")
	}

	deps, err := app.BuildDeps(cfg.WriteDatabaseURL, cfg.ReadDatabaseURL)
	if err != nil {
		log.Fatal("ERROR: Dependency injection failed:", err)
	}

	if err := db.RunMigrations(deps.WriteDB); err != nil {
		log.Printf("WARNING: Migration failed: %v", err)
	}

	sqlReadDb, err := deps.ReadDB.DB()
	if err != nil {
		log.Fatal("ERROR: Read db connection failed")
	}
	defer sqlReadDb.Close()

	sqlWriteDb, err := deps.WriteDB.DB()
	if err != nil {
		log.Fatal("ERROR: Write db conneciton failed")
	}
	defer sqlWriteDb.Close()

	cb := app.BuildCommandBus(deps)
	router := app.NewRouter(cb)

	srv := infra.NewServer(cfg, router)
	log.Printf("INFO: Server starting on %s", cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
