package gojahost_test

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/gojahost"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("goja: Node", func() {
	Describe("Prototype", func() {
		It("Should have function installed", func() {
			engine := gojahost.New()
			window := html.NewWindow(html.WindowOptions{ScriptHost: engine})
			Expect(window.Run("const proto = Node.prototype")).To(Succeed())
			Expect(window.Eval("typeof proto.contains")).To(Equal("function"), "contains")
			Expect(window.Eval("typeof proto.insertBefore")).To(Equal("function"), "insertBefore")
			Expect(window.Eval("typeof proto.appendChild")).To(Equal("function"), "appendChild")
			Expect(window.Eval("typeof proto.removeChild")).To(Equal("function"), "removeChild")
		})
	})
})
