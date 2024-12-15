package debt

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type Debt struct {
	ID          uuid.UUID    `gorm:"primary_key;type:uuid;column:id" json:"id"`
	Name        string       `gorm:"column:name" json:"name" binding:"required"`
	Description string       `gorm:"column:description" json:"description" binding:"required"`
	Value       float64      `gorm:"column:value" json:"value" binding:"required"`
	DueDate     carbon.Date  `gorm:"column:due_date;type:date" json:"dueDate" binding:"required"`
	PaymentDate *carbon.Date `gorm:"column:payment_date;type:date" json:"paymentDate"`
}
