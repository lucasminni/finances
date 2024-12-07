package handlers

import (
	"financas/infra/db"
	"financas/internal/schemas/debt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func GetDebts(c *gin.Context) {
	var debts []debt.Debt

	debts = db.GetDebts()

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}

func InsertDebt(c *gin.Context) {
	json := &debt.DebtInputSchema{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	debt := debt.Debt{
		uuid.NewV4(),
		json.Name,
		json.Descrition,
		json.Value,
	}

	err = db.InsertDebt(debt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting debt error - " + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"debt": debt.Name, "debtId": debt.DebtId})
	}

}
