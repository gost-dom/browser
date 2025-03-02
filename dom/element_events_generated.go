// This file is generated. Do not edit.

package dom

type elementEvents struct {
	target EventTarget
}

type ElementEvents interface {
	Auxclick() bool
	Click() bool
	Contextmenu() bool
}

func (e *elementEvents) Auxclick() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("auxclick", EventBubbles(true), EventCancelable(true)),
	)
}

func (e *elementEvents) Click() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("click", EventBubbles(true), EventCancelable(true)),
	)
}

func (e *elementEvents) Contextmenu() bool {
	return e.target.DispatchEvent(
		NewPointerEvent("contextmenu", EventBubbles(true), EventCancelable(true)),
	)
}
