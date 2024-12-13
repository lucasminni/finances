package routes

import (
	"financas/pkg/routes/v1"
	financialV1 "financas/pkg/routes/v1"
	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()

	v1.Check(r)
	Grouper(r)

	r.Run(":8080")
}

func Grouper(r *gin.Engine) {
	group := r.Group("/api")

	v1 := group.Group("/v1")

	financialV1.Register(v1.Group("/financials"))
}
