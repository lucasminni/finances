package routes

import (
	"financas/pkg/handlers/debt"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {

	r.GET("/debt", debt.GetDebts)
	r.POST("/debt", debt.InsertDebt)
	r.DELETE("/debt:id", debt.DeleteDebtById)
	r.PUT("/debt", debt.UpdateDebt)

}
