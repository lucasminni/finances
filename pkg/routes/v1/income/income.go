package routes

import (
	h "finances/pkg/handlers/income"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {

	r.GET("/income", h.List)
	r.POST("/income", h.Create)
	r.DELETE("/income/:id", h.Delete)
	r.PUT("/income", h.Update)

}
