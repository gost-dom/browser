package event

import "github.com/gost-dom/browser/internal/entity"

type Entity interface {
	ObjectId() entity.ObjectId
}

/* -------- event -------- */

// Event corresponds to a [DOM Event] dispatched by a [DOM EventTarget].
// Different types of events carry different data. The event-specific data exist
// in the Data property, which must be a valid "EventInit" type. The data
// contains the event-specific values of the 2nd constructor argument.
//
//	// JS:
//	const event = new CustomEvent("my-custom", { bubbles: true, details: "Something else" })
//
// The corresponding Go event would be:
//
//	event := Event {
//	  Type: "My-custom",
//	  Bubbles: true,
//	  Data: CustomEventInit{ Details: "Something else" },
//	}
//
// The "EventInit" postfix reflect naming in IDL specifications.
//
// The Go Event representation stores the value for Bubbles on the event itself.
// The other properties on the event options are data communicated between the
// event dispatcher and the event listener, which Gost doesn't care about, and
// as such is stored as an interface{} type.
//
// In JavaScript, events are represented by concrete subclasses of the base
// Event class. The concrete class used will be determined by the Data property.
//
// [DOM Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event
// [DOM EventTarget]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type Event struct {
	entity.Entity
	Type       string
	Bubbles    bool
	Cancelable bool
	Data       any

	EventPhase       EventPhase
	DefaultPrevented bool
	stopped          bool
	Target           EventTarget
	CurrentTarget    EventTarget
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

func (e *Event) StopPropagation() { e.stopped = true }
func (e *Event) PreventDefault()  { e.DefaultPrevented = true }
