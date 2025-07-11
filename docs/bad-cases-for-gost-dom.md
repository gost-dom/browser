# When not to use Gost-DOM

Gost-DOM was written with hypermedia frameworks like HTMX and Datastar in mind.
But the more general pattern is web applications that serve server-rendered
HTML, yet a significant part of the application behaviour depends on JavaScript.

When diverging from this scenario, Gost-DOM may be less ideal.

## Single-Page Applications

I would not recommend Gost-DOM to drive the implementation of SPAs. An SPA
contain a significant amount of completexity that is much better developed by
applying a TDD process to the front-end code.

However, a _reasonable use case_ when working with SPAs is verify how the
front-end consumes the APIs, e.g., setting up an initial state containing som
entities, and verify that they are shown on the page meant to display those
entities.

## Applications without JavaScript

Gost-DOM is still an early project, and not everything is implemented correctly.
On top of that, starting a JavaScript runtime has an overhead attached.

While Gost-DOM _can_ verify things like clicking buttons navigate to specific
pages, and form submission works correctly, to a large part I believe just using
a library to parse the HTML, and perform request/response verification is a
better strategy, for example [x/net/html](https://pkg.go.dev/golang.org/x/net/html).[^1]

An exception would be if you foresee that you might want to add a hypermedia
framework, and potentially refactor parts of the application to use these
frameworks. In those cases, Gost-DOM would be a sensible choice, as you have
written the test cases in a way that can verify the future refactoring.

## Crawling random pages

It is not a priority of Gost-DOM to be fully specs-compliant. It is a priority
to be a usable tool for testing modern web applications. Crawling random web
site is likely to visit applications using unsupported web standards, and thus
not work in Gost-DOM.

Using a real browser in headless mode is a much better alternative to this case.

Crawling your own applications can be a viable use case, if you know they don't
use deprecated features unsupported by Gost-DOM.

But ... does the features of Gost-DOM benefit you? You don't need back-door
manipulation of a Go-based back-end, and the state of the system before/after
exercising it. And you don't depend that heavily on the speed gained by running
the browser running in the same thread of execution as the test case (unless you
perform extremely chatty communication).

[^1]: This is actually what Gost-DOM uses internally to parse HTML; but that's
    an implementation detail not exposed in the API.
