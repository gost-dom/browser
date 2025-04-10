package v8host_test

import (
	"github.com/gost-dom/browser/html"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 Append (element, document, document fragment)", func() {
	It("Should accept multiple values", func() {
		win := html.NewWindow(html.WindowOptions{ScriptHost: host})
		Expect(win.Eval(`
			const d = document.createElement("div")
			d.append(
				document.createElement("p"),
				document.createElement("p"),
			);
			d.outerHTML`)).To(Equal("<div><p></p><p></p></div>"))
	})

	It("Should accept text values", func() {
		win := html.NewWindow(html.WindowOptions{ScriptHost: host})
		Expect(win.Eval(`
			const d = document.createElement("div")
			d.append(
				document.createElement("p"),
				"Foo",
				"bar",
				document.createElement("p"),
			);
			d.outerHTML`)).To(Equal("<div><p></p>Foobar<p></p></div>"))
	})
})
