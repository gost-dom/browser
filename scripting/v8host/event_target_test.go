package v8host_test

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestV8EventTargetAddRemoveListeners(t *testing.T) {
	g := gomega.NewWithT(t)
	win := scriptTestSuite.NewWindow()
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
