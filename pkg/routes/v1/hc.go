package v1

import "github.com/gin-gonic/gin"

func Check(r *gin.Engine) {
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
