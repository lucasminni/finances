package income

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type Income struct {
	ID           uuid.UUID   `gorm:"primary_key;type:uuid;column:id" json:"id"`
	Name         string      `gorm:"column:name" json:"name" binding:"required"`
	Description  string      `gorm:"column:description" json:"description" binding:"required"`
	Value        float64     `gorm:"column:value" json:"value" binding:"required"`
	ReceivedDate carbon.Date `gorm:"column:received_date;type:date" json:"receivedDate" binding:"required" swaggertype:"string" example:"2024-12-25"`
}
