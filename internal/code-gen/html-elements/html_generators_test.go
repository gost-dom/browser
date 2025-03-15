package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"

	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
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

func exp(t *testing.T) func(any, ...any) gomega.GomegaAssertion {
	return func(actual interface{}, extras ...interface{}) gomega.GomegaAssertion {
		return gomega.NewWithT(t).Expect(actual, extras...)
	}
}

func TestGenerateTabindex(t *testing.T) {
	// This verifies that 'long' becomes an 'int'
	expect := exp(t)
	expect(getIdlInterfaceGenerator("html", "HTMLOrSVGElement")).To(HaveRenderedSubstring(
		"TabIndex() int\n\tSetTabIndex(int)"),
		"Custom override of attribute type from long->int")
}

func TestGenerateNoFocusOptions(t *testing.T) {
	// Verify that the focusoptions are not generated
	expect := exp(t)
	expect(getIdlInterfaceGenerator("html", "HTMLOrSVGElement")).To(HaveRenderedSubstring(
		"\tFocus() bool\n"), "Focus doesn't have options")
}

func TestGenerateEventHandlerFunction(t *testing.T) {
	expect := exp(t)
	expect(getIdlInterfaceGenerator("html", "HTMLOrSVGElement")).To(HaveRenderedSubstring(
		"\tBlur() bool\n"), "Blur returns a bool")

}
