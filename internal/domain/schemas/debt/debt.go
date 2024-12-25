package debt

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type InputUpdatePaymentStatus struct {
	ID   uuid.UUID `json:"id"`
	Paid bool      `json:"paid"`
}

type InputNewDebt struct {
	Name        string      `gorm:"column:name" json:"name" binding:"required"`
	Description string      `gorm:"column:description" json:"description" binding:"required"`
	Value       float64     `gorm:"column:value" json:"value" binding:"required"`
	DueDate     carbon.Date `gorm:"column:due_date;type:date" json:"dueDate" binding:"required" swaggertype:"string" example:"2024-12-25"`
}
