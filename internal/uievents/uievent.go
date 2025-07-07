package uievents

import "github.com/gost-dom/browser/dom/event"

type UIEvent = *event.Event

type UIEventInit struct {
	view event.EventTarget
}

type MouseEventInit struct {
	UIEventInit
	ScreenX int
	ScreenY int
}

type PointerEventInit struct {
	MouseEventInit
	PointerId int
}

type FocusEventInit struct {
	UIEventInit
	RelatedTarget event.EventTarget
}

type KeyboardEventInit struct {
	UIEventInit
	Key string
}

type InputEventInit struct {
	UIEventInit
}

func NewUIEvent(type_ string) *event.Event {
	return event.New(type_, UIEventInit{})
}

func NewPointerEvent(type_ string, init PointerEventInit) *event.Event {
	return event.New(type_, init)
}
