package router

import "github.com/gin-gonic/gin"

func RouterInit() {

	// Instacioando o router
	r := gin.Default()

	// Rota pra health check
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Inicia o router na porta 8080
	r.Run(":8080")
}
