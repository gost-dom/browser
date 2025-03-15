package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"

	"github.com/stretchr/testify/assert"
)

func TestDocumentFragmentClone(t *testing.T) {

	doc := ParseHtmlString("")
	f := doc.CreateDocumentFragment()
	dNode, _ := f.AppendChild(doc.CreateElement("div"))
	div := dNode.(dom.Element)
	div.SetAttribute("class", "foo")
	div.SetTextContent("Text")
	clone := f.CloneNode(true).(dom.DocumentFragment)

	assert.Equal(t, `<div class="foo">Text</div>`,
		clone.FirstChild().(dom.Element).OuterHTML(),
	)
}
