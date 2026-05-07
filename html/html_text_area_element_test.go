package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/controller"
	"github.com/gost-dom/browser/input/key"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTMLTextArea(t *testing.T) {
	win := htmltest.NewWindowHelper(t, nil)
	doc := win.HTMLDocument()
	elm := doc.CreateHTMLElement("textarea").HTMLElement
	textArea, ok := elm.(html.HTMLTextAreaElement)
	require.True(t, ok, "element is an HTMLTextArea")

	textArea.SetValue("foo")
	assert.Equal(t, "foo", textArea.Value())
	textArea.SetValue("")

	require.NoError(t, doc.Body().Append(elm))

	ctrl := controller.KeyboardController{Window: win}

	textArea.Focus()
	ctrl.SendKeys(key.StringToKeys("Hello, World!"))

	assert.Equal(t, "Hello, World!", textArea.Value())
}
