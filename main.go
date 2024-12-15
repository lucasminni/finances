package main

import (
	"finances/internal/infra/db"
	"finances/pkg/routes"
	"log"
)

func main() {

	log.Println("Staring databse connection...")
	db.ConnectDatabase()

	log.Println("Starting server on port 8080...")
	routes.StartServer()

}
