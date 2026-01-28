package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/testing/gosttest"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func testEventTarget(t *testing.T, e html.ScriptEngine) {
	w := browsertest.InitWindow(t, e)

	t.Run("Prototype", func(t *testing.T) {
		w.MustRun(`
			gost.assertTypeOf(new EventTarget(), 'object')
		`)
	})

	t.Run("Cancelable", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		Expect(w.MustEval(`
			const target = new EventTarget();
			target.addEventListener("custom", e => { e.preventDefault() });
			target.dispatchEvent(new CustomEvent("custom"))
		`)).To(BeTrue(), "Event shouldn't be cancelable by default")

		Expect(w.MustEval(`
			const target2 = new EventTarget();
			target2.addEventListener("custom", e => { e.preventDefault() });
			target2.dispatchEvent(new CustomEvent("custom", {cancelable: true }))
		`)).To(BeFalse())
	})

	t.Run("CallListener", func(t *testing.T) {
		w := browsertest.InitWindow(t, e)
		w.MustRun(`
			var callCount = 0
			function listener() { callCount++ };
			const target = new EventTarget();
			target.addEventListener('custom', listener);
			target.dispatchEvent(new CustomEvent('custom'));
		`)
		assert.EqualValues(t, 1, w.MustEval("callCount"))
	})

	t.Run("Event bubbling not specified", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		w := browsertest.InitWindow(t, e, browsertest.WithHtml(
			`<div id="parent"><div id="target"></div></div>`,
		))
		w.MustRun(`
			var targetCalled = false;
			var parentCalled = false;
			const target = document.getElementById("target")
			target.addEventListener("go:home", e => { targetCalled = true });
			document.getElementById("parent").addEventListener(
				"go:home",
				e => { parentCalled = true });
			target.dispatchEvent(new CustomEvent("go:home", {}))
		`)
		Expect(w.MustEval("targetCalled")).To(BeTrue(), "Target handler called")
		Expect(w.MustEval("parentCalled")).To(BeFalse(), "Parent handler called")
	})

	t.Run("Event Bubble", func(t *testing.T) {
		w := browsertest.InitWindow(t, e, browsertest.WithHtml(
			`<div id="parent"><div id="target"></div></div>`))
		w.MustRun(`
			var targetCalled = false;
			var parentCalled = false;
			const target = document.getElementById("target")
			target.addEventListener("go:home", e => { targetCalled = true });
			document.getElementById("parent").addEventListener(
				"go:home",
				e => { parentCalled = true });
			target.dispatchEvent(new CustomEvent("go:home", { bubbles: true }))
		`)
		assert.Equal(t, true, w.MustEval("targetCalled"))
		assert.Equal(t, true, w.MustEval("parentCalled"))
	})

	t.Run("Propagate event to JS", func(t *testing.T) {
		w := browsertest.InitWindow(t, e)

		w.MustRun(`
		var callCount = 0
		var event;
		function listener(e) { event = e; callCount++ };
		const target = window;
		target.addEventListener('custom', listener);
	`)
		w.DispatchEvent(event.NewCustomEvent("custom", event.CustomEventInit{}))
		assert.EqualValues(t, 1, w.MustEval("callCount"))
		assert.Equal(t, true, w.MustEval(`Object.getPrototypeOf(event) === CustomEvent.prototype`))
		assert.Equal(t, "custom", w.MustEval(`event.type`), "type of actual event")
	})

	t.Run("Add/Remove event listeners", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		win := browsertest.InitWindow(t, e)
		Expect(win.Eval(`
			let events = [];
			let noOfEvents = []
			const handler = e => { events.push(e) }
			window.addEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			window.removeEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			window.addEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			noOfEvents
		`)).To(HaveExactElements([]any{
			BeEquivalentTo(1),
			BeEquivalentTo(1),
			BeEquivalentTo(2),
		}))
	})

	t.Run("Event Capture", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		win := browsertest.InitWindow(t, e)
		Expect(win.Eval(`
			let events = [];
			let noOfEvents = []
			const div = document.createElement("div")
			document.body.appendChild(div)
			window.addEventListener("gost", e => { events.push("Window bubble. Phase: " + e.eventPhase)})
			window.addEventListener("gost", e => { events.push("Window capture. Phase: " + e.eventPhase)}, true)
			div.addEventListener("gost", e => { events.push("Div bubble. Phase: " + e.eventPhase)})
			div.addEventListener("gost", e => { events.push("Div capture. Phase: " + e.eventPhase)}, {capture:true})
			div.dispatchEvent(new CustomEvent("gost", { bubbles: true }))

			events
		`)).To(HaveExactElements(
			"Window capture. Phase: 1",
			"Div capture. Phase: 2",
			"Div bubble. Phase: 2",
			"Window bubble. Phase: 3",
		))
	})

	t.Run("Event option Once", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			let callCount = 0
			window.addEventListener("gost", e => { callCount++ }, { once: true })
			window.dispatchEvent(new CustomEvent("gost"))
			window.dispatchEvent(new CustomEvent("gost"))
			gost.assertEqual(callCount, 1)
		`)
	})

	t.Run("Error handler returns error", func(t *testing.T) {
		// Avoid a stack overflow, an error event handler throws an error, raising an
		// error event ...
		win := browsertest.InitWindow(t, e, browsertest.WithLogOption(gosttest.AllowErrors()))
		err := win.Run(`
			let callCount = 0
			window.addEventListener("error", () => {
				callCount++
				throw new Error()
			})
			window.addEventListener("custom", () => {
				throw new Error
			})
			window.dispatchEvent(new CustomEvent("custom"))
		`)
		assert.NoError(t, err)
		win.Clock().RunAll()
		win.MustRun(`gost.assertEqual(1, callCount)`)
	})

	t.Run("Event target is same object as Window", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			let target = null
			window.addEventListener('custom', e => { target = e.target })
			window.dispatchEvent(new Event('custom'))
			gost.assertNotNull(target)
			gost.assertEqual(target, window)
		`)
	})

	t.Run("Return false is treated as cancel", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			window.addEventListener("cancelled", () => { return false })
			window.addEventListener("normal", () => {})
		`)
		assert.True(t, win.DispatchEvent(&event.Event{Type: "normal", Cancelable: true}))
		assert.False(t, win.DispatchEvent(&event.Event{Type: "cancelled", Cancelable: true}))

	})
}
