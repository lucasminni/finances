package balance

import (
	"finances/internal/domain/services/debt"
	"finances/internal/domain/services/income"
)

type Balance struct {
}

var d debt.Debt

func (b *Balance) GetCurrentBalance(incomeTotal float64, totalPaidDebts float64) float64 {
	return incomeTotal - totalPaidDebts
}

func (b *Balance) GetEstimatedNetBalance(incomeTotal float64, debtTotal float64) float64 {
	return incomeTotal - debtTotal
}

func (b *Balance) GetTotalDebt() float64 {
	return d.GetTotalDebtValue()
}

func (b *Balance) GetTotalIncome() float64 {
	return income.GetTotalIncomeValue()
}

func (b *Balance) GetTotalPaidDebts() float64 {
	return d.GetTotalPaidDebtValue()
}
