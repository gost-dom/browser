package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/onsi/gomega"
)

type DocumentTestSuite struct {
	ScriptHostSuite
}

func NewDocumentSuite(h html.ScriptHost) *DocumentTestSuite {
	return &DocumentTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *DocumentTestSuite) TestPrototype() {
	s.Assert().Equal("HTMLDocument",
		s.mustEval("Object.getPrototypeOf(document).constructor.name"),
		"Global document instance")

	s.Assert().Equal(true, s.mustEval(
		`Object.getOwnPropertyNames(Document.prototype).includes("createElement")`),
		"createElement exists on Document.prototype")

	s.Assert().Equal("Node",
		s.mustEval(`Object.getPrototypeOf(Document.prototype).constructor.name`),
		"Document inherits from Node")
	s.Assert().Equal("Document",
		s.mustEval(`Object.getPrototypeOf(HTMLDocument.prototype).constructor.name`),
		"HTMLDocument inherits from Document")

	s.Assert().Equal(false, s.mustEval(
		`Object.getOwnPropertyNames(Document).includes("createElement")`),
		"createElement exist on Document (static method)")

	s.Assert().Contains(
		s.mustEval("Object.getOwnPropertyNames(document)"), "location",
		"location should exist on document instance")
	s.Assert().NotContains(
		s.mustEval("Object.getOwnPropertyNames(Document.prototype)"), "location",
		"location should not exist on Document prototype")
}

func (s *DocumentTestSuite) TestCreateElement() {
	s.Assert().Equal(true,
		s.mustEval(`document.createElement("base") instanceof HTMLElement`),
		"Element is an HTMLElement instance")
}

func (s *DocumentTestSuite) TestGetElementByID() {
	s.mustLoadHTML(
		`<body><div id='elm-1'>Elm: 1</div><div id='elm-2'>Elm: 2</div></body>`,
	)
	s.Expect(s.eval(`
		const e = document.getElementById("elm-2")
		e.outerHTML
	`)).To(Equal(`<div id="elm-2">Elm: 2</div>`))

	s.Expect(s.eval(`Object.getPrototypeOf(e).constructor.name`)).To(Equal("HTMLDivElement"))
}

func (s *DocumentTestSuite) TestNewDocument() {
	s.mustRun("const actual = new Document()")

	s.Assert().Equal(false,
		s.mustEval(`actual === document`),
		"New document must not equal global document")

	s.Assert().EqualValues(9, s.mustEval("actual.nodeType"), "new Document().nodeType")
	s.Assert().Equal("Document",
		s.mustEval("Object.getPrototypeOf(actual).constructor.name"),
		"Actual constructor")
}
