package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func testDocument(t *testing.T, e html.ScriptEngine) {
	t.Run("Global scope configuration", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)

		win.Assert().InstanceOf("document", "HTMLDocument")

		win.MustRun(`
			const docPrototype = Document.prototype
			const htmlDocPrototype = HTMLDocument.prototype
			gost.assertEqual(docPrototype.constructor.name, "Document")
			gost.assertEqual(htmlDocPrototype.constructor.name, "HTMLDocument")
			gost.assertTrue(Object.getOwnPropertyNames(docPrototype).includes("createElement"))
			gost.assertFalse(Object.getOwnPropertyNames(htmlDocPrototype).includes("createElement"))
			gost.assertInheritsFrom(HTMLDocument, Document)
			gost.assertEqual(Document.prototype.constructor.name, "Document")

			gost.assertUndefined(Document.createElement)

			gost.assertHasOwnProperty(document, "location")
		`)
	})

	t.Run("location", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			gost.assertInstanceOf(document.location, Location)
			gost.assertEqual(document.location, window.location)
		`)
	})

	t.Run("Document structure", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			gost.assertEqual(document.head.tagName, "HEAD")
			gost.assertEqual(document.body.tagName, "BODY")
			gost.assertInstanceOf(document.documentElement, HTMLHtmlElement)
		`)
	})

	t.Run("new Document()", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`
			const doc = new Document()
			gost.assertFalse(doc === document)

			gost.assertEqual(doc.nodeType, 9)
			gost.assertEqual(Document.prototype.constructor.name, "Document")
		`)
	})

	t.Run("createElement", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		win.MustRun(`gost.assertInstanceOf(document.createElement("base"), HTMLElement)`)
	})

	t.Run("getElementById", func(t *testing.T) {
		win := browsertest.InitWindow(t, e, browsertest.WithHtml(
			`<body><div id='elm-1'>Elm: 1</div><div id='elm-2'>Elm: 2</div></body>`,
		))
		g := gomega.NewGomegaWithT(t)
		g.Expect(win.Eval(`
		const e = document.getElementById("elm-2")
		e.outerHTML
		`)).To(Equal(`<div id="elm-2">Elm: 2</div>`))

		g.Expect(win.Eval(`Object.getPrototypeOf(e).constructor.name`)).To(Equal("HTMLDivElement"))
	})

	t.Run("querySelectorAll", func(t *testing.T) {
		win := browsertest.InitWindow(t, e, browsertest.WithHtml(
			`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div></body>`,
		))

		g := gomega.NewGomegaWithT(t)
		g.Expect(
			win.MustEval(
				"Array.from(document.querySelectorAll('[data-key]')).map(x => x.outerHTML).join(',')",
			),
		).To(Equal(`<div data-key="1">1</div>,<div data-key="2">2</div>`))
	})

	t.Run("querySelector", func(t *testing.T) {
		win := browsertest.InitWindow(t, e, browsertest.WithHtml(
			`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div><body>`,
		))
		g := gomega.NewGomegaWithT(t)
		g.Expect(
			win.MustEval("document.querySelector('[data-key]').outerHTML"),
		).To(Equal(`<div data-key="1">1</div>`))
		g.Expect(
			win.MustEval(`document.querySelector('[data-key="2"]').outerHTML`),
		).To(Equal(`<div data-key="2">2</div>`))
		g.Expect(
			win.MustEval(`document.querySelector('script')`),
		).To(BeNil())
	})
}
