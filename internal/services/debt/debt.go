package debt

import (
	db "finances/internal/infra/db/repositories/debt"
	d "finances/internal/schemas/debt"
	uuid "github.com/satori/go.uuid"
)

func GetDebts() []d.Debt {

	return db.GetDebts()

}

func CreateDebt(debt d.Debt) (*d.Debt, error) {

	debt.ID = uuid.NewV4()

	err := db.InsertDebt(debt)

	if err != nil {
		return nil, err
	}

	return &debt, nil

}

func UpdateDebt(debt d.Debt) (*d.Debt, error) {

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
