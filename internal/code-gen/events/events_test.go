package events_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/events"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/gost-dom/webref/events"
	"github.com/onsi/gomega"
)

func TestGenerateElementEventMethod(t *testing.T) {
	g := gomega.NewWithT(t)
	r, err := CreateMethodGenerator(EventGeneratorSpecs{
		Api:        "uievents",
		SourceType: "Element",
		EventName:  "click",
	})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(r).To(HaveRendered(
		`func (e *elementEvents) Click() bool {
	data := PointerEventInit{}
	event := &event.Event{
		Data: data,
		Type: "click",
	}
	event.Bubbles = true
	event.Cancelable = true
	return e.target.DispatchEvent(event)
}`,
	))
}

func TestInterfaceGeneration(t *testing.T) {
	g := gomega.NewWithT(t)
	api, err := events.Load("uievents")
	g.Expect(err).ToNot(HaveOccurred())
	res := EventInterfaceGenerator{
		Element: "Element",
		Events:  api.EventsForType("Element"),
	}
	g.Expect(res).To(HaveRenderedSubstring("type ElementEvents interface {"))
	g.Expect(res).To(HaveRendered(gomega.MatchRegexp("(?m)^\tClick\\(\\) bool$")))
}

func TestGenerateEventDispatcher(t *testing.T) {
	g := gomega.NewWithT(t)
	res, err := CreateEventSourceGenerator("uievents", "Element")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res).To(HaveRendered(gomega.HavePrefix(`type elementEvents struct {
	target event.EventTarget
}`)))
	g.Expect(res).To(HaveRenderedSubstring("func (e *elementEvents) Click() bool {"))
	// g.Expect(res).To(HaveRenderedSubstring("type ElementEvents interface {"))
	// g.Expect(res).To(HaveRendered(gomega.MatchRegexp("(g:)^\tClick()$")))
}
