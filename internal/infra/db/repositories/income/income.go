package income

import (
	"errors"
	"financas/internal/infra/db"
	"financas/internal/schemas/income"
	"fmt"
	"log"
)

func InsertIncome(income income.Income) error {
	result := db.SQLConnector.Create(&income)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetIncomes() []income.Income {
	var incomes []income.Income
	result := db.SQLConnector.Find(&incomes)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return incomes
}

func GetincomeById(id string) (*income.Income, error) {
	var result income.Income
	query := db.SQLConnector.First(&result, "id = ?", id)

	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateIncome(income income.Income) (*income.Income, error) {
	result := db.SQLConnector.Save(&income)

	if result.Error != nil {
		return nil, result.Error
	} else {
		log.Println("income id " + income.ID.String() + " updated")
	}

	updatedincome, err := GetincomeById(income.ID.String())

	if err != nil {
		return nil, err
	} else {
		return updatedincome, nil
	}

}

func DeleteIncomeByID(id string) error {

	if db.SQLConnector.First(&income.Income{}, "id = ?", id).RowsAffected > 0 {
		db.SQLConnector.Delete(&income.Income{}, "id = ?", id)
		log.Println("income id " + id + " deleted")
		return nil
	} else {
		err := errors.New("income id " + id + " not found")
		return err
	}

}
