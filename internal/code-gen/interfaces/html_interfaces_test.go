package interfaces_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/code-gen/interfaces"
	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
)

func generateHtmlIntf(t testing.TB, name string) g.Generator {
	t.Helper()
	gen, err := interfaces.GenerateInterface("html", "htmlinterfaces", name)
	require.NoError(t, err)
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
	Expect(gen).ToNot(HaveRenderedSubstring("\tOrigin()"))
	Expect(gen).To(HaveRenderedSubstring("\tSetTimeout(TimerHandler, time.Duration"))
}
func TestMessagePortDoesntEmbedEventTarget(t *testing.T) {
	// The specialized EventTarget IDL interfaces contains on... attributes for
	// the events. We want to expose those to JavaScript, but not the
	// implementation interfaces.
	Expect := gomega.NewWithT(t).Expect
	gen := generateHtmlIntf(t, "MessagePort")

	Expect(gen).To(HaveRenderedSubstring("\tevent.EventTarget"))
	Expect(gen).To(HaveRenderedSubstring("\tentity.Components"))
	Expect(gen).ToNot(HaveRenderedSubstring("\tMessageEventTarget"))
}
