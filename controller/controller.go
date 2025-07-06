package controller

import (
	"iter"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/key"
)

type KeyboardController struct {
	Window html.Window
}

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

func (c KeyboardController) SendKeys(keys iter.Seq[key.Key]) {
	for k := range keys {
		c.SendKey(k)
	}
}
