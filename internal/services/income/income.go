package debt

import (
	db "finances/internal/infra/db/repositories/income"
	i "finances/internal/schemas/income"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetIncomes() []i.Income {

	return db.GetIncomes()

}

func InsertIncome(income i.Income) (*i.Income, error) {

	income.ID = uuid.NewV4()

	err := db.InsertIncome(income)

	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}

	return &income, nil

}

func UpdateIncome(income i.Income) (*i.Income, error) {

	err := db.UpdateIncome(income)

	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}

	result, err := db.GetincomeById(income.ID.String())

	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}

	return result, nil
}

func DeleteIncome(id string) error {
	err := db.DeleteIncomeByID(id)

	if err != nil {
		log.Panic(err.Error())
		return err
	}

	return nil
}
