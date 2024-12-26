package routes

import (
	"finances/internal/domain/services/debt"
	"finances/internal/domain/services/income"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.RouterGroup) {
	r.GET("/balance", balance)
}

func balance(c *gin.Context) {
	balance := income.GetTotalIncomeValue() - debt.GetTotalDebtValue()

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
