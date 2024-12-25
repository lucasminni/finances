package main

import (
	"finances/internal/infra/db"
	"finances/internal/infra/server"
	"log"
)

func main() {
	log.Println("Starting database connection...")
	db.ConnectDatabase()

	log.Println("Starting server on port 8080...")
	server.StartServer()
}
