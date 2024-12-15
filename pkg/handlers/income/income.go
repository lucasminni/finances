package income

import (
	db "finances/internal/infra/db/repositories/income"
	"finances/internal/schemas/income"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func InsertIncome(c *gin.Context) {
	json := &income.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	newIncome := income.Income{
		ID:           uuid.NewV4(),
		Name:         json.Name,
		Description:  json.Description,
		Value:        json.Value,
		ReceivedDate: json.ReceivedDate,
	}

	err = db.InsertIncome(newIncome)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting income error - " + err.Error()})
		log.Panic("Inserting income error - " + err.Error())
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id":           newIncome.ID,
			"name":         newIncome.Name,
			"description":  newIncome.Description,
			"value":        newIncome.Value,
			"receivedDate": newIncome.ReceivedDate,
		})
	}

}

func GetIncomes(c *gin.Context) {
	var incomes []income.Income

	incomes = db.GetIncomes()

	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}

func UpdateIncome(c *gin.Context) {
	json := &income.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	updateIncome := income.Income{
		ID:           json.ID,
		Name:         json.Name,
		Description:  json.Description,
		Value:        json.Value,
		ReceivedDate: json.ReceivedDate,
	}

	updatedIncome, err := db.UpdateIncome(updateIncome)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating income error - " + err.Error()})
		log.Panic("Updating income error - " + err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":           updatedIncome.ID,
			"name":         updatedIncome.Name,
			"description":  updatedIncome.Description,
			"value":        updatedIncome.Value,
			"receivedDate": updatedIncome.ReceivedDate,
		})
	}
}

func DeleteIncomeById(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = db.DeleteIncomeByID(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting income error - " + err.Error()})
		log.Panic("Deleting income error - " + err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
