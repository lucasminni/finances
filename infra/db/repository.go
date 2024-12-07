package db

import (
	"financas/internal/schemas/debt"
	"fmt"
)

func GetDebts() []debt.Debt {
	var debts []debt.Debt
	result := db.Find(&debts)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return debts
}

func InsertDebt(debt debt.Debt) error {
	result := db.Create(&debt)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
