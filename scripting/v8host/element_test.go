package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/scripting/v8host"
	"github.com/stretchr/testify/suite"
)

type ElementTestSuite struct {
	*scripttests.ScriptHostSuite
}

func TestElement(t *testing.T) {
	logger := gosttest.NewTestLogger(t)
	host := v8host.New(v8host.WithLogger(logger))
	suite.Run(t,
		&ElementTestSuite{scripttests.NewScriptHostSuite(host)},
	)
}

func (s *ElementTestSuite) TestBasicElementProperties() {
	s.Expect(
		s.Eval("Object.getPrototypeOf(Element.prototype) === Node.prototype"),
	).To(BeTrue(), "Element by be a direct descendant of node")

	s.Expect(s.Eval("document.body.nodeType")).To(BeEquivalentTo(1), "Element must have nodetype 1")
	s.Expect(s.Eval("document.body.nodeType === Node.ELEMENT_NODE")).
		To(BeTrue(), "Element.nodeType must equal Node.ELEMENT_NODE")
}

func (s *ElementTestSuite) TestTextContent() {
	s.MustRunScript("document.body.textContent = 'text content'")
	s.Assert().Equal(
		"text content",
		s.Window.Document().Body().TextContent())
}

func (s *ElementTestSuite) TestAttributes() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.Expect(s.Eval(
		`document.getElementById("1").getAttribute("class")`,
	)).To(Equal("foo"))
	s.Expect(s.Eval(`
		const e = document.getElementById("1")
		e.setAttribute("data-foo", "bar");
		e.getAttribute("data-foo")`,
	)).To(Equal("bar"))

	s.Expect(s.Eval(`e.getAttribute('non-existing') === null`)).
		To(BeTrue(), "Reading a non-existing attribute should return null")

	s.Expect(s.Eval(`e.hasAttribute("data-foo")`)).
		To(BeTrue(), "hasAttribute of existing attribute")

	s.Expect(s.Eval(`e.hasAttribute("non-existing-attribute")`)).
		To(BeFalse(), "hasAttribute of non-existing attribute")
}

func (s *ElementTestSuite) TestInsertAdjacentHTML() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.Expect(s.Eval(
		`document.getElementById("1").insertAdjacentHTML("beforebegin", "<p>foo</p>")`,
	)).Error().ToNot(HaveOccurred())
	s.Expect(
		s.Window.Document().Body().OuterHTML(),
	).To(Equal(`<body><p>foo</p><div id="1" class="foo"></div></body>`))
}

func (s *ElementTestSuite) TestQuerySelector() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.MustRunScript(`const e = document.getElementById("1")`)
	s.Expect(s.Eval(`typeof e.querySelector`)).To(Equal("function"))
	s.Expect(s.Eval(`typeof e.querySelectorAll`)).To(Equal("function"))
}
