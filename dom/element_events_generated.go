// This file is generated. Do not edit.

package dom

import event "github.com/gost-dom/browser/dom/event"

type elementEvents struct {
	target event.EventTarget
}

type ElementEvents interface {
	Auxclick() bool
	Click() bool
	Contextmenu() bool
}

func (e *elementEvents) Auxclick() bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "auxclick", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Click() bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Contextmenu() bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "contextmenu", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}
