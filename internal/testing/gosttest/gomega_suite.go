package gosttest

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/suite"
)

type GomegaSuite struct {
	suite.Suite
	gomega gomega.Gomega
}

func (s *GomegaSuite) Expect(actual interface{}, extra ...interface{}) types.Assertion {
	if s.gomega == nil {
		s.gomega = gomega.NewWithT(s.T())
	}
	return s.gomega.Expect(actual, extra...)
}
