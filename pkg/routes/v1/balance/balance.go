package routes

import (
	sd "finances/internal/domain/services/debt"
	si "finances/internal/domain/services/income"
	"github.com/gin-gonic/gin"
	"net/http"
)

//internal/domain/services/debt/debt.go:84

func Register(r *gin.RouterGroup) {

	r.GET("/balance", balance)

}

func balance(c *gin.Context) {

	balance := si.GetTotalIncomeValue() - sd.GetTotalDebtValue()

	c.JSON(http.StatusOK, gin.H{"balance": balance})

}
