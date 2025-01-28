package debt

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DebtTestSuite struct {
	suite.Suite
	service *Debt
}

func SetupTest(t *testing.T) {

}

func TestDebtTestSuite(t *testing.T) {
	suite.Run(t, new(DebtTestSuite))
}
