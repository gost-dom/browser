package interfaces_test

import (
	"testing"

	"github.com/gost-dom/code-gen/interfaces"
	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func generateHtmlIntf(t testing.TB, name string) g.Generator {
	t.Helper()
	gen, err := interfaces.GenerateInterface("html", "htmlinterfaces", name)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	return gen
}

func TestHTMLHistoryInterface(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	gen := generateHtmlIntf(t, "History")

	Expect(gen).To(HaveRenderedSubstring("\tLength() int\n"))
	Expect(gen).ToNot(HaveRenderedSubstring("ScrollRestoration()"))

	Expect(
		gen,
	).To(HaveRenderedSubstring("\tState() HistoryState\n"), "History state 'overrides' default type")
}

func TestHTMLInterfaceWindowOrWorkerGlobalScope(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	gen := generateHtmlIntf(t, "WindowOrWorkerGlobalScope")

	Expect(gen).ToNot(HaveRenderedSubstring("CreateImageBitmap"))
	Expect(gen).ToNot(HaveRenderedSubstring("\tOrigin("))
}
