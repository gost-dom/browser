package event

import "github.com/gost-dom/browser/internal/entity"

type Entity interface {
	ObjectId() entity.ObjectId
}

/* -------- event -------- */

type Init interface {
	GetBubbles() bool
	GetCancelable() bool
}

type EventInit struct {
	Bubbles    bool
	Cancelable bool
}

func (d EventInit) GetBubbles() bool {
	return d.Bubbles
}

func (d EventInit) GetCancelable() bool {
	return d.Cancelable
}

type Event struct {
	entity.Entity
	Init
	Type       string
	Bubbles    bool
	Cancelable bool

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
func New(eventType string, eventInit Init) *Event {
	return &Event{
		Init: eventInit,
		Type: eventType,
	}
}

func (e *Event) bubbles() bool              { return e.Bubbles || e.Init.GetBubbles() }
func (e *Event) cancelable() bool           { return e.Cancelable || e.Init.GetCancelable() }
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
