package handlers

import (
	"financas/infra/db"
	"financas/internal/model/debt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDebts(c *gin.Context) {
	var debts []debt.Debt

	debts = db.GetDebts()

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}
