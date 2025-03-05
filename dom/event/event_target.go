package event

import (
	"slices"

	"github.com/gost-dom/browser/internal/log"
)

type EventPhase int

const (
	EventPhaseNone     EventPhase = 0
	EventPhaseCapture  EventPhase = 1
	EventPhaseAtTarget EventPhase = 2
	EventPhaseBubbline EventPhase = 3
)

func Capture(o *EventListener) { o.Capture = true }
func Once(o *EventListener)    { o.Once = true }

type EventTarget interface {
	AddEventListener(eventType string, listener EventHandler, options ...func(*EventListener))
	RemoveEventListener(eventType string, listener EventHandler, options ...func(*EventListener))
	DispatchEvent(event *Event) bool
	// Adds a listener that will receive _all_ dispatched event. This listener
	// will not be removed from the window when navigating. This makes it useful
	// for a test to setup event listeners _before_ navigating, as by the time the
	// Navigate function returns, the DOMContentLoaded event _has_ fired, and
	// subscribed listeners have been called.
	SetCatchAllHandler(listener EventHandler)
	SetParentTarget(EventTarget)
	RemoveAll()
	// Unexported
	dispatchError(err error)
	dispatchEvent(event *Event, capture bool)
	dispatchOnParent(*Event, bool)
	setSelf(e EventTarget)
}

type EventListener struct {
	Handler EventHandler
	Capture bool
	Once    bool
}

type eventTarget struct {
	parentTarget    EventTarget
	lmap            map[string][]EventListener
	catchAllHandler EventHandler
	self            EventTarget
}

func NewEventTarget() EventTarget {
	res := newEventTarget()
	return &res
}

func (t *eventTarget) SetParentTarget(parent EventTarget) {
	t.parentTarget = parent
}

func SetEventTargetSelf(t EventTarget) {
	t.setSelf(t)
}

func newEventTarget() eventTarget {
	return eventTarget{
		lmap: make(map[string][]EventListener),
	}
}

func (e *eventTarget) setSelf(self EventTarget) {
	e.self = self
}

func (t *eventTarget) createListener(
	handler EventHandler,
	options []func(*EventListener),
) EventListener {
	listener := EventListener{Handler: handler}
	for _, o := range options {
		o(&listener)
	}
	return listener
}

func (e *eventTarget) AddEventListener(
	eventType string,
	handler EventHandler,
	options ...func(*EventListener),
) {
	listener := e.createListener(handler, options)
	log.Debug("AddEventListener", "EventType", eventType)
	// TODO: Handle options
	// - once
	// - passive. Defaults to false,
	// - signal - TODO: Implement AbortSignal
	// Browser specific
	// - Safari
	//   - passive defaults to true for `wheel`, `mousewheel` `touchstart`, `tourchmove` events
	// - Firefox (Gecko), receives an extra boolean argument, `wantsUntrusted`
	//   - https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener#wantsuntrusted
	listeners := e.lmap[eventType]
	for _, l := range listeners {
		if l.Handler.Equals(handler) && l.Capture == listener.Capture {
			return
		}
	}
	e.lmap[eventType] = append(listeners, listener)
}

func (e *eventTarget) RemoveAll() {
	e.lmap = make(map[string][]EventListener)
}

func (e *eventTarget) RemoveEventListener(
	eventType string,
	handler EventHandler,
	options ...func(*EventListener),
) {
	listener := e.createListener(handler, options)
	listeners := e.lmap[eventType]
	for i, l := range listeners {
		if l.Handler.Equals(handler) && l.Capture == listener.Capture {
			e.lmap[eventType] = append(listeners[:i], listeners[i+1:]...)
			return
		}
	}
}

func (e *eventTarget) SetCatchAllHandler(handler EventHandler) {
	e.catchAllHandler = handler
}

func (e *eventTarget) DispatchEvent(event *Event) bool {
	event.reset(e.self)
	log.Debug("Dispatch event", "EventType", event.Type())

	event.setEventPhase(EventPhaseCapture)
	e.dispatchOnParent(event, true)

	event.setEventPhase(EventPhaseAtTarget)
	e.dispatchEvent(event, true)
	e.dispatchEvent(event, false)
	event.setEventPhase(EventPhaseBubbline)

	e.dispatchOnParent(event, false)

	event.setEventPhase(EventPhaseNone)

	return !event.isCancelled()
}

func (e *eventTarget) dispatchEvent(event *Event, capture bool) {
	if event.isStopped() {
		return
	}
	event.setCurrentTarget(e.self)
	defer func() { event.setCurrentTarget(nil) }()
	if e.catchAllHandler != nil && !capture {
		if err := e.catchAllHandler.HandleEvent(event); err != nil {
			log.Debug("Error occurred", "error", err.Error())
			e.dispatchError(err)
		}
	}

	eventType := event.Type()
	listeners := e.lmap[eventType]
	for i := 0; i < len(listeners); i++ {
		l := listeners[i]
		log.Debug("eventTarget.dispatchEvent: Calling event handler", "type", event.Type())
		if l.Capture == capture {
			if err := l.Handler.HandleEvent(event); err != nil {
				e.handleError(err)
			}
			if l.Once {
				listeners = slices.Delete(listeners, i, i+1)
				i--
				e.lmap[eventType] = listeners
			}
		}
	}
}

func (e *eventTarget) handleError(err error) {
	log.Error(
		"eventTarget.dispatchEvent: Error occurred in event handler",
		"error",
		err.Error(),
	)
	e.dispatchError(err)
}

func (e *eventTarget) dispatchOnParent(event *Event, capture bool) {
	if e.parentTarget != nil {
		if capture {
			e.parentTarget.dispatchOnParent(event, capture)
			e.parentTarget.dispatchEvent(event, capture)
		} else {
			if event.Bubbles() {
				e.parentTarget.dispatchEvent(event, capture)
				e.parentTarget.dispatchOnParent(event, capture)
			}
		}
	}
}

func (e *eventTarget) dispatchError(err error) {
	event := NewErrorEvent(err)
	if e.parentTarget == nil {
		e.DispatchEvent(event)
	} else {
		e.parentTarget.dispatchError(err)
	}
}

/* -------- Event & CustomEvent -------- */

// type Event interface {
// 	entity.Entity
// 	Cancelable() bool
// 	Bubbles() bool
// 	PreventDefault()
// 	Type() string
// 	StopPropagation()
// 	Target() EventTarget
// 	CurrentTarget() EventTarget
// 	reset(t EventTarget)
// 	EventPhase() EventPhase
//
// 	isCancelled() bool
// 	isStopped() bool
// 	setEventPhase(phase EventPhase)
// 	setCurrentTarget(t EventTarget)
// }

// type CustomEvent interface {
// 	Event
// }

type EventOption func(*EventInitDict)

func EventOptions(options ...EventOption) EventOption {
	return func(e *EventInitDict) {
		for _, o := range options {
			o(e)
		}
	}
}

func EventBubbles(bubbles bool) EventOption {
	return func(e *EventInitDict) { e.Bubbles = bubbles }
}

func EventCancelable(cancelable bool) EventOption {
	return func(e *EventInitDict) { e.Cancelable = cancelable }
}
