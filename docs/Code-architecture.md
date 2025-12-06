# Code architecture

It is a very clear design decision that the API of the library must match the
Web IDL specifications, for two reasons:

- This is the API that a web developer knows, so defining a new API would seem wrong.
- JavaScript needs to access this API.

The packages in general reflect the different [Web
APIs](https://developer.mozilla.org/en-US/docs/Web/API)

## DOM

The DOM is the core of the code base. This describes the tree of elements that
exists in a document.

This includes the core event dispatching mechanism, including bubbling, event
capture.

Everything else is build on top of the DOM.

## HTML

The DOM defines the basic `Element` type, but all the concrete interactible
elements exist in the `html` package.

This also includes the `Window` object itself.

The HTML package defines an interface the script host must implement.

This decouples the DOM and HTML implementation from the actual JavaScript
engine. It is possible to create a window without a script engine, in which case
script elements will not be executed.

## Clock

The code uses virtual time that will not proceed until client code tells to
clock to move forward. All JavaScript `setTimeout` or `setInterval` callbacks
will therefore not execute, until test code explicitly


## Scripting

Two independent script engines exist, but only one is complete.

- V8, based on v8go - this is the working script engine.
- Sobek is a pure Go JavaScript engine. This isn't yet working, but being worked
  on to get up to date.

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
