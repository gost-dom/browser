package event_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/eventtest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
)

type TestSuite struct {
	gosttest.GomegaSuite
	clickCount   int
	target       EventTarget
	clickHandler EventHandler
}

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.clickCount = 0
	s.target = NewEventTarget()
	s.clickHandler = NewTestHandler(func(e *Event) { s.clickCount++ })
	s.target.AddEventListener("click", s.clickHandler)
}

func (s *TestSuite) TestDispatchEvent() {
	s.target.DispatchEvent(NewCustomEvent("click", CustomEventInit{}))
	s.Expect(s.clickCount).To(Equal(1))
}

func (s *TestSuite) TestDispatchDifferentEvent() {
	s.target.DispatchEvent(NewCustomEvent("auxclick", CustomEventInit{}))
	s.Expect(s.clickCount).To(Equal(0))
}

func (s *TestSuite) TestRemoveEventHandler() {
	s.target.DispatchEvent(NewCustomEvent("click", CustomEventInit{}))
	s.Expect(s.clickCount).To(Equal(1))

	s.target.RemoveEventListener("click", s.clickHandler)
	s.target.DispatchEvent(NewCustomEvent("click", CustomEventInit{}))
	s.Expect(s.clickCount).To(Equal(1))
}

func (s *TestSuite) TestAddingSameHandlerTwice() {
	s.target.AddEventListener("click", s.clickHandler)
	s.target.AddEventListener("click", s.clickHandler)
	s.target.DispatchEvent(NewCustomEvent("click", CustomEventInit{}))
	s.Expect(s.clickCount).To(Equal(1))
}

func (s *TestSuite) TestEventHandlersCalledInOrder() {
	var order []string
	s.target.AddEventListener("auxclick",
		NewTestHandler(func(e *Event) { order = append(order, "A") }),
	)
	s.target.AddEventListener("auxclick",
		NewTestHandler(func(e *Event) { order = append(order, "B") }),
	)
	s.target.AddEventListener("auxclick",
		NewTestHandler(func(e *Event) { order = append(order, "C") }),
	)
	s.target.DispatchEvent(NewCustomEvent("auxclick", CustomEventInit{}))
	s.Expect(order).To(Equal([]string{"A", "B", "C"}))
}

type EventPropagationTestSuiteBase struct {
	gosttest.GomegaSuite
	window html.Window
	target dom.Element
}

type EventPropagationTestSuite struct {
	EventPropagationTestSuiteBase
}

func TestEventPropagation(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(EventPropagationTestSuite))
}

func (s *EventPropagationTestSuiteBase) SetupTest() {
	var err error
	s.window, err = html.NewWindowReader(
		strings.NewReader(`<body><div id="target"></div></body>`),
	)
	s.Expect(err).ToNot(HaveOccurred())
	s.target = s.window.Document().GetElementById("target")
}

func (s *EventPropagationTestSuite) TestRemoveCorrectPhase() {
	var events []string
	h := NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Phase: %d", e.EventPhase))
	})
	event := &event.Event{Type: "gost:remove", Bubbles: true}
	s.window.AddEventListener("gost:remove", h)
	s.window.AddEventListener("gost:remove", h, Capture)

	s.target.DispatchEvent(event)
	s.Assert().Equal([]string{
		"Phase: 1",
		"Phase: 3",
	}, events)

	s.window.RemoveEventListener("gost:remove", h)

	events = nil
	s.target.DispatchEvent(event)
	s.Assert().Equal([]string{
		"Phase: 1",
	}, events)

	s.window.AddEventListener("gost:remove", h)
	s.window.RemoveEventListener("gost:remove", h, Capture)

	events = nil
	s.target.DispatchEvent(event)
	s.Assert().Equal([]string{
		"Phase: 3",
	}, events)
}

func (s *EventPropagationTestSuite) TestEventOnce() {
	var events []string
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Handler A"))
	}))
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Handler B"))
	}), Once)

	s.target.DispatchEvent(&event.Event{Type: "custom", Bubbles: true})
	s.target.DispatchEvent(&event.Event{Type: "custom", Bubbles: true})

	s.Assert().Equal([]string{
		"Handler A",
		"Handler B",
		"Handler A",
	}, events)
}

