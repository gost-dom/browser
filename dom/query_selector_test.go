package dom_test

import (
	"testing"

	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

func TestDocumentQuerySelector(t *testing.T) {
	parse := func(t testing.TB, html string) htmltest.HTMLDocumentHelper {
		return htmltest.NewHTMLDocumentHelper(t, ParseHtmlString(html))
	}
	t.Run("Tag name", func(t *testing.T) {
		doc := parse(t, "<body><div>hello</div><p>world!</p><div>Selector</div></body>")
		Expect(t,
			doc.QuerySelectorHTML("div")).
			To(HaveOuterHTML("<div>hello</div>"))

		Expect(t,
			doc.QuerySelectorHTML("DIV")).
			To(HaveOuterHTML("<div>hello</div>"))
	})

	t.Run("attribute value", func(t *testing.T) {
		doc := parse(t,
			`<body><div>hello</div><p>world!</p><div data-foo="bar">Selector</div></body>`,
		)
		Expect(t,
			doc.QuerySelectorHTML("div[data-foo='bar']")).
			To(HaveOuterHTML(`<div data-foo="bar">Selector</div>`))
	})
}
