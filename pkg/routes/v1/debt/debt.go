package routes

import (
	model "finances/internal/domain/models/debt"
	schema "finances/internal/domain/schemas/debt"
	"finances/internal/domain/services/debt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/debt", list)
	r.POST("/debt", create)
	r.DELETE("/debt/:id", delete)
	r.PUT("/debt", update)
	r.POST("/debt/payment/", pay)
}

var d debt.Debt

// @Summary      Show debts
// @Description  Lists all debts
// @Tags         debts
// @Accept       json
// @Produce      json
// @Param        overdue   query   string  false  "Debt id"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /debt [get]
func list(c *gin.Context) {
	var overdue *bool

	param := c.Query("overdue")
	err := c.ShouldBindUri(param)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if param == "" {
		overdue = nil
	} else {
		parsedOverdue, _ := strconv.ParseBool(param)
		overdue = &parsedOverdue
	}

	debts := d.GetDebts(overdue)

	c.JSON(http.StatusOK, gin.H{"debts": debts})
}

// @Summary      Create a debt
// @Description  Creates a new debt
// @Tags         debts
// @Accept       json
// @Produce      json
// @Param		 debt body		debt.BodyNewDebt	true	"Request body to add a debt"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /debt [post]
func create(c *gin.Context) {
	json := &model.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	debt, err := d.CreateDebt(*json)

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
// @Param        debt body      debt.BodyUpdateDebt  true  "Request body to update debt"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /debt [put]
func update(c *gin.Context) {
	json := &model.Debt{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding JSON error - " + err.Error()})
		log.Panic("Binding JSON error - " + err.Error())
	}

	debt, err := d.UpdateDebt(*json)

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
// @Param        id   path      string  true  "Debt id"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /debt/{id} [delete]
func delete(c *gin.Context) {
	uri := c.Param("id")

	err := c.ShouldBindUri(uri)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding param error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	err = d.DeleteDebt(uri)

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
// @Param		 debt body		debt.BodyUpdateDebtPaymentStatus	true	"Request body to update a debt payment status"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /debt/payment/ [post]
func pay(c *gin.Context) {

	json := &schema.BodyUpdateDebtPaymentStatus{}

	err := c.ShouldBindJSON(json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding URI error - " + err.Error()})
		log.Panic("Binding URI error - " + err.Error())
	}

	debt, err := d.UpdatePaymentStatus(json.ID.String(), json.Paid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update debt error - " + err.Error()})
		log.Panic("Update debt error - " + err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"id": debt.ID,
			"paid":        debt.Paid,
			"paymentDate": debt.PaymentDate})
	}

}
