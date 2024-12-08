package db

import (
	"errors"
	"financas/internal/schemas/debt"
	"fmt"
	"log"
)

func InsertDebt(debt debt.Debt) error {
	result := db.Create(&debt)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetDebts() []debt.Debt {
	var debts []debt.Debt
	result := db.Find(&debts)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return debts
}

func GetDebtById(id string) (*debt.Debt, error) {
	var result debt.Debt
	query := db.First(&result, "id = ?", id)

	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}

	return &result, nil
}

func UpdateDebt(debt debt.Debt) (*debt.Debt, error) {
	result := db.Save(&debt)

	if result.Error != nil {
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

	if db.First(&debt.Debt{}, "id = ?", id).RowsAffected > 0 {
		db.Delete(&debt.Debt{}, "id = ?", id)
		log.Println("Debt id " + id + " deleted")
		return nil
	} else {
		err := errors.New("Debt id " + id + " not found")
		return err
	}

}
