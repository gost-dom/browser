package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/v8host"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestV8EventTargetAddRemoveListeners(t *testing.T) {
	t.Parallel()
	host := v8host.New()
	t.Cleanup(func() { host.Close() })

	g := gomega.NewWithT(t)
	win := html.NewWindow(html.WindowOptionHost(host))
	t.Cleanup(win.Close)
	g.Expect(win.Eval(`
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
	`)).To(HaveExactElements([]int32{1, 1, 2}))
}

func TestV8EventCapture(t *testing.T) {
	g := gomega.NewWithT(t)
	win := html.NewWindow(html.WindowOptionHost(host))
	t.Cleanup(win.Close)
	g.Expect(win.Eval(`
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
	`)).To(gomega.HaveExactElements(
		"Window capture. Phase: 1",
		"Div capture. Phase: 2",
		"Div bubble. Phase: 2",
		"Window bubble. Phase: 3",
	))
}
