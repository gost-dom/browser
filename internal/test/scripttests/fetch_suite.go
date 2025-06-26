package scripttests

import (
	"context"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
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
	s.Expect(s.Eval(`typeof Response`)).To(Equal("function"), "Response is a constructor")
	s.Expect(s.Eval(`typeof Request`)).To(Equal("function"), "Request is a constructor")
}

type option struct {
	logOptions []gosttest.TestLoggerOption
}

type InitOption func(*option)

func WithLogOption(lo gosttest.TestLoggerOption) InitOption {
	return func(o *option) { o.logOptions = append(o.logOptions, lo) }
}

func WithMinLogLevel(lvl slog.Level) InitOption {
	return WithLogOption(gosttest.MinLogLevel(lvl))
}

func initWindow(
	t *testing.T,
	host html.ScriptHost,
	h http.Handler,
	opts ...InitOption,
) htmltest.WindowHelper {
	var o option
	for _, opt := range opts {
		opt(&o)
	}
	logger := gosttest.NewTestLogger(t, o.logOptions...)
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithLogger(logger),
		browser.WithHandler(h),
		browser.WithScriptHost(host),
	))
	return b.OpenWindow("https://example.com/index.html")
}

func testFetch(t *testing.T, host html.ScriptHost) {
	t.Parallel()

	t.Run("Fetch resource async/await", func(t *testing.T) { testFetchJSONAsync(t, host) })
	t.Run("Fetch invalid JSON", func(t *testing.T) { testFetchInvalidJSON(t, host) })
	t.Run("404 for not found resource", func(t *testing.T) { testNotFound(t, host) })
}

func testFetchJSONAsync(t *testing.T, host html.ScriptHost) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	delayedHandler := &gosttest.PipeHandler{T: t}
	handler := gosttest.StaticFileServer{
		"/index.html":     gosttest.StaticHTML(`<body>dummy</body>`),
		"/data.json":      gosttest.StaticJSON(`{"foo": "Foo value"}`),
		"/slow-data.json": delayedHandler,
	}

	g := gomega.NewWithT(t)
	win := initWindow(t, host, handler, WithMinLogLevel(slog.LevelDebug))
	win.MustRun(`
		let gotStatus
		let gotJson = "uninitialized"
		let err
		(async () => {
			try {
				const response = await fetch("slow-data.json")
				gotStatus = response.status
				globalThis.js = await response.json()
				gotJson = JSON.stringify(js)
			} catch (e) {
				err = e
			}
		})()
	`)
	g.Expect(win.Eval("gotStatus")).To(BeNil(), "status before fetch settles")
	delayedHandler.WriteHeader(200)

	assert.NoError(t, win.Clock().ProcessEventsWhile(ctx, func() bool {
		return win.MustEval("gotStatus") == nil
	}))

	g.Expect(win.Eval("gotStatus")).To(BeEquivalentTo(200), "status after fetch settles")
	delayedHandler.Print(`{"foo": "Foo value"}`)
	delayedHandler.Flush()
	g.Expect(win.Eval("gotJson")).To(Equal("uninitialized"), "json before response closes")
	delayedHandler.Close()
	assert.NoError(t, win.Clock().ProcessEventsWhile(ctx, func() bool {
		res, err := win.Eval("gotJson === 'uninitialized'")
		if err != nil {
			t.Error("Error evaluating gotJson")
			return false
		}
		return res.(bool)
	}))
	g.Expect(win.Eval("gotJson")).ToNot(Equal("uninitialized"), "json after response closes")
	g.Expect(win.Eval("err")).To(BeNil())
}

func testFetchInvalidJSON(t *testing.T, host html.ScriptHost) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	handler := gosttest.StaticFileServer{
		"/index.html":    gosttest.StaticHTML(`<body>dummy</body>`),
		"/bad-data.json": gosttest.StaticJSON(`{"foo": "Foo value",`),
	}
	g := gomega.NewWithT(t)
	win := initWindow(t, host, handler)

	win.MustRun(`
			let code = 0
			let resolved = null
			let rejected = null
			let resolvedAfterJson = false
			fetch("bad-data.json")
				.then(r => { code = r.status; return r })
				.then(r => { return r.json().then(
					r => { resolved = r },
					r => { rejected = r })
				})
		`)
	assert.NoError(t, win.Clock().ProcessEvents(ctx))
	g.Expect(win.Eval("code")).To(BeEquivalentTo(200))
	g.Expect(win.Eval("resolved")).To(BeNil(), "resolved")
	g.Expect(win.Eval("!!rejected")).To(BeTrue())
}

func testNotFound(t *testing.T, host html.ScriptHost) {
	handler := gosttest.StaticFileServer{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
	}
	g := gomega.NewWithT(t)
	win := initWindow(t, host, handler)
	win.MustRun(`
			fetch("non-existing.json")
				.then(response => { globalThis.got = response.status })
		`)
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()
	assert.NoError(t, win.Clock().ProcessEvents(ctx))
	g.Expect(win.Eval("got")).To(BeEquivalentTo(404))
}
