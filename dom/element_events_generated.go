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
	init.Bubbles = true
	init.Cancelable = true
	return e.target.DispatchEvent(
		NewPointerEvent("auxclick", init),
	)
}

func (e *elementEvents) Click() bool {
	init := PointerEventInit{}
	init.Bubbles = true
	init.Cancelable = true
	return e.target.DispatchEvent(
		NewPointerEvent("click", init),
	)
}

func (e *elementEvents) Contextmenu() bool {
	init := PointerEventInit{}
	init.Cancelable = true
	init.Bubbles = true
	return e.target.DispatchEvent(
		NewPointerEvent("contextmenu", init),
	)
}
