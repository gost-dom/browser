package controller_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

type Key string

func KeyChar(r rune) Key {
	return Key(string(r))
}

type KeyboardController struct {
	Window html.Window
}

func (c KeyboardController) SendKey(k Key) {
	active := c.Window.Document().ActiveElement()
	switch e := active.(type) {
	case html.HTMLInputElement:
		e.SetAttribute("value", e.Value()+string(k))
	}
}

func TestKeyboardController(t *testing.T) {
	g := gomega.NewWithT(t)
	html := `
		<body>
			<input id="input" type="text" />
		</body>
	`
	win := htmltest.NewWindowHTML(t, html)
	ctrl := KeyboardController{win}

	input := win.HTMLDocument().GetHTMLElementById("input")

	ctrl.SendKey(KeyChar('a'))
	g.Expect(input).To(HaveIDLValue(""), "Keypress when input does not have focus")

	input.Focus()

	ctrl.SendKey(KeyChar('a'))
	g.Expect(input).To(HaveIDLValue("a"), "Keypress when input does not have focus")
	g.Expect(input).To(HaveAttribute("value", "a"), "Keypress when input does not have focus")

	ctrl.SendKey(KeyChar('b'))
	g.Expect(input).To(HaveIDLValue("ab"), "Keypress when input does not have focus")
	g.Expect(input).To(HaveAttribute("value", "ab"), "Keypress when input does not have focus")
}
