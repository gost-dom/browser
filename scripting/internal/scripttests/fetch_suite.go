package scripttests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/browseroptions"
	"github.com/gost-dom/browser/html"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testFetch(t *testing.T, e html.ScriptEngine) {
	t.Parallel()

	t.Run("Prototypes", func(t *testing.T) {
		w := browsertest.InitWindow(t, e)
		assert.Equal(t, "function", w.MustEval(`typeof fetch`), "fetch is a function")
		assert.Equal(t, "function", w.MustEval(`typeof Response`), "Response is a constructor")
		assert.Equal(t, "function", w.MustEval(`typeof Request`), "Request is a constructor")
	})

	t.Run(
		"Abort using AbortController and AbortSignal",
		func(t *testing.T) { testFetchAbortSignal(t, e) },
	)
	t.Run("Fetch resource async/await", func(t *testing.T) { testFetchJSONAsync(t, e) })
	t.Run("Fetch invalid JSON", func(t *testing.T) { testFetchInvalidJSON(t, e) })
	t.Run("404 for not found resource", func(t *testing.T) { testNotFound(t, e) })
	t.Run("ReadableStream body", func(t *testing.T) { testReadableStream(t, e) })
	t.Run("Headers", func(t *testing.T) { testHeaders(t, e) })
	t.Run("Response surface", func(t *testing.T) { testResponseSurface(t, e) })
	t.Run("Request", func(t *testing.T) { testRequest(t, e) })
	t.Run("Request-based Programmable delays", func(t *testing.T) { testProgrammableDelays(t, e) })
	t.Run("Simple Programmable delays", func(t *testing.T) { testSimpleProgrammableDelays(t, e) })
	t.Run("Test zero-delay responses", func(t *testing.T) { testZeroDelayResponses(t, e) })
}

func testFetchAbortSignal(t *testing.T, e html.ScriptEngine) {
	// This test has been seen failing on the build server.
	// Add some random logging to try to diagnose the issue
	t.Log("testFetchAbortSignal: start")
	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()

	g := gomega.NewWithT(t)
	t.Log("testFetchAbortSignal: child context")

	delayedHandler := &gosttest.PipeHandler{T: t}
	handler := gosttest.HttpHandlerMap{
		"/index.html":     gosttest.StaticHTML(`<body>dummy</body>`),
		"/slow-data.json": delayedHandler,
	}
	win := openWindow(t, e, handler, "https://example.com/index.html")
	t.Log("testFetchAbortSignal: window initialized")
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

	t.Log("testFetchAbortSignal: script run")
	win.Clock().ProcessEvents(ctx)
	t.Log("testFetchAbortSignal: events processed")

	t.Log("testFetchAbortSignal: eval ctrl")
	ctrl := win.MustEval("ctrl").(dominterfaces.AbortController)
	assert.NotNil(t, ctrl, "AbortController nil")
	t.Log("testFetchAbortSignal: eval typeof signal")
	g.Expect(win.Eval(`typeof signal`)).To(Equal("object"), "signal is an object")
	t.Log("testFetchAbortSignal: eval rejected")
	g.Expect(win.Eval(`rejected`)).To(Equal("abort-reason"))
}

func getRequestOptions(req *http.Request, o *browseroptions.FetchRoundtripOptions) {
	switch req.URL.Path {
	case "/data1.json":
		o.Delay = 2 * time.Millisecond
	case "/data2.json":
		o.Delay = 10 * time.Millisecond
	}
}

func testProgrammableDelays(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), defaultTimeout)
	defer cancel()

	// This is bad test practice, the test sets a global value; And to make
	// matters worse, it's in a parallel test. However, this should be the
	// _only_ test in the entire codebase depending on the global default value,
	// so ... unless the bad pattern is accidentally duplicated, we're good, and
	// we get to test how default values affect the test.
	browseroptions.SetDefaultFetchDelay(5 * time.Millisecond)

	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"/data1.json": gosttest.StaticJSON(`{"foo": "bar"}`),
		"/data2.json": gosttest.StaticJSON(`{"foo": "bar"}`),
		"/data3.json": gosttest.StaticJSON(`{"foo": "bar"}`),
	}
	b := browser.New(
		browser.WithScriptEngine(e),
		browser.WithHandler(handler),
		browseroptions.FetchRequestOptions(getRequestOptions),
	)
	win, err := b.Open("https://example.com/index.html")
	require.NoError(t, err)

	require.NoError(t, win.Run(`
		let msgs = [];
		setTimeout(() => { msgs.push("after 1ms") }, 1);
		setTimeout(() => { msgs.push("after 3ms") }, 3);
		setTimeout(() => { msgs.push("after 9ms") }, 9);
		setTimeout(() => { msgs.push("after 11ms") }, 11);
		(async () => {
			const response = await fetch("data1.json")
			msgs.push("after response 1: " + response.status)
		})();
		(async () => {
			const response = await fetch("data2.json")
			msgs.push("after response 2: " + response.status)
		})();
		(async () => {
			const response = await fetch("data3.json")
			msgs.push("after response 3: " + response.status)
		})();
	`))

	win.Clock().ProcessEvents(ctx)

	msgs, err := win.Eval("msgs")
	assert.NoError(t, err)
	assert.Equal(t, []any{
		"after 1ms",
		"after response 1: 200",
		"after 3ms",
		"after response 3: 200",
		"after 9ms",
		"after response 2: 200",
		"after 11ms",
	}, msgs)
}

