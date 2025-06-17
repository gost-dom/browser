package scripttests

import (
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

func testFetch(t *testing.T, host html.ScriptHost) {
	g := gomega.NewWithT(t)
	logger := gosttest.NewTestLogger(t)
	handler := gosttest.StaticFileServer{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"data.json":   gosttest.StaticJSON(`{"foo": "Foo value"}`),
	}
	b := browser.New(
		browser.WithLogger(logger),
		browser.WithHandler(handler),
		browser.WithScriptHost(host),
	)
	win := htmltest.NewBrowserHelper(t, b).OpenWindow("https://example.com/index.html")
	g.Expect(win.Eval(`typeof fetch`)).To(Equal("function"), "fetch is a function")
	win.MustRun(`const response = fetch("data.json")`)
	g.Expect(win.Eval(`typeof response.then`)).To(Equal("function"), "response is a promise")
}
