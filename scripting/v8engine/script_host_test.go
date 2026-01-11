package v8engine

import (
	"strings"
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/onsi/gomega"
)

func TestScriptHostDocumentScriptLoading(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	reader := strings.NewReader(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
	host := New()
	t.Cleanup(host.Close)
	options := html.WindowOptions{ScriptHost: host, Clock: host.clock}
	win, err := html.NewWindowReader(reader, nil, options)
	Expect(err).ToNot(HaveOccurred())
	defer win.Close()
	ctx := win.ScriptContext()
	Expect(
		ctx.Eval("window.sut"),
	).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
}

func TestBasics(t *testing.T) {
	scripttests.RunBasicSuite(t, assertEngine)
}
