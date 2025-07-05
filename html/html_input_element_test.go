package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
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

func (s *HTMLInputElementTestSuite) TestNameIDLAttribute() {
	e := html.NewHTMLInputElement(nil)
	e.SetName("Foo")

	s.Expect(e).To(HaveAttribute("name", "Foo"), "Name set as IDL attribute")

	e.SetName("Bar")
	s.Expect(e.Name()).To(Equal("Bar"))
}

func (s *HTMLInputElementTestSuite) TestValueDefaultsToContentAttribute() {
	e := html.NewHTMLInputElement(nil)
	e.SetAttribute("value", "foo")

	s.Expect(e.Value()).To(Equal("foo"))

	e.SetValue("bar")
	s.Expect(e.Value()).To(Equal("bar"))
	s.Expect(e).To(HaveAttribute("value", "foo"),
		"Set IDL value attribute doesn't affect content attribute")

	e.SetAttribute("value", "baz")
	s.Expect(e.Value()).To(Equal("bar"),
		"Changing value attribute doesn't affect IDL attribute once set")

	s.Expect(e).To(HaveAttribute("value", "baz"),
		"Set IDL value attribute doesn't affect content attribute")
}
