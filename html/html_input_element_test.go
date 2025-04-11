package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/suite"
)

type HTMLInputElementTestSuite struct {
	gosttest.GomegaSuite
}

func TestHTMLInputElement(t *testing.T) {
	suite.Run(t, new(HTMLInputElementTestSuite))
}

func (s *HTMLInputElementTestSuite) TestDefaultValue() {
	e := html.NewHTMLInputElement(nil)
	s.Expect(e.Type()).To(Equal("text"))
}
