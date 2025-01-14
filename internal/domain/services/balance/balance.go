package balance

import (
	"finances/internal/domain/services/debt"
	"finances/internal/domain/services/income"
)

func GetNetBalance() float64 {
	debtTotal := debt.GetTotalOutstandingDebtValue()
	incomeTotal := income.GetTotalIncomeValue()

	return incomeTotal - debtTotal

}

func GetEstimatedNetBalance() float64 {
	debtTotal := debt.GetTotalDebt()
	incomeTotal := income.GetTotalIncomeValue()

	return incomeTotal - debtTotal
}
