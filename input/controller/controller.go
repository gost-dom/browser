package controller

import (
	"iter"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/key"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	"github.com/gost-dom/browser/internal/uievents"
)

// KeyboardController simulates the user typing a sequence of keys.
type KeyboardController struct {
	Window html.Window
}

func (c KeyboardController) handleKey(active dom.Element, k key.Key) {
	switch e := active.(type) {
	case htmlinterfaces.HTMLInputtableElement:
		e.SetValue(e.Value() + k.Letter)
		uievents.Input(e)
	}
}

// SendKey simulates the input of a single key.
//
// Ignored if no Window is assigned.
func (c KeyboardController) SendKey(k key.Key) {
	if c.Window == nil {
		return
	}
	eventInit := k.EventInit()
	active := c.Window.Document().ActiveElement()
	if k.Down {
		if uievents.KeydownInit(active, eventInit) {
			c.handleKey(active, k)
		}
	}
	if k.Up {
		uievents.KeyupInit(active, eventInit)
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