func (s *EventPropagationTestSuite) TestEventCapture() {
	var events []string
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Window capture. Phase: %d", e.EventPhase))
	}), Capture)
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Window bubble. Phase: %d", e.EventPhase))
	}))
	s.target.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Target capture. Phase: %d", e.EventPhase))
	}), Capture)
	s.target.AddEventListener("custom", NewTestHandler(func(e *Event) {
		events = append(events, fmt.Sprintf("Target bubble. Phase: %d", e.EventPhase))
	}))

	event := &Event{Type: "custom", Bubbles: true}
	s.Assert().Equal(EventPhaseNone, event.EventPhase, "Phase before dispatch")
	s.target.DispatchEvent(event)
	expected := []string{
		"Window capture. Phase: 1",
		"Target capture. Phase: 2",
		"Target bubble. Phase: 2",
		"Window bubble. Phase: 3",
	}

	s.Assert().Equal(expected, events)
	s.Assert().Equal(EventPhaseNone, event.EventPhase, "Phase after dispatch")

	events = nil
	s.target.DispatchEvent(&Event{Type: "custom"})
	s.Assert().Equal(
		[]string{
			"Window capture. Phase: 1",
			"Target capture. Phase: 2",
			"Target bubble. Phase: 2",
		}, events, "Event handlers when event doesn't bubble")

	events = nil
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		e.StopPropagation()
	}), Capture)
	s.target.DispatchEvent(event)
	s.Assert().Equal(
		[]string{
			"Window capture. Phase: 1",
		}, events, "Event handlers when propagation stopped")

}

func (s *EventPropagationTestSuite) TestDefaultEventPropagation() {
	called := false
	var l EventHandler = NewTestHandler(func(e *Event) {
		called = true
	})
	s.window.Document().Body().AddEventListener("custom", l)
	s.target.DispatchEvent(&Event{Type: "custom"})
	s.Expect(called).To(BeFalse())
}

func (s *EventPropagationTestSuite) TestPropagateToWindow() {
	called := false

	var l EventHandler = NewTestHandler(func(e *Event) {
		called = true
	})
	s.window.AddEventListener("custom", l)
	s.target.DispatchEvent(&Event{Type: "custom", Bubbles: true})
	s.Expect(called).To(BeTrue())
}

func (s *EventPropagationTestSuite) TestTargetOrCurrentTarget() {
	var actualEvent *Event
	var actualTarget EventTarget
	var actualCurrentTarget EventTarget

	var l EventHandler = NewTestHandler(func(e *Event) {
		actualEvent = e
		actualTarget = e.Target
		actualCurrentTarget = e.CurrentTarget
	})
	s.window.AddEventListener("custom", l)
	s.target.DispatchEvent(&Event{Type: "custom", Bubbles: true})
	s.Expect(actualTarget).To(Equal(s.target), "Event target")
	s.Expect(actualCurrentTarget).To(Equal(s.window), "CurrentEvent target")
	s.Expect(actualEvent.CurrentTarget).To(BeNil(), "CurrentTarget after event")
	s.Expect(actualEvent.Target).To(Equal(s.target), "Target after event")
}

func (s *EventPropagationTestSuite) TestPropagateToWindowBubbles() {
	called := false

	// window.Document()
	var l EventHandler = NewTestHandler(func(e *Event) {
		called = true
	})
	s.window.AddEventListener("custom", l)
	s.target.DispatchEvent(&Event{Type: "custom", Bubbles: true})
	s.Expect(called).To(BeTrue())
}

func (s *EventPropagationTestSuite) TestStopPropagation() {
	calledA := false
	calledB := false
	s.window.Document().Body().
		AddEventListener("custom", NewTestHandler(func(e *Event) {
			calledA = true
			e.StopPropagation()
		}))
	s.window.AddEventListener("custom", NewTestHandler(func(e *Event) {
		calledB = true
	}))
	s.target.DispatchEvent(&Event{Type: "custom", Bubbles: true})
	s.Expect(calledA).To(BeTrue(), "Event dispatched on body")
	s.Expect(calledB).To(BeFalse(), "Event dispatched on window")
}

func (s *EventPropagationTestSuite) TestReturnValueForPreventDefault() {
	s.target.AddEventListener("custom", NewTestHandler(func(e *Event) {
		e.PreventDefault()
	}))
	s.Assert().False(
		s.target.DispatchEvent(&Event{Type: "custom", Cancelable: true}),
		"DispatchEvent return value with default prevented, Cancelable: true",
	)
	s.Assert().True(
		s.target.DispatchEvent(&Event{Type: "custom", Cancelable: false}),
		"DispatchEvent return value with default prevented, Cancelable: false",
	)
	s.Assert().True(
		s.target.DispatchEvent(&Event{Type: "custom"}),
		"DispatchEvent return value with default prevented, Cancelable not set",
	)
}

func (s *EventPropagationTestSuite) TestEventHandlerGeneratesError() {
	var errorOnWindow bool
	s.window.AddEventListener("error", NewTestHandler(func(e *Event) {
		errorOnWindow = true
	}))
	var errorOnTarget bool
	s.target.AddEventListener("error", NewTestHandler(func(e *Event) {
		errorOnTarget = true
	}))
	s.target.AddEventListener("custom", NewEventHandlerFunc(func(e *Event) error {
		return errors.New("Error")
	}))
	s.target.DispatchEvent(&Event{Type: "custom"})
	s.Assert().True(errorOnWindow, "Error event dispached on Window")
	s.Assert().False(errorOnTarget, "Error event dispached on original target")
}
