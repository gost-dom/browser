package controller

import (
	"iter"
	"log/slog"

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
// If c has no Window, this function will produce no effects.
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

// dispatchKeyUp dispatches the keyup event if applicable. After that; it
// proceeds to handling the remaining keys in a sequence of keys.
func (c KeyboardController) dispatchKeyUp(getNext func() (key.Key, bool), k key.Key, stop func()) {
	eventInit := k.EventInit()
	if k.Up {
		active := c.Window.Document().ActiveElement()
		uievents.KeyupInit(active, eventInit)
		c.Window.SetTimeout(func() error {
			c.processNextKey(getNext, stop)
			return nil
		}, k.KeyupDelay)
	} else {
		c.processNextKey(getNext, stop)
	}
}

func (c KeyboardController) processNextKey(getNext func() (key.Key, bool), stop func()) {
	k, ok := getNext()
	if !ok {
		return
	}
	eventInit := k.EventInit()
	if k.KeyupDelay < 0 || k.KeydownDelay < 0 {
		c.Window.Logger().Error(
			"KeyboardController: Negative KeyupDelay - aborting key sequence",
			slog.Any("key", k),
		)
		stop()
		return
	}
	if k.Down {
		active := c.Window.Document().ActiveElement()
		if uievents.KeydownInit(active, eventInit) {
			c.handleKey(active, k)
		}
		c.Window.SetTimeout(func() error {
			c.dispatchKeyUp(getNext, k, stop)
			return nil
		}, k.KeydownDelay)
	} else {
		c.dispatchKeyUp(getNext, k, stop)
	}
}

// SendKeys simulates the user typing a sequence of keys. The key package
// contains functionality to generate sequences of keys from an input string.
//
// Ignored if no Window is assigned.
func (c KeyboardController) SendKeys(keys iter.Seq[key.Key]) {
	if c.Window == nil {
		return
	}
	next, stop := iter.Pull(keys)
	c.processNextKey(next, stop)
	c.Window.Clock().Advance(0)
}

func (c KeyboardController) SendText(text string) {
	c.SendKeys(key.StringToKeys(text))
}
