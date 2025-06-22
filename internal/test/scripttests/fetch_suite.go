package scripttests

import (
	"context"
	"fmt"
	"log/slog"
	"maps"
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

	handler := gosttest.StaticFileServer{
		"/index.html":    gosttest.StaticHTML(`<body>dummy</body>`),
		"/data.json":     gosttest.StaticJSON(`{"foo": "Foo value"}`),
		"/bad-data.json": gosttest.StaticJSON(`{"foo": "Foo value",`),
	}

	t.Run("Fetch resource", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer cancel()

		g := gomega.NewWithT(t)
		h2 := maps.Clone(handler)
		fs := make(chan func(http.ResponseWriter))
		h2["/slow-data.json"] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for f := range fs {
				f(w)
			}
		})
		win := initWindow(t, host, h2, WithMinLogLevel(slog.LevelDebug))
		win.MustRun(`
			let gotStatus
			let gotJson = "uninitialized"
			let err
			(async () => {
				try {
					const response = await fetch("slow-data.json")
					gotStatus = response.status
					// window.dispatchEvent(new CustomEvent("gost-response"))
					globalThis.js = await response.json()
					gotJson = JSON.stringify(js)
					// window.dispatchEvent(new CustomEvent("gost-response"))
				} catch (e) {
					err = e
				}
			})()
		`)
		g.Expect(win.Eval("gotStatus")).To(BeNil())
		fs <- func(w http.ResponseWriter) {
			w.WriteHeader(200)
			fmt.Fprint(w, `{"foo": "Foo value"}`)
			w.(http.Flusher).Flush()
		}
		assert.NoError(t, win.Clock().ProcessEventsWhile(ctx, func() bool {
			res, err := win.Eval("gotStatus")
			if err != nil {
				t.Error(err)
			}
			return res == nil
		}))
		g.Expect(win.Eval("gotStatus")).To(BeEquivalentTo(200))
		g.Expect(win.Eval("gotJson")).To(Equal("uninitialized"), "json before response closes")
		close(fs)
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
	})

	t.Run("Fetch invalid JSON", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer cancel()

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
	})

	t.Run("404 for not found resource", func(t *testing.T) {
		g := gomega.NewWithT(t)
		win := initWindow(t, host, handler)
		win.MustRun(`
			fetch("non-existing.json")
				.then(response => { globalThis.got = response.status })
		`)
		ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer cancel()
		t.Log("test: Process events")
		assert.NoError(t, win.Clock().ProcessEvents(ctx))
		g.Expect(win.Eval("got")).To(BeEquivalentTo(404))
	})
}