func testZeroDelayResponses(t *testing.T, e html.ScriptEngine) {
	// HTTP responses with zero delay has the capacity to cause deadlock issues.
	//
	// After running scripts, the _current default_ behaviour is to run all
	// microtasks and macrotasks scheduled for the current time, i.e., all
	// setTimeout(..., 0) callbacks.
	//
	// When a fetch response is scheduled to be processed at this time, but we
	// haven't had the ability to send the response yet, we'd wait indefinitely
	// for the response to arrive.
	//
	// The WithAsync() option tells the clock, that it can return from Advance()
	// calls, even if this task was scheduled for the current simulated time.
	// Removing WithAsync() in the fetch implementation will cause this test to block;
	//
	// This tests two scenarios
	//
	//   - The response is not coupled to the test case itself. Tested code calls
	//     HTTP handlers that themselves are capable of generating responses.
	//   - The response is generated in the test code after JavaScript code has sent
	//     the request, i.e., there is no way for the response to be generated
	//     before exiting JavaScript execution scope.
	t.Run("When response is generated before script returns", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), defaultTimeout)
		defer cancel()

		handler := gosttest.HttpHandlerMap{
			"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
			"/data.json":  gosttest.StaticJSON(`{"foo": "bar"}`),
		}
		b := browser.New(
			browser.WithScriptEngine(e),
			browser.WithHandler(handler),
			browseroptions.FetchDelay(0),
		)
		win, err := b.Open("https://example.com/index.html")
		require.NoError(t, err)

		require.NoError(t, win.Run(`
			let msgs = [];
			setTimeout(() => { msgs.push("after 1ms") }, 1);
			(async () => {
				const response = await fetch("data.json")
				msgs.push("after response: " + response.status)
			})();
		`))

		win.Clock().ProcessEvents(ctx)

		msgs, err := win.Eval("msgs")
		assert.NoError(t, err)
		assert.Equal(t, []any{
			"after response: 200",
			"after 1ms",
		}, msgs)
	})

	t.Run("When response is generated _after_ the script returns", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(t.Context(), defaultTimeout)
		defer cancel()

		delayedHandler := &gosttest.PipeHandler{T: t}

		handler := gosttest.HttpHandlerMap{
			"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
			"/data.json":  delayedHandler,
		}
		b := browser.New(
			browser.WithScriptEngine(e),
			browser.WithHandler(handler),
			browseroptions.FetchDelay(0),
		)
		win, err := b.Open("https://example.com/index.html")
		require.NoError(t, err)

		require.NoError(t, win.Run(`
			let msgs = [];
			setTimeout(() => { msgs.push("after 1ms") }, 1);
			(async () => {
				const response = await fetch("data.json")
				msgs.push("after response: " + response.status)
			})();
		`))

		msgs, err := win.Eval("msgs")
		require.NoError(t, err)
		assert.Empty(t, msgs)

		delayedHandler.WriteHeader(200)
		delayedHandler.Close()
		win.Clock().ProcessEvents(ctx)

		msgs, err = win.Eval("msgs")
		assert.NoError(t, err)
		assert.Equal(t, []any{
			"after response: 200",
			"after 1ms",
		}, msgs)

	})
}

