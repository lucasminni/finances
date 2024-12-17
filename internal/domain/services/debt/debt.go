package debt

import (
	d "finances/internal/domain/models/debt"
	db "finances/internal/infra/db/repositories/debt"
	uuid "github.com/satori/go.uuid"
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
