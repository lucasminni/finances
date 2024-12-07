package debt

import uuid "github.com/satori/go.uuid"

type Debt struct {
	DebtId      uuid.UUID `gorm:"primary_key;type:uuid"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Value       float64   `gorm:"column:value"`
}

type DebtInputSchema struct {
	Name       string  `json:"name" binding:"required"`
	Descrition string  `json:"description" binding:"required"`
	Value      float64 `json:"value" binding:"required"`
}
