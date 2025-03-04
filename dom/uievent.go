package dom

import "github.com/gost-dom/browser/dom/events"

type UIEvent struct {
	events.Event
}

func NewUIEvent(type_ string, options ...events.EventOption) UIEvent {
	return UIEvent{events.NewEvent(type_, options...)}
}

type MouseEvent struct{ UIEvent }

type PointerEvent struct{ MouseEvent }

func NewMouseEvent(type_ string, options ...events.EventOption) MouseEvent {
	return MouseEvent{NewUIEvent(type_, options...)}
}

func NewPointerEvent(type_ string, options ...events.EventOption) PointerEvent {
	return PointerEvent{NewMouseEvent(type_, options...)}
}
