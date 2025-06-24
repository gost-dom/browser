package html_test

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestHTMLAnchorElement(t *testing.T) {
	win := html.NewWindow(html.WindowOptions{
		BaseLocation: "http://example.com/",
	})
	doc := win.Document()

	Expect := gomega.NewWithT(t).Expect

	{
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		Expect(a.Href()).To(Equal(""), "href IDL attribute when no href is set")
		Expect(a.Hostname()).To(Equal(""), "Hostname IDL attribute")
	}
	{
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetAttribute("href", "/foo/bar")
		Expect(a.Hostname()).To(Equal("example.com"))
		Expect(a.Hostname()).To(Equal(win.Location().Hostname()))
		Expect(
			a.Href()).To(Equal("http://example.com/foo/bar"),
			"href IDL attribute returns resolved URL")
	}
	{
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetPathname("/local")
		Expect(a.Pathname()).To(Equal(""), "pathname is NOT updated if there is no HREF attribute")

		a.SetHref("/")
		a.SetPathname("/local")
		Expect(
			a.Pathname(),
		).To(Equal("/local"), "pathname IS updated if there is no HREF attribute")
		Expect(a).To(HaveAttribute("href",
			"http://example.com/local"), "href is updated to reflect pathname")
	}
	{
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetHref("/")
		a.SetPathname("/local")
		Expect(a).To(HaveAttribute("href", "http://example.com/local"))
	}

	t.Run("SetHref() with full URL", func(t *testing.T) {
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetHref("http://other-site.example.com/local")
		Expect(a).To(HaveAttribute("href", "http://other-site.example.com/local"))

	})

	t.Run("String()", func(t *testing.T) {
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetHref("http://example.com/local")
		Expect(a.(fmt.Stringer).String()).To(Equal("http://example.com/local"))
	})

	t.Run("RelList", func(t *testing.T) {
		a := doc.CreateElement("a").(html.HTMLAnchorElement)
		a.SetRel("a b c")
		Expect(a).To(HaveAttribute("rel", "a b c"), "Attribute value")
		Expect(a.RelList().All()).To(ContainElements("a", "b", "c"), "Elements in RelList()")
	})
}
