package gojahost_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/gojahost"
	. "github.com/gost-dom/browser/scripting/gojahost"
)

var _ = Describe("GojaDriver", func() {
	It("Starts with a test", func() {
		engine := New()
		window := html.NewWindow(html.WindowOptions{ScriptHost: engine})
		Expect(window.Eval("Window.name")).To(Equal("Window"), "Window.name")
		Expect(window.Eval("typeof globalThis")).To(Equal("object"))
		Expect(window.Eval("window === globalThis.window")).To(BeTrue())
		Expect(window.Eval("typeof Window")).To(Equal("function"))
		Expect(
			window.Eval("Object.getPrototypeOf(globalThis).constructor.name"),
		).To(Equal("Window"), "Global object constructor name")
		Expect(
			window.Eval(
				"Object.getPrototypeOf(globalThis).constructor === Window",
			),
		).To(BeTrue(), "Global object has Window as the constructor")

		Expect(window.Eval("window instanceof Window")).To(BeTrue(), "window instanceof Window")
		Expect(
			window.Eval("window instanceof EventTarget"),
		).To(BeTrue(), "window instanceof EventTarget")
	})
})

func TestGojaHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, gojahost.New())
}
