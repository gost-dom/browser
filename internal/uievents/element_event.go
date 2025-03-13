package uievents

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
)

// Dispatches a click event. Returns the return value from
// [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Click(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}
