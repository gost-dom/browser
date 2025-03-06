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
	event := &event.Event{
		Data: data,
		Type: "auxclick",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Click() bool {
	data := PointerEventInit{}
	event := &event.Event{
		Data: data,
		Type: "click",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Contextmenu() bool {
	data := PointerEventInit{}
	event := &event.Event{
		Data: data,
		Type: "contextmenu",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}
