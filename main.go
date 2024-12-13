package main

import (
	"financas/internal/infra/db"
	"financas/pkg/routes"
	"fmt"
)

func main() {

	fmt.Println("Staring databse connection...")
	db.ConnectDB()

	fmt.Println("Starting server on port 8080...")
	routes.Start()

}
