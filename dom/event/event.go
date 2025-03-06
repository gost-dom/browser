package event

import "github.com/gost-dom/browser/internal/entity"

type Entity interface {
	ObjectId() entity.ObjectId
}

/* -------- event -------- */

// Event corresponds to a [DOM Event] dispatched by a [DOM EventTarget].
// Different types of events carry different types of data, which is represented
// by the Data property.
//
// The different event types have different data types, all carrying having the
// prefix, `EventInit`. This naming replect the naming in the IDL
// specifications. When types in JavaScript are constructed, concrete subclasses
// of `Event` in JavaScript are created based on the concrete type of the Data.
//
// Properties are arranged slightly differently on this type, than the DOM
// version, where the properties that affect the event dispatching behaviour,
// such as Bubbles, and Cancelable, are part of the EventInit, or options
// argument.
//
//	const = new CustomEvent("my-custom", { bubbles: true, details: "Something else" })
//
// The Go Event representation stores the value for Bubbles on the event itself.
// The other properties on the event options are data communicated between the
// event dispatcher and the event listener, which Gost doesn't case about, and
// as such is stored as an interface{} type.
//
// [DOM Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event
// [DOM EventTarget]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type Event struct {
	entity.Entity
	Type       string
	Bubbles    bool
	Cancelable bool
	Data       any

	phase         EventPhase
	cancelled     bool
	stopped       bool
	target        EventTarget
	currentTarget EventTarget
}

// New creates a new Event object passing a specific event type and event data.
//
// Deprecated: Calling New was originally necessary to handle object
// initialization, but this is no longer necessary, and it's suggested to just
// create an Event objects directly
func New(eventType string, eventInit any) *Event {
	return &Event{
		Data: eventInit,
		Type: eventType,
	}
}

func (e *Event) cancelable() bool           { return e.Cancelable }
func (e *Event) bubbles() bool              { return e.Bubbles }
func (e *Event) StopPropagation()           { e.stopped = true }
func (e *Event) PreventDefault()            { e.cancelled = true }
func (e *Event) EventPhase() EventPhase     { return e.phase }
func (e *Event) Target() EventTarget        { return e.target }
func (e *Event) CurrentTarget() EventTarget { return e.currentTarget }

func (e *Event) reset(t EventTarget) {
	e.target = t
	e.stopped = false
	e.cancelled = false
}
