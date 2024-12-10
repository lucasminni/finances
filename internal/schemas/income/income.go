package income

import uuid "github.com/satori/go.uuid"

type Income struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name" binding:"required"`
	Description string    `gorm:"column:description" json:"description" binding:"required"`
	Value       float64   `gorm:"column:value" json:"value" binding:"required"`
}
