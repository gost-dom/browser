package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
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
