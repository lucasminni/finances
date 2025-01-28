package balance

import (
	"finances/internal/domain/models/debt"
	"finances/internal/domain/models/income"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BalanceTestSuite struct {
	suite.Suite
	income  income.Income
	debt    debt.Debt
	service *Balance
}

func TestBalanceTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceTestSuite))
}

func (s *BalanceTestSuite) SetupTest() {
	incomeValue := 1000.00
	debtValue := 100.00

	s.income = income.Income{
		Value: &incomeValue,
	}

	s.debt = debt.Debt{
		Value:   &debtValue,
		Overdue: false,
		Paid:    false,
	}
}

func (s *BalanceTestSuite) TestVerifyEstimatedBalance() {
	// act
	result := s.service.GetEstimatedNetBalance(*s.income.Value, *s.debt.Value)

	//assert
	s.Equal(result, 900.0)
}

func (s *BalanceTestSuite) TestVerifyCurrentBalance() {
	// act
	result := s.service.GetCurrentBalance(*s.income.Value, *s.debt.Value)

	//assert
	s.Equal(result, 900.0)
}
