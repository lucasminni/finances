package income

import "github.com/dromara/carbon/v2"

type InputNewIncome struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Value        float64     `json:"value"`
	ReceivedDate carbon.Date `json:"receivedDate" swaggertype:"string" example:"2024-12-25"`
}
