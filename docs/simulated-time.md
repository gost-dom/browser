# Simulated Time

In a real browser, all JavaScript in the Window realm, and all DOM manipulation
runs in the "main thread". The browser _may_ run _some tasks_ in concurrent
threads, e.g., video rendering, downloading files.

Because the intended use case of Gost-DOM requires the client code (a test
case) to inspect and interact with the DOM, this places a hard constraint on the
design:

The "main thread" of the browser is the goroutine that the test runs in. As a
result, `thread.Sleep` or `<- time.After(...)` will neither have an effect on
event loop in the browser. Test code _actively must call_ the browser to tell it
to process events.

> [!NOTE]
> I use the term "main thread" as this is how browsers are described, and as
> such is implicitly a domain concept, though not one present in code. I use the
> term "goroutine" when describing how Gost-DOM operates in a Go. The
> "goroutine" in effect is the language-specific construct that models the "main
> thread" domain-level concept

## Clock and Simulated Time.

One aspect of Gost-DOM is to be able to test what happens when time passes;
without actually having to wait. E.g., verifying a behaviour throttled by e.g.,
200ms doesn't require the test to have that delay. You tell Gost-DOM to simulate
that time passes, allowing you to verify the absence of an effect before 200ms
has passed, and the effect after 200ms has passed. All running without any
delays in tests.

> [!WARNING]
> Actually querying `Date` for current time isn't supported yet. `Date` will
> return the system time. V8 has the ability to control the return value of
> `Date` - however, that part is not exposed in v8go - and would require some
> extensive changes. Goja/Sobek does expose the ability to control the value;
> but it was never implemented.

## Waiting for the System to Settle

A huge pain when maintaining browser automation tests for front-end heavy
application is being able to wait for the system to _settle_ and be ready for
the next step in the test. Browser automation tools typically return control to
the test when the page has loaded, which includes that initial JavaScript has
been executed. 

When control is returned to the test case, HTTP responses have not yet been
received and processed. Test code typically poll parts of the DOM to determine
if it's in the right state. This is typically very noisy code in test cases,
sometimes coupled to UI layout completely unrelated to the actual business logic
we want to verify. Depending on unrelated functionality also contribute to
fragile tests; i.e., test cases often need rework after refactorings despite the
app wasn't broken. And because tests are slow to run, this is discovered late.

It is even worse when you need to verify the _absence of an effect_. Your only
option is to wait before verifying, causing delays to test cases. If you reduce
the delay to improve test run-time, you risk false positives.

Gost-DOM completely solves this problem with no unnecessary delays. The function
`ProcessEvents` runs until nothing is pending, no `setTimeout` callbacks are
waiting, no HTTP requests are in-flight, etc. When the function returns; there
is nothing the browser is waiting for. There are variations of this function.
E.g., fast forward a certain amount of time; Fast forward while a specific
condition evaluates to true, etc.

> [!NOTE]
> There is one case that `ProcessEvents` and friends do not handle: Reacting to
> messages pushed by the server through SSE or WebSockets, as the browser itself
> has no way of expecting that such a message should arrive in the future.[^1] A
> solution to this problem is to use Go's `synctest` package, as it can reliably
> wait for all pending goroutines to have completed.

> [!NOTE]
> The name `ProcessEvents` is slightly misleading, as "event" is a slightly
> ambiguous term. Hopefully a better unambiguous name will be discovered.

## HTTP Requests

When using `fetch` HTTP requests are performed in a separate goroutine; however
the response is received, the JavaScript response object must be created, and
passed to the promise resolver in the "main thread", i.e., test goroutine.

With a simulated time advancing automatically; this leads to a potential error,
simulated time pass so fast, that request timeout handlers may be called before
we provide the request to the script.

To solve that, fetch responses promises are designed settle at a well-defined
point in simulated time. There is a default simulated delay, a value arbitrarily
set to 5ms (but is subject to change, so do not rely on this value).

Test code can explicitly control this value using functions in the
`browseroptions` package. The options provide two means of controlling this, a
simple version where the test supplies a single value; and one where the test
can provide a callback function that can be used to control this for each
individual request; allowing you to control the order responses are received by
client code.

[^1]: WebSocket are mentioned in the context of what the clock can potentially
    handle. But they aren't supported yet in Gost-DOM.
