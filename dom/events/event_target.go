package events

import (
	"slices"

	"github.com/gost-dom/browser/internal/entity"
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
	// Adds a listener that will receive _all_ dispatched events. This listener
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
	return func(e *EventInitDict) { e.bubbles = bubbles }
}

func EventCancelable(cancelable bool) EventOption {
	return func(e *EventInitDict) { e.cancelable = cancelable }
}

/* -------- event -------- */

type EventInit interface {
	Bubbles() bool
	Cancelable() bool
}

type EventInitDict struct {
	bubbles    bool
	cancelable bool
}

func (d EventInitDict) Bubbles() bool {
	return d.bubbles
}

func (d EventInitDict) Cancelable() bool {
	return d.cancelable
}

type Event struct {
	entity.Entity
	EventInit
	phase         EventPhase
	cancelled     bool
	eventType     string
	stopped       bool
	target        EventTarget
	currentTarget EventTarget
}

func newEvent(eventType string, eventInit EventInit) *Event {
	return &Event{
		Entity:    entity.New(),
		EventInit: eventInit,
		eventType: eventType,
	}
}

func NewEventInitDict(options ...EventOption) EventInitDict {
	var init EventInitDict
	for _, o := range options {
		o(&init)
	}
	return init
}

func NewEvent(eventType string, options ...EventOption) *Event {
	var init EventInitDict
	for _, o := range options {
		o(&init)
	}
	e := newEvent(eventType, init)
	return e
}
func NewEventInit(eventType string, init EventInit) *Event {
	e := newEvent(eventType, init)
	return e
}

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
		ErrorEventInitDict{Err: err, EventInitDict: EventInitDict{bubbles: true}},
	)
	return e
}

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
