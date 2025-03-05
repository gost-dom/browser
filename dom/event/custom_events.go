package event

/* -------- customEvent -------- */

type CustomEventInitDict struct {
	EventInitDict
	Details interface{}
}

func NewCustomEvent(eventType string, options ...EventOption) *Event {
	var init CustomEventInitDict
	for _, o := range options {
		o(&init.EventInitDict)
	}
	e := newEvent(eventType, init)
	return e
}

/* -------- errorEvent -------- */

type ErrorEventInitDict struct {
	EventInitDict
	Err error
}

func NewErrorEvent(err error) *Event {
	e := newEvent(
		"error",
		ErrorEventInitDict{Err: err, EventInitDict: EventInitDict{Bubbles: true}},
	)
	return e
}
