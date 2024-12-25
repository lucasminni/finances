package debt

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type BodyNewDebt struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Value       float64     `json:"value"`
	DueDate     carbon.Date `json:"dueDate" swaggertype:"string" example:"2024-12-25"`
}

type BodyUpdateDebt struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Value       float64     `json:"value"`
	DueDate     carbon.Date `json:"dueDate" swaggertype:"string" example:"2024-12-25"`
	PaymentDate carbon.Date `json:"paymentDate" swaggertype:"string" example:"2024-12-25"`
}
type BodyUpdateDebtPaymentStatus struct {
	ID   uuid.UUID `json:"id"`
	Paid bool      `json:"paid"`
}
