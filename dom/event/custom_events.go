package event

/* -------- customEvent -------- */

type CustomEventInit struct {
	EventInit
	Details interface{}
}

func NewCustomEvent(eventType string, options ...EventOption) *Event {
	var init CustomEventInit
	for _, o := range options {
		o(&init.EventInit)
	}
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
