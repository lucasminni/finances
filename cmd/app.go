package cmd

import (
	"finances/internal/infra/db"
	"finances/internal/infra/server"
	"log"
)

func Execute() {
	log.Println("Starting database connection...")
	db.ConnectDatabase()

	log.Println("Starting server on port 8080...")
	server.StartServer()
}
