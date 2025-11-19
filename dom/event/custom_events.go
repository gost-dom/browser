package event

/* -------- customEvent -------- */

type CustomEventInit struct {
	Detail interface{}
}

func NewCustomEvent(eventType string, init CustomEventInit) *Event {
	return &Event{Type: eventType, Data: init}
}

/* -------- errorEvent -------- */

type ErrorEventInit struct {
	Err error
}

func NewErrorEvent(err error) *Event {
	e := &Event{
		Type:    "error",
		Bubbles: true,
		Data:    ErrorEventInit{Err: err},
	}
	return e
}
