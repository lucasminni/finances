package router

import (
	"financas/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RouterInit() {

	// Instacioando o router
	r := gin.Default()

	// Rota pra health check
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Rota para recuperar todos os debts
	r.GET("/", handlers.GetDebts)

	// Inicia o router na porta 8080
	r.Run(":8080")
}
