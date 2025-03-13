// This file is generated. Do not edit.

package dom

import event "github.com/gost-dom/browser/dom/event"

type elementEvents struct {
	target event.EventTarget
}

// Deprecated: ElementEvents expose methods that are not part of the Element specification
type ElementEvents interface {
	// Deprecated: auxclick is not a method defined on Element in the DOM
	Auxclick() bool
	// Deprecated: click is not a method defined on Element in the DOM
	Click() bool
	// Deprecated: contextmenu is not a method defined on Element in the DOM
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
	init.Bubbles = true
	init.Cancelable = true
	return e.target.DispatchEvent(
		NewPointerEvent("contextmenu", init),
	)
}
