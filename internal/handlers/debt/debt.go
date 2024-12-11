package debt

import (
	db "financas/infra/db/debt"
	"financas/internal/schemas/debt"
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
		log.Println("Binding JSON error - " + err.Error())
	}

	newDebt := debt.Debt{
		ID:          uuid.NewV4(),
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
	}

	err = db.InsertDebt(newDebt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting debt error - " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"debtId":      newDebt.ID,
			"debt":        newDebt.Name,
			"description": newDebt.Description,
			"value":       newDebt.Value,
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
		log.Println("Binding JSON error - " + err.Error())
	}

	updateDebt := debt.Debt{
		ID:          json.ID,
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
	}

	log.Println(updateDebt.ID, updateDebt.Name, updateDebt.Description, updateDebt.Value)

	updatedDebt, err := db.UpdateDebt(updateDebt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating debt error - " + err.Error()})
		log.Println("Updating debt error - " + err.Error())
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"debtId":      updatedDebt.ID,
			"debt":        updatedDebt.Name,
			"description": updatedDebt.Description,
			"value":       updatedDebt.Value,
		})
	}
}

func DeleteDebtById(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Println("Binding URI error - " + err.Error())
	}

	err = db.DeleteDebtByID(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting debt error - " + err.Error()})
		log.Println("Deleting debt error - " + err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
