package v8host_test

import (
	"strings"
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestScriptHostDocumentScriptLoading(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
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

}
