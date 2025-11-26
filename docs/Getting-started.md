# Getting started with Gost-DOM

This document will try to explain the most essential concepts of Gost-DOM to get
started, as well as limitations you should be aware of.

## Creating a browser, an example

The recommended way is to export the root `http.Handler`, and connecting to this
from Gost-DOM. The root handler is the handler you would pass to `http.Server`
or `http.ListenAndServe`, but optionally configured without irrellevant
middlewares like logging/monitoring.

```go
package server

import "net/http"

var RootHandler http.Handler

func init() {
    mux := http.NewServeMux()
    configureRoutes(mux)
    RootHandler = mux
}
```

The easiest way to start is to create `browser.Browser` passing the HTTP
handler. In this mode, _you don't need to start the server_. Gost will call your
`http.Handler` function directly, resulting in every call into your code being
in a stack trace originating from your tests.

```go
package server_test

import (
    "testing"
    "myapp/server"
    "github.com/gost-dom/browser"
    "github.com/gost-dom/html"
    "github.com/gost-dom/testing/gosttest"
)

func TestWebBrowser(t *testing.T) {
    t.Parallel() 
    b := browser.New(
        browser.WithScriptEngine(v8engine.DefaultEngine()),
        // Alternatively (adjust imports accordingly), using Sobek, a pure Go
        // JavaScript engine:
        // browser.WithScriptEngine(sobekengine.DefaultEngine()),
        browser.WithHandler(server.RootHttpHandler),
        browser.WithLogger(gosttest.NewTestingLogger(t)),
        browser.WithContext(t.Context()),
    )
    window, err := b.Open("http://example.com/") // Host is ignored
    assert.NoError(t, err)
    win.Document().GetElementById("test-button").(html.HTMLElement).Click()

    resultField := win.Document().GetElementById("output-element")
    assert.Equal(t, "The button was clicked", resultField.TextContent())
}
```

Breakdown of the code:

- `browser.New` obviously creates a new browser instance. By default, this will
  be configured with a V8 script engine.
  - `browser.WithScriptEngine` passes a script engine to use.
  - `browser.WithHandler` is the recommended way, connect the browser directly
    to the root HTTP handler, bypassing the TCP stack. This is not necessary.
    Without it, you need to start the server on a TCP port, as well as remember
    to close it afterwards.
  - `browser.WithLogger` accepts an `*slog.Logger` from the std [slog] package.
    `gosttest.NewTestingLogger(t)` returns a `*Logger` that writes log output to
    the `testing.TB` instance.
  - `browser.WithContext(t.Context())` will cause the browser to automatically
    close when the test context is cancelled. This will free resources, and
    allow the V8 instance to be reused for a new test, reducing the overhead of
    configuring V8.
- `b.Open("http://example.com")` opens the page, and returns an `html.Window`.
  The host name is ignored when having used `browser.WithHandler`. In fact,
  `Open("/")` will work, but cookies don't, as they are associated with an
  origin.
- The window implements the DOM API, adjusted for Go naming conventions (upper
  case letters. Attributes are accessed througn methods).
- `GetElementById` returns an `Element` from the DOM specification. But
  `Focus()` is a method on `HTMLElement` from the HTML DOM specification, so
  this requires a type assertion.


> [!NOTE]
>
> The `gosttest.NewTestLogger()` could be replaced by something like [slogt].
> But be sure to close the browser _before_ the test completes, as closing the
> browser can write log statements, and writing the test output _after_ the test
> completes will panic.

> [!WARNING]
>
> When using the `browser.WithHandler` option, **NO** outgoing HTTP requests
> will be performed. This effectively means that you **must** serve all content
> locally; JavaScript served from CDN will not work.

[HTML DOM API Window]: https://developer.mozilla.org/en-US/docs/Web/API/Window
[slog]: https://pkg.go.dev/log/slog
[slogt]: https://github.com/neilotoole/slogt

## Read next

- [Timeouts and the event loop](./event-loop.md) for information about how
  delayed callbacks and the "event loop" works.
- [Simulating user input](./simulating-user-input.md) to simulate user
  interactions.

## A mixed bag of information

There are some crucial points you should be aware of:

- JavaScript served from CDN doesn't work (in the recommended usage)
- The host name is ignored, but not quite! (in the recommended usage)
- `setTimeout` and `setInterval` handlers require you to "advance the clock"

But cookies are associated with a host name, so the second version will not use
cookies.

Be aware of which origin is considered "secure context". `https` and `localhost`
are considered secure, but non-local `http` are not.

Some Web APIs, e.g. Location services, require a secure context.

This doesn't affect Gost-DOM at the moment, but it's advised to use a secure
origin, in order to not have to rewrite a lot of tests in the future.

