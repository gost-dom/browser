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
	init := PointerEventInit{}
	event := &event.Event{
		Init: init,
		Type: "auxclick",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Click() bool {
	init := PointerEventInit{}
	event := &event.Event{
		Init: init,
		Type: "click",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}

func (e *elementEvents) Contextmenu() bool {
	init := PointerEventInit{}
	event := &event.Event{
		Init: init,
		Type: "contextmenu",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}
