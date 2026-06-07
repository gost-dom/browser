package controller_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/input/controller"
	"github.com/gost-dom/browser/input/key"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/internal/uievents"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestKeyboardController(t *testing.T) {
	suite := initKeyboardControllerSuite(t)

	input := suite.input

	suite.SendKey(key.RuneToKey('a'))
	suite.Expect(input).
		To(HaveIDLValue(""), "input field IDL value without focus after keyboard input")

	input.Focus()

	suite.SendKey(key.RuneToKey('a'))
	suite.Expect(input).
		To(HaveIDLValue("a"), "input field IDL value with focus after keyboard input")
	suite.Expect(input).
		ToNot(HaveAttribute("value", nil), "keyboard input should not set value content attribute")

	suite.SendKey(key.RuneToKey('b'))
	suite.Expect(input).To(HaveIDLValue("ab"), "subsequent input is appended to the value")
	suite.Expect(input).
		ToNot(HaveAttribute("value", nil), "subsequent input does not set value content attribute")
}

func TestKeyboardControllerPreventDefault(t *testing.T) {
	suite := initKeyboardControllerSuite(t)
	input := suite.input
	input.Focus()

	r := &EventRecorder{}

	input.AddEventListener("keydown", r)
	input.AddEventListener("keyup", r)
	input.AddEventListener("input", r)
	input.AddEventListener("change", r)

	input.AddEventListener("keydown", event.NewEventHandlerFunc(event.NoError(func(e *event.Event) {
		e.PreventDefault()
	})))

	suite.SendKey(key.RuneToKey('a'))

	suite.Expect(input).To(HaveIDLValue(""))
	suite.Expect(r).To(HaveRecordedEvents(
		&MatchEvent{Type: "keydown"},
		&MatchEvent{Type: "keyup"},
	))
}

func TestInputEventIsDispatchedAfterInputUpdates(t *testing.T) {
	suite := initKeyboardControllerSuite(t)

	input := suite.input
	input.Focus()
	var eventFired bool

	input.AddEventListener("input", event.NewEventHandlerFunc(func(e *event.Event) error {
		suite.Expect(e.Target).To(Equal(input))
		suite.Expect(e.Target).To(HaveIDLValue("a"))
		eventFired = true
		return nil
	}))

	suite.SendKey(key.RuneToKey('a'))

	suite.Expect(eventFired).To(BeTrue())
}

func TestEventsDispatched(t *testing.T) {
	suite := initKeyboardControllerSuite(t)
	input := suite.input
	input.Focus()
	r := &EventRecorder{}

	input.AddEventListener("keydown", r)
	input.AddEventListener("keyup", r)
	input.AddEventListener("input", r)
	input.AddEventListener("change", r)

	suite.SendKeys(key.StringToKeys("ab"))

	suite.Expect(r).To(HaveRecordedEvents(
		&MatchEvent{Type: "keydown"},
		&MatchEvent{Type: "input"},
		&MatchEvent{Type: "keyup"},
		&MatchEvent{Type: "keydown"},
		&MatchEvent{Type: "input"},
		&MatchEvent{Type: "keyup"},
	))
}

func TestStreamOfEventsWithShiftKey(t *testing.T) {
	suite := initKeyboardControllerSuite(t)
	input := suite.input
	input.Focus()
	r := &EventRecorder{}

	input.AddEventListener("keydown", r)
	input.AddEventListener("keyup", r)

	suite.SendKeys(key.StringToKeys("aBc"))

	suite.Expect(r).To(HaveRecordedEvents(
		&MatchEvent{Type: "keydown", Key: "a"},
		&MatchEvent{Type: "keyup", Key: "a"},
		&MatchEvent{Type: "keydown", Key: "Shift"},
		&MatchEvent{Type: "keydown", Key: "B"},
		&MatchEvent{Type: "keyup", Key: "B"},
		&MatchEvent{Type: "keyup", Key: "Shift"},
		&MatchEvent{Type: "keydown", Key: "c"},
		&MatchEvent{Type: "keyup", Key: "c"},
	))
	suite.Expect(input).To(HaveIDLValue("aBc"))
}

func getKey(e *event.Event) string {
	eventInit, _ := e.Data.(uievents.KeyboardEventInit)
	return eventInit.Key
}

