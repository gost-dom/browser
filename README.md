# Gost-DOM - A headless browser for Go

**The Go-to headless browser for TDD workflows.**

Gost-DOM is a headless browser written in Go designed simplify and speed up
testing of web application in Go, particularly when a the application behaviour
relies on JavaScript. Properties of Gost-DOM-based tests:

- Tests run in parallel due to complete _complete isolation_[^1]
- No erratic behaviour due to 100% predictable UI reactions.
- "Blazingly fast". No out-of-process calls, not even thread boundaries for web
  API calls. Web application code runs in the test thread, so a panic in your
  code keeps a full stack trace to the test case. [^2]
- Dependencies can be replaced while testing.

Yet Gost-DOM still uses HTTP request and responses for verification, testing the
entire stack, including how middlewares. 

I wrote a [longer document](./docs/why-gost.md) further describing the use case
for such a tool.

> [!NOTE]
>
> This is 0.x version still, and breaking API changes do occur, but will be
> announced before release in the [Gost-DOM discussions] (do say Hi! 👋)

[Gost-DOM discussions]: https://github.com/orgs/gost-dom/discussions/categories/announcements

## Looking for sponsors

This tool has reached a level where it can be used to test some web applications
with JavaScript, e.g., simple HTMX applications. But there is still a lot to
build to support just the most relevant Web APIs.

I've made good progress because of too much spare time; but that will not last.
If I could find enough sponsors, it could mean the difference between continued
development, or death 😢

For companies wanting to sponsor, I can send formal invoices too. More
information on the project's [Sponsor page](https://gostdom.net/sponsor)

## Getting started

- Read [Getting Started]
- Familiarize yourself with the [Feature list] to know what is implemented.
- [Join my discord server] to chat with me, and stay up-to-date on progress.
- [say hi!] on the github discussions page.
- Read the [contribution guide](./CONTRIBUTING.md) to see how you can help.

[Getting Started]: ./docs/Getting-started.md
[Feature list]: ./docs/Features.md
[say hi!]: https://github.com/orgs/gost-dom/discussions
[Join my discord server]: https://discord.gg/rPBRt8Rf

## Project status

This is still in an early phase, but it's approaching a design that seems
promising for the purpose.

At the moment there's an emphasis on high-risk features that can expose poor
design choices, but the "primary API" has been reasonably stable for a good
amount of time.

### Upcoming work

The near-future work is prioritised around

- Provide distinct error messages when using unsupported JS functions
- Improving `fetch` implementation
- Fix incorrectly implemented features
- Simulate user interaction, such as typing

#### Notify when using unsupported JS functions

Most of the JS mapping layer is auto-generated, including operations and
attribute getter/setters that are not yet supported. This generates explicit
error messages in the callback functions, making it clear when using, that the
failed test case is caused by lack of support in Gost-DOM - not a bug in your
code.

This should be improved further - including when calling functions accepting an
options object, and clien code pass a valid, but unsupported option.

#### Implement more of `fetch`

The current `fetch` implementation is extremely basic, having only support for
simple `GET` requests.

This will be improved in the future for more advanced scenarios, including
correct headers support, cookies, body - including streaming response bodies,
etc.

#### Fix incorrectly implemented features

Many of the already implemented features or APIs are not completely implemented,
a few examples.

- Assigning to the `history` doesn't navigate
- [Live collections] are not live.
- Submit buttons cannot override form method and action.

To give users a better chance of predicting what works, and what doesn't, it is
an aim to make sure that existing features work as they would in a real browser.

[Live collections]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection

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

#### Pure Go script engine

V8 is a "feature complete" JavaScript environment, so V8 support will never go
away. But it has some overhead, and it depends on Cgo. Also, the current V8
layer [leaks memory in the scope of a browsing context](./docs/V8.md).

[Goja](https://github.com/dop251/goja) is a pure Go JavaScript engine, and is
alsmost fully supported as an alternative. JavaScript bindings in code target a
layer of abstraction, allowing the script engine to be replaced.

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

