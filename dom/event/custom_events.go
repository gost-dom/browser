package event

/* -------- customEvent -------- */

type CustomEventInit struct {
	Detail interface{}
}

func NewCustomEvent(eventType string, init CustomEventInit) *Event {
	e := New(eventType, init)
	return e
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
