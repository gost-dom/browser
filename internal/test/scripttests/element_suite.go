package scripttests

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

type ElementSuite struct {
	ScriptHostSuite
}

func NewElementSuite(h html.ScriptEngine) *ElementSuite {
	return &ElementSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ElementSuite) TestBasicElementProperties() {
	s.Expect(
		s.Eval("Object.getPrototypeOf(Element.prototype) === Node.prototype"),
	).To(BeTrue(), "Element by be a direct descendant of node")

	s.Expect(s.Eval("document.body.nodeType")).To(BeEquivalentTo(1), "Element must have nodetype 1")
	s.Expect(s.Eval("document.body.nodeType === Node.ELEMENT_NODE")).
		To(BeTrue(), "Element.nodeType must equal Node.ELEMENT_NODE")
}

func (s *ElementSuite) TestAppendMultipleElements() {
	s.Expect(s.Eval(`
		const d = document.createElement("div")
		d.append(
			document.createElement("p"),
			document.createElement("p"),
		);
		d.outerHTML`)).To(Equal("<div><p></p><p></p></div>"))

	s.Expect(s.Eval(`
		const d2 = document.createElement("div")
		d2.append(
			document.createElement("p"),
			"Foo",
			"bar",
			document.createElement("p"),
		);
		d2.outerHTML`)).To(Equal("<div><p></p>Foobar<p></p></div>"))
}

func (s *ElementSuite) TestSetOuterHTML() {
	s.Expect(
		s.Eval(`
			const d = document.createElement("div")
			document.body.replaceChildren(d)
			d.outerHTML = "<div id='TEST'>Data</div>"

			document.body.innerHTML`)).
		To(Equal(`<div id="TEST">Data</div>`))
	s.Expect(
		s.Window.Document().GetElementById("TEST"),
	).To(HaveTextContent("Data"))

	s.Expect(
		s.Eval(`
			document.body.replaceChildren(d)
			d.outerHTML = "<div id='TEST'>Data</div>"

			document.body.innerHTML`)).
		To(Equal(`<div id="TEST">Data</div>`))
}

func (s *ElementSuite) TestAttributes() {
	s.MustLoadHTML(`<body><div foo="foo-value" bar="bar-value"></div><body>`)
	s.Assert().Equal("foo-value", s.MustEval(`document.body.firstChild.getAttribute("foo")`))
	s.MustEval(`
		const div = document.querySelector("div")
		const attrs = Array.from(div.attributes)
		const names = attrs.map(x => x.name).join(",")
		const values = attrs.map(x => x.value).join(",")
	`)
	s.Assert().Equal("foo,bar", s.MustEval("names"))
	s.Assert().Equal("foo-value,bar-value", s.MustEval("values"))

	s.Expect(s.Eval(`document.body.getAttribute('non-existing') === null`)).
		To(BeTrue(), "Reading a non-existing attribute should return null")

	s.Expect(s.Eval(`div.hasAttribute("foo")`)).
		To(BeTrue(), "hasAttribute of existing attribute")
	s.Expect(s.Eval(`div.hasAttribute("non-existing-attribute")`)).
		To(BeFalse(), "hasAttribute of non-existing attribute")
}

func (s *ElementSuite) TestIDLInterfaceNamesForElements() {
	s.MustRunScript(`
		gost.assertInstanceOf(document.createElement('a'), HTMLAnchorElement);
		gost.assertInstanceOf(document.createElement('p'), HTMLParagraphElement);
		gost.assertInstanceOf(document.createElement('div'), HTMLDivElement);
	`)
}

func (s *ElementSuite) TestChildren() {
	s.MustLoadHTML(`
		<body>
			<div id="target">
				Initial text
				<div id="child-1">Child 1</div>
				Some text
				<div id="child-2">Child 2</div>
				Final text
			</div>
		</body>`,
	)
	s.MustRunScript(`
		const target = document.getElementById("target")
		const child1ByItem = target.children.item(0)
		const child1ByIndex = target.children[0]
		const child2ByItem = target.children.item(1)
		const child2ByIndex = target.children[1]
		const length = target.children.length
		const arr = Array.from(target.children)
		const contents = arr.map(x => x.getAttribute("id")).join(",")
	`)
	assert := s.Assert()
	assert.Equal("Child 1", s.MustEval("child1ByItem.textContent"))
	assert.Equal("Child 1", s.MustEval("child1ByIndex.textContent"))
	assert.Equal("Child 2", s.MustEval("child2ByItem.textContent"))
	assert.Equal("Child 2", s.MustEval("child2ByIndex.textContent"))

	s.Expect(s.MustEval("length")).To(BeEquivalentTo(2))

	assert.Equal("child-1,child-2", s.MustEval("contents"))
}

func (s *ElementSuite) TestInsertAdjacentHTML() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.Expect(s.Eval(
		`document.getElementById("1").insertAdjacentHTML("beforebegin", "<p>foo</p>")`,
	)).Error().ToNot(HaveOccurred())
	s.Expect(
		s.Window.Document().Body().OuterHTML(),
	).To(Equal(`<body><p>foo</p><div id="1" class="foo"></div></body>`))
}

func (s *ElementSuite) TestInsertAdjacentHTMLBadPosition() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.MustEval(`
		let err
		try {
			document.getElementById("1").insertAdjacentHTML("invalid", "<p>foo</p>")
		} catch (e) {
			err = e
		}
	`)

	s.Expect(s.Eval("err instanceof DOMException")).To(BeTrue())
	s.Expect(s.Eval("err.code")).To(BeEquivalentTo(12))
}

func (s *ElementSuite) TestQuerySelector() {
	s.MustLoadHTML(`<div id="1" class="foo"></div>`)
	s.MustRunScript(`const e = document.getElementById("1")`)
	s.Expect(s.Eval(`typeof e.querySelector`)).To(Equal("function"))
	s.Expect(s.Eval(`typeof e.querySelectorAll`)).To(Equal("function"))
}

func (s *ElementSuite) TestElementSiblings() {
	win := s.Window
	s.MustLoadHTML(`
  	<body>
		<div id="e-1"></div>
		text node
		<div id="e-2"></div>
		text node
		<div id="e-3"></div>
		<script>const e = document.getElementById("e-2")</script>
	</body>`)
	s.Expect(win.Eval(`e.previousSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeText), "nextSibling node type")
	s.Expect(win.Eval(`e.previousElementSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeElement), "nextElementSibling node type")
	s.Expect(win.Eval(`e.previousElementSibling.id`)).To(Equal("e-1"), "nextElementSibling id")

	s.Expect(win.Eval(`e.nextSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeText), "nextSibling node type")
	s.Expect(win.Eval(`e.nextElementSibling.nodeType`)).
		To(BeEquivalentTo(dom.NodeTypeElement), "nextElementSibling node type")
	s.Expect(win.Eval(`e.nextElementSibling.id`)).To(Equal("e-3"), "nextElementSibling id")
}
