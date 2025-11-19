package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/onsi/gomega"
)

type ElementTestSuite struct {
	gosttest.GomegaSuite
	doc dom.Document
}

func (s *ElementTestSuite) SetupTest() {
	s.doc = CreateHTMLDocument()
}

func (s *ElementTestSuite) TestGetSetAttribute() {
	elm := s.doc.CreateElement("div")
	s.Expect(elm.Attributes().Length()).To(Equal(0))
	elm.SetAttribute("id", "1")
	s.Expect(elm.Attributes().Length()).To(Equal(1))
}

func (s *ElementTestSuite) TestOverwriteExistingAttribute() {
	elm := s.doc.CreateElement("div")
	elm.SetAttribute("id", "1")
	elm.SetAttribute("id", "2")
	s.Expect(elm).To(HaveAttribute("id", "2"))
	s.Expect(elm.Attributes().Length()).To(Equal(1))
}

func (s *ElementTestSuite) TestGetMissingAttribute() {
	elm := s.doc.CreateElement("div")
	_, ok := elm.GetAttribute("non-existing")
	s.Expect(ok).To(BeFalse())
}

func (s *ElementTestSuite) TestGetMissingAttributeNode() {
	elm := s.doc.CreateElement("div")
	attr := elm.GetAttributeNode("class")
	s.Expect(attr).To(BeNil())
}

func (s *ElementTestSuite) TestAttributeNodesAreMutable() {
	elm := s.doc.CreateElement("div")
	elm.SetAttribute("class", "foo")
	attr := elm.GetAttributeNode("class")
	s.Expect(attr).ToNot(BeNil())
	s.Expect(attr.Value()).To(Equal("foo"), "Attribute value before mutation")
	s.Expect(attr.ParentNode()).To(Equal(elm), "Parent on attribute node")
	attr.SetValue("bar")
	actual := elm.GetAttributeNode("class")
	s.Expect(actual.Value()).To(Equal("bar"), "Attribute value after mutation")
}

func (s *ElementTestSuite) TestSetAttributeNodeAddsNew() {
	elm := s.doc.CreateElement("div")
	attr := s.doc.CreateAttribute("class")
	attr.SetValue("foo")
	result, err := elm.SetAttributeNode(attr)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(result).To(BeNil())
	s.Expect(elm.Attributes().Length()).To(Equal(1))
	actual, _ := elm.GetAttribute("class")
	s.Expect(actual).To(Equal("foo"))
}

func (s *ElementTestSuite) TestAttributeNodeUpdatesExisting() {
	elm := s.doc.CreateElement("div")
	elm.SetAttribute("class", "bar")
	attr := s.doc.CreateAttribute("class")
	attr.SetValue("foo")
	result, err := elm.SetAttributeNode(attr)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(result.Name()).To(Equal("class"))
	s.Expect(result.Value()).To(Equal("bar"))
	s.Expect(elm.Attributes().Length()).To(Equal(1))
	actual, _ := elm.GetAttribute("class")
	s.Expect(actual).To(Equal("foo"))
}

func (s *ElementTestSuite) TestReturnDOMErrorWhenMovingToWrongElement() {
	elm := s.doc.CreateElement("div")
	elm2 := s.doc.CreateElement("div")
	elm2.SetAttribute("class", "bar")
	attributeFromAnotherElement := elm2.GetAttributeNode("class")
	s.Expect(
		elm.SetAttributeNode(attributeFromAnotherElement),
	).Error().To(BeADOMError())
	s.Expect(elm.Attributes().Length()).To(Equal(0), "Target elm received attribute?")
	s.Expect(
		elm2.Attributes().Length(),
	).To(Equal(1), "Target attribute still on source elm")
}

func (s *ElementTestSuite) TestRemoveExistingAttribute() {
	elm := s.doc.CreateElement("div")
	elm.SetAttribute("class", "bar")
	nodeToRemove := elm.GetAttributeNode("class")
	removedNode, err := elm.RemoveAttributeNode(nodeToRemove)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(elm.Attributes().Length()).To(Equal(0))
	s.Expect(removedNode).To(Equal(nodeToRemove))
	s.Expect(removedNode.ParentNode()).To(BeNil(), "Attribute parent")
}

