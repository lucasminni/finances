package routes

import (
	"finances/pkg/handlers/income"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {

	r.GET("/income", income.List)
	r.POST("/income", income.Create)
	r.DELETE("/income/:id", income.Delete)
	r.PUT("/income", income.Update)

}
