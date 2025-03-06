package testing

import "github.com/gost-dom/browser/dom/event"

func NewTestHandler(
	f func(*event.Event),
) event.EventHandler {
	return event.NewEventHandlerFunc(event.NoError(f))
}
