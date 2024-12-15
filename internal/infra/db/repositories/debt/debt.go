package debt

import (
	"errors"
	"finances/internal/infra/db"
	"finances/internal/schemas/debt"
	"log"
)

func InsertDebt(debt debt.Debt) error {
	result := db.SQLConnector.Create(&debt)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	}

	return nil
}

func GetDebts() []debt.Debt {
	var debts []debt.Debt
	result := db.SQLConnector.Find(&debts)

	if result.Error != nil {
		log.Panic(result.Error)
		return nil
	}

	return debts
}

func GetDebtById(id string) (*debt.Debt, error) {
	var result debt.Debt
	query := db.SQLConnector.First(&result, "id = ?", id)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateDebt(debt debt.Debt) (*debt.Debt, error) {
	result := db.SQLConnector.Save(&debt)

	if result.Error != nil {
		log.Panic(result.Error)
		return nil, result.Error
	} else {
		log.Println("Debt id " + debt.ID.String() + " updated")
	}

	updatedDebt, err := GetDebtById(debt.ID.String())

	if err != nil {
		return nil, err
	} else {
		return updatedDebt, nil
	}

}

func DeleteDebtByID(id string) error {

	if db.SQLConnector.First(&debt.Debt{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&debt.Debt{}, "id = ?", id)
		log.Println("Debt id " + id + " deleted")
		return nil
	} else {
		err := errors.New("Debt id " + id + " not found")
		log.Panic(err.Error())
		return err
	}

}
