package html_test

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/suite"

	. "github.com/onsi/gomega"
)

type HTMLFormElementTestSuite struct {
	gosttest.GomegaSuite
	win  html.Window
	doc  html.HTMLDocument
	form html.HTMLFormElement
}

func TestHTMLFormElement(t *testing.T) {
	suite.Run(t, new(HTMLFormElementTestSuite))
}

func (s *HTMLFormElementTestSuite) SetupTest() {
	s.win = NewWindowHelper(s.T(), NewWindow(WindowOptions{
		BaseLocation: "http://example.com/forms/example-form.html?original-query=original-value",
	}))

	s.doc = NewHTMLDocument(s.win)
	s.form = s.doc.CreateElement("form").(HTMLFormElement)
}

func (s *HTMLFormElementTestSuite) createForm() HTMLFormElement {
	return s.doc.CreateElement("form").(HTMLFormElement)
}

func (s *HTMLFormElementTestSuite) TestMethodIDLAttribute() {
	form := s.createForm()

	s.Assert().Equal("get", form.Method(), "Default value for HTMLFormElement.method")

	form.SetMethod("new value")
	s.Expect(form).To(HaveAttribute("method", "new value"),
		"Content attribute when setting IDL to invalid value")
	s.Expect(form.Method()).To(Equal("get"),
		"IDL method attribute when set to an invalid value")

	for _, value := range []string{"post", "POST", "PoSt"} {
		form.SetMethod(value)
		s.Expect(form.Method()).To(Equal("post"),
			fmt.Sprintf("IDL method attribute, set to %s", value),
		)
	}
}

func (s *HTMLFormElementTestSuite) TestActionIDLAttribute() {
	form := s.createForm()

	s.Assert().Equal(s.win.LocationHREF(), form.Action(), "Default value for action IDL attribute")

	form.SetAction("/foo-bar")
	s.Expect(form).
		To(HaveAttribute("action", "/foo-bar"), "Action content attribute with absolute path")
	s.Expect(form.Action()).
		To(Equal("http://example.com/foo-bar"), "Action IDL attribute when set to an absolute path")

	form.SetAttribute("action", "submit-target")
	s.Expect(form).To(
		HaveAttribute("action", "submit-target"),
		"Action content attribute with relative path")

	s.Expect(form.Action()).To(
		Equal("http://example.com/forms/submit-target"),
		"Action IDL attribute when set to a relative path")
}

// TestFormDataWithUnnamedInput tests that form data construction skips input
// elements without a name attribute, conforming to the HTML spec:
// https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#constructing-the-form-data-set
func (s *HTMLFormElementTestSuite) TestFormDataWithUnnamedInput() {
	form := s.createForm()

	// Create input with name
	namedInput := s.doc.CreateElement("input").(HTMLInputElement)
	namedInput.SetName("username")
	namedInput.SetValue("john")
	form.AppendChild(namedInput)

	// Create input without name attribute
	unnamedInput := s.doc.CreateElement("input").(HTMLInputElement)
	unnamedInput.SetValue("should-be-skipped")
	form.AppendChild(unnamedInput)

	// Create another input with name
	anotherNamedInput := s.doc.CreateElement("input").(HTMLInputElement)
	anotherNamedInput.SetName("email")
	anotherNamedInput.SetValue("john@example.com")
	form.AppendChild(anotherNamedInput)

	// Construct form data - should not panic and should skip unnamed input
	formData := html.NewFormDataForm(form)

	s.Expect(formData.Has("username")).To(BeTrue(), "Form data includes named input 'username'")
	s.Expect(formData.Get("username")).To(Equal(html.FormDataValue("john")), "Form data has correct value for 'username'")
	s.Expect(formData.Has("email")).To(BeTrue(), "Form data includes named input 'email'")
	s.Expect(formData.Get("email")).To(Equal(html.FormDataValue("john@example.com")), "Form data has correct value for 'email'")
	s.Expect(len(formData.Entries)).To(Equal(2), "Form data has exactly 2 entries, skipping the unnamed input")
}
