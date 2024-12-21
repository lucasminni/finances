package debt

import (
	"github.com/dromara/carbon/v2"
	uuid "github.com/satori/go.uuid"
)

type Debt struct {
	ID          uuid.UUID   `gorm:"primary_key;type:uuid;column:id" json:"id"`
	Name        string      `gorm:"column:name" json:"name" binding:"required"`
	Description string      `gorm:"column:description" json:"description" binding:"required"`
	Value       float64     `gorm:"column:value" json:"value" binding:"required"`
	DueDate     carbon.Date `gorm:"column:due_date;type:date" json:"dueDate" binding:"required"`
	PaymentDate carbon.Date `gorm:"column:payment_date;type:date;default:null" json:"paymentDate"`
	Overdue     bool        `gorm:"column:overdue;default:null" json:"overdue"`
	Paid        bool        `gorm:"column:paid;default:false" json:"paid"`
}

func (d *Debt) SetPaid() {
	d.Paid = true
	d.PaymentDate = carbon.Date{Carbon: carbon.Now()}
}

func (d *Debt) SetUnpaid() {
	d.Paid = false
	d.PaymentDate = carbon.Date{}
}

func (d *Debt) SetOverdue() bool {
	if carbon.Now().ToDateString() > d.DueDate.String() {
		return true
	}

	return false
}
