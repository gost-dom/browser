package htmlelements_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"

	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
)

func GenerateHtmlAnchor() (g.Generator, error) {
	return getIdlInterfaceGenerator("html", "HTMLAnchorElement")
}

func TestHTMLAnchor(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect

	Expect(GenerateHtmlAnchor()).To(HaveRendered(MatchRegexp(
		`type HTMLAnchorElement interface {\n\tHTMLElement`)),
		"Rendered HTMLAnchor interface declaration")

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

var _ = Describe("ElementGenerator", func() {
	It("Should generate a getter and setter", func() {
		Skip("Atrribute implementation currently disabled because of RelList()")
		Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring(
			`func (e *htmlAnchorElement) Target() string {`)))
	})

	It("Should NOT sanitize type", func() {
		Skip("Atrribute implementation currently disabled because of RelList()")
		Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring(
			`func (e *htmlAnchorElement) Type() string`)))
	})

	It("Should generate a struct with embedded htmlElement", func() {
		Skip("Disabled for HTMLAnchorElement, custom data in struct")
		Expect(GenerateHtmlAnchor()).To(HaveRendered(MatchRegexp(
			`type htmlAnchorElement struct {\n\tHTMLElement`)))
	})

	It("Should generate an interface ", func() {
	})

	It("Should generate a constructor", func() {
		Skip("Disabled for HTMLAnchorElement, custom construction")
		Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring(
			`func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{NewHTMLElement("a", ownerDoc)}
	result.SetSelf(result)
	return result
}`)))
	})

	Describe("HTMLAnchorElement", func() {
		It("Should have embedded HyperLinkUtils", func() {
			Expect(GenerateHtmlAnchor()).To(HaveRendered(
				MatchRegexp("(?m)^\tHTMLHyperlinkElementUtils$")))

		})
		It("Should have interface for URL properties", func() {
			Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring("\n\tTarget() string\n")))
		})

		It("Should not setters for read-only properties", func() {
			Expect(GenerateHtmlAnchor()).ToNot(HaveRendered(ContainSubstring("\tSetOrigin(")))
		})

		It("Should not have an implementation for URL properties", func() {
			Expect(
				GenerateHtmlAnchor(),
			).ToNot(HaveRendered(MatchRegexp(`func \([^(]+\) Host\(\) string`)))
		})
	})
})
