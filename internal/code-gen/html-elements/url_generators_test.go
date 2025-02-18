package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestURLSearchParamsShouldEmbedStringer(t *testing.T) {
	// The stringifier is an unnamed operation and shouldn't be included in the
	// go interface; but will generate an error if not handled in code
	expect := newGomega(t)
	g, err := getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")
	expect(err).NotTo(gomega.HaveOccurred())
	expect(g).To(HaveRenderedSubstring("interface {\n\tfmt.Stringer\n"))
}
