package scripttests

import (
	"fmt"
	"maps"
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom/event"
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
	t.Parallel()

	handler := gosttest.StaticFileServer{
		"/index.html":    gosttest.StaticHTML(`<body>dummy</body>`),
		"/data.json":     gosttest.StaticJSON(`{"foo": "Foo value"}`),
		"/bad-data.json": gosttest.StaticJSON(`{"foo": "Foo value",`),
	}

	t.Run("Fetch resource", func(t *testing.T) {
		g := gomega.NewWithT(t)
		h2 := maps.Clone(handler)
		fs := make(chan func(http.ResponseWriter))
		handlerDone := make(chan bool, 1)
		h2["/slow-data.json"] = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for f := range fs {
				f(w)
			}
			handlerDone <- true
		})
		win := initWindow(t, host, h2)
		done := make(chan bool)
		win.AddEventListener("gost-response", event.NewEventHandlerFunc(func(*event.Event) error {
			go func() { done <- true }()
			return nil
		}))
		win.MustRun(`
			let gotStatus
			let err
			(async () => {
				try {
					const response = await fetch("slow-data.json")
					gotStatus = response.status
					window.dispatchEvent(new CustomEvent("gost-response"))
					globalThis.js = await response.json()
					globalThis.gotJson = JSON.stringify(js)
					window.dispatchEvent(new CustomEvent("gost-response"))
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
		<-done
		g.Expect(win.Eval("gotStatus")).To(BeEquivalentTo(200))
		close(fs)
		<-done
		g.Expect(win.Eval("err")).To(BeNil())
	})

	t.Run("Fetch invalid JSON", func(t *testing.T) {
		g := gomega.NewWithT(t)
		win := initWindow(t, host, handler)

		done := make(chan bool)
		win.AddEventListener("gost-done", event.NewEventHandlerFunc(func(*event.Event) error {
			go func() { done <- true }()
			return nil
		}))
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
				.then(() => { window.dispatchEvent(new CustomEvent("gost-done")) })
		`)
		<-done
		g.Expect(win.Eval("code")).To(BeEquivalentTo(200))
		g.Expect(win.Eval("resolved")).To(BeNil(), "resolved")
		g.Expect(win.Eval("!!rejected")).To(BeTrue())
	})

	t.Run("404 for not found resource", func(t *testing.T) {
		g := gomega.NewWithT(t)
		win := initWindow(t, host, handler)
		done := make(chan bool)
		win.AddEventListener("gost-done", event.NewEventHandlerFunc(func(*event.Event) error {
			go func() { done <- true }()
			return nil
		}))
		win.MustRun(`
			fetch("non-existing.json")
				.then(response => { globalThis.got = response.status })
				.finally(() => { 
					window.dispatchEvent(new CustomEvent("gost-done")) 
				})
		`)
		win.Clock().RunAll()
		<-done
		g.Expect(win.Eval("got")).To(BeEquivalentTo(404))
	})
}
