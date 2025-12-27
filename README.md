<h1 align="center"><img src="https://avatars.githubusercontent.com/u/196428063?s=50&v=4" alt="logo" /><br />Gost-DOM<br /><span font-size="0.05em">A headless browser for Go</span></h1>

<p align="center">
<em>The Go-to solution for a TDD workflow.</em>
</p>

<div align="center">
<blockquote margin="100">As a developer<br />
In order to work effeciently with code<br />
I want a fast feedback loop for <em>all</em> of my code.</blockquote></div>

Gost-DOM is a development tool to help building web applications in Go by
providing a blazingly fast, completely reliable, feedback loop.

By simulating a browser envinronment in the test itself, including JavaScript
execution[^1], Gost-DOM allows you to write an automated test suite, and run
them so fast enough that they become a useful feedback tool _while implementing
behaviour_.

<p align="center">
<em>

Gost-DOM is **the headless browser** that provides sub-second feedback with 100%
predictable code execution.

</em>
</p>

To learn more, read [Why Gost-DOM?](./docs/why-gost.md)

## Why not use ...

The typical solution relies on browser automation (e.g., playwright). These have a
significant overhead controlling a remote browser. In addition, developers
often struggle with erratic tests due to unpredictable code execution.

As a result, these types of tests are typically written _after the fact_, when
the system already works.

These tests didn't provide any value as a feedback tool _during development_.


### Not a replacement

Gost-DOM allows verification of individual pieces of behaviour _during
development_, but it's not a full browser, e.g., you cannot export screen shots.

In addition, there are sensible tests to write _after_ features were developed,
e.g., test a complete web-shop order flow from login to check-out. Using a real
browser for these types of tests would be sensible; and verifying in _all_
browsers you support.

## Other Benefits of Gost-DOM

Gost-DOM has a few additional benefits over browser automation:

- Tests run in parallel due to _complete isolation_[^2]
- No erratic behaviour; 100% predictable UI reactions.
- _Blazingly fast_.[^3] No out-of-process calls, not even thread boundaries for web
  API calls as web application code runs in the test thread.[^4]
- Dependencies can be replaced in tests.
- Write tests at a higher level of abstraction, expressing the expected
  _behaviour_ of a system, decoupled from implementation details.

Gost-DOM still uses HTTP request and responses for verification, testing the
entire stack, including how middlewares affect the behaviour, verifying, and
supporting refactoring of e.g., authentication logic.

## Getting started

- Read [Getting Started]
- Familiarize yourself with the [Feature list] to know what is implemented.
- [Join my discord server] to chat with me, and stay up-to-date on progress.
- [say hi!] on the github discussions page.
- Read the [contribution guide](./CONTRIBUTING.md) to see how you can help.

Also, check out the [Shaman module] which provides capabilities of querying the
DOM at a higher level of abstraction, e.g., find an element with a specific
_label_ / _accessibility name_, allowing tests to be more expressive

[Shaman module]: https://github.com/gost-dom/shaman

> [!NOTE]
>
> This is 0.x version still, and breaking API changes do occur, but will be
> announced before release in the [Gost-DOM discussions] (do say Hi! 👋)

[Gost-DOM discussions]: https://github.com/orgs/gost-dom/discussions/categories/announcements
[Getting Started]: ./docs/Getting-started.md
[Feature list]: ./docs/Features.md
[say hi!]: https://github.com/orgs/gost-dom/discussions
[Join my discord server]: https://discord.gg/rPBRt8Rf

## Looking for sponsors

This project is the spare time project for a single developer making good
progress because of too much spare time; but that will not last.

If I could find enough sponsors, it could mean the difference between continued
development, or death 😢

For companies wanting to sponsor, I can send formal invoices too. More
information on the project's [Sponsor page](https://gostdom.net/sponsor)

## Attribution / 3rd party included code.

Some web APIs are implemented [by embedding
polyfills](./scripting/internal/polyfills) from other open-source
JavaScript libraries.

- XPath evaluator uses [code from](./scripting/internal/polyfills/xpath)
  the [jsdom project](https://github.com/jsdom/jsdom) distributed under the MIT
  license.
- [FastestSmallestTextEncoderDecoder](./scripting/internal/polyfills/FastestSmallestTextEncoderDecoder)
  distributed under the Creative Commons Zero v1.0 Universal license.

In addition, to verify compatibility with 3rd party JavaScript libraries, the
[test code] in repository contains compiled versions of:

- [HTMX](https://htmx.org) distributed under the Zero-Clause BSD license.
- [Datastar](https://data-star.dev/) distributed under the MIT license.

[test code]: ./internal/test/integration/test-app/content/public/datastar/datastar.rc6.js

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=gost-dom/browser&type=Date)](https://www.star-history.com/#gost-dom/browser&Date)

---

[^1]: Gost-DOM can use V8 - the same JavaScript engine that powers Chrome; but a
    native Go alternative also exists.
[^2]: Complete isolation depends on _your code_, e.g., if you don't replace
    database dependencies, tests are not isolated.
[^3]: Cliché, I know! But it is!
[^4]: This depends on how you configure Gost-DOM.

