package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/suite"
)

type NodeTestSuite struct {
	gosttest.GomegaSuite
}

func (s *NodeTestSuite) TestShallowClone() {
	doc := ParseHtmlString(`<body id='body'><div>First</div><div id="1">1</div></body>`)
	clone := doc.Body().CloneNode(false)
	s.Expect(clone).ToNot(BeNil())
	s.Expect(clone).To(HaveTag("BODY"))
	elm := clone.(html.HTMLElement)
	s.Expect(elm).To(HaveAttribute("id", "body"))
}

func (s *NodeTestSuite) TestDeepClone() {
	doc := ParseHtmlString(`<body id="body"><div>First</div><div id="1">1</div></body>`)
	clone := doc.Body().CloneNode(true)
	s.Expect(
		clone,
	).To(HaveOuterHTML(`<body id="body"><div>First</div><div id="1">1</div></body>`))
	s.Expect(
		doc.Body(),
	).To(HaveOuterHTML(`<body id="body"><div>First</div><div id="1">1</div></body>`),
		"Original node was not mutated",
	)
}

func (s *NodeTestSuite) TestInsertBeforeOfNewElement() {
	doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
	div := doc.GetElementById("1")
	s.Expect(div).ToNot(BeNil())
	newElm := doc.CreateElement("p")
	s.Expect(doc.Body().InsertBefore(newElm, div)).Error().To(Succeed())
	s.Expect(
		doc.Body(),
	).To(HaveOuterHTML(`<body><div>First</div><p></p><div id="1">1</div></body>`))
	s.Expect(newElm.ParentNode()).To(Equal(doc.Body()))
}

func (s *NodeTestSuite) TestInsertBeforeAppendsWithNilReference() {
	doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
	newElm := doc.CreateElement("p")
	s.Expect(doc.Body().InsertBefore(newElm, nil)).Error().ToNot(HaveOccurred())
	s.Expect(
		doc.Body(),
	).To(HaveOuterHTML(`<body><div>First</div><div id="1">1</div><p></p></body>`))
	s.Expect(newElm.ParentNode()).To(Equal(doc.Body()))
}

func (s *NodeTestSuite) TestInsertDocumentFragmentOrder() {
	doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
	fragment := doc.CreateDocumentFragment()
	d1 := doc.CreateElement("div")
	d2 := doc.CreateElement("div")
	d1.SetAttribute("id", "c-1")
	d2.SetAttribute("id", "c-2")
	fragment.Append(d1)
	fragment.Append(d2)
	ref := doc.GetElementById("1")

	result, err := doc.Body().InsertBefore(fragment, ref)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(
		doc.Body(),
	).To(HaveOuterHTML(`<body><div>First</div><div id="c-1"></div><div id="c-2"></div><div id="1">1</div></body>`))
	s.Expect(result.ChildNodes().All()).To(BeEmpty())
}

func TestNode(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(NodeTestSuite))
}

func (s *NodeTestSuite) TestContains() {
	doc := ParseHtmlString(`<body>
	<div id="parent-1">
		<div id="1">
			<div id="2">
				<div id="3">
		</div></div></div></div>
	<div id="sibling"></div>
</body>`)
	parent := doc.GetElementById("parent-1")
	s.Assert().True(parent.Contains(doc.GetElementById("1")), "Direct child node")
	s.Assert().True(parent.Contains(doc.GetElementById("3")), "Great grand child node")
	s.Assert().False(parent.Contains(doc.GetElementById("sibling")), "Sibling node")
}

func (s *NodeTestSuite) TestTextContents() {
	doc := ParseHtmlString(
		`<body><div style="display: none">Hidden text</div><div>Visible text</div></body>`,
	)
	s.Assert().Equal("Hidden textVisible text", doc.Body().TextContent())
}

func (s *NodeTestSuite) TestIsConnected() {
	doc := ParseHtmlString(`<body></body>`)
	div := doc.CreateElement("div")
	s.Expect(div.IsConnected()).To(BeFalse())

	_, err := doc.DocumentElement().AppendChild(div)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(div.IsConnected()).To(BeTrue())
}

