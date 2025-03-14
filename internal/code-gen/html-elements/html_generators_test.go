package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"

	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
)

func GenerateHtmlAnchor() (g.Generator, error) {
	return getIdlInterfaceGenerator("html", "HTMLAnchorElement")
}

func TestHTMLAnchorInterface(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect

	Expect(GenerateHtmlAnchor()).To(HaveRendered(MatchRegexp(
		`type HTMLAnchorElement interface {\n\tHTMLElement`)),
		"Rendered HTMLAnchorElement interface declaration")

	Expect(GenerateHtmlAnchor()).To(HaveRendered(
		MatchRegexp("(?m)^\tHTMLHyperlinkElementUtils$")),
		"Rendered included interface HTMLHyperlinkElementUtils",
	)

	Expect(GenerateHtmlAnchor()).To(
		HaveRenderedSubstring("\n\tTarget() string\n"),
		"Rendered interface for URL properties")

	Expect(GenerateHtmlAnchor()).ToNot(HaveRenderedSubstring("\tSetOrigin("),
		"Rendered setters for read-only property")

	Expect(
		GenerateHtmlAnchor(),
	).ToNot(HaveRendered(MatchRegexp(`func \([^(]+\) Host\(\) string`)),
		"Rendered implementation for URL properties")
}
