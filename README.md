<h1 align="center"><img src="https://avatars.githubusercontent.com/u/196428063?s=50&v=4" alt="logo" /><br />Gost-DOM<br /><span font-size="0.05em">A headless browser for Go</span></h1>

<p align="center">
<em>The Go-to solution for a TDD workflow.</em>
</p>

<div align="center">
<blockquote margin="100">As a developer<br />
In order to work effeciently with code<br />
I want a fast feedback loop for <em>all</em> of my code.</blockquote></div>

Gost-DOM was born from the philosophy that the fast feedback loop provided by
TDD makes it the most effective way to work with _the majority of the code
base_.

Web UIs are typical exceptions to this rule. They rely on _real browsers_ for
testing which introduce excessive overhead and slows down the feedback loop;
reducing, or even eliminating, the effectiveness of TDD. In addition, developers
often struggle with erratic tests due to unpredictable code execution.

Gost-DOM aims to solve that problem for web projects using Go. Gost-DOM
simulates a browser environment, using a JavaScript engine to execute client
script, allowing you to write test cases in Go to verify application behaviour,
and apply an iterative process; supporting refactoring.[^1]

To learn more, read [Why Gost-DOM?](./docs/why-gost.md)

## Benefits of Gost-DOM

Compared to browser automation, Gost-DOM provides the benefits:

- Tests run in parallel due to complete _complete isolation_[^2]
- No erratic behaviour; 100% predictable UI reactions.
- _Blazingly fast_.[^3] No out-of-process calls, not even thread boundaries for web
  API calls as web application code runs in the test thread.[^4]
- Dependencies can be replaced in tests.
- Write tests at a higher level of abstraction, expressing the expected
  _behaviour_ of a system, decoupled from implementation details.

Gost-DOM still uses HTTP request and responses for verification, testing the
entire stack, including how middlewares affect the behaviour, verifying, and
supporting refactoring of e.g., authentication logic. 

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

> [!NOTE]
>
> This is 0.x version still, and breaking API changes do occur, but will be
> announced before release in the [Gost-DOM discussions] (do say Hi! 👋)

[Gost-DOM discussions]: https://github.com/orgs/gost-dom/discussions/categories/announcements
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

Gost-DOM supports basic HTMX interactions, as well as basic Datastar cases. 

### Future goals

There is much to do, which includes (but this is not a full list):

- Support web-sockets and server-sent events.
- Implement all standard JavaScript classes that a browser should support; but
  not part of the ECMAScript standard itself.
  - JavaScript polyfills would be a good starting point; which is how xpath is
    implemented at the moment.
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

Parsing CSS woule be nice, allowing test code to verify the style properties of
HTML elements; in particular whether an element is visible; but having a working
DOM with a JavaScript engine is higher priority.

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

Some specs don't really have any usage in modern web applications, like
`document.write` or depending on quirks mode.

### Accessibility tree

It is not currently planned that this library should maintain the accessibility
tree.

The [Shaman module] provides capabilities of querying the DOM at a higher level
of abstraction, e.g., find an element with a specific _label_ / _accessibility
name_, allowing tests to be more expressive. This is inspired by the
capabilities that [Testing Library] provides for JavaScript.

[Shaman module]: https://github.com/gost-dom/shaman
[Testing Library]: https://testing-library.com

### Visual Rendering

It is not a goal to be able to provide a visual rendering of the DOM. 

But just like the accessibility tree, this could be implemented in a new library
depending only on the interface from here.

## Terminology

Some words inherntly have multiple meanings.

- **Interface**. The IDL Specification defines _interfaces_; which are exposed
in certain scopes, implemented by "classes" in JavaScript. 
  - The interfaces can be composed of _partial_ or _mixin_ interfaces.
  - IDL Interfaces and mixin interfaces are represented in Go, and typically
    exposed as Go `interface` types.

## Attribution / 3rd party included code.

Some web APIs are implemented [by embedding
polyfills](./scripting/internal/polyfills) from other open-source
JavaScript libraries.

- XPath evaluator uses [code from](./scripting/internal/polyfills/xpath)
  the [jsdom project](https://github.com/jsdom/jsdom) distributed under the MIT
  license.
- [FastestSmallestTextEncoderDecoder](./scripting/internal/polyfills/FastestSmallestTextEncoderDecoder)
  distributed under the Creative Commons Zero v1.0 Universal license.

In addition, for testing compatibility, test code of this repository contains
compiled versions of:

- [HTMX](https://htmx.org)
- [Datastar](https://data-star.dev/)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=gost-dom/browser&type=Date)](https://www.star-history.com/#gost-dom/browser&Date)

---

[^1]: Gost-DOM, by default, embeds a V8 engine - the same JavaScript engine that powers Chrome.
[^2]: Complete isolation depends on _your code_, e.g., if you don't replace
    database dependencies, they are not isolated.
[^3]: Cliché, I know! But it is!
[^4]: This depends on how you configure Gost-DOM. 

