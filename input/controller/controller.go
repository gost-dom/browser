package controller

import (
	"iter"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/key"
	"github.com/gost-dom/browser/internal/uievents"
)

// KeyboardController simulates the user typing a sequence of keys.
type KeyboardController struct {
	Window html.Window
}

// SendKey simulates the input of a single key.
//
// Ignored if no Window is assigned.
func (c KeyboardController) SendKey(k key.Key) {
	if c.Window == nil {
		return
	}
	active := c.Window.Document().ActiveElement()
	switch e := active.(type) {
	case html.HTMLInputElement:
		if !uievents.Keydown(e) {
			return
		}
		e.SetValue(e.Value() + k.Letter)
		uievents.Input(e)
		uievents.Keyup(e)
	}
}

// SendKeys simulates the user typing a sequence of keys. The key package
// contains functionality to generate sequences of keys from an input string.
//
// Ignored if no Window is assigned.
func (c KeyboardController) SendKeys(keys iter.Seq[key.Key]) {
	for k := range keys {
		c.SendKey(k)
	}
}
