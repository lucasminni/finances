package routes

import (
	d "finances/internal/domain/models/debt"
	sd "finances/internal/domain/schemas/debt"
	s "finances/internal/domain/services/debt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/debt", list)
	r.POST("/debt", create)
	r.DELETE("/debt/:id", delete)
	r.PUT("/debt", update)
	r.POST("/debt/payment/", pay)
}

// @Summary      Show debts
// @Description  Lists all debts
// @Tags         debts
// @Accept       json
// @Produce      json
// @Success      200  {object}  d.Debt
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /debt [get]
func list(c *gin.Context) {
	debts := s.GetDebts()

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}

// @Summary      Create a debt
// @Description  Creates a new debt
// @Tags         debts
// @Accept       json
// @Produce      json
// @Param		 debt body		sd.InputNewDebt	true	"Add debt"
// @Success      200  {object}  d.Debt
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /debt [post]
func create(c *gin.Context) {
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

// @Summary      Update a debt
// @Description  Updates an existing debt
// @Tags         debts
// @Accept       json
// @Produce      json
// @Success      200  {object}  d.Debt
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /debt [put]
func update(c *gin.Context) {

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

// @Summary      Delete a debt
// @Description  Deletes an existing debt
// @Tags         debts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Debt ID"
// @Success      200  {object}  d.Debt
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /debt/{id} [delete]
func delete(c *gin.Context) {

	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = s.DeleteDebt(uri)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting debt error - " + err.Error()})
		log.Panic("Deleting debt error - " + err.Error())
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}

// @Summary      Set a debt paid/unpaid
// @Description  Sets an existing debt paid/unpaid
// @Tags         debts
// @Accept       json
// @Produce      json
// @Success      200  {object}  d.Debt
// @Failure      400  {object}  schemas.ErrorResponse
// @Failure      404  {object}  schemas.ErrorResponse
// @Failure      500  {object}  schemas.ErrorResponse
// @Router       /debt/payment/ [post]
func pay(c *gin.Context) {

	json := &sd.InputUpdatePaymentStatus{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding URI error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	debt, err := s.UpdatePaymentStatus(json.ID.String(), json.Paid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update debt error - " + err.Error()})
		log.Panic("Update debt error - " + err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"id": debt.ID,
			"paid":        debt.Paid,
			"paymentDate": debt.PaymentDate})
	}

}
