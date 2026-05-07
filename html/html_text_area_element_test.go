package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTMLTextArea(t *testing.T) {
	doc := htmltest.NewHTMLDocumentHelper(t, nil)
	elm := doc.CreateHTMLElement("textarea").HTMLElement
	textArea, ok := elm.(html.HTMLTextAreaElement)
	require.True(t, ok, "element is an HTMLTextArea")

	textArea.SetValue("foo")
	assert.Equal(t, "foo", textArea.Value())
}
