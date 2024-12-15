package debt

import (
	db "finances/internal/infra/db/repositories/debt"
	d "finances/internal/schemas/debt"
	s "finances/internal/services/debt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func List(c *gin.Context) {
	debts := s.GetDebts()

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}

func Create(c *gin.Context) {
	json := &d.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	debt, err := s.CreateDebt(*json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting debt error - " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": debt.ID,
		})
	}

}

func Update(c *gin.Context) {
	json := &d.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	debt, err := s.UpdateDebt(*json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating debt error - " + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":          debt.ID,
			"name":        debt.Name,
			"description": debt.Description,
			"value":       debt.Value,
			"dueDate":     debt.DueDate,
			"paymentDate": debt.PaymentDate,
		})
	}
}

func Delete(c *gin.Context) {
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
