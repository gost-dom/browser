package html_test

import (
	"testing"

	. "github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

func TestEmptyHTMLDocument(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	doc := NewHTMLDocument(nil)
	docElm := doc.DocumentElement()
	body := doc.Body()
	head := doc.Head()
	assert.Equal("HTML", docElm.TagName(), "Document has an <html> root")
	assert.Equal("HEAD", head.TagName(), "Document.Head() is a <head>")
	assert.Equal("BODY", body.TagName(), "Document.Body() is a <body>")

	assert.Equal(docElm, head.Parent(), "<head> is child of <html>")
	assert.Equal(docElm, body.Parent(), "<body> is child of <html>")
}

func TestHTMLDocumentCreateElement(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	doc := NewHTMLDocument(nil)
	{
		// Separate scopes to isolates cases
		_, e1IsAnchor := doc.CreateElement("a").(HTMLAnchorElement)
		assert.True(e1IsAnchor, "CreateElement('a') returns an HTMLAnchorElement")
	}
	{
		_, elIsAnchor := doc.CreateElement("A").(HTMLAnchorElement)
		assert.True(elIsAnchor, "CreateElement('a') returns an HTMLAnchorElement")
	}
}
