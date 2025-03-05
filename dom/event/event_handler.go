package event

import "github.com/gost-dom/browser/internal/entity"

/* -------- EventHandler -------- */

// EventHandler is the interface for an event handler. In JavaScript; an event
// handler can be a function; or an object with a `handleEvent` function. In Go
// code, you can provide your own implementation, or use [NewEventHandlerFunc]
// to create a valid handler from a function.
//
// Multiple EventHandler instances can represent the same underlying event
// handler. E.g., when JavaScript code calls RemoveEventListener, a new Go
// struct is created wrapping the same underlying handler.
//
// The Equals function must return true when the other event handler is the same
// as the current value, so event handlers can properly be removed, and avoiding
// duplicates are added by AddEventListener.
type EventHandler interface {
	// HandleEvent is called when the the event occurrs.
	//
	// An non-nil error return value will dispatch an error event on the global
	// object in a normally configured environment.
	HandleEvent(event *Event) error
	// Equals must return true, if they underlying event handler of the other
	// handler is the same as this handler.
	Equals(other EventHandler) bool
}

type HandlerFuncWithoutError = func(*Event)
type HandlerFunc = func(*Event) error

type eventHandlerFunc struct {
	handlerFunc func(*Event) error
	id          entity.ObjectId
}

// NewEventHandlerFunc creates an [EventHandler] implementation from a compatible
// function.
//
// Note: Calling this twice for the same Go-function will be treated as
// different event handlers. Be sure to use the same instance returned from this
// function when removing.
func NewEventHandlerFunc(handler HandlerFunc) EventHandler {
	return eventHandlerFunc{handler, entity.NewObjectId()}
}

// NoError takes a function accepting a single argument and has no return value,
// and transforms it into a function that can be used where an error return
// value is expected.
func NoError[T func(U), U any](f T) func(U) error {
	return func(u U) error {
		f(u)
		return nil
	}
}

func NewEventHandlerFuncWithoutError(handler HandlerFuncWithoutError) EventHandler {
	return eventHandlerFunc{func(event *Event) error {
		handler(event)
		return nil
	}, entity.NewObjectId()}
}

func (e eventHandlerFunc) HandleEvent(event *Event) error {
	return e.handlerFunc(event)
}

func (e eventHandlerFunc) Equals(handler EventHandler) bool {
	x, ok := handler.(eventHandlerFunc)
	return ok && x.id == e.id
}
