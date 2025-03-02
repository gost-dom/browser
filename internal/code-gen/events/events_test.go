package events_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/events"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	. "github.com/gost-dom/generators/testing/matchers"
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
