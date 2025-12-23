package interfaces_test

import (
	"testing"

	"github.com/gost-dom/code-gen/interfaces"
	. "github.com/gost-dom/generators/testing/matchers"
	. "github.com/onsi/gomega"
)

func TestGenerationOfExplicitVariadicArgument(t *testing.T) {
	t.Parallel()

	gom := NewGomegaWithT(t)

	g, err := interfaces.GenerateInterface("dom", "dominterfaces", "MutationObserver")
	gom.Expect(err).ToNot(HaveOccurred())
	gom.Expect(g).To(HaveRenderedSubstring("type MutationObserver interface {\n"))
	gom.Expect(
		g,
	).To(HaveRenderedSubstring("\n\tObserve(dom.Node, ...ObserveOption) error\n"))
}

func TestGenerationOfEventTarget(t *testing.T) {
	t.Parallel()

	expect := NewGomegaWithT(t).Expect

	g, err := interfaces.GenerateInterface("dom", "dominterfaces", "AbortSignal")
	expect(err).ToNot(HaveOccurred())
	expect(g).To(HaveRenderedSubstring("{\n\tevent.EventTarget\n"))
	expect(g).To(HaveRenderedSubstring("\n\tOnabort() event.EventHandler\n"))
}
