package controller

import (
	"iter"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/key"
)

// KeyboardController simulates the user typing a sequence of keys.
type KeyboardController struct {
	Window html.Window
}

// SendKey simulates the input of a single key.
func (c KeyboardController) SendKey(k key.Key) {
	active := c.Window.Document().ActiveElement()
	switch e := active.(type) {
	case html.HTMLInputElement:
		e.DispatchEvent(&event.Event{Type: "keydown"})
		e.SetValue(e.Value() + k.Letter)
		e.DispatchEvent(&event.Event{Type: "input"})
		e.DispatchEvent(&event.Event{Type: "keyup"})
	}
}

// SendKeys simulates the user typing a sequence of keys. The key package
// contains functionality to generate sequences of keys from an input string.
func (c KeyboardController) SendKeys(keys iter.Seq[key.Key]) {
	for k := range keys {
		c.SendKey(k)
	}
}
