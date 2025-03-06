package event

/* -------- customEvent -------- */

type CustomEventInit struct {
	EventInit
	Details interface{}
}

func NewCustomEvent(eventType string, init CustomEventInit) *Event {
	e := New(eventType, init)
	return e
}

/* -------- errorEvent -------- */

type ErrorEventInit struct {
	EventInit
	Err error
}

func NewErrorEvent(err error) *Event {
	e := New(
		"error",
		ErrorEventInit{Err: err, EventInit: EventInit{Bubbles: true}},
	)
	return e
}
