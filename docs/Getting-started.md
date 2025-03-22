# Getting started with Gost-DOM

This document will try to explain the most essential concepts of Gost-DOM to get
started, as well as limitations you should be aware of.

## Creating a browser

The recommended way is to export the root `http.Handler`, and connecting to this
from Gost-DOM. This is generally the the handler passed to an `http.Server`
or `http.ListenAndServe`, but optionally without irrellevant middlewares like
logging/monitoring.

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
)

func TestWebBrowser(t *testing.T) {
    t.Parallel()
    b := browser.NewFromHandler(server.RootHttpHandler)
    window, err := b.Open("http://example.com/") // Host is ignored
    assert.NoError(t, err)
    // Interact with the document
    win.Document().GetElementById("test-button").Click()
    // Inspect the state of the document
    resultField := win.Document().GetElementById("output-element")
    assert.Equal(t, "The button was clicked", resultField.TextContent())
}
```

The returned `Window` represents an [HTML DOM API Window](https://developer.mozilla.org/en-US/docs/Web/API/Window), mening you can navigate the DOM just as you would in client-side script.

## Important notes

There are some crucial points you should be aware of:

- JavaScript served from CDN doesn't work (in the recommended usage)
- The host name is ignored, but not quite! (in the recommended usage)
- `setTimeout` and `setInterval` handlers require you to "advance the clock"

### JavaScript served from CDN doesn't work

When you using the recommended approach, to connect directly to an
`http.Handler`, _no outgoing_ HTTP requests will be made, meaning that CDN content will not be downloaded.

[Let us know](https://github.com/gost-dom/browser/issues/53) if you need this,
it's not difficult; so far the assumption has been that your server serves all
required assets.

### The host name is ignored, but not quite!

The call to open the your application in the Gost-DOM browser ignores the host,
so the following two lines will both open the "index path"

```go
b.Open("http://example.com/")
b.Open("/")
```

But cookies are associated with a host name, so the second version will not use
cookies.

Be aware of which origin is considered "secure context". `https` and `localhost`
are considered secure, but non-local `http` are not.

Some Web APIs, e.g. Location services, require a secure context.

This doesn't affect Gost-DOM at the moment, but it's advised to use a secure
origin, in order to not have to rewrite a lot of tests in the future.

### Time and timeouts

Gost-DOM's "event loop" runs in the test thread. That means that callbacks
registered by `setTimeout` and `setInterval` are not necessarily executed. They
are controlled by a "virtual clock", that test code controls

This means that a test of a behaviour that is throttled for e.g. 300ms, doesn't
actually need to wait for 300ms - you just tell the clock to advance 300ms.

When JavaScript code is executed, any callbacks that are registered for the
same time, i.e., microtasks, or `setTimeout` called with a callback of zero
milliseconds will execute before returning to Go code.

But other invocations of `setTimeout` or `setInterval` requires test code to
explicitly forward time to execute.

- `Window.Clock().Advance(time.Duration)` advances time for a certain amount of
  time, running all timeout and interval callbacks that should run in that
  period.
- `Window.Clock().RunAll()` will run until all `setTimeout` callbacks are
  called.

Both versions will panic if the number of registered callbacks does not
decrease. So `RunAll()` will currently panic if there are any `setInterval`
handlers that doesn't get cleared. Likewise, if `setInterval` is always called
with zero delay, it will too ([which is a missing
behaviour](https://github.com/gost-dom/browser/issues/45)

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

Each `Browser` instance is not currently safe to use from multiple threads, and
there is some overhead to creating an isolated V8 instance and setting up global
scope, so browsers them _may_ lead to better performance, but it is strongly
recommended to not do that, unless performance _is_ a problem. There are
optimisations planned, such as reusing V8 instances from discarded browsers.

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

- Most JavaScript API calls will log a debug statement.
- Some internal Go calls will log at the debug level.
- Some high level functions log at info level, e.g., `Window.Navigate`.
- `console` functions will log with the appropriate level.
- Errors, including unhandled JavaScript errors will generate error logs
