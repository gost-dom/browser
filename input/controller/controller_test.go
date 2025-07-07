package controller_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/input/controller"
	"github.com/gost-dom/browser/input/key"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

type keyboardControllerSuite struct {
	gomega.Gomega
	KeyboardController
	win htmltest.WindowHelper
	doc htmltest.HTMLDocumentHelper
}

func initKeyboardControllerSuite(t *testing.T) *keyboardControllerSuite {
	html := `
		<body>
			<input id="input" type="text" />
		</body>
	`
	win := htmltest.NewWindowHTML(t, html)
	return &keyboardControllerSuite{
		Gomega:             gomega.NewWithT(t),
		KeyboardController: KeyboardController{win},
		win:                win,
		doc:                win.HTMLDocument(),
	}
}

func TestKeyboardController(t *testing.T) {
	suite := initKeyboardControllerSuite(t)

	input := suite.doc.GetHTMLElementById("input")

	suite.SendKey(key.RuneToKey('a'))
	suite.Expect(input).To(HaveIDLValue(""), "Keydown when input does not have focus")

	input.Focus()

	suite.SendKey(key.RuneToKey('a'))
	suite.Expect(input).To(HaveIDLValue("a"), "Keydown when input does not have focus")
	suite.Expect(input).ToNot(HaveAttribute("value", nil), "Keydown when input does not have focus")

	suite.SendKey(key.RuneToKey('b'))
	suite.Expect(input).To(HaveIDLValue("ab"), "Keydown when input does not have focus")
	suite.Expect(input).ToNot(HaveAttribute("value", nil), "Keydown when input does not have focus")
}

func TestInputEventIsDispatchedAfterInputUpdates(t *testing.T) {
	suite := initKeyboardControllerSuite(t)

	input := suite.doc.GetHTMLElementById("input")
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
	input := suite.doc.GetHTMLElementById("input")
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
