package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type NamedNodeMapSuite struct {
	ScriptHostSuite
}

func NewNamedNodeMapSuite(h html.ScriptEngine) *NamedNodeMapSuite {
	return &NamedNodeMapSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *NamedNodeMapSuite) TestInheritance() {
	s.Expect(
		s.Eval("Object.getPrototypeOf(NamedNodeMap.prototype) === Object.prototype"),
	).To(BeTrue(), "NamedNodeMap extends Object")
}

func (s *NamedNodeMapSuite) TestIterateAttributes() {
	s.MustLoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
	s.MustRunScript(`
		const elm = document.getElementById("foo");
		const attributes = elm.attributes;
		let idAttribute;
		for (let i = 0; i < attributes.length; i++) {
			const attr = attributes.item(i)
			if (attr.name === "id") { idAttribute = attr }
		}`)
	s.Expect(s.Eval("attributes.length")).To(BeEquivalentTo(3))
	s.Expect(s.Eval("idAttribute.value")).To(Equal("foo"))
	s.MustRunScript("idAttribute.value = 'bar'")
	s.Expect(s.Eval("idAttribute.value")).To(Equal("bar"))
}

func (s *NamedNodeMapSuite) TestIndex() {
	s.MustLoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
	s.Expect(s.Eval(`
		const elm = document.getElementById("foo");
		const attributes = elm.attributes;
		attributes[0] instanceof Attr &&
		attributes[1] instanceof Attr &&
		attributes[2] instanceof Attr
	`)).To(BeTrue())
}

func (s *NamedNodeMapSuite) TestIndexOutOfRange() {
	s.MustLoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
	s.Expect(s.Eval(`
		const elm = document.getElementById("foo");
		const attributes = elm.attributes;
		attributes[3]
	`)).To(BeNil())

	s.Expect(s.Eval("document.body.attributes[42] === undefined")).
		To(BeTrue(), "Out of range indexer returns undefined")
	s.Expect(s.Eval("document.body.attributes.item(42) === null")).
		To(BeTrue(), "item() with out of range index returns null")
}

func (s *NamedNodeMapSuite) TestNodeTypeOfAttributes() {
	s.MustLoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
	s.Expect(s.Eval(`
		const elm = document.getElementById("foo");
		const attribute = elm.attributes.item(0);
		attribute.nodeType
	`)).To(BeEquivalentTo(2))
}
