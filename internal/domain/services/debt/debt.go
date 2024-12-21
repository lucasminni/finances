package debt

import (
	d "finances/internal/domain/models/debt"
	db "finances/internal/infra/db/repositories/debt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetDebts() []d.Debt {
	debts := db.GetDebts()

	debts = db.UpdateOverdueDebt(debts)

	return debts
}

func CreateDebt(debt d.Debt) (*d.Debt, error) {
	debt.ID = uuid.NewV4()
	debt.Overdue = debt.SetOverdue()

	err := db.InsertDebt(debt)

	if err != nil {
		return nil, err
	}

	return &debt, nil
}

func UpdateDebt(debt d.Debt) (*d.Debt, error) {
	debt.Overdue = debt.SetOverdue()

	err := db.UpdateDebt(debt)

	if err != nil {
		return nil, err
	}

	result, err := db.GetDebtById(debt.ID.String())

	return result, nil
}

func DeleteDebt(id string) error {
	err := db.DeleteDebtByID(id)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePaymentStatus(id string, status bool) (*d.Debt, error) {
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

func GetTotalDebtValue() float64 {

	var totalDebtValue float64

	debts, _ := db.GetOutStandingDebts()

	for _, debt := range debts {
		totalDebtValue += debt.Value
	}
	return totalDebtValue
}
