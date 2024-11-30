package debt

import "time"

type Debt struct {
	Name        string
	Description string
	Value       float64
	DueDate     time.Time
	PaymentDate time.Time
}