func TestKeyboardControllerEventTiming(t *testing.T) {
	suite := initKeyboardControllerSuite(t)
	win := suite.win
	input := suite.input

	var msgs []string

	input.AddEventListener("keydown", event.NewEventHandlerFuncWithoutError(func(e *event.Event) {
		msgs = append(msgs, fmt.Sprintf("keydown: %s", getKey(e)))
		win.SetTimeout(func() error {
			msgs = append(msgs, fmt.Sprintf("after keydown: %s", getKey(e)))
			return nil
		}, 0)
	}))
	input.AddEventListener("keyup", event.NewEventHandlerFuncWithoutError(func(e *event.Event) {
		msgs = append(msgs, fmt.Sprintf("keyup: %s", getKey(e)))
		win.SetTimeout(func() error {
			msgs = append(msgs, fmt.Sprintf("after keyup: %s", getKey(e)))
			return nil
		}, 0)
	}))

	input.Focus()

	t.Run("Allows setTimeout with zero delay to be processed first", func(t *testing.T) {
		suite.SendKeys(key.StringToKeys("abc"))
		suite.Expect(msgs).To(Equal([]string{
			"keydown: a",
			"after keydown: a",
			"keyup: a",
			"after keyup: a",
			"keydown: b",
			"after keydown: b",
			"keyup: b",
			"after keyup: b",
			"keydown: c",
			"after keydown: c",
			"keyup: c",
			"after keyup: c",
		}))
	})

	msgs = nil
	t.Run("Allow simulating keyboard delays", func(t *testing.T) {
		const one_ms = time.Millisecond
		const two_ms = 2 * one_ms
		const three_ms = 3 * one_ms
		suite.SendKeys(key.StringToKeys("aB",
			key.WithKeydownDelay(one_ms),
			key.WithKeyupDelay(two_ms),
		))
		clock := win.Clock()

		suite.Expect(msgs).To(Equal([]string{
			"keydown: a",
			"after keydown: a",
		}))

		// One ms after the 'a' keydown event, the keyup should be dispatched
		msgs = nil
		clock.Advance(one_ms)
		suite.Expect(msgs).To(Equal([]string{
			"keyup: a",
			"after keyup: a",
		}))

		msgs = nil
		// Next key should be two milliseconds later, shift down
		clock.Advance(one_ms)
		suite.Expect(msgs).To(BeEmpty())
		clock.Advance(one_ms)
		suite.Expect(msgs).To(Equal([]string{
			"keydown: Shift",
			"after keydown: Shift",
		}))

		// One millisecond after Shift keydown, we get "B" keydown
		msgs = nil
		clock.Advance(one_ms)
		suite.Expect(msgs).To(Equal([]string{
			"keydown: B",
			"after keydown: B",
		}))

		// One millisecond after "B" keydown, we get "B" keyup
		msgs = nil
		clock.Advance(one_ms)
		suite.Expect(msgs).To(Equal([]string{
			"keyup: B",
			"after keyup: B",
		}))

		// Finally, the Shift keyup is two milliseconds after "B" keyup
		msgs = nil
		clock.Advance(one_ms)
		suite.Expect(msgs).To(BeEmpty())
		clock.Advance(one_ms)
		suite.Expect(msgs).To(Equal([]string{
			"keyup: Shift",
			"after keyup: Shift",
		}))
	})
}

func TestKeyboardControllerFocusChangeBetweenUpDownEvents(t *testing.T) {
	win := htmltest.NewWindowHTML(
		t,
		`<body><input id="input-1" type="text" /><input id="input-2" type="text" /></body>`,
	)
	input1 := win.HTMLDocument().GetHTMLElementById("input-1")
	input2 := win.HTMLDocument().GetHTMLElementById("input-2")
	input1.AddEventListener("keydown", event.NewEventHandlerFuncWithoutError(func(e *event.Event) {
		input2.Focus()
	}))
	r1 := &EventRecorder{}
	r2 := &EventRecorder{}
	input1.AddEventListener("keydown", r1)
	input2.AddEventListener("keyup", r2)

	input1.Focus()
	KeyboardController{win}.SendKeys(key.StringToKeys("a"))

	g := gomega.NewGomegaWithT(t)
	g.Expect(r1).To(HaveRecordedEvents(
		&MatchEvent{Type: "keydown", Key: "a"},
	))
	g.Expect(r2).To(HaveRecordedEvents(
		&MatchEvent{Type: "keyup", Key: "a"},
	))
}

func initKeyboardControllerSuite(t *testing.T) *keyboardControllerSuite {
	win := htmltest.NewWindowHTML(t, `<body><input id="input" type="text" /></body>`)
	input := win.HTMLDocument().GetHTMLElementById("input")

	return &keyboardControllerSuite{
		Gomega:             gomega.NewWithT(t),
		KeyboardController: KeyboardController{win},
		win:                win,
		doc:                win.HTMLDocument(),
		input:              input.(html.HTMLInputElement),
	}
}

type keyboardControllerSuite struct {
	gomega.Gomega
	KeyboardController
	win   htmltest.WindowHelper
	doc   htmltest.HTMLDocumentHelper
	input html.HTMLInputElement
}