func testSimpleProgrammableDelays(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), defaultTimeout)
	defer cancel()

	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"/data.json":  gosttest.StaticJSON(`{"foo": "bar"}`),
	}
	b := browser.New(
		browser.WithScriptEngine(e),
		browser.WithHandler(handler),
		browseroptions.FetchDelay(2*time.Millisecond),
	)
	win, err := b.Open("https://example.com/index.html")
	require.NoError(t, err)

	require.NoError(t, win.Run(`
		let msgs = [];
		setTimeout(() => { msgs.push("after 1ms") }, 1);
		setTimeout(() => { msgs.push("after 3ms") }, 3);
		(async () => {
			const response = await fetch("data.json")
			msgs.push("after response: " + response.status)
		})();
	`))

	win.Clock().ProcessEvents(ctx)

	msgs, err := win.Eval("msgs")
	assert.NoError(t, err)
	assert.Equal(t, []any{
		"after 1ms",
		"after response: 200",
		"after 3ms",
	}, msgs)
}

func testFetchJSONAsync(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	delayedHandler := &gosttest.PipeHandler{T: t}
	handler := gosttest.HttpHandlerMap{
		"/index.html":     gosttest.StaticHTML(`<body>dummy</body>`),
		"/slow-data.json": delayedHandler,
	}

	g := gomega.NewWithT(t)
	win := openWindow(t, e, handler, "https://example.com/index.html", WithContext(ctx))
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

func testFetchInvalidJSON(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	handler := gosttest.HttpHandlerMap{
		"/index.html":    gosttest.StaticHTML(`<body>dummy</body>`),
		"/bad-data.json": gosttest.StaticJSON(`{"foo": "Foo value",`),
	}
	g := gomega.NewWithT(t)
	win := openWindow(t, e, handler, "https://example.com/index.html")

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

func testNotFound(t *testing.T, e html.ScriptEngine) {
	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
	}
	g := gomega.NewWithT(t)
	win := openWindow(t, e, handler, "https://example.com/index.html")
	win.MustRun(`
			fetch("non-existing.json")
				.then(response => { globalThis.got = response.status })
		`)
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()
	assert.NoError(t, win.Clock().ProcessEvents(ctx))
	g.Expect(win.Eval("got")).To(BeEquivalentTo(404))
}

func testReadableStream(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	pipe := gosttest.NewPipeHandler(t)
	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"/piped":      pipe,
	}
	g := gomega.NewWithT(t)
	win := openWindow(t, e, handler, "https://example.com/index.html")
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

func testRequest(t *testing.T, e html.ScriptEngine) {
	t.Run("Uses document location", func(t *testing.T) {
		b := browsertest.InitBrowser(t, nil, e)
		w := b.OpenWindow("https://example.com/pages/page-1")
		assert.Equal(t,
			"https://example.com/pages/page-2",
			w.MustEval(`
				const req = new Request("page-2")
				req.url
		`))
	})
	t.Run("Headers return SameObject", func(t *testing.T) {
		win := initWindow(t, e, nil)
		res := win.MustEval(`
			const h = new Request("http://example.com")
			const a = h.headers
			const b = h.headers
			a === b
	  `)
		assert.True(t, res.(bool))
	})

	t.Run("Request parse headers", func(t *testing.T) {
		win := initWindow(t, e, nil)
		res := win.MustEval(`
			const h = new Request("http://example.com", { headers: {"foo":"bar"} })
			h.headers.get("foo")
		`)
		if assert.NotNil(t, res) {
			assert.Equal(t, "bar", res.(string))
		}
	})
}

// testResponseSurface exercises the Response.ok/status/statusText getters and
// the Body.text/json/bytes/arrayBuffer consumers against both a 200 and a 404
// response.
func testResponseSurface(t *testing.T, e html.ScriptEngine) {
	ctx, cancel := context.WithTimeout(t.Context(), defaultTimeout)
	defer cancel()

	g := gomega.NewWithT(t)
	handler := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`<body>dummy</body>`),
		"/data.json":  gosttest.StaticJSON(`{"foo":"bar"}`),
	}
	win := openWindow(t, e, handler, "https://example.com/index.html")
	win.MustRun(`
		globalThis.res = {};
		(async () => {
			const ok = await fetch("/data.json");
			res.ok = ok.ok;
			res.status = ok.status;
			res.statusText = ok.statusText;
			res.text = await (await fetch("/data.json")).text();
			res.json = JSON.stringify(await (await fetch("/data.json")).json());
			const buf = await (await fetch("/data.json")).arrayBuffer();
			res.arrayBufferIsArrayBuffer = buf instanceof ArrayBuffer;
			res.arrayBuffer = buf.byteLength;
			const bytes = await (await fetch("/data.json")).bytes();
			res.bytesIsUint8Array = bytes instanceof Uint8Array;
			res.bytes = bytes.length;
			const notFound = await fetch("/does-not-exist");
			res.notOk = notFound.ok;
			res.notStatus = notFound.status;
			res.notStatusText = notFound.statusText;
		})();
	`)
	assert.NoError(t, win.Clock().ProcessEvents(ctx))

	g.Expect(win.Eval("res.ok")).To(BeTrue(), "Response.ok for 200")
	g.Expect(win.Eval("res.status")).To(BeEquivalentTo(200), "Response.status")
	g.Expect(win.Eval("res.statusText")).To(Equal("OK"), "Response.statusText")
	g.Expect(win.Eval("res.text")).To(Equal(`{"foo":"bar"}`), "Body.text()")
	g.Expect(win.Eval("res.json")).To(Equal(`{"foo":"bar"}`), "Body.json()")
	g.Expect(win.Eval("res.arrayBufferIsArrayBuffer")).To(BeTrue(), "Body.arrayBuffer() returns an ArrayBuffer")
	g.Expect(win.Eval("res.arrayBuffer")).To(BeEquivalentTo(13), "Body.arrayBuffer() byteLength")
	g.Expect(win.Eval("res.bytesIsUint8Array")).To(BeTrue(), "Body.bytes() returns a Uint8Array")
	g.Expect(win.Eval("res.bytes")).To(BeEquivalentTo(13), "Body.bytes() length")
	g.Expect(win.Eval("res.notOk")).To(BeFalse(), "Response.ok for 404")
	g.Expect(win.Eval("res.notStatus")).To(BeEquivalentTo(404), "Response.status for 404")
	g.Expect(win.Eval("res.notStatusText")).To(Equal("Not Found"), "Response.statusText for 404")
}

