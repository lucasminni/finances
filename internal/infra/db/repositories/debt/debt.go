package debt

import (
	"errors"
	"finances/internal/infra/db"
	d "finances/internal/schemas/debt"
	"log"
)

func InsertDebt(debt d.Debt) error {
	result := db.SQLConnector.Create(&debt)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	}

	return nil
}

func GetDebts() []d.Debt {
	var debts []d.Debt
	result := db.SQLConnector.Find(&debts)

	if result.Error != nil {
		log.Panic(result.Error)
		return nil
	}

	return debts
}

func GetDebtById(id string) (*d.Debt, error) {
	var result d.Debt
	query := db.SQLConnector.First(&result, "id = ?", id)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateDebt(debt d.Debt) error {
	result := db.SQLConnector.Save(&debt)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	} else {
		log.Println("Debt id " + debt.ID.String() + " updated")
		return nil
	}

}

func DeleteDebtByID(id string) error {

	if db.SQLConnector.First(&d.Debt{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&d.Debt{}, "id = ?", id)
		log.Println("Debt id " + id + " deleted")
		return nil
	} else {
		err := errors.New("Debt id " + id + " not found")
		log.Panic(err.Error())
		return err
	}

}
