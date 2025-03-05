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
	return e.target.DispatchEvent(
		NewPointerEvent("auxclick", event.EventBubbles(true), event.EventCancelable(true)),
	)
}

func (e *elementEvents) Click() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("click", event.EventBubbles(true), event.EventCancelable(true)),
	)
}

func (e *elementEvents) Contextmenu() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("contextmenu", event.EventBubbles(true), event.EventCancelable(true)),
	)
}