func (s *ElementTestSuite) TestRemoveNonExistingAttributeNode() {
	elm := s.doc.CreateElement("div")
	elm2 := s.doc.CreateElement("div")
	elm2.SetAttribute("class", "bar")
	attributeFromAnotherElement := elm2.GetAttributeNode("class")
	s.Expect(
		elm.RemoveAttributeNode(attributeFromAnotherElement),
	).Error().To(BeADOMError())
}

func (s *ElementTestSuite) TestMatchesTagName() {
	d := s.doc.CreateElement("div")
	p := s.doc.CreateElement("p")
	d.Append(p)
	s.Assert().True(d.Matches("div"))
	s.Assert().False(d.Matches("p"))
	s.Assert().False(p.Matches("div"))
}

func (s *ElementTestSuite) TestMatchesAttribute() {
	d := s.doc.CreateElement("div")
	d.SetAttribute("known-attribute", "")
	s.Assert().True(d.Matches("[known-attribute]"))
	s.Assert().False(d.Matches("[unknown-attribute]"))
}

func (s *ElementTestSuite) TestMatchesAttributeNameValue() {
	d := s.doc.CreateElement("div")
	d.SetAttribute("a", "good")
	s.Expect(d.Matches(`div[a="good"]`)).To(BeTrue())
	s.Expect(d.Matches(`div[a="bad"]`)).To(BeFalse())
}

func (s *ElementTestSuite) TestTextContent() {
	d := s.doc.CreateElement("div")
	p := s.doc.CreateElement("p")
	p.AppendChild(s.doc.CreateTextNode("Original paragraph"))
	d.AppendChild(p)
	s.Expect(d).To(HaveTextContent("Original paragraph"))
	d.SetTextContent("Replace the p")
	s.Expect(d).To(HaveTextContent("Replace the p"))
	s.Assert().Equal(1, d.ChildNodes().Length())
	s.Assert().Equal(NodeTypeText, d.ChildNodes().Item(0).NodeType())
}

func TestElement(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ElementTestSuite))
}

func TestElementInsertAdjacentHTML(t *testing.T) {
	specs := map[string]string{
		"beforebegin": `<body>
  <div id="1">El 1</div>
  <div id="new">1st new child</div>Text node<div>2nd new child</div><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`,
		"afterbegin": `<body>
  <div id="1">El 1</div>
  <div id="2"><div id="new">1st new child</div>Text node<div>2nd new child</div>El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`,
		"beforeend": `<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  <div id="new">1st new child</div>Text node<div>2nd new child</div></div>
  <div id="3">El 1</div>
</body>`,
		"afterend": `<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div><div id="new">1st new child</div>Text node<div>2nd new child</div>
  <div id="3">El 1</div>
</body>`,
	}

	for position, want := range specs {
		t.Run(position, func(t *testing.T) {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el := doc.GetElementById("2")
			assert.NoError(t, el.InsertAdjacentHTML(
				position,
				`<div id="new">1st new child</div>Text node<div>2nd new child</div>`,
			))
			got := doc.Body().OuterHTML()
			assert.Equal(t, want, got)
		})
	}
}

func TestInsertAdjacentElement(t *testing.T) {
	specs := map[string]string{
		"beforebegin": `<body><div>New child</div><div id="target">Text node</div></body>`,
		"afterbegin":  `<body><div id="target"><div>New child</div>Text node</div></body>`,
		"beforeend":   `<body><div id="target">Text node<div>New child</div></div></body>`,
		"afterend":    `<body><div id="target">Text node</div><div>New child</div></body>`,
	}

	for position, want := range specs {
		t.Run(position, func(t *testing.T) {
			doc := ParseHtmlString(`<body><div id="target">Text node</div></body>`)
			e := doc.GetElementById("target")
			d := doc.CreateElement("div")
			d.SetTextContent("New child")
			res, err := e.InsertAdjacentElement(position, d)
			assert.Equal(t, d, res)
			assert.NoError(t, err)

			assert.Equal(t, want, doc.Body().OuterHTML())
		})
	}
}

func TestElementOuterHTML(t *testing.T) {
	gomega := gomega.NewWithT(t)
	// Whitespace is part of the parsed HTML as #text nodes
	doc := ParseHtmlString(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`)
	gomega.Expect(doc.Body().OuterHTML()).To(Equal(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`))
	gomega.Expect(doc.Body().InnerHTML()).To(Equal(`<div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>`))
}

