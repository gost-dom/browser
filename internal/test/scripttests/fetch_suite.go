package scripttests

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
)

type FetchSuite struct {
	ScriptHostSuite
}

func NewFetchSuite(h html.ScriptHost) *FetchSuite {
	return &FetchSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *FetchSuite) TestRequestURL() {
	s.NewWindowLocation("https://example.com/pages/page-1")
	s.Expect(s.Eval(`
		const req = new Request("page-2")
		req.url
	`)).To(Equal("https://example.com/pages/page-2"))
}

func (s *FetchSuite) TestPrototypes() {
	s.Expect(s.Eval(`typeof fetch`)).To(Equal("function"), "fetch is a function")
}

func initWindow(t *testing.T, host html.ScriptHost, h http.Handler) htmltest.WindowHelper {
	logger := gosttest.NewTestLogger(t)
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithLogger(logger),
		browser.WithHandler(h),
		browser.WithScriptHost(host),
	))
	return b.OpenWindow("https://example.com/index.html")
}

func testFetch(t *testing.T, host html.ScriptHost) {
	handler := gosttest.StaticFileServer{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"data.json":   gosttest.StaticJSON(`{"foo": "Foo value"}`),
	}
	t.Run("Fetch resource", func(t *testing.T) {
		g := gomega.NewWithT(t)
		win := initWindow(t, host, handler)
		win.MustRun(`
			let resolved;
			let rejected;
			const response = fetch("data.json")
			response.then(r => { resolved = r }, r => { rejected = r });
		`)
		g.Expect(
			win.Eval("resolved && Object.getPrototypeOf(resolved) === Response.prototype"),
		).To(BeTrue())
	})

	t.Run("404 for not found resource", func(t *testing.T) {
		g := gomega.NewWithT(t)
		win := initWindow(t, host, handler)
		win.MustRun(`
			(async () => {
				const response = await fetch("non-existing.json")
				globalThis.got = response.status
			})()
		`)
		g.Expect(win.Eval("got")).To(BeEquivalentTo(404))
	})
}