## The DOM API

Not all functions are implemented on the Window and the DOM objects, but they do
adhere to the [DOM
API](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model),
and [DOM HTML
API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API), but adopted
for Go:

- Functions start with upper-case letters to be exportable.
- [IDL Attribute](#idl-attributes-vs-data-attributes) getters become functions with the same name. E.g., `form.method`
     becomes `form.Method()`
- IDL Attribute setters becomes functions prefixed with `Set`. `form.method =
  "post"` becomes `form.SetMethod("post")`
- Functions thay may throw an error in JavaScript return an extra `error` value.
  E.g., `querySelector` throws an error if the pattern is invalid. Go's version
  has two return values `QuerySelector(pattern string) (Element, error)`.

So your existing knowledge about navigating and manipulating the DOM applies to
Gost as well.

## Test isolation and parallelism.

Each `Browser` instance has it's own isolated V8 instance, and each window
created from the browser has it's own context, i.e. it's own global scope,
allowing all tests to run in parallel.

## IDL Attributes vs. Content Attributes

There are two sets of attributes in the DOM, IDL Attributes and Content Attributes. 

Content Attributes are what most would consider attributes. They exist in the HTML,
like the `class` content attribute in `<div class="container"></div>`. The data
attributes are accessible in JavaScript using e.g.,
`getAttribute`/`setAttribute`.

IDL attributes are properties on objects in the DOM, i.e., the JavaScript
objects, and they normally reflect a corresponding content attribute. E.g., an
anchor element has an `href` attribute.

```html
<a id="my-link" href="https://example.com">Example</a>
```

```javascript
const a = document.getElementById("my-link")
console.log(a.getAttribute("href") // Logs https://example.com
console.log(a.href) // Logs https://example.com
```

Normally, setting an IDL attribute updates the corresponding content attribute. But
other behaviour can be attaced.

E.g., the link has IDL attributes for the different URL components, and
assigning a new value to them will update the href content attribute.

```javascript
a.patnhame = "/idl-attributes"
console.log(a.getAttribute("href")) // Logs https://example.com/idl-attributes
```

Some IDL attributes have different names, e.g., the IDL attribute that
corresponds to the `class` content attribute is not called `class` but `className`.

## Logging

By default, Gost does not write to stdout. You can inject a global logger
`*log/slog.Logger` calling `SetDefaultLogger` in the `browser/logger` package.

Each browser also supports you to inject a browser scoped logger.

```go
browser := browser.New(
    browser.WithLogger(logger), // *slog.Logger instance
    browser.WithHandler(rootHTTPHandler),
)
win, _ := browser.Open(url)
```

### Log levels and verbosity

There isn't a concrete logging strategy, except all error cases should be
logged.

- Most JavaScript API calls will mostly log a debug statement.
- Some internal Go calls will log at the debug level.
- Some high level functions log at info level, e.g., `Window.Navigate`.
- `console` functions will log with the appropriate level.
- Errors, including unhandled JavaScript errors will generate error logs

### Piping logs to `testing.T`

Gost-DOM is written with testing in mind, so piping log output to the
`testing.T` log can be helpful. Furthermore, error logs _typically_ represent
scenarios where your code is behaving unexpectedly, and you'd want the test to
fail - even if assertions are passing.

Gost-DOM log errors in these cases:

- A JavaScript error is unhandled.
- A network error occurs.
- JavaScript code calls functions that are not yet supported.

Whether you want error level logs to automatically fail the test or not, the
last part provides crucial information. The test is failing, not because of a
bug in your code, but you use a feature not yet implemented in Gost-DOM (it's
not you, it's us). The error message will include the URL where you can submit
an issue.

A simple log handler could look like this.

```go
type TestingLogHandler struct { testing.TB; allowErrors bool }

func (l TestingLogHandler) Enabled(_ context.Context, lvl slog.Level) bool { return lvl >= slog.LevelInfo }

func (l TestingLogHandler) Handle(_ context.Context, r slog.Record) error {
	h.TB.Context().Err() != nil {
        // Check if the context is cancelled to detect if the test has
        // completed to avoid calling t.Log/t.Error (which panics).
        // This can happen when constructing a browser using t.Context(). The
        // browser will dispose resources _after_ the context is closed, and log
        // statements written during cleanup would result in a panic
        return nil
    }
	l.TB.Helper()
	if r.Level < slog.LevelError || l.allowErrors {
		l.TB.Logf("%v: %s", r.Level, r.Message)
	} else {
		l.TB.Errorf("%v: %s", r.Level, r.Message)
	}
	return nil
}

func (l TestingLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return l }
func (l TestingLogHandler) WithGroup(name string) slog.Handler { return l }
```

3rd party modules also exist to provide similar behaviour.