func (s *NodeTestSuite) TestRemoveChild() {
	doc := ParseHtmlString(
		`<body><div id="parent"><div id="child">child</div></div></body>`,
	)
	parent := doc.GetElementById("parent")
	child := doc.GetElementById("child")
	removed, err := parent.RemoveChild(child)
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(parent).To(HaveOuterHTML(`<div id="parent"></div>`))
	s.Expect(removed).To(HaveOuterHTML(`<div id="child">child</div>`))
}

func (s *NodeTestSuite) TestInvalidRemoveChild() {
	doc := ParseHtmlString(
		`<body><div id="sibling-1"></div><div id="sibling-2"></div></body>`,
	)
	node1 := doc.GetElementById("sibling-1")
	node2 := doc.GetElementById("sibling-2")
	s.Expect(node1.RemoveChild(node2)).Error().To(BeADOMError())
}

func (s *NodeTestSuite) TestGetRootNode() {
	doc := ParseHtmlString(
		`<body><div id="parent"><div id="child">child</div></div></body>`,
	)
	root := doc.GetElementById("child").GetRootNode()
	s.Expect(root).To(Equal(doc))
}

func (s *NodeTestSuite) TestRootNodeOfDisconnectedRoot() {
	doc := CreateHTMLDocument()
	div := doc.CreateElement("div")
	s.Expect(div.GetRootNode()).To(Equal(div))
}

type NodeMoveToNewParentTest struct {
	gosttest.GomegaSuite
	doc dom.Document
}

func (s *NodeMoveToNewParentTest) SetupTest() {
	s.doc = ParseHtmlString(
		`<body>
  <div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>
  <div id="parent-2"><div id="ref"></div></div>
</body>`,
	)
}
func (s *NodeMoveToNewParentTest) TestMovedNodeAreRemoveFromPrevParent() {
	elm := s.doc.GetElementById("2")
	ref := s.doc.GetElementById("ref")
	oldParent := s.doc.GetElementById("parent-1")
	newParent := s.doc.GetElementById("parent-2")
	s.Expect(
		oldParent,
	).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>`))
	s.Expect(newParent.InsertBefore(elm, ref)).Error().ToNot(HaveOccurred())
	s.Expect(
		oldParent,
	).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="3">3</div></div>`))
	s.Expect(
		newParent,
	).To(HaveOuterHTML(`<div id="parent-2"><div id="2">2</div><div id="ref"></div></div>`))
}

func (s *NodeMoveToNewParentTest) TestOldParentIsUpdatedWhenUsingAppend() {
	elm := s.doc.GetElementById("2")
	oldParent := s.doc.GetElementById("parent-1")
	newParent := s.doc.GetElementById("parent-2")
	s.Expect(
		oldParent,
	).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>`))
	newParent.AppendChild(elm)
	s.Expect(
		oldParent,
	).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="3">3</div></div>`))
	s.Expect(
		newParent,
	).To(HaveOuterHTML(`<div id="parent-2"><div id="ref"></div><div id="2">2</div></div>`))
}

func TestNodeInsertBefore(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(NodeMoveToNewParentTest))
}

type NodeOwnerDocumentTestSuite struct {
	suite.Suite
	DocumentSuite
}

func (s *NodeOwnerDocumentTestSuite) TestNewElementHasSourceDocumentAsOwner() {
	element := s.Document().CreateElement("div")
	s.Assert().Same(element.OwnerDocument(), s.Document())
}

func (s *NodeOwnerDocumentTestSuite) TestMoveNewElementToDifferentDocument() {
	element := s.Document().CreateElement("div")

	otherDoc := s.CreateDocument()
	s.Assert().
		Same(otherDoc.DocumentElement().OwnerDocument(), otherDoc, "Owner doc of new doc's document element")

	otherDoc.DocumentElement().AppendChild(element)

	s.Assert().NotSame(element.OwnerDocument(), s.Document(), "Owner doc of moved element")
	s.Assert().Same(element.OwnerDocument(), otherDoc, "Owner doc of moved element")
}

func (s *NodeOwnerDocumentTestSuite) TestOwnerDocumentOfDocumentNode() {
	s.Assert().Nil(s.Document().OwnerDocument())
}

func TestNodeOwnerDocument(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(NodeOwnerDocumentTestSuite))
}
