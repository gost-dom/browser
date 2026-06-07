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
	if k.Down {
		active := c.Window.Document().ActiveElement()
		if uievents.KeydownInit(active, eventInit) {
			c.handleKey(active, k)
		}
	}
	if k.Up {
		active := c.Window.Document().ActiveElement()
		uievents.KeyupInit(active, eventInit)
	}
}

func (c KeyboardController) next(getNext func() (key.Key, bool)) {
	if c.Window == nil {
		return
	}
	c.Window.SetTimeout(func() error {
		if k, ok := getNext(); ok {
			eventInit := k.EventInit()
			if k.Down {
				active := c.Window.Document().ActiveElement()
				if uievents.KeydownInit(active, eventInit) {
					c.handleKey(active, k)
				}
			}
			if k.Up {
				c.Window.SetTimeout(func() error {
					active := c.Window.Document().ActiveElement()
					uievents.KeyupInit(active, eventInit)
					return nil
				}, 0)
			}
			c.Window.SetTimeout(func() error {
				c.next(getNext)
				return nil
			}, 0)
		}
		return nil
	}, 0)
}

// SendKeys simulates the user typing a sequence of keys. The key package
// contains functionality to generate sequences of keys from an input string.
//
// Ignored if no Window is assigned.
func (c KeyboardController) SendKeys(keys iter.Seq[key.Key]) {
	next, _ := iter.Pull(keys)
	c.next(next)
	c.Window.Clock().Advance(0)
}

func (c KeyboardController) SendText(text string) {
	c.SendKeys(key.StringToKeys(text))
}
