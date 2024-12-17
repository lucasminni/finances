package income

import (
	"errors"
	i "finances/internal/domain/models/income"
	"finances/internal/infra/db"
	"log"
)

func InsertIncome(income i.Income) error {
	result := db.SQLConnector.Create(&income)

	if result.Error != nil {
		log.Panic(result.Error)
		return result.Error
	}

	return nil
}

func GetIncomes() []i.Income {
	var incomes []i.Income
	result := db.SQLConnector.Find(&incomes)

	if result.Error != nil {
		log.Panic(result.Error)
		return nil
	}

	return incomes
}

func GetincomeById(id string) (*i.Income, error) {
	var result i.Income
	query := db.SQLConnector.First(&result, "id = ?", id)

	if query.Error != nil {
		log.Panic(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateIncome(income i.Income) error {
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

	if db.SQLConnector.First(&i.Income{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&i.Income{}, "id = ?", id)
		log.Println("income id " + id + " deleted")
		return nil
	} else {
		err := errors.New("income id " + id + " not found")
		log.Panic(err.Error())
		return err
	}

}
