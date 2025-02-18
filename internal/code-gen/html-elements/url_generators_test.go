package htmlelements_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestStringifierShouldNotMakeInterfaceFail(t *testing.T) {
	// The stringifier is an unnamed operation and shouldn't be included in the
	// go interface; but will generate an error if not handled in code
	expect := newGomega(t)
	g, err := getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")
	expect(err).NotTo(gomega.HaveOccurred())
	expect(render(g)).Error().ToNot(HaveOccurred())
}
