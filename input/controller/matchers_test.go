package controller_test

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// Gomega matchers used in this test package alone for now. Will probably move
// to a general helpers package.

// EventRecorder implements an [event.EventHandler] that records the dispatched
// events.
//
// The events are not stored as pointers to make a copy, so cancelling or
// stopPropagation will not affect the real event.
type EventRecorder struct {
	Events []event.Event
}

// assert that EventRecorder is a valid EventHandler
var _ = event.EventHandler(&EventRecorder{})

// HandleEvent implements HandleEvent of [event.EventHandler]
func (r *EventRecorder) HandleEvent(e *event.Event) error {
	r.Events = append(r.Events, *e)
	return nil
}

// Equals implements Equals of [event.EventHandler]
func (r *EventRecorder) Equals(other event.EventHandler) bool {
	h, ok := other.(*EventRecorder)
	return ok && h == r
}

func HaveRecordedEvents(expected ...types.GomegaMatcher) types.GomegaMatcher {
	return gomega.WithTransform(func(rec *EventRecorder) []event.Event {
		return rec.Events
	}, gomega.HaveExactElements(expected))
	// m := gomega.HaveExactElements(expected)
	// return gcustom.MakeMatcher(func(rec *EventRecorder) (bool, error) {
	// 	return m.Match(rec.Events)
	// })
}

type MatchEvent struct {
	Type   string
	Key    string
	actual event.Event
}

func (e *MatchEvent) Match(actual any) (success bool, err error) {
	var (
		isEvent bool
		ptr     *event.Event
	)
	e.actual, isEvent = actual.(event.Event)
	if !isEvent {
		if ptr, isEvent = actual.(*event.Event); isEvent {
			e.actual = *ptr
		}
	}
	fmt.Println("Testing one two")
	if !isEvent {
		fmt.Println("Not an event")
		return false, fmt.Errorf("Value is not an event")
	}
	data := e.actual.Data
	if e.Key != "" {
		eventInit, ok := data.(uievents.KeyboardEventInit)
		if !ok {
			fmt.Println("Nit a keyboard event")
			return true, fmt.Errorf("Expected a KeyboardEvent")
		}
		if eventInit.Key != e.Key {
			fmt.Println("Key doesn't match", e.Key)
			return false, nil
		}
	}
	return e.actual.Type == e.Type, nil
}

func (e *MatchEvent) FailureMessage(actual any) (message string) {
	return fmt.Sprintf("Expected event type: %s. Got: %s", e.Type, e.actual.Type)
}

func (e *MatchEvent) NegatedFailureMessage(actual any) (message string) {
	return fmt.Sprintf("Expected event type to not be: %s", e.Type)
}
