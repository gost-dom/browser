package events_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/events"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	"github.com/gost-dom/code-gen/packagenames"
	gen "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/gost-dom/webref/events"
	"github.com/onsi/gomega"
)

func CreateMethodGenerator(specs EventGeneratorSpecs) (res gen.Generator, err error) {
	api, err := events.Load(specs.Api)
	target := EventTargetType{specs.SourceType, packagenames.Dom}
	x := ElementEventGenerator{
		Api:        api,
		TargetType: target,
	}
	for _, e := range x.Events() {
		if e.Type == specs.EventName {
			res = EventDispatchMethodGenerator{
				TargetType: target,
				Event:      e,
				Type:       x.Type(),
			}
			break
		}
	}
	return
}

func TestClickEventDispatcher(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	r, err := CreateMethodGenerator(EventGeneratorSpecs{
		Api:        "uievents",
		SourceType: "Element",
		EventName:  "click",
	})
	Expect(err).ToNot(HaveOccurred())
	Expect(r).To(HaveRendered(
		`// Dispatches a [click event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [click event]: https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value
func Click(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}`,
	))
}
