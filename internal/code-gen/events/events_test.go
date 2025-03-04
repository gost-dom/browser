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
	return e.target.DispatchEvent(
		NewPointerEvent("click", EventBubbles(true), EventCancelable(true)),
	)
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
	target EventTarget
}`)))
	g.Expect(res).To(HaveRenderedSubstring("func (e *elementEvents) Click() bool {"))
	// g.Expect(res).To(HaveRenderedSubstring("type ElementEvents interface {"))
	// g.Expect(res).To(HaveRendered(gomega.MatchRegexp("(g:)^\tClick()$")))
}
