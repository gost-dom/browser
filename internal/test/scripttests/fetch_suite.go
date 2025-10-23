package scripttests

import (
	"context"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
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
	logOptions []gosttest.HandlerOption
	ctx        context.Context
}

type InitOption func(*option)

func WithLogOption(lo gosttest.HandlerOption) InitOption {
	return func(o *option) { o.logOptions = append(o.logOptions, lo) }
}

func WithMinLogLevel(lvl slog.Level) InitOption {
	return WithLogOption(gosttest.MinLogLevel(lvl))
}

func WithContext(ctx context.Context) InitOption {
	return func(o *option) { o.ctx = ctx }
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
	ctx := o.ctx
	if ctx == nil {
		ctx = t.Context()
	}
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithContext(ctx),
		browser.WithLogger(logger),
		browser.WithHandler(h),
		browser.WithScriptHost(host),
	))
	return b.OpenWindow("https://example.com/index.html")
}

func testFetch(t *testing.T, host html.ScriptHost) {
	t.Parallel()

	t.Run(
		"Abort using AbortController and AbortSignal",
		func(t *testing.T) { testFetchAbortSignal(t, host) },
	)
	t.Run("Fetch resource async/await", func(t *testing.T) { testFetchJSONAsync(t, host) })
	t.Run("Fetch invalid JSON", func(t *testing.T) { testFetchInvalidJSON(t, host) })
	t.Run("404 for not found resource", func(t *testing.T) { testNotFound(t, host) })
	t.Run("ReadableStream body", func(t *testing.T) { testReadableStream(t, host) })
}

func testFetchAbortSignal(t *testing.T, host html.ScriptHost) {
	g := gomega.NewWithT(t)
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	delayedHandler := &gosttest.PipeHandler{T: t}
	handler := gosttest.HttpHandlerMap{
		"/index.html":     gosttest.StaticHTML(`<body>dummy</body>`),
		"/slow-data.json": delayedHandler,
	}
	win := initWindow(t, host, handler, WithMinLogLevel(slog.LevelDebug), WithContext(ctx))
	win.MustRun(`
		let resolved;
		let rejected;
		const ctrl = new AbortController()
		const signal = ctrl.signal
		fetch("/slow-data.json", { signal })
			.then(r => r.json())
			.then(r => { resolved = r }, r => { rejected = r })
		ctrl.abort("abort-reason")
	`)
	win.Clock().ProcessEvents(ctx)

	ctrl := win.MustEval("ctrl").(dominterfaces.AbortController)
	assert.NotNil(t, ctrl, "AbortController nil")
	g.Expect(win.Eval(`typeof signal`)).To(Equal("object"), "signal is an object")
	g.Expect(win.Eval(`rejected`)).To(Equal("abort-reason"))
}

func testFetchJSONAsync(t *testing.T, host html.ScriptHost) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	delayedHandler := &gosttest.PipeHandler{T: t}
	handler := gosttest.HttpHandlerMap{
		"/index.html":     gosttest.StaticHTML(`<body>dummy</body>`),
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

	handler := gosttest.HttpHandlerMap{
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
	handler := gosttest.HttpHandlerMap{
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

func testReadableStream(t *testing.T, host html.ScriptHost) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	pipe := gosttest.NewPipeHandler(t)
	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"/piped":      pipe,
	}
	g := gomega.NewWithT(t)
	win := initWindow(t, host, handler)
	win.MustRun(`
		let response;
		let rejected;
		let body;
		let reader;
		fetch("/piped")
			.then(r => { 
				response = r;
				body = response.body;
				reader = body.getReader()
			})
	`)
	pipe.WriteHeader(200)
	win.Clock().ProcessEvents(ctx)

	g.Expect(win.MustEval("typeof response")).To(Equal("object"), "Response is an object")
	assert.Equal(t, "ReadableStream", win.MustEval("Object.getPrototypeOf(body).constructor.name"),
		"body is a ReadableStream",
	)

	win.MustEval(`
		let readResult
		const dummy = reader.read().then(x => {
			readResult = new TextDecoder().decode(x.value)
		}, r => { rejected = r })
	`)
	pipe.Print("Hello, world!")
	win.Clock().ProcessEvents(ctx)

	assert.Equal(t, "Hello, world!", win.MustEval("readResult"))
}
