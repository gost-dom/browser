# Script engine

JavaScript execution is performed by an _optional script engine_. I.e., you can
create a browser without one, and it'll just execute behaviour for `<script>`
blocks.

There are two script engines in the code base

- `scripting/v8engine` - Uses V8
- `scripting/sobekengine` - 

## Pros and cons

> [!NOTE]
> While the pros and cons measures factors affecting performance, both build and
> runtime, nothing has been measured yet.

### V8

- Pros
  - Based on Chrome's V8 JavaScript engine, this is more or less guaranteed to
    be maintained, supporting new JavaScript features
  - Supports creating "templates", avoiding recreating global scope for every
    test or when the browser navigates/reloads
- Cons
  - Requires CGo (although only for test code, not production code for the
    intended use case)
  - JavaScript objects to not integrate to Go's garbage collector. As a Go
    object can keep a reference to a JS Object, and a JS object can keep a
    reference to a Go object, the v8 engine effectively leaks object _in the
    context of a script context_. It shouldn't be a problem for the intended use
    case where contexts are created repeatedly.
  - V8 is a massive project, causing longer build times

### Sobek

[Sobek], a fork of [Goja] is a pure Go JavaScript engine. While Goja doesn't
have ESM support, Sobek was forked for this purpose.

- Pros
  - As JS objects are implemented by Go objects, garbage collection works as it
    should.
  - Pure Go dependency makes for significantly less build-time headaches,
    including potential future updates, thay may require CGo build changes.
  - Build times should _potentially_ be much faster, but this hasn't been
    measured.
- Cons
  - No support for creating a "template" of global scope. Every new JS scope
    needs to be initialized with all global scope.
    - The runtime cust of this setup hasn't been measured.

[Goja]: https://github.com/dop251/goja
[Sobak]: https://github.com/grafana/sobek

---

## Appendix

### About v8go

V8 is the JavaScript engine in Chrome. It has a C++ API, and requires a lot of
CGo code to work. However, this library is only intended for _testing_, not
production. So your production system does not inherit a CGo dependency.

V8 is based on the v8go project, Originally created by [Github user
rogchap](https://github.com/rogchap/v8go), but left unmaintaind. Now
[Tommie's branch](https://github.com/tommie/v8go) is the best maintained, e.g.,
it automatically pulls latest V8 sources from the chromium repository

Unfortunately, many v8 necessary v8 features were not implemented. I have added
support for  these in [my own v8go fork], and working with Tommie to get them
merged into his.

[my own v8go fork]: (https://github.com/stroiman/v8go/tree/go-dom-feature-dev)
