// This file is generated. Do not edit.

package dom

import "github.com/gost-dom/browser/dom/events"

type elementEvents struct {
	target events.EventTarget
}

type ElementEvents interface {
	Auxclick() bool
	Click() bool
	Contextmenu() bool
}

func (e *elementEvents) Auxclick() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("auxclick", events.EventBubbles(true), events.EventCancelable(true)),
	)
}

func (e *elementEvents) Click() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("click", events.EventBubbles(true), events.EventCancelable(true)),
	)
}

func (e *elementEvents) Contextmenu() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("contextmenu", events.EventBubbles(true), events.EventCancelable(true)),
	)
}
