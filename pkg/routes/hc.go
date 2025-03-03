package routes

import (
	"finances/internal/infra/db"

	"github.com/gin-gonic/gin"
)

func CheckDatabase(r *gin.RouterGroup) {
	r.GET("/db", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"db": db.CheckDatabase(),
		})
	})
}

func CheckApp(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
		})
	})
}
