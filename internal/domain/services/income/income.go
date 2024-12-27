package income

import (
	"finances/internal/domain/models/income"
	db "finances/internal/infra/db/repositories/income"
	uuid "github.com/satori/go.uuid"
	"log"
)

func GetIncomes() []income.Income {
	return db.GetIncomes()
}

func CreateIncome(income income.Income) (*income.Income, error) {
	income.ID = uuid.NewV4()

	err := db.InsertIncome(income)

	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}
	return &income, nil
}

func UpdateIncome(income income.Income) (*income.Income, error) {
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

func GetTotalIncomeValue() float64 {
	var totalIncomeValue float64

	incomes := db.GetIncomes()

	for _, income := range incomes {
		totalIncomeValue += *income.Value
	}

	return totalIncomeValue
}
