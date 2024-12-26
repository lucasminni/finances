package income

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type BodyNewIncome struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Value        float64     `json:"value"`
	ReceivedDate carbon.Date `json:"receivedDate" swaggertype:"string" example:"2024-12-25"`
}

type BodyUpdateIncome struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Value        float64       `json:"value"`
	ReceivedDate carbon.Carbon `json:"receivedDate" swaggertype:"string" example:"2024-12-25"`
}
