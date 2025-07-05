package controller_test

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

type Key string

func KeyChar(r rune) Key {
	return Key(string(r))
}

type KeyboardController struct {
	Window html.Window
}

func (c KeyboardController) SendKey(k Key) {
	active := c.Window.Document().ActiveElement()
	switch e := active.(type) {
	case html.HTMLInputElement:
		e.DispatchEvent(&event.Event{Type: "keydown"})
		e.DispatchEvent(&event.Event{Type: "keypress"})
		e.DispatchEvent(&event.Event{Type: "input"})
		e.DispatchEvent(&event.Event{Type: "keyup"})
		e.SetAttribute("value", e.Value()+string(k))
	}
}

func TestKeyboardController(t *testing.T) {
	g := gomega.NewWithT(t)
	html := `
		<body>
			<input id="input" type="text" />
		</body>
	`
	win := htmltest.NewWindowHTML(t, html)
	ctrl := KeyboardController{win}

	input := win.HTMLDocument().GetHTMLElementById("input")

	ctrl.SendKey(KeyChar('a'))
	g.Expect(input).To(HaveIDLValue(""), "Keypress when input does not have focus")

	input.Focus()

	ctrl.SendKey(KeyChar('a'))
	g.Expect(input).To(HaveIDLValue("a"), "Keypress when input does not have focus")
	g.Expect(input).To(HaveAttribute("value", "a"), "Keypress when input does not have focus")

	ctrl.SendKey(KeyChar('b'))
	g.Expect(input).To(HaveIDLValue("ab"), "Keypress when input does not have focus")
	g.Expect(input).To(HaveAttribute("value", "ab"), "Keypress when input does not have focus")
}

func TestEventsDispatched(t *testing.T) {
	g := gomega.NewWithT(t)
	html := `
		<body>
			<input id="input" type="text" />
		</body>
	`
	win := htmltest.NewWindowHTML(t, html)
	input := win.HTMLDocument().GetHTMLElementById("input")
	input.Focus()
	r := &EventRecorder{}

	input.AddEventListener("keydown", r)
	input.AddEventListener("keyup", r)
	input.AddEventListener("keypress", r)
	input.AddEventListener("input", r)
	input.AddEventListener("change", r)

	ctrl := KeyboardController{win}
	ctrl.SendKey(Key('a'))

	g.Expect(r).To(HaveRecordedEvents(
		&MatchEvent{Type: "keydown"},
		&MatchEvent{Type: "keypress"},
		&MatchEvent{Type: "input"},
		&MatchEvent{Type: "keyup"},
	))
}

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
	m := gomega.HaveExactElements(expected)
	return gcustom.MakeMatcher(func(rec *EventRecorder) (bool, error) {
		return m.Match(rec.Events)
	})
}

type MatchEvent struct {
	Type   string
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
	if isEvent {
		return e.actual.Type == e.Type, nil
	} else {
		return false, fmt.Errorf("Value is not an event")
	}
}

func (e *MatchEvent) FailureMessage(actual any) (message string) {
	return fmt.Sprintf("Expected event type: %s. Got: %s", e.Type, e.actual.Type)
}

func (e *MatchEvent) NegatedFailureMessage(actual any) (message string) {
	return fmt.Sprintf("Expected event type to not be: %s", e.Type)
}
