package router

import (
	"financas/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RouterInit() {

	r := gin.Default()

	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.GET("/", handlers.GetDebts)
	r.POST("/", handlers.InsertDebt)
	r.DELETE("/:id", handlers.DeleteDebtById)
	r.PUT("/", handlers.UpdateDebt)

	r.Run(":8080")
}
