package dom_test

import (
	"testing"

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

func TestOuterHTML(t *testing.T) {
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
