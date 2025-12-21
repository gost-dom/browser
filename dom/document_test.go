package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
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

func TestGetElementsByTagNameReturnsALiveCollection(t *testing.T) {
	doc := ParseHtmlString(`<body><div id="1"></div><div id="2"></div></body>`)
	divs := doc.GetElementsByTagName("div")
	assert.Equal(t, 2, divs.Length())

	t.Run("Adding a non-matching element", func(t *testing.T) {
		doc.Body().Append(doc.CreateElement("p"))
		assert.Equal(t, 2, divs.Length())
	})

	t.Run("Adding a matching element is placed in the right position", func(t *testing.T) {
		newDiv := doc.CreateElement("div")
		newDiv.SetID("Test")
		doc.Body().InsertBefore(newDiv, doc.Body().FirstChild())
		assert.Equal(t, 3, divs.Length())
		assert.Equal(t, "Test", divs.Item(0).ID())
	})

	t.Run("Adding elements further down", func(t *testing.T) {
		newDiv := doc.CreateElement("div")
		newDiv.SetID("Test2")

		doc.GetElementById("2").Append(
			newDiv,
			doc.CreateElement("p"),
		)
		assert.Equal(t, 4, divs.Length())
		assert.Equal(t, "Test2", divs.Item(3).ID())
	})

	t.Run("Doesn't return the node itself", func(t *testing.T) {
		el := doc.GetElementById("2")
		divs := el.GetElementsByTagName("div")
		assert.Equal(t, 1, divs.Length())
		assert.Equal(t, "Test2", divs.Item(0).ID())
	})
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
	newDoc := CreateHTMLDocument()
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

func TestDocumentCreateElementNS(t *testing.T) {
	doc := dom.NewDocument(nil)
	el := doc.CreateElementNS("ns", "Name")
	assert.Equal(t, "ns", el.Namespace())
	assert.Equal(t, "name", el.LocalName())
}
