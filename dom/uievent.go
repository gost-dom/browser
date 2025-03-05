package dom

import "github.com/gost-dom/browser/dom/event"

type UIEvent = *event.Event

type UIEventInit struct {
	event.EventInit
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

func NewUIEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.New(type_, UIEventInit{
		EventInit: event.NewEventInit(options...),
	})
}

// type MouseEvent struct{ UIEvent }
//
// type PointerEvent struct{ MouseEvent }

func NewMouseEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.New(type_, MouseEventInit{UIEventInit: UIEventInit{
		EventInit: event.NewEventInit(options...),
	}})
}

func NewPointerEvent(type_ string, options ...event.EventOption) *event.Event {
	return event.New(
		type_,
		PointerEventInit{
			MouseEventInit: MouseEventInit{UIEventInit: UIEventInit{
				EventInit: event.NewEventInit(options...),
			}},
		},
	)
}
