# Gost-DOM - A headless browser for Go

**The Go-to headless browser for TDD workflows.**

Gost-DOM is a headless browser written in Go intended to write tests of web
application in Go that relies on JavaScript. Properties of Gost-DOM-based tests:

- Tests run in parallel due to complete _complete isolation_[^1]
- No erratic behaviour due to 100% predictable UI reactions.
- "Blazingly fast". No out-of-process calls, not even thread boundaries. Web
  application code runs in the test thread, so a panic in your code keeps a full
  stack trace to the test case. [^2]
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
with JavaScript, e.g., simple HTMX applications. But there is still a lot to
build to support just the most relevant Web APIs.

I've made good progress because of too much spare time; but that will not last.
If I could find enough sponsors, it could mean the difference between continued
development, or death 😢

For companies wanting to sponsor, I can send formal invoices too. More
information on the project's [Sponsor page](https://gostdom.net/sponsor)

## Join the "community"

- [Join my discord server](https://discord.gg/rPBRt8Rf) to chat with me, and stay
up-to-date on progress
- Participate in the [github discussions](https://github.com/orgs/gost-dom/discussions), and [say
hi!](https://github.com/orgs/gost-dom/discussions)
- Want to contribute to the project, there's a very rough early [guide describing the overall project structure](./CONTRIBUTING.md)

## Versioning and breaking changes

While still versioned as 0.x, this library will _generally_ follow the
convention, that a minor version increment indicates a breaking change.

Breaking changes will be announced up-front in the [anouncements
discussions](https://github.com/orgs/gost-dom/discussions/categories/announcements).
When feasible withing a reasonable amount of work, an alternate solution will be
made available prior to releasing the breaking change. Reply to the anouncements
if this would affect you.

## Project status

This still early pre-release, and only the core web APIs are supported, and not
100%. Check the [Feature list](./docs/Features.md) for an overview.

The original priority was to support session based login-flow using in an HTMX
app. This is working, and Gost-DOM supports basic HTMX behaviour, including
content swapping, XHR, forms, and cookies.

In order to identify risks and architectural flaws, focus has moved to support 
[Datastar](https://data-star.dev/)

### Upcoming changes

Currently, the script binding layer is undergoing a refactoring process. 

In order to support DataStar, a lot of new JS APIs were needed. But the
JavaScript layer has some unsolved problems problems in the JavaScript were a
becoming a greater and greater burden:

- Supporting different script engines without rebuilding the entire script layer
  from scratch.
- Design a modular approach, allowing non-core web APIs to be implemented as 3rd
  party modules; without coupling them to V8 and the v8go project.
- Provide better code structure in the JS wrapping layer; rather than all files
  in one package.

This task shows more details: https://github.com/gost-dom/browser/issues/92

#### JS Features

To support DataStar, I need the following features:

- ESM support (requires new features in  v8go)
- `HTMLElement.dataset` (requires new features in v8go)
- [MutationObserver](https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver)
  support.
- `fetch` API, including `AbortController` and `ReadableStream`.

I have experimental working features for ESM, `dataset`. `MutationObserver` is
included in the latest release, but the Go implementation is stil in an
`internal/` package, as I don't think the Go interface is just right yet.

Only `fetch` is lacking completely. I tried a few polyfills to build it on the
JS side on top if XHR - but I was unable to find a working combination of
polyfills that works with Datastar's streaming. So a native fetch implementation
is necessary.

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

### Memory Leaks

The current implementation is leaking memory for the scope of a browsing
context. all DOM nodes created and deleted for the lifetime of the context will
stay in memory until navigating, or the browser is actively disposed.

**This is not a problem for the intended use case**. Web scraping SPAs _might_
become a problem.

Actions that will reset the browsing context are:

- Clicking a link (default behaviour not prevented by JavaScript)
- Submitting a form (default behaviour not prevented by JavaScript)
- Refresh or navigating history entries generated by these operations.

Examples that **will not reset** the browsing context are (not necessarily a
complete list):

- Navigating an SPA.
- Navigating an HTML app with boosted links.
- Navigating history entries generated by these types of actions.

#### Why memory leaks

This codebase is a marriage between two garbage collected runtimes, and what is
conceptually _one object_ is split into two, a Go object and a JavaScript
wrapper. As long of them is reachable; so must the other be.

I could join them into one; but that would result in an undesired coupling; the
DOM implementation being coupled to the JavaScript execution engine.

Since this project started, weak references has been introduced to Go 1.24 which
I beleve provides a solution to the problem, but this hasn't been prioritised 

For that reason; and because it's not a problem for the intended use case, I
have postponed dealing with that issue.


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
- Avoid tests breaking due to a new UI in your external dependency.
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

Some web APIs are implemented [by embedding
polyfills](./scripting/internal/polyfills) from other open-source
JavaScript libraries.

- XPath evaluator uses [code from](./scripting/internal/polyfills/xpath)
  the [jsdom project](https://github.com/jsdom/jsdom) distributed under the MIT
  license.
- [FastestSmallestTextEncoderDecoder](./scripting/internal/polyfills/FastestSmallestTextEncoderDecoder)
  distributed under the Creative Commons Zero v1.0 Universal license.

---

[^1]: Complete isolation depends on _your code_, e.g., if you don't replace
    database dependencies, they are not isolated.
[^2]: This depends on how you configure Gost-DOM. 

[dataset]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset
[MutationObserver]: https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver
