package dom

import "github.com/gost-dom/browser/dom/event"

type UIEvent = *event.Event

type UIEventInitDict struct {
	event.EventInitDict
	view event.EventTarget
}

type MouseEventInitDict struct {
	UIEventInitDict
	ScreenX int
	ScreenY int
}

type PointerEventInitDict struct {
	MouseEventInitDict
	PointerId int
}

func NewUIEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.NewEventInit(type_, UIEventInitDict{
		EventInitDict: event.NewEventInitDict(options...),
	})
}

// type MouseEvent struct{ UIEvent }
//
// type PointerEvent struct{ MouseEvent }

func NewMouseEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.NewEventInit(type_, MouseEventInitDict{UIEventInitDict: UIEventInitDict{
		EventInitDict: event.NewEventInitDict(options...),
	}})
}

func NewPointerEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.NewEventInit(
		type_,
		PointerEventInitDict{
			MouseEventInitDict: MouseEventInitDict{UIEventInitDict: UIEventInitDict{
				EventInitDict: event.NewEventInitDict(options...),
			}},
		},
	)
}
