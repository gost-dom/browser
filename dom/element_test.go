package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/suite"

	. "github.com/onsi/ginkgo/v2"
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
	s.Expect(attr.Parent()).To(Equal(elm), "Parent on attribute node")
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
	s.Expect(removedNode.Parent()).To(BeNil(), "Attribute parent")
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
	p.AppendChild(s.doc.CreateText("Original paragraph"))
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

var _ = Describe("Element", func() {
	Describe("InsertAdjacentHTML", func() {
		It("Should insert correctly 'beforeBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el := doc.GetElementById("2")
			gomega.Expect(el.InsertAdjacentHTML(
				"beforebegin",
				"<div>1st new child</div><div>2nd new child</div>",
			)).To(Succeed())
			gomega.Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div>1st new child</div><div>2nd new child</div><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			gomega.Expect(err).ToNot(HaveOccurred())
			gomega.Expect(
				el.InsertAdjacentHTML(
					"afterbegin",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			gomega.Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2"><div>1st new child</div><div>2nd new child</div>El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'beforeEnd'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			gomega.Expect(err).ToNot(HaveOccurred())
			gomega.Expect(
				el.InsertAdjacentHTML(
					"beforeend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			gomega.Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  <div>1st new child</div><div>2nd new child</div></div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterend'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			gomega.Expect(err).ToNot(HaveOccurred())
			gomega.Expect(
				el.InsertAdjacentHTML(
					"afterend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			gomega.Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div><div>1st new child</div><div>2nd new child</div>
  <div id="3">El 1</div>
</body>`))
		})
	})

	Describe("HTML Rendering", func() {
		It("Should support OuterHTML", func() {
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
		})
	})

	Describe("Append/Prepend/ReplaceChildren/Children", func() {
		// These should ideaaly be tested on both Element, Document, and
		// DocumentFragment. The functions are defined in the ParentNode
		// mixin interface that all 3 share.

		It("Should add elements in 'Append'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.Append(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			gomega.Expect(b).To(HaveOuterHTML(`<body>a<div>b</div>cd<div>e</div>f</body>`))
		})

		It("Should add elements first in 'Prepend'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.Prepend(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			gomega.Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>fa<div>b</div>c</body>`))
		})

		It("Should replace elements first in 'ReplaceChildren'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.ReplaceChildren(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			gomega.Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>f</body>`))
		})

		It("Should iterate elements in 'Children'", func() {
			doc := ParseHtmlString(
				`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
			)
			b := doc.Body()
			c := b.Children()

			gomega.Expect(c.Length()).To(Equal(3))
			gomega.Expect(c.Item(0)).To(HaveOuterHTML(`<div id="el-1">b</div>`))
			gomega.Expect(c.Item(1)).To(HaveAttribute("id", "el-2"))
			gomega.Expect(c.Item(2)).To(HaveAttribute("name", "el-3"))

			gomega.Expect(c.Item(-1)).To(BeNil())
			gomega.Expect(c.Item(3)).To(BeNil())

			gomega.Expect(c.NamedItem("el-1")).To(HaveAttribute("id", "el-1"))
			gomega.Expect(c.NamedItem("el-2")).To(HaveAttribute("id", "el-2"))
			gomega.Expect(c.NamedItem("el-3")).To(HaveAttribute("name", "el-3"))
		})

		Describe("First/Last element", func() {
			It("Should return elements when they exist", func() {
				doc := ParseHtmlString(
					`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
				)
				b := doc.Body()
				gomega.Expect(b.FirstElementChild()).To(HaveAttribute("id", "el-1"))
				gomega.Expect(b.LastElementChild()).To(HaveOuterHTML(`<div name="el-3">f</div>`))
			})

			It("Should nil when there are only non-element children", func() {
				doc := ParseHtmlString(
					`<body>body text</body>`,
				)
				b := doc.Body()
				gomega.Expect(b.FirstElementChild()).To(BeNil())
				gomega.Expect(b.LastElementChild()).To(BeNil())
			})
		})

		It("Should handle empty string correctly in 'NamedItem'", func() {
			Skip("Need to research")
		})
	})
})
