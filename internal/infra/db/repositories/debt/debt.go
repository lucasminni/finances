package debt

import (
	"errors"
	d "finances/internal/domain/models/debt"
	"finances/internal/infra/db"
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

func GetOutStandingDebts() ([]d.Debt, error) {

	var debts []d.Debt

	query := db.SQLConnector.Where("paid = ?", false).Find(&debts)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	} else {
		return debts, nil
	}

}

func UpdateOverdueDebt(debts []d.Debt) []d.Debt {

	var updatedDebts []d.Debt

	for _, debt := range debts {
		debt.Overdue = debt.SetOverdue()
		db.SQLConnector.Model(&debt).Update("overdue", debt.Overdue)
		updatedDebts = append(updatedDebts, debt)
	}

	return updatedDebts

}
