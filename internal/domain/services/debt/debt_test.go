package debt

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DebtTestSuite struct {
	suite.Suite
	debt    *Debt
	service *Debt
}

func TestDebtTestSuite(t *testing.T) {
	suite.Run(t, new(DebtTestSuite))
}

func (s *DebtTestSuite) SetupTest() {
	s.debt = &Debt{}
}

//
//func (s *DebtTestSuite) TestGetDebts_WhenOverdueTrue() {
//	overdue := true
//	expectedDebts := []debt.Debt{
//		{
//			ID:      uuid.NewV4(),
//			Value:   100.00,
//			Overdue: true,
//			Paid:    false,
//		},
//		{
//			ID:      uuid.NewV4(),
//			Value:   200.00,
//			Overdue: true,
//			Paid:    false,
//		},
//	}
//
//	// Mock the repository functions
//	debt.GetDebts = func(overdue *bool) []debt.Debt {
//		return expectedDebts
//	}
//
//	debt.UpdateOverdueDebt = func(debts []debt.Debt) []debt.Debt {
//		return debts
//	}
//
//	// Act
//	actualDebts := s.service.GetDebts(&overdue)
//
//	// Assert
//	s.Equal(expectedDebts, actualDebts)
//}
