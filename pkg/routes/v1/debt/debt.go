package routes

import (
	h "finances/pkg/handlers/debt"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {

	r.GET("/debt", h.List)
	r.POST("/debt", h.Create)
	r.DELETE("/debt/:id", h.Delete)
	r.PUT("/debt", h.Update)

}
