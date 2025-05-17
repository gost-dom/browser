package htmlelements_test

import (
	"testing"

	. "github.com/gost-dom/generators/testing/matchers"
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

func TestURLSearchParamsHaveSliceReturnValue(t *testing.T) {
	expect := newGomega(t)
	g, err := getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")
	expect(err).NotTo(gomega.HaveOccurred())
	expect(g).To(HaveRenderedSubstring("\tGetAll(string) []string\n"))
}

func TestURLSearchParamsReturnFoundOnNullableReturnValues(t *testing.T) {
	expect := newGomega(t)
	g, err := getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")
	expect(err).NotTo(gomega.HaveOccurred())
	expect(g).To(HaveRenderedSubstring("\tGet(string) (string, bool)\n"))
}

func TestURLSearchParamsOptionalArgs(t *testing.T) {
	expect := newGomega(t)
	g, err := getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")
	expect(err).NotTo(gomega.HaveOccurred())
	expect(g).To(HaveRenderedSubstring("\tHas(string) bool\n\tHasValue(string, string) bool\n"))
	expect(
		g,
	).To(HaveRenderedSubstring("\tDelete(string)\n\tDeleteValue(string, string)\n"))
}

func TestURLSearchParamsIterable(t *testing.T) {
	expect := newGomega(t)
	expect(getIdlInterfaceGenerator("urlinterfaces", "URLSearchParams")).To(HaveRenderedSubstring(
		"\tAll() iter.Seq2[string, string]\n"), "URLSearchParams is iterable")

}