func TestSetElementInnerHTML(t *testing.T) {
	gomega := gomega.NewWithT(t)
	doc := ParseHtmlString(`<body>body text</body>`)
	t.Logf("Owner: %T", doc.Body().OwnerDocument())
	t.Log("SetInnterHTML")
	err := doc.Body().SetInnerHTML(`<div id="1">Foo</div>Bar<div id="2">Baz</div>`)
	gomega.Expect(err).ToNot(HaveOccurred())
	gomega.
		Expect(doc.Body().InnerHTML()).
		To(Equal(`<div id="1">Foo</div>Bar<div id="2">Baz</div>`))
	gomega.
		Expect(doc.Body().TagName()).
		To(Equal("BODY"))
	gomega.
		Expect(doc.Body().FirstElementChild().TagName()).
		To(Equal("DIV"))
}

// ParentElementTestSuite describes functionality in the ParentNode IDL
// interface mixin, which is used by both Element, Document, and
// DocumentFragment.
//
// Ideally, these tests should be executed on all 3 types
type ParentElementTestSuite struct {
	gosttest.GomegaSuite
}

func TestParentElement(t *testing.T) {
	suite.Run(t, new(ParentElementTestSuite))
}

func (s *ParentElementTestSuite) TestAppend() {
	doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
	b := doc.Body()
	divE := doc.CreateElement("div")
	divE.SetTextContent("e")
	b.Append(
		doc.CreateTextNode("d"),
		divE,
		doc.CreateTextNode("f"),
	)
	s.Expect(b).To(HaveOuterHTML(`<body>a<div>b</div>cd<div>e</div>f</body>`))
}

func (s *ParentElementTestSuite) TestPrepend() {
	doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
	b := doc.Body()
	divE := doc.CreateElement("div")
	divE.SetTextContent("e")
	b.Prepend(
		doc.CreateTextNode("d"),
		divE,
		doc.CreateTextNode("f"),
	)
	s.Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>fa<div>b</div>c</body>`))
}

func (s *ParentElementTestSuite) TestReplaceChildren() {
	doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
	b := doc.Body()
	divE := doc.CreateElement("div")
	divE.SetTextContent("e")
	b.ReplaceChildren(
		doc.CreateTextNode("d"),
		divE,
		doc.CreateTextNode("f"),
	)
	s.Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>f</body>`))

	b.ReplaceChildren()
	s.Expect(b).To(HaveOuterHTML(`<body></body>`))
}

func (s *ParentElementTestSuite) TestIterateChildren() {
	doc := ParseHtmlString(
		`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
	)
	b := doc.Body()
	c := b.Children()

	s.Expect(c.Length()).To(Equal(3))
	s.Expect(c.Item(0)).To(HaveOuterHTML(`<div id="el-1">b</div>`))
	s.Expect(c.Item(1)).To(HaveAttribute("id", "el-2"))
	s.Expect(c.Item(2)).To(HaveAttribute("name", "el-3"))

	s.Expect(c.Item(-1)).To(BeNil())
	s.Expect(c.Item(3)).To(BeNil())

	s.Expect(c.NamedItem("el-1")).To(HaveAttribute("id", "el-1"))
	s.Expect(c.NamedItem("el-2")).To(HaveAttribute("id", "el-2"))
	s.Expect(c.NamedItem("el-3")).To(HaveAttribute("name", "el-3"))
}

func (s *ParentElementTestSuite) TestFirstLastOnExistingElement() {
	doc := ParseHtmlString(
		`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
	)
	b := doc.Body()
	s.Expect(b.FirstElementChild()).To(HaveAttribute("id", "el-1"))
	s.Expect(b.LastElementChild()).To(HaveOuterHTML(`<div name="el-3">f</div>`))
}

func (s *ParentElementTestSuite) TestFirstLastWithNonElement() {
	doc := ParseHtmlString(
		`<body>body text</body>`,
	)
	b := doc.Body()
	s.Expect(b.FirstElementChild()).To(BeNil())
	s.Expect(b.LastElementChild()).To(BeNil())
}

func (s *ParentElementTestSuite) TestEmptyString() {
	s.T().Skip("How to handle empty string in NamedItem")
}
