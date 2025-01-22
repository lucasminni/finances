package debt

import (
	"errors"
	"finances/internal/domain/models/debt"
	"finances/internal/infra/db"
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

func GetDebts(overdue *bool) []debt.Debt {
	var debts []debt.Debt

	if overdue == nil {
		query := db.SQLConnector.Find(&debts)

		if query.Error != nil {
			log.Panic(query.Error)
			return nil
		}

		return debts
	}

	if overdue != nil {
		switch *overdue {
		case true:
			query := db.SQLConnector.Find(&debts, "overdue = ?", true)

			if query.Error != nil {
				log.Panic(query.Error)
				return nil
			}

		case false:
			query := db.SQLConnector.Find(&debts, "overdue = ?", false)

			if query.Error != nil {
				log.Panic(query.Error)
				return nil
			}
		}

		return debts
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

func UpdateDebt(debt debt.Debt) error {
	result := db.SQLConnector.Save(&debt)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	} else {
		return nil
	}
}

func DeleteDebtByID(id string) error {
	if db.SQLConnector.First(&debt.Debt{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&debt.Debt{}, "id = ?", id)
		log.Println("Debt id " + id + " deleted")
		return nil
	} else {
		err := errors.New("Debt id " + id + " not found")
		log.Println(err.Error())
		return err
	}
}

func GetPaidDebts() ([]debt.Debt, error) {
	var debts []debt.Debt

	if db.SQLConnector == nil {
		log.Println("db.SQLConnector is nil")
	}

	query := db.SQLConnector.Where("paid = ?", true).Find(&debts)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	} else {
		return debts, nil
	}
}

func UpdateOverdueDebt(debts []debt.Debt) []debt.Debt {
	var updatedDebts []debt.Debt

	for _, debt := range debts {
		debt.Overdue = debt.SetOverdue()
		db.SQLConnector.Model(&debt).Update("overdue", debt.Overdue)
		updatedDebts = append(updatedDebts, debt)
	}

	return updatedDebts
}

func SetDebtPaid(debt debt.Debt) error {
	err := db.SQLConnector.Model(&debt).Updates(map[string]interface{}{"paid": debt.Paid, "payment_date": debt.PaymentDate})

	if err != nil {
		return err.Error
	}

	return nil
}

func SetDebtUnpaid(debt debt.Debt) error {
	err := db.SQLConnector.Model(&debt).Updates(map[string]interface{}{"paid": debt.Paid, "payment_date": nil})

	if err != nil {
		return err.Error
	}

	err = db.SQLConnector.Model(&debt).Update("payment_date", nil)

	if err != nil {
		return err.Error
	}

	return nil
}
