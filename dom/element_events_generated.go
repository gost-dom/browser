// This file is generated. Do not edit.

package dom

import event "github.com/gost-dom/browser/dom/event"

type elementEvents struct {
	target event.EventTarget
}

type ElementEvents interface {
	Auxclick() bool
	Blur() bool
	Click() bool
	Contextmenu() bool
	Focus() bool
	Focusin() bool
	Focusout() bool
}

func (e *elementEvents) Auxclick() bool {
	data := e.defaultPointerEventInit()
	event := &event.Event{Type: "auxclick", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Blur() bool {
	data := e.defaultFocusEventInit()
	event := &event.Event{Type: "blur", Data: data}
	event.Bubbles = false
	event.Cancelable = false
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Click() bool {
	data := e.defaultPointerEventInit()
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Contextmenu() bool {
	data := e.defaultPointerEventInit()
	event := &event.Event{Type: "contextmenu", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Focus() bool {
	data := e.defaultFocusEventInit()
	event := &event.Event{Type: "focus", Data: data}
	event.Bubbles = false
	event.Cancelable = false
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Focusin() bool {
	data := e.defaultFocusEventInit()
	event := &event.Event{Type: "focusin", Data: data}
	event.Bubbles = true
	event.Cancelable = false
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Focusout() bool {
	data := e.defaultFocusEventInit()
	event := &event.Event{Type: "focusout", Data: data}
	event.Bubbles = true
	event.Cancelable = false
	return e.target.DispatchEvent(event)
}
