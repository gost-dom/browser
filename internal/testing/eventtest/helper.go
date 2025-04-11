package eventtest

import "github.com/gost-dom/browser/dom/event"

func NewTestHandler(
	f func(*event.Event),
) event.EventHandler {
	return event.NewEventHandlerFunc(event.NoError(f))
}

// PreventDefaultHandler creates an event handler that will call PreventDefault.
func PreventDefaultHandler() event.EventHandler {
	return NewTestHandler(func(e *event.Event) {
		e.PreventDefault()
	})
}
