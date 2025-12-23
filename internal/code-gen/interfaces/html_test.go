package interfaces_test

import (
	"testing"

	"github.com/gost-dom/code-gen/interfaces"
	. "github.com/gost-dom/code-gen/internal/gomega-matchers"
	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
)

func GenerateHtmlAnchor() (g.Generator, error) {
	return interfaces.GenerateInterface("html", "html", "HTMLAnchorElement")
}
func GenerateLocation() (g.Generator, error) {
	return interfaces.GenerateInterface("html", "html", "Location")
}
func GenerateHtmlOrSvgElement() (g.Generator, error) {
	return interfaces.GenerateInterface("html", "html", "HTMLOrSVGElement")
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

func TestLocationIsEntity(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect

	Expect(GenerateLocation()).To(HaveRendered(MatchRegexp(
		`type Location interface {\n\tentity.Components`)),
	)
}

func TestGenerateTabindex(t *testing.T) {
	// This verifies that 'long' becomes an 'int'
	expect := exp(t)
	expect(GenerateHtmlOrSvgElement()).To(HaveRenderedSubstring(
		"TabIndex() int\n\tSetTabIndex(int)"),
		"Custom override of attribute type from long->int")
}

func TestGenerateGetterReturnsStruct(t *testing.T) {
	// Verify that the focusoptions are not generated
	expect := exp(t)
	expect(GenerateHtmlOrSvgElement()).To(HaveRenderedSubstring(
		"\tDataset() *DOMStringMap\n"), "HTMLOrSVGElement should be a pointer")
}

func TestGenerateNoFocusOptions(t *testing.T) {
	// Verify that the focusoptions are not generated
	expect := exp(t)
	expect(GenerateHtmlOrSvgElement()).To(HaveRenderedSubstring(
		"\tFocus()\n"), "Focus doesn't have options")
}

func TestGenerateEventHandlerFunction(t *testing.T) {
	expect := exp(t)
	expect(GenerateHtmlOrSvgElement()).To(HaveRenderedSubstring(
		"\tBlur()\n"), "Blur returns a bool")

}

func exp(t *testing.T) func(any, ...any) gomega.GomegaAssertion {
	return func(actual any, extras ...any) gomega.GomegaAssertion {
		return gomega.NewWithT(t).Expect(actual, extras...)
	}
}
