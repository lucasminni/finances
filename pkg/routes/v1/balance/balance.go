package routes

import (
	"finances/internal/domain/services/balance"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/balance", calculate)
}

var b balance.Balance

// @Summary      Show balance
// @Tags         balance
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /balance [get]
func calculate(c *gin.Context) {
	current := b.GetCurrentBalance(b.GetTotalIncome(), b.GetTotalPaidDebts())
	estimated := b.GetEstimatedNetBalance(b.GetTotalIncome(), b.GetTotalDebt())

	c.JSON(http.StatusOK, gin.H{"current": current, "estimated": estimated})
}
