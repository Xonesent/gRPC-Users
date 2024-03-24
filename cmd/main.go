package main

import (
	"log"

	"users/config"
	"users/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	psqlDB, err := server.NewDB(cfg.Postgres)
	if err != nil {
		log.Fatalf("psqlDB: %v", err)
	}

	s := server.NewServer(
		cfg,
		psqlDB,
	)

	if err = s.Run(); err != nil {
		log.Printf("Cannot start server: %v", err)
		return
	}
}
