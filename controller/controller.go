package controller

import (
	"iter"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
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
		e.DispatchEvent(&event.Event{Type: "keydown"})
		e.SetValue(e.Value() + string(k))
		e.DispatchEvent(&event.Event{Type: "input"})
		e.DispatchEvent(&event.Event{Type: "keyup"})
	}
}

func (c KeyboardController) SendKeys(keys iter.Seq[Key]) {
	for k := range keys {
		c.SendKey(k)
	}
}

func KeysOfString(s string) iter.Seq[Key] {
	return func(yield func(Key) bool) {
		for _, r := range s {
			if !yield(KeyChar(r)) {
				return
			}
		}
	}
}
