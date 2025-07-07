package controller_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/input/controller"
	"github.com/gost-dom/browser/input/key"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
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
