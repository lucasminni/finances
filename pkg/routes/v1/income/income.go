package routes

import (
	i "finances/internal/domain/models/income"
	s "finances/internal/domain/services/income"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(r *gin.RouterGroup) {
	r.GET("/income", list)
	r.POST("/income", create)
	r.DELETE("/income/:id", delete)
	r.PUT("/income", update)
}

func list(c *gin.Context) {
	incomes := s.GetIncomes()

	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}

func create(c *gin.Context) {
	json := &i.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	income, err := s.CreateIncome(*json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting income error - " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": income.ID,
		})
	}
}

func update(c *gin.Context) {
	json := &i.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	income, err := s.UpdateIncome(*json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating income error - " + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":           income.ID,
			"name":         income.Name,
			"description":  income.Description,
			"value":        income.Value,
			"receivedDate": income.ReceivedDate,
		})
	}
}

func delete(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = s.DeleteIncome(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting income error - " + err.Error()})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
