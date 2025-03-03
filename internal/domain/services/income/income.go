package income

import (
	"finances/internal/domain/models/income"
	db "finances/internal/infra/db/repositories/income"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Income struct {
}

func (i *Income) GetIncomes() []income.Income {
	return db.GetIncomes()
}

func (i *Income) CreateIncome(income income.Income) (*income.Income, error) {
	income.ID = uuid.NewV4()

	err := db.InsertIncome(income)

	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}
	return &income, nil
}

func (i *Income) UpdateIncome(income income.Income) (*income.Income, error) {
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

func (i *Income) DeleteIncome(id string) error {
	err := db.DeleteIncomeByID(id)

	if err != nil {
		log.Panic(err.Error())
		return err
	}
	return nil
}

func (i *Income) GetTotalIncomeValue() float64 {
	var totalIncomeValue float64

	incomes := db.GetIncomes()

	for _, income := range incomes {
		totalIncomeValue += *income.Value
	}

	return totalIncomeValue
}
