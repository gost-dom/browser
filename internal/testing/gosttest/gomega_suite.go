package gosttest

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/suite"
)

// GomegaSuite is a specialised [suite.Suite] that provides support for the
// [gomega] assertion library by providing a [GomegaSuite.Expect] function
// automatically initializes gomega with
type GomegaSuite struct {
	suite.Suite
	gomega gomega.Gomega
}

// Expect is wraps [gomega/Gomega.Expect], passing the correct [testing/T]. This
// supports gomega matchers for verification; which can in some cases provide
// more expressive tests.
func (s *GomegaSuite) Expect(actual interface{}, extra ...interface{}) types.Assertion {
	if s.gomega == nil {
		s.gomega = gomega.NewWithT(s.T())
	}
	return s.gomega.Expect(actual, extra...)
}
