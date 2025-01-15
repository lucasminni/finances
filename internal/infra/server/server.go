package server

import (
	"finances/pkg/routes"
	v1 "finances/pkg/routes/v1"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	v1.Check(r)
	routes.Grouper(r)
	routes.Docs(r)

	serverPort := os.Getenv("SERVER_PORT")

	r.Run(fmt.Sprintf(":%s", serverPort))
}
