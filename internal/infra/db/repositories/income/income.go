package income

import (
	"errors"
	"finances/internal/domain/models/income"
	"finances/internal/infra/db"
	"log"
)

func InsertIncome(income income.Income) error {
	result := db.SQLConnector.Create(&income)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	}

	return nil
}

func GetIncomes() []income.Income {
	var incomes []income.Income
	query := db.SQLConnector.Find(&incomes)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil
	}

	return incomes
}

func GetincomeById(id string) (*income.Income, error) {
	var result income.Income
	query := db.SQLConnector.First(&result, "id = ?", id)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateIncome(income income.Income) error {
	result := db.SQLConnector.Save(&income)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	} else {
		log.Println("income id " + income.ID.String() + " updated")
		return nil
	}
}

func DeleteIncomeByID(id string) error {

	if db.SQLConnector.First(&income.Income{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&income.Income{}, "id = ?", id)
		log.Println("income id " + id + " deleted")
		return nil
	} else {
		err := errors.New("income id " + id + " not found")
		log.Panic(err.Error())
		return err
	}

}
