package cmd

import (
	"finances/internal/infra/db"
	"finances/internal/infra/server"
	"finances/internal/infra/settings"
	"log"
)

func Execute() {
	log.Println("Loading environment variables...")
	settings.LoadEnvs()

	log.Println("Starting database connection...")
	db.ConnectDatabase()

	log.Println("Starting server on port 8080...")
	server.StartServer()
}
