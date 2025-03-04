package dom

import "github.com/gost-dom/browser/dom/events"

type UIEvent = *events.Event

type UIEventInitDict struct {
	events.EventInitDict
	view events.EventTarget
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

func NewUIEvent(type_ string, options ...events.EventOption) *events.Event {
	return events.NewEventInit(type_, UIEventInitDict{
		EventInitDict: events.NewEventInitDict(options...),
	})
}

// type MouseEvent struct{ UIEvent }
//
// type PointerEvent struct{ MouseEvent }

func NewMouseEvent(type_ string, options ...events.EventOption) *events.Event {
	return events.NewEventInit(type_, MouseEventInitDict{UIEventInitDict: UIEventInitDict{
		EventInitDict: events.NewEventInitDict(options...),
	}})
}

func NewPointerEvent(type_ string, options ...events.EventOption) *events.Event {
	return events.NewEventInit(
		type_,
		PointerEventInitDict{
			MouseEventInitDict: MouseEventInitDict{UIEventInitDict: UIEventInitDict{
				EventInitDict: events.NewEventInitDict(options...),
			}},
		},
	)
}
