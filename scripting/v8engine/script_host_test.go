package v8engine

import (
	"context"
	"log/slog"
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/internal/js"
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

type dummyContext struct {
	*entity.Entity
	ctx context.Context
}

func (c dummyContext) Context() context.Context { return c.ctx }
func (c dummyContext) HTTPClient() http.Client  { return *http.DefaultClient }
func (c dummyContext) LocationHREF() string     { return "http://example.com" }
func (c dummyContext) Logger() *slog.Logger     { return nil }

func TestV8Engine(t *testing.T) {
	scripttests.RunScriptEngineSuites(t,
		func(c js.Configurer[jsTypeParam]) html.ScriptEngine { return newEngine(c) },
	)
}
