package main

import (
	"financas/infra/db"
	"financas/internal/routers"
	"fmt"
)

func main() {

	fmt.Println("Staring databse connection...")
	db.ConnectDB()

	fmt.Println("Starting server on port 8080...")
	routers.RouterInit()

}
