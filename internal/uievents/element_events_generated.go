// This file is generated. Do not edit.

package uievents

import (
	dom "github.com/gost-dom/browser/dom"
	event "github.com/gost-dom/browser/dom/event"
)

// Dispatches a [auxclick event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [auxclick event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/auxclick_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Auxclick(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "auxclick", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}

// Dispatches a [blur event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [blur event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/blur_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Blur(e dom.Element) bool {
	data := FocusEventInit{}
	event := &event.Event{Type: "blur", Data: data}
	event.Bubbles = false
	event.Cancelable = false
	return e.DispatchEvent(event)
}

// Dispatches a [click event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [click event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Click(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}

// Dispatches a [contextmenu event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [contextmenu event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/contextmenu_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Contextmenu(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "contextmenu", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}

// Dispatches a [focus event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [focus event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/focus_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Focus(e dom.Element) bool {
	data := FocusEventInit{}
	event := &event.Event{Type: "focus", Data: data}
	event.Bubbles = false
	event.Cancelable = false
	return e.DispatchEvent(event)
}

// Dispatches a [focusin event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [focusin event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/focusin_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Focusin(e dom.Element) bool {
	data := FocusEventInit{}
	event := &event.Event{Type: "focusin", Data: data}
	event.Bubbles = true
	event.Cancelable = false
	return e.DispatchEvent(event)
}

// Dispatches a [focusout event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [focusout event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/focusout_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Focusout(e dom.Element) bool {
	data := FocusEventInit{}
	event := &event.Event{Type: "focusout", Data: data}
	event.Bubbles = true
	event.Cancelable = false
	return e.DispatchEvent(event)
}
