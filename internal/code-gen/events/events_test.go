package events_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/events"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	gen "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/gost-dom/webref/events"
	"github.com/onsi/gomega"
)

func CreateMethodGenerator(specs EventGeneratorSpecs) (res gen.Generator, err error) {
	api, err := events.Load(specs.Api)
	x := ElementEventGenerator{
		Api:     api,
		Element: specs.SourceType,
	}
	for _, e := range x.Events() {
		if e.Type == specs.EventName {
			res = EventDispatchMethodGenerator{
				SourceTypeName: specs.SourceType,
				Event:          e,
				Type:           x.Type(),
			}
			break
		}
	}
	return
}

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
	data := e.defaultPointerEventInit()
	event := &event.Event{Type: "click", Data: data}
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
	x := ElementEventGenerator{
		Api:     api,
		Element: "Element",
	}
	res := EventInterfaceGenerator{
		Element: "Element",
		Events:  x.Events(),
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
