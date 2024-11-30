package router

import "github.com/gin-gonic/gin"

// doc:https://gin-gonic.com/docs/quickstart/

func RouterInit() {
	r := gin.Default()
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.Run(":8080")
}
