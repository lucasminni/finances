package server

import (
	"finances/pkg/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	routes.Grouper(r)
	routes.Docs(r)

	serverPort := os.Getenv("SERVER_PORT")

	r.Run(fmt.Sprintf(":%s", serverPort))
}
