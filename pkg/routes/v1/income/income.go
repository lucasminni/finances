package routes

import (
	"finances/pkg/handlers/income"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {

	r.GET("/income", income.GetIncomes)
	r.POST("/income", income.InsertIncome)
	r.DELETE("/income:id", income.DeleteIncomeById)
	r.PUT("/income", income.UpdateIncome)

}
