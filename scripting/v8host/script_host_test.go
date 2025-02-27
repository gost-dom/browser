package v8host_test

import (
	"strings"

	"github.com/gost-dom/browser/html"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScriptHost", func() {
	Describe("Script host", func() {
		Describe("Load document with script", func() {
			It("Runs the script when connected to DOM", func() {
				reader := strings.NewReader(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
				options := html.WindowOptions{ScriptHost: host}
				win, err := html.NewWindowReader(reader, options)
				defer win.Close()
				Expect(err).ToNot(HaveOccurred())
				ctx := win.ScriptContext()
				Expect(
					ctx.Eval("window.sut"),
				).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
			})
		})
	})
})
