package event

import "github.com/gost-dom/browser/internal/entity"

/* -------- event -------- */

type Init interface {
	bubbles() bool
	cancelable() bool
}

type EventInit struct {
	Bubbles    bool
	Cancelable bool
}

func (d EventInit) bubbles() bool {
	return d.Bubbles
}

func (d EventInit) cancelable() bool {
	return d.Cancelable
}

type Event struct {
	entity.Entity
	Init
	phase         EventPhase
	cancelled     bool
	eventType     string
	stopped       bool
	target        EventTarget
	currentTarget EventTarget
}

func New(eventType string, eventInit Init) *Event {
	return &Event{
		Entity:    entity.New(),
		Init:      eventInit,
		eventType: eventType,
	}
}

func NewEventInit(options ...EventOption) EventInit {
	var init EventInit
	for _, o := range options {
		o(&init)
	}
	return init
}

func (e *Event) Bubbles() bool                  { return e.Init.bubbles() }
func (e *Event) Cancelable() bool               { return e.Init.cancelable() }
func (e *Event) Type() string                   { return e.eventType }
func (e *Event) StopPropagation()               { e.stopped = true }
func (e *Event) PreventDefault()                { e.cancelled = true }
func (e *Event) isStopped() bool                { return e.stopped }
func (e *Event) isCancelled() bool              { return e.Cancelable() && e.cancelled }
func (e *Event) EventPhase() EventPhase         { return e.phase }
func (e *Event) Target() EventTarget            { return e.target }
func (e *Event) CurrentTarget() EventTarget     { return e.currentTarget }
func (e *Event) setEventPhase(phase EventPhase) { e.phase = phase }
func (e *Event) setCurrentTarget(t EventTarget) { e.currentTarget = t }

func (e *Event) reset(t EventTarget) {
	e.target = t
	e.stopped = false
	e.cancelled = false
}
