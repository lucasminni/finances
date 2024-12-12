package income

import (
	db "financas/infra/db/repositories/income"
	"financas/internal/schemas/income"
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
		log.Println("Binding JSON error - " + err.Error())
	}

	newIncome := income.Income{
		ID:          uuid.NewV4(),
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
	}

	err = db.InsertIncome(newIncome)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting income error - " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"incomeId":    newIncome.ID,
			"income":      newIncome.Name,
			"description": newIncome.Description,
			"value":       newIncome.Value,
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
		log.Println("Binding JSON error - " + err.Error())
	}

	updateIncome := income.Income{
		ID:          json.ID,
		Name:        json.Name,
		Description: json.Description,
		Value:       json.Value,
	}

	log.Println(updateIncome.ID, updateIncome.Name, updateIncome.Description, updateIncome.Value)

	updatedIncome, err := db.UpdateIncome(updateIncome)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating income error - " + err.Error()})
		log.Println("Updating income error - " + err.Error())
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"incomeId":    updatedIncome.ID,
			"income":      updatedIncome.Name,
			"description": updatedIncome.Description,
			"value":       updatedIncome.Value,
		})
	}
}

func DeleteIncomeById(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Println("Binding URI error - " + err.Error())
	}

	err = db.DeleteIncomeByID(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting income error - " + err.Error()})
		log.Println("Deleting income error - " + err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
