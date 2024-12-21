package debt

import uuid "github.com/satori/go.uuid"

type InputUpdatePaymentStatus struct {
	ID   uuid.UUID `json:"id"`
	Paid bool      `json:"paid"`
}
