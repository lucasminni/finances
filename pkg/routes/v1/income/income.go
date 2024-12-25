package routes

import (
	i "finances/internal/domain/models/income"
	s "finances/internal/domain/services/income"
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
// @Success      200  {object}  i.Income
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /income [get]
func list(c *gin.Context) {
	incomes := s.GetIncomes()

	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
}

// @Summary      Create an income
// @Description  Creates a new income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Success      200  {object}  i.Income
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /income [post]
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

// @Summary      Update an income
// @Description  Updates an existing income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Success      200  {object}  i.Income
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /income [put]
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

// @Summary      Delete an income
// @Description  Deletes an existing income
// @Tags         incomes
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Debt ID"
// @Success      200  {object}  i.Income
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /income/{id} [delete]
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
