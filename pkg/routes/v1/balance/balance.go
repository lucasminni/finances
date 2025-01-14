package routes

import (
	"finances/internal/domain/services/balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.RouterGroup) {
	r.GET("/balance", calculate)
}

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
	netBalance := balance.GetNetBalance()
	estimatedBalance := balance.GetEstimatedNetBalance()

	c.JSON(http.StatusOK, gin.H{"netBalance": netBalance, "estimatedBalance": estimatedBalance})
}
