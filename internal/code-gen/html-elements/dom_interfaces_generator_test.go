package htmlelements_test

import (
	"testing"

	htmlelements "github.com/gost-dom/code-gen/html-elements"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	. "github.com/gost-dom/generators/testing/matchers"
)

func TestGenerateStructOfReadOnlyProperties(t *testing.T) {
	t.Parallel()

	expect := newGomega(t)
	g, err := getFileGenerator("dominterfaces", "mutation_record")
	expect(err).ToNot(HaveOccurred())
	expect(g).To(HaveRendered(`type MutationRecord struct {
	Type               string
	Target             dom.Node
	AddedNodes         dom.NodeList
	RemovedNodes       dom.NodeList
	PreviousSibling    dom.Node
	NextSibling        dom.Node
	AttributeName      *string
	AttributeNamespace *string
	OldValue           *string
}`))
}

func TestGenerationOfExplicitVariadicArgument(t *testing.T) {
	t.Parallel()

	expect := newGomega(t)

	g, err := htmlelements.GenerateInterface("dom", "dominterfaces", "MutationObserver")
	expect(err).ToNot(HaveOccurred())
	expect(g).To(HaveRenderedSubstring("type MutationObserver interface {\n"))
	expect(
		g,
	).To(HaveRenderedSubstring("\n\tObserve(dom.Node, ...ObserveOption) error\n"))
}

func TestGenerationOfEventTarget(t *testing.T) {
	t.Parallel()

	expect := newGomega(t)

	g, err := htmlelements.GenerateInterface("dom", "dominterfaces", "AbortSignal")
	expect(err).ToNot(HaveOccurred())
	expect(g).To(HaveRenderedSubstring("{\n\tevent.EventTarget\n"))
	expect(g).To(HaveRenderedSubstring("\n\tOnabort() event.EventHandler\n"))
}
