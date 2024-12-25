package server

import (
	routes "finances/pkg/routes"
	v1 "finances/pkg/routes/v1"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	v1.Check(r)
	routes.Grouper(r)
	routes.Docs(r)

	r.Run(":8080")
}
