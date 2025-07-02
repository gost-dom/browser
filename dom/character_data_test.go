package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestCommentNode(t *testing.T) {
	g := gomega.NewWithT(t)
	node := CreateHTMLDocument().CreateComment("dummy")
	g.Expect(node.NodeType()).To(Equal(dom.NodeType(8)), "Comment node type")

	node = CreateHTMLDocument().CreateComment("A sequence of 27 characters")
	g.Expect(node.Data()).To(Equal("A sequence of 27 characters"))
	g.Expect(node.Length()).To(Equal(27), "Length of 27 character string")

	node = CreateHTMLDocument().CreateComment("𐀀")
	g.Expect(
		node.Length(),
	).To(Equal(1), "Length of single character requiring multiple bytes of encoding")
}

func TestCharacterDataRemove(t *testing.T) {
	doc := htmltest.ParseHTMLDocument(t, "<body><p>Para1</p>Text node<p>Para2</p></body>")
	para1 := doc.QuerySelectorHTML("p")
	textNode := para1.NextSibling().(dom.Text)
	textNode.Remove()

	assert.Equal(t, "<p>Para1</p><p>Para2</p>", doc.Body().InnerHTML())
	assert.Equal(t, 2, doc.Body().ChildNodes().Length())
}
