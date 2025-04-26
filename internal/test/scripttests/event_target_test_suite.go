package scripttests

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/onsi/gomega"
)

type EventTargetTestSuite struct {
	ScriptHostSuite
}

func NewEventTargetTestSuite(h html.ScriptHost) *EventTargetTestSuite {
	return &EventTargetTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *EventTargetTestSuite) TestPrototype() {
	s.Expect(s.MustEval("typeof (new EventTarget())")).To(Equal("object"))
	s.Expect(s.MustEval("(new EventTarget()) instanceof EventTarget")).To(BeTrue())
}

func (s *EventTargetTestSuite) TestCancelable() {
	s.Expect(s.MustEval(`
		const target = new EventTarget();
		target.addEventListener("custom", e => { e.preventDefault() });
		target.dispatchEvent(new CustomEvent("custom"))
	`)).To(BeTrue(), "Event shouldn't be cancelable by default")

	s.Expect(s.MustEval(`
		const target2 = new EventTarget();
		target2.addEventListener("custom", e => { e.preventDefault() });
		target2.dispatchEvent(new CustomEvent("custom", {cancelable: true }))
	`)).To(BeFalse())
}

func (s *EventTargetTestSuite) TestDOMEventBubbleNotSpecified() {
	s.MustLoadHTML(`<div id="parent"><div id="target"></div></div>`)
	s.MustRunScript(`
		var targetCalled = false;
		var parentCalled = false;
		const target = document.getElementById("target")
		target.addEventListener("go:home", e => { targetCalled = true });
		document.getElementById("parent").addEventListener(
			"go:home",
			e => { parentCalled = true });
		target.dispatchEvent(new CustomEvent("go:home", {}))
	`)
	s.Expect(s.MustEval("targetCalled")).To(BeTrue(), "Target handler called")
	s.Expect(s.MustEval("parentCalled")).To(BeFalse(), "Parent handler called")
}

func (s *EventTargetTestSuite) TestDOMEventBubble() {
	s.MustLoadHTML(`<div id="parent"><div id="target"></div></div>`)
	s.MustRunScript(`
		var targetCalled = false;
		var parentCalled = false;
		const target = document.getElementById("target")
		target.addEventListener("go:home", e => { targetCalled = true });
		document.getElementById("parent").addEventListener(
			"go:home",
			e => { parentCalled = true });
		target.dispatchEvent(new CustomEvent("go:home", { bubbles: true }))
	`)
	s.Assert().Equal(true, s.MustEval("targetCalled"))
	s.Assert().Equal(true, s.MustEval("parentCalled"))
}

func (s *EventTargetTestSuite) TestCallingEventListener() {
	s.MustRunScript(`
		var callCount = 0
		function listener() { callCount++ };
		const target = new EventTarget();
		target.addEventListener('custom', listener);
		target.dispatchEvent(new CustomEvent('custom'));
	`)
	s.Assert().EqualValues(1, s.MustEval("callCount"))
}

func (s *EventTargetTestSuite) TestPropagateGoEventToJS() {
	s.MustRunScript(`
		var callCount = 0
		var event;
		function listener(e) { event = e; callCount++ };
		const target = window;
		target.addEventListener('custom', listener);
	`)
	s.Window.DispatchEvent(event.NewCustomEvent("custom", event.CustomEventInit{}))
	s.Assert().EqualValues(1, s.MustEval("callCount"))
	s.Assert().Equal(true, s.MustEval(`Object.getPrototypeOf(event) === CustomEvent.prototype`))
	s.Assert().Equal("custom", s.MustEval(`event.type`), "type of actual event")
}
