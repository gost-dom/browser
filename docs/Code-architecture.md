# Code architecture

It is a very clear design decision that the API of the library must match the
Web IDL specifications, for two reasons:

- This is the API that a web developer knows, so defining a new API would seem wrong.
- JavaScript needs to access this API.

The packages in general reflect the different [Web
APIs](https://developer.mozilla.org/en-US/docs/Web/API)

## DOM

The DOM represents the fundamental document structure of elements, attributes,
and text content. It also includes core event dispatch behaviour, including
bubbling, cancallation, and event capture.

Everything else is build on top of the DOM.

## HTML DOM

The HTML DOM adds HTML specific behaviour, such as form submission. It includes
the `Window` object, loading HTML documents from an HTTP server,

The HTML package also orchestrates execution of client-side script code provided
by an optional script engine.

## Clock

The code uses virtual time that will not proceed until client code tells to
clock to move forward. All JavaScript `setTimeout` or `setInterval` callbacks
will therefore not execute, until test code explicitly

## Scripting

JavaScript is executed by a "script engine". Two implementations exist

- `v8engine`, based on v8go. Uses the V8 engine powering the Chrome browser.
- `sobek`, based on [sobek][^1], a JavaScript engine written in Go. Less battle
  tested, but avoids some headaches of V8 linking C++ code.

[sobek]: https://github.com/grafana/sobek

## Browser

The `browser` package is a simple helper for creating a `html.Window` with the
a default script engine, wrapping an `http.Handler`.

## HTTP requests

HTTP requests, either because of navigating, or from XHR requests, are executed
using Go's `http.Client`. This provides an abstraction over the transport layer.

Gost provides an implementation that bypasses the TCP, allowing requests to
connect directly to the `http.Handler`. When using this, outgoing requests from
the browser are just pure method calls the the root HTTP handler, eliminating
the need to start a server, listening on a port.

When using the _default `http.Handler`_, all real HTTP requests over a TCP
port, requiring you to manually start a server. Note that this should
theoretically work, I've never used the system in this mode.

## Modularisation

Although the code isn't modularised yet, it is an idea that you should be able
to include the modules relevant to your app. E.g., if your app deals with
location services, you can add a module implementing location services.

This helps keep the size of the dependencies down for client projects; keeping
build times down for the TDD loop.

It also provides the option of alternate implementations. E.g., for location
services, the simple implementation can provide a single function to set the
current location / accuracy. The advanced implementation can replay a GPX track.

[^1]: Sobek is a fork of [Goja](https://github.com/dop251/goja) aiming to bring
    ESM suport, which Goja doesn't support.
