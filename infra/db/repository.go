package db

import (
	"financas/internal/model/debt"
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
