package debt

import (
	db "finances/internal/infra/db/repositories/debt"
	"finances/internal/schemas/debt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func InsertDebt(c *gin.Context) {
	json := &debt.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	newDebt := debt.Debt{
		ID:          uuid.NewV4(),
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
		DueDate:     json.DueDate,
		PaymentDate: json.PaymentDate,
	}

	err = db.InsertDebt(newDebt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting debt error - " + err.Error()})
		log.Panic("Inserting debt error - " + err.Error())
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id":          newDebt.ID,
			"name":        newDebt.Name,
			"description": newDebt.Description,
			"value":       newDebt.Value,
			"dueDate":     newDebt.DueDate,
			"paymentDate": newDebt.PaymentDate,
		})
	}

}

func GetDebts(c *gin.Context) {
	var debts []debt.Debt

	debts = db.GetDebts()

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}

func UpdateDebt(c *gin.Context) {
	json := &debt.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	updateDebt := debt.Debt{
		ID:          json.ID,
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
		DueDate:     json.DueDate,
		PaymentDate: json.PaymentDate,
	}

	updatedDebt, err := db.UpdateDebt(updateDebt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating debt error - " + err.Error()})
		log.Panic("Updating debt error - " + err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":          updatedDebt.ID,
			"name":        updatedDebt.Name,
			"description": updatedDebt.Description,
			"value":       updatedDebt.Value,
			"dueDate":     updateDebt.DueDate,
			"paymentDate": updatedDebt.PaymentDate,
		})
	}
}

func DeleteDebtById(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = db.DeleteDebtByID(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting debt error - " + err.Error()})
		log.Panic("Deleting debt error - " + err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
