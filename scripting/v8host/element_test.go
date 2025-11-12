package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/v8host"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type ElementTestSuite struct {
	*scripttests.ScriptHostSuite
}

func TestElement(t *testing.T) {
	suite.Run(t,
		&ElementTestSuite{scripttests.NewScriptHostSuite(v8host.NewEngine())},
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

func TestElementSiblings(t *testing.T) {
	g := gomega.NewWithT(t)
	win := browsertest.InitBrowser(t, nil).NewWindow()
	win.LoadHTML(`
  	<body>
		<div id="e-1"></div>
		text node
		<div id="e-2"></div>
		text node
		<div id="e-3"></div>
		<script>const e = document.getElementById("e-2")</script>
	</body>`)
	g.Expect(win.Eval(`e.previousSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeText), "nextSibling node type")
	g.Expect(win.Eval(`e.previousElementSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeElement), "nextElementSibling node type")
	g.Expect(win.Eval(`e.previousElementSibling.id`)).To(Equal("e-1"), "nextElementSibling id")

	g.Expect(win.Eval(`e.nextSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeText), "nextSibling node type")
	g.Expect(win.Eval(`e.nextElementSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeElement), "nextElementSibling node type")
	g.Expect(win.Eval(`e.nextElementSibling.id`)).To(Equal("e-3"), "nextElementSibling id")
}
