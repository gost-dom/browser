package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"

	"github.com/onsi/gomega"
)

func TestDocumentCreateTextNode(t *testing.T) {
	doc := ParseHtmlString("")
	text := doc.CreateTextNode("data")
	assert.Equal(t, doc, text.OwnerDocument())
}

func TestDocumentCreateElement(t *testing.T) {
	doc := ParseHtmlString("")
	e := doc.CreateElement("div")
	assert.Equal(t, doc, e.OwnerDocument())
}

func TestDocumentFindElementById(t *testing.T) {
	doc := ParseHtmlString(`<body>
  <div id="uncle></div>
  <div id="parent">
    <div id="child">
      <div id="dummy"></div>
      <div id="grand-child"></div>
    </div>
  </div></body>`)
	elm := doc.GetElementById("grand-child")
	gomega.NewWithT(t).Expect(elm).To(HaveOuterHTML(`<div id="grand-child"></div>`))
}

func TestGetElementsByTagName(t *testing.T) {
	doc := ParseHtmlString(`<body><div id="1"></div><div id="2"></div></body>`)
	divs := doc.GetElementsByTagName("div")
	assert.Equal(t, 2, divs.Length())

	nosuchtags := doc.GetElementsByTagName("nosuchtag")
	assert.Equal(t, 0, nosuchtags.Length())
}

func TestDocumentImportNode(t *testing.T) {
	doc := ParseHtmlString(`
		<body>
			<div id="uncle"></div>
			<div id="parent">
				<div id="child">
					<div id="dummy"></div>
					<div id="grand-child"></div>
				</div>
			</div>
		</body>`,
	)
	parent := doc.GetElementById("parent")
	newDoc := html.NewHTMLDocument(html.NewWindow())
	if !assert.NotNil(t, parent) {
		return
	}
	clone := newDoc.ImportNode(parent, true)
	assert.Equal(t, newDoc, clone.OwnerDocument())
}

func TestDocumentOuterHTML(t *testing.T) {
	// Parsing an empty HTML doc generates both head and body
	doc := ParseHtmlString("")
	assert.Equal(t,
		"<html><head></head><body></body></html>",
		doc.DocumentElement().OuterHTML())
}

func TestDocumentDocumentElementIsHTMLElement(t *testing.T) {
	expect := gomega.NewWithT(t).Expect
	doc := ParseHtmlString("")
	expect(doc.DocumentElement()).To(BeHTMLElement())
}

func TestDocumentCloneNode(t *testing.T) {
	doc := ParseHtmlString("<div>foo</div>")
	node := doc.CloneNode(true)
	newDoc := node.(dom.Document)
	assert.Equal(t, doc.DocumentElement().OuterHTML(), newDoc.DocumentElement().OuterHTML())
}
