package server

import (
	"finances/pkg/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	routes.Grouper(r)
	routes.Docs(r)

	serverPort := os.Getenv("SERVER_PORT")
	protocol := os.Getenv("SERVER_PROTOCOL")

	switch protocol {
	case "http":
		if err := r.Run(fmt.Sprintf(":%s", serverPort)); err != nil {
			log.Fatal("Error on server start: ", err)
		}
	case "https":
		if err := r.RunTLS(fmt.Sprintf(":%s", serverPort), os.Getenv("SERVER_CERT"), os.Getenv("SERVER_KEY")); err != nil {
			log.Fatal("Error on server start: ", err)
		}
	}

}