func testHeaders(t *testing.T, e html.ScriptEngine) {
	t.Run("Throws on invalid value", func(t *testing.T) {
		win := initWindow(t, e, nil)
		res := win.MustEval(`
			var err
			const h = new Headers()
			try { h.append("\uFFFF", "value") } catch(e) { err = e }
			err instanceof TypeError
		`)
		assert.True(t, res.(bool), "TypeError is thrown on append")

		res = win.MustEval(`
			err = null
			try { h.set("\uFFFF", "value") } catch(e) { err = e }
			err instanceof TypeError
		`)
		assert.True(t, res.(bool), "TypeError is thrown on set")
	})

	t.Run("Are iterable", func(t *testing.T) {
		win := initWindow(t, e, nil)
		if !(assert.Equal(
			t, "true",
			win.MustEval("Headers.prototype.hasOwnProperty(Symbol.iterator).toString()"),
			"Headers are iterable") &&
			assert.Equal(t, "function", win.MustEval("typeof (new Headers().keys)"))) {
			return
		}

		win.MustRun(`
			const h = new Headers(new Proxy({ foo: "foo-value", bar: "bar-value" },{}))
			var keys = []
			var values = []
			for(const [key,val] of h.entries()) {
				keys.push(key)
				values.push(val)
			}
			keys.sort()
			values.sort()
		`)

		assert.Equal(t, "bar,foo", win.MustEval("keys.join(',')"))
		assert.Equal(t, "bar-value,foo-value", win.MustEval("values.join(',')"))
	})

	t.Run("Construct with null is a TypeError", func(t *testing.T) {
		b := browser.New(browser.WithScriptEngine(e))
		win := htmltest.NewWindowHelper(t, b.NewWindow())
		win.MustRun(`
			var err
			try { new Headers(null) } catch(e) { err = e }
		`)
		assert.True(t, win.MustEval("err instanceof TypeError").(bool))
		assert.Equal(t, "TypeError", win.MustEval("Object.getPrototypeOf(err).constructor.name"))
	})

	t.Run("Construct with invalid key", func(t *testing.T) {
		t.Run("Key is outside ASCII range", func(t *testing.T) {
			win := initWindow(t, e, nil)
			res := win.MustEval(`
				var err
				try { new Headers({ "\uFFFF": "foo" }) } catch(e) { err = e }
				err instanceof TypeError
		`)
			assert.True(t, res.(bool))
		})

		t.Run("Key is a symbol", func(t *testing.T) {
			win := initWindow(t, e, nil)
			res := win.MustEval(`
				var err
				const sym = Symbol()
				try { new Headers({ [sym]: "foo" }) } catch(e) { err = e }
				err instanceof TypeError
		`)
			assert.True(t, res.(bool))
		})
	})
}
