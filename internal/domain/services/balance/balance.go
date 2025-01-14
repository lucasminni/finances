package balance

import (
	"finances/internal/domain/services/debt"
	"finances/internal/domain/services/income"
)

func GetCurrentBalance() float64 {
	debtTotal := debt.GetTotalPaidDebtValue()
	incomeTotal := income.GetTotalIncomeValue()

	return incomeTotal - debtTotal

}

func GetEstimatedNetBalance() float64 {
	debtTotal := debt.GetTotalDebt()
	incomeTotal := income.GetTotalIncomeValue()

	return incomeTotal - debtTotal
}
