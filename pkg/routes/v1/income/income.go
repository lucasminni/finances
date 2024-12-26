package routes

import (
	model "finances/internal/domain/models/income"
	service "finances/internal/domain/services/income"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/income", list)
	r.POST("/income", create)
	r.DELETE("/income/:id", delete)
	r.PUT("/income", update)
}

// @Summary      Show incomes
// @Description  Lists all incomes
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /income [get]
func list(c *gin.Context) {
	incomes := service.GetIncomes()

	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}

// @Summary      Create an income
// @Description  Creates a new income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param		 income body	income.BodyNewIncome	true	"Request body to add an income"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /income [post]
func create(c *gin.Context) {
	json := &model.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	income, err := service.CreateIncome(*json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inserting income error - " + err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id": income.ID,
		})
	}
}

// @Summary      Update an income
// @Description  Updates an existing income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param		 income body	income.BodyUpdateIncome	true	"Request body to update an income"
// @Success      20
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /income [put]
func update(c *gin.Context) {
	json := &model.Income{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	income, err := service.UpdateIncome(*json)

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

// @Summary      Delete an income
// @Description  Deletes an existing income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Income id"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /income/{id} [delete]
func delete(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = service.DeleteIncome(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting income error - " + err.Error()})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
