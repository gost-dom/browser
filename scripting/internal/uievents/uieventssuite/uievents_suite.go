package uieventssuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/onsi/gomega"
)

func RunUieventsSuite(t *testing.T, e html.ScriptEngine) {
	t.Run("Inheritance", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		w := browsertest.InitWindow(t, e)
		w.MustRun(
			"const getSuperclassName = (o) => Object.getPrototypeOf(o.prototype).constructor.name",
		)
		Expect(w.MustEval(`getSuperclassName(PointerEvent)`)).
			To(Equal("MouseEvent"), "Pointer event superclass")
		Expect(w.MustEval(`getSuperclassName(MouseEvent)`)).
			To(Equal("UIEvent"), "MouseEvent event superclass")
		Expect(w.MustEval(`getSuperclassName(UIEvent)`)).
			To(Equal("Event"), "UIEvent event superclass")
		Expect(w.MustEval(`getSuperclassName(KeyboardEvent)`)).
			To(Equal("UIEvent"), "KeyboardEvent event superclass")
	})

	t.Run("KeyboardEvent", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		w := browsertest.InitWindow(t, e, browsertest.WithHtml(`<body><div id="foo"></div></body>`))
		w.MustRun(`
			let event
			document.getElementById("foo").addEventListener("keydown", e => { event = e })
		`)
		uievents.KeydownInit(
			w.HTMLDocument().GetHTMLElementById("foo"),
			uievents.KeyboardEventInit{Key: "a"},
		)
		Expect(
			w.MustEval(`event instanceof KeyboardEvent`),
		).To(BeTrue(), "Event is a KeyboardEvent")
		Expect(w.MustEval(`event.key`)).To(Equal("a"), "Event has key: 'a'")
	})

	t.Run("Click event is PointerEvent", func(t *testing.T) {
		Expect := gomega.NewGomegaWithT(t).Expect
		w := browsertest.InitWindow(t, e, browsertest.WithHtml(`<body><div id="foo"></div></body>`))
		w.MustRun(`
			let event
			document.getElementById("foo").addEventListener("click", e => { event = e })
		`)
		w.HTMLDocument().GetHTMLElementById("foo").Click()
		Expect(w.MustEval(`event instanceof PointerEvent`)).To(BeTrue(), "Event is a PointerEvent")
	})

	t.Run("KeyboardEvent initialization", func(t *testing.T) {
		w := browsertest.InitWindow(t, e)
		w.MustRun(`
			const event = new KeyboardEvent("dummy", {
				key: "k"
			})
			gost.assertEqual(event.key, "k")
		`)
	})
}
