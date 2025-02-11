package debt

import (
	"finances/internal/domain/models/debt"
	db "finances/internal/infra/db/repositories/debt"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Debt struct {
}

func (d *Debt) GetDebts(overdue *bool) []debt.Debt {
	return db.GetDebts(overdue)
}

func (d *Debt) CreateDebt(debt debt.Debt) (*debt.Debt, error) {
	debt.ID = uuid.NewV4()

	err := db.InsertDebt(debt)

	if err != nil {
		return nil, err
	}

	return &debt, nil
}

func (d *Debt) UpdateDebt(debt debt.Debt) (*debt.Debt, error) {
	debt.Overdue = debt.SetOverdue()

	err := db.UpdateDebt(debt)

	if err != nil {
		return nil, err
	}

	result, err := db.GetDebtById(debt.ID.String())

	return result, nil
}

func (d *Debt) DeleteDebt(id string) error {
	err := db.DeleteDebtByID(id)

	if err != nil {
		return err
	}

	return nil
}

func (d *Debt) UpdatePaymentStatus(id string, status bool) (*debt.Debt, error) {
	debt, err := db.GetDebtById(id)

	if err != nil {
		return nil, err
	}

	if status {

		debt.SetPaid()

		err := db.SetDebtPaid(*debt)

		if err != nil {
			return nil, err
		}

		log.Println("Debt ", debt.ID, "set paid")

		return debt, nil
	} else {

		debt.SetUnpaid()

		err := db.SetDebtUnpaid(*debt)

		if err != nil {
			return nil, err
		}

		log.Println("Debt ", debt.ID, "set unpaid")

		return debt, nil
	}
}

func (d *Debt) GetTotalPaidDebtValue() float64 {
	var totalDebtValue float64

	debts, _ := db.GetPaidDebts()

	for _, debt := range debts {
		totalDebtValue += *debt.Value
	}
	return totalDebtValue
}

func (d *Debt) GetTotalDebtValue() float64 {
	var totalDebtValue float64

	debts := db.GetDebts(nil)

	for _, debt := range debts {
		totalDebtValue += *debt.Value
	}

	return totalDebtValue
}
