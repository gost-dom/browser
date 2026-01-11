package v8engine

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/onsi/gomega"
)

func TestScriptHostDocumentScriptLoading(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	win := browsertest.InitWindow(t, defaultEngine)
	err := win.LoadHTML(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
	Expect(err).ToNot(HaveOccurred())
	Expect(
		win.Eval("window.sut"),
	).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
}

func TestBasics(t *testing.T) {
	scripttests.RunBasicSuite(t, assertEngine)
}
