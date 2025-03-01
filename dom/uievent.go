package dom

type UIEvent struct {
	Event
}

func NewUIEvent(type_ string, options ...EventOption) UIEvent {
	return UIEvent{NewEvent(type_, options...)}
}

type MouseEvent struct{ UIEvent }

type PointerEvent struct{ MouseEvent }

func NewMouseEvent(type_ string, options ...EventOption) MouseEvent {
	return MouseEvent{NewUIEvent(type_, options...)}
}

func NewPointerEvent(type_ string, options ...EventOption) PointerEvent {
	return PointerEvent{NewMouseEvent(type_, options...)}
}
