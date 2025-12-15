package event

import (
	"log/slog"
	"slices"
	"strconv"

	"github.com/gost-dom/browser/internal/log"
)

type EventPhase int

const (
	EventPhaseNone     EventPhase = 0
	EventPhaseCapture  EventPhase = 1
	EventPhaseAtTarget EventPhase = 2
	EventPhaseBubbling EventPhase = 3
)

func (i EventPhase) String() string {
	switch i {
	case EventPhaseNone:
		return "none"
	case EventPhaseCapture:
		return "capture"
	case EventPhaseAtTarget:
		return "atTarget"
	case EventPhaseBubbling:
		return "bubbling"
	default:
		return strconv.Itoa(int(i))
	}
}

// Deprecated: Prefer using [WithCapture] - it provides a more consistent API
func Capture(o *EventListener) { o.Capture = true }

// Deprecated: Prefer using [WithOnce] - it provides a more consistent API
func Once(o *EventListener) { o.Once = true }

type EventListenerOption = func(*EventListener)

// WithCapture sets whether the event handler is called during the event capture
// phase.
func WithCapture(c bool) EventListenerOption {
	return func(e *EventListener) { e.Capture = c }
}

// WitnOnce sets whether the event handler is removed automatically at first
// invocation.
func WithOnce(c bool) EventListenerOption {
	return func(e *EventListener) { e.Once = c }
}

type EventTarget interface {
	AddEventListener(eventType string, listener EventHandler, options ...EventListenerOption)
	RemoveEventListener(eventType string, listener EventHandler, options ...EventListenerOption)
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

func (t *eventTarget) SetParentTarget(parent EventTarget) { t.parentTarget = parent }

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

// AddEventListener implements core DOM event dispatch functionality.
//
// The following options are not implemented yet:
// - passive - https://github.com/gost-dom/browser/issues/39
// - signal - https://github.com/gost-dom/browser/issues/40
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (e *eventTarget) AddEventListener(
	eventType string,
	handler EventHandler,
	options ...func(*EventListener),
) {
	listener := e.createListener(handler, options)
	e.logger().Debug("AddEventListener",
		slog.Group("event",
			slog.String("type", eventType)),
		slog.Group("options",
			slog.Bool("capture", listener.Capture),
			slog.Bool("once", listener.Once)),
	)
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
			e.lmap[eventType] = slices.Delete(listeners, i, i+1)
			return
		}
	}
}

func (e *eventTarget) SetCatchAllHandler(handler EventHandler) {
	e.catchAllHandler = handler
}

func (e *eventTarget) DispatchEvent(event *Event) bool {
	e.logger().Debug("Dispatch event", "EventType", event.Type)
	event.Target = e.self
	event.stopped = false
	event.DefaultPrevented = false

	event.EventPhase = EventPhaseCapture
	e.dispatchOnParent(event, true)

	event.EventPhase = EventPhaseAtTarget
	e.dispatchEvent(event, true)
	e.dispatchEvent(event, false)
	event.EventPhase = EventPhaseBubbling

	e.dispatchOnParent(event, false)

	event.EventPhase = EventPhaseNone

	return !(event.Cancelable && event.DefaultPrevented)
}

func (e *eventTarget) dispatchEvent(event *Event, capture bool) {
	log := e.logger().With(newLogAttrForEvent(event))
	log.Debug("eventTarget.dispatchEvent")

	if event.stopped {
		return
	}
	event.CurrentTarget = e.self
	defer func() { event.CurrentTarget = nil }()
	if e.catchAllHandler != nil && !capture {
		if err := e.catchAllHandler.HandleEvent(event); err != nil {
			e.handleError(event, err, log.With(slog.String("handler", "catchAll")))
		}
	}

	listeners := e.lmap[event.Type]
	for i := 0; i < len(listeners); i++ {
		l := listeners[i]
		if l.Capture == capture {
			childLog := log.With(newLogAttrForListener(l))
			childLog.Debug("eventTarget.dispatchEvent: call handler")
			if l.Once {
				listeners = slices.Delete(listeners, i, i+1)
				i--
				e.lmap[event.Type] = listeners
			}
			if err := l.Handler.HandleEvent(event); err != nil {
				e.handleError(event, err, childLog)
			}
		}
	}
}

func (e *eventTarget) handleError(event *Event, err error, l *slog.Logger) {
	l.Error(
		"eventTarget.dispatchEvent: Error occurred in event handler",
		log.ErrAttr(err),
	)

	if event.Type != "error" {
		e.dispatchError(err)
	}
}

func (e *eventTarget) dispatchOnParent(event *Event, capture bool) {
	if e.parentTarget != nil {
		if capture {
			e.parentTarget.dispatchOnParent(event, capture)
			e.parentTarget.dispatchEvent(event, capture)
		} else {
			if event.Bubbles {
				e.parentTarget.dispatchEvent(event, capture)
				e.parentTarget.dispatchOnParent(event, capture)
			}
		}
	}
}

func (e *eventTarget) dispatchError(err error) {
	if e.parentTarget == nil {
		event := NewErrorEvent(err)
		e.DispatchEvent(event)
	} else {
		e.parentTarget.dispatchError(err)
	}
}

func (e *eventTarget) logger() (res *slog.Logger) {
	if source, ok := e.self.(log.LogSource); ok {
		res = source.Logger()
	}
	if res == nil {
		res = log.Default()
	}
	return
}
