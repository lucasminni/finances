package router

import (
	"financas/internal/handlers/debt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {

	r := gin.Default()

	r.GET("/hc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.GET("/", debt.GetDebts)
	r.POST("/", debt.InsertDebt)
	r.DELETE("/:id", debt.DeleteDebtById)
	r.PUT("/", debt.UpdateDebt)

	r.Run(":8080")
}
