package htmlelements_test

import (
	"testing"

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
