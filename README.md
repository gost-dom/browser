# Gost-DOM - A headless browser for Go

**The Go-to headless browser for TDD workflows.**

Gost-DOM is a headless browser written in Go intended to write tests of web
application in Go that relies on JavaScript. Properties of Gost-DOM-based tests:

- Tests run in parallel due to complete _complete isolation_[^2]
- No erratic behaviour due to 100% predictable UI reactions.
- "Blazingly fast". No out-of-process calls, not even thread boundaries. Web
  application code runs in the test thread, so a panic in your code keeps a full
  stack trace to the test case. [^3]
- Dependencies can be replaced while testing.

Yet Gost-DOM still uses HTTP request and responses for verification, testing the
entire stack, as well as middlewares. 

[Read this to get started](./docs/Getting-started.md), and see a quick example
of usage.

> [!NOTE]
>
> This is 0.x version still, and breaking API changes do occur, but will be
> announced before release in the [Gost-DOM
> discussions](https://github.com/orgs/gost-dom/discussions/categories/announcements) (do say Hi! 👋)

## Looking for sponsors

This tool has reached a level where it can be used to test some web applications
with JavaScript, e.g., simple HTMX applications in Go, particularly combined
with HTMX, a tech combination which is becoming increasingly popular.

But there is still a lot to build to support just the most relevant Web APIs.

I've bade good progress because of too much spare time; but that will not last. 

If I could find enough sponsors, it could mean the difference between continued
development, or death 😢

## Join the "community"

- [Join my discord server](https://discord.gg/rPBRt8Rf) to chat with me, and stay
up-to-date on progress
- Participate in the [github discussions](https://github.com/orgs/gost-dom/discussions), and [say
hi!](https://github.com/orgs/gost-dom/discussions)
- Want to contribute to the project, there's a very rough early [guide describing the overall project structure](./CONTRIBUTING.md)

<!---
## Project background

Go and HTMX is gaining in popularity as a stack.

While Go has great tooling for verifying request/responses of HTTP applications,
but for HTMX, or just client-side scripting with server side rendering, you need
browser automation to test the behaviour.

This introduces a significant overhead; not only from out-of-process
communication with the browser, but also the necessity of launching your server.

This overhead discourages a TDD loop.

The purpose of this project is to enable a fast TDD feedback loop these types of
project, where verification depend on

- Behaviour of client-side scripts.
- Browser behaviour when interacting with browser elements, e.g., clicking the
  submit button submits a form, and redirects are followed.

### Unique features

Being written in Go, this library supports consuming an
[`http.Handler`](https://pkg.go.dev/net/http#Handler) directly. This removes the
necessity managing TCP ports, and start a server on a real port. Your HTTP
server is consumed by test code, like any other Go component would, also
allowing you to replace dependencies for the test if applicable.

This also makes it easy to run parallel tests in isolation as each can create
their own _instance_ of the HTTP handler.

### Drawbacks to Browser automation

- You cannot verify how it look; e.g. you cannot get a screenshot of a failing
test, nor use such screenshots for snapshot tests.
- The verification doesn't prove that it works as intended in _all browsers_ you
want to support.

This isn't intended as a replacement for the cases where an end-2-end test is
the right choice. It is intended as a tool to help when you want a smaller
isolated test, e.g. mocking out part of the behaviour;

--->

## Project status

This still early pre-release, and only the core web APIs are supported, and not
100%. Check the [Feature list](./README_FEATURE_LIST.md) for a list.

The 0.1 focus was to support a common session based login-flow using HTMX,
meaning to support content swapping, XHR, forms, and cookies; in order to
identify risks and architectural flaws.

But many features were not fully implemented, e.g., you cannot navigate by
assigning `history.href`, and redirect responses are not followed.

### Memory Leaks

The current implementation is leaking memory for the scope of a browser
`Window`. I.e., all DOM nodes created and deleted for the lifetime of the
window will stay in memory until the window is actively disposed.

**This is not a problem for the intended use case**

#### Why memory leaks

This codebase is a marriage between two garbage collected runtimes, and what is
conceptually _one object_ is split into two, a Go object and a JavaScript
wrapper. As long of them is reachable; so must the other be.

I could join them into one; but that would result in an undesired coupling; the
DOM implementation being coupled to the JavaScript execution engine. Eventually,
a native Go JavaScript runtime will be supported.

A solution to this problem involves the use of weak references. This exists as
an `internal` but [was accepted](https://github.com/golang/go/issues/67552) as a
feature.

For that reason; and because it's not a problem for the intended use case, I
have postponed dealing with that issue.

### Next up

Currently there are two main focus areas

- Element focus
- Fix already implemented features

A side project worked on in parallel is to support Goja as an alternate script engine.

#### Element focus

Implement focus behaviour, including `focus()` and `blur()` methods, and their
Go counterparts, including the relevant events.

This is primarily a priority not because just adding `autofocus` on an input
element that is swapped in by HTMX causes an JavaScript error to be thrown.

#### Fix already implemented features

Many of the already implemented features or APIs are not completely implemented,
a few examples.

- Assigning to the `history` doesn't navigate
- [Live collections](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection) are static.
- Submit buttons cannot override form method and action.

To give users a better chance of predicting what works, and what doesn't, it is
an aim to make sure that existing features work as they would in a real browser.

#### Goja support

V8 depends on Cgo, but [Goja](https://github.com/dop251/goja) is a pure Go
JavaScript engine. While it may not be as complete as V8, it could be a usable
alternative for many projects providing a pure Go option.


V8 support will not go away, so there's always a fallback, if important JS
features are lacking from Goja.

### Future goals

There is much to do, which includes (but this is not a full list):

- Support web-sockets and server events.
- Implement all standard JavaScript classes that a browser should support; but
  not part of the ECMAScript standard itself.
  - JavaScript polyfills would be a good starting point; which is how xpath is
    implemented at the moment.
    - Conversion to native go implementations would be prioritized on usage, e.g.
      [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) 
      would be high in the list of priorities.
- Implement default browser behaviour for user interaction, e.g. pressing 
  <key>enter</key> when an input field has focus should submit the form.

### Long Term Goals

#### CSS Parsing

Parsing CSS woule be nice, allowing test code to verify the resulting styles of
an element; but having a working DOM with a JavaScript engine is higher
priority.

#### Mock external sites

The system may depend on external sites in the browser, most notably identity
providers (IDP), where your app redirects to the IDP, which redirects on
successful login; but could be other services such as map providers, etc.

For testing purposes, replacing this with a dummy replacement would have some
benefits:

- The verification of your system doesn't depend on the availability of an
  external service; when working offline
- Avoid tests breaking because of changes to the external system.
- For an identity provider
  - Avoid pollution of dummy accounts to run your test suite.
  - Avoid locking out test accounts due to _"suspiscious activity"_.
  - The IDP may use a Captcha or 2FA that can be impossible; or difficult to
    control from tests, and would cause a significant slowdown to the test
    suite.
- For applications like map providers
  - Avoid being billed for API use during testing.

## Out of scope.

### Full Spec Compliance

> A goal is not always meant to be reached, it often serves simply as something
> to aim at.
> 
> - Bruce Lee

While it is a goal to reach whatwg spec compliance, the primary goal is to have
a useful tool for testing modern web applications. 

Some specs don't really have any usage in modern web applications. For example,
you generally wouldn't write an application that depends on quirks mode.

Another example is `document.write`. I've yet to work on any application that
depends on this. However, implementing support for this feature require a
complete rewrite of the HTML parser. You would need a really good case (or
sponsorship level) to have that prioritised.

### Accessibility tree

It is not currently planned that this library should maintain the accessibility
tree; nor provide higher level testing capabilities like what
[Testing Library](https://testing-library.com) provides for JavaScript.

These problems _should_ eventually be solved, but could easily be implemented in
a different library with dependency to the DOM alone.

### Visual Rendering

It is not a goal to be able to provide a visual rendering of the DOM. 

But just like the accessibility tree, this could be implemented in a new library
depending only on the interface from here.

## Terminology

Some words inherntly have multiple meanings.

- **Interface**. The IDL Specification defines _interfaces_; which are exposed
in certain scopes, implemented by "classes" in JavaScript. 
  - The interfaces can be composed of _partial_ or _mixin_ interfaces.
  - IDL Interfaces and mixin interfaces are represented in Go, and typically exposed as Go `interface` types.

## Attribution / 3rd party included code.

This library contains [code derived](./scripting/v8host/polyfills/xpath) from the [jsdom project](https://github.com/jsdom/jsdom) distributed under the MIT license.

---

[^1]: Current focus is to support HTMX apps, but eventually React/Angular,
    whatever front-end framework you use, will work in Gost-DOM.
[^2]: Complete isolation depends on _your code_, e.g., if you don't replace
    database dependencies, they are not isolated.
[^3]: This depends on how you configure Gost-DOM. 
