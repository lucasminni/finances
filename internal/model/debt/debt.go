package debt

import "time"

type Debt struct {
	Name        string
	Description string
	Value       float64
	DueDate     time.Time `gorm:"column:due_date"`
	PaymentDate time.Time `gorm:"column:payment_date"`
}
