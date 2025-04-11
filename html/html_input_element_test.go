package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/eventtest"
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

func (s *HTMLInputElementTestSuite) TestClickCheckbox() {
	e := html.NewHTMLInputElement(nil)
	e.SetType("checkbox")
	s.Assert().False(e.Checked(), "Checkbox should be false by default")

	e.Click()
	s.Assert().True(e.Checked(), "Checkbox should be true when clicked")

	e.AddEventListener("click", eventtest.PreventDefaultHandler())
	e.Click()
	s.Assert().True(e.Checked(), "Checked should not change when default is prevented")
}
