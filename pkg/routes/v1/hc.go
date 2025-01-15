package v1

import (
	"finances/internal/infra/db"
	"github.com/gin-gonic/gin"
)

func Check(r *gin.Engine) {
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "true",
			"db":     CheckDatabase(),
		})
	})
}

func CheckDatabase() bool {
	return db.CheckDatabase()
}
