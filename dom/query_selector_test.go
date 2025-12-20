package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

func TestDocumentQuerySelector(t *testing.T) {
	win := browsertest.InitWindow(t, nil)
	t.Run("Tag name", func(t *testing.T) {
		win.LoadHTML("<body><div>hello</div><p>world!</p><div>Selector</div></body>")
		Expect(t,
			win.HTMLDocument().QuerySelectorHTML("div")).
			To(HaveOuterHTML("<div>hello</div>"))

		Expect(t,
			win.HTMLDocument().QuerySelectorHTML("DIV")).
			To(HaveOuterHTML("<div>hello</div>"))
	})

	t.Run("attribute value", func(t *testing.T) {
		win.LoadHTML(
			`<body><div>hello</div><p>world!</p><div data-foo="bar">Selector</div></body>`,
		)
		Expect(t,
			win.HTMLDocument().QuerySelectorHTML("div[data-foo='bar']")).
			To(HaveOuterHTML(`<div data-foo="bar">Selector</div>`))
	})
}
