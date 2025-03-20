# Feature list

Gost-DOM is still in version 0.x and not all features are supported.

Don't hesitate to [Start a
discussion](https://github.com/orgs/gost-dom/discussions) or [Submit an
issue](https://github.com/gost-dom/browser/issues?q=sort%3Aupdated-desc+is%3Aissue+is%3Aopen)
if you need any of the features missing.

User feedback helps setting the direction and prioritising development.

For missing features, described here, check the linked issues to see if this
document might be missing updates updates.

## General browser behaviour

Opening windows with isolated script context, sharing a cookie store.

## Script support

The browser supports plain JavaScript, either as content in a `<script>`
element, or `<script src="">`.

- `async` or `defer` attributes are ignored, and the script is executed when
mounted in the DOM. 
- ESM is not supported, [#64](https://github.com/gost-dom/browser/issues/64).
- WASM is not supported.

### Script runtime

- Time does not advance by itself, e.g., for `setTimeout`/`setCallback` to
execute. Test code controls the passing of time.
- The `Date` object does not reflect passing of simulated time.

## Supported APIs

The web specifications are grouped into different [Web
APIs](https://developer.mozilla.org/en-US/docs/Web/API). Of these, the following
are partially implemented:

- [Document Object Model (DOM)](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model)
- [HTML DOM](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API)
- [URL](https://developer.mozilla.org/en-US/docs/Web/API/URL_API)
- [XMLHttpRequest](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest_API)
- [Console](https://developer.mozilla.org/en-US/docs/Web/API/console) (sort of)
- [History](https://developer.mozilla.org/en-US/docs/Web/API/History_API)

So no location API, no session storage, cookies API, etc. (Cookies work, but not
exposed to JavaScript)

Note that an modular approach is envisioned for these APIs, allowing
implementations to be build independently.

### DOM API

- Only supports HTML documents. No support for XML or XHTML documents, as well as namespaces.
- `Element.outerHTML` and `innerHTML` output is not properly escaped.
- No support for shadow DOM.
- MutationObserver - [In progress](https://github.com/gost-dom/browser/issues/65)

### HTML API

- Opening a window returns an error if the server response with a non-200 status code.
- Redirects are not followed.
- All HTML elements have the correct class in JavaScript, but only a few have
element-specific behaviour is implemented..

Note: The result can still be used when there is a bad HTTP status code. But
there is no differentiation between a network error, and an HTTP status error.

#### Supported element types

Specific HTML elements have specific behaviour.

- `<a>` Click will navigate to the resource. Fragments are not handled (i.e., a
  fragment link generates an HTTP request, but shouldn't)
- `<form>` (be aware of missing redirects)
    - `<input>`
    - `<button>`/`<input type="submit">` When clicked inside a form, submits. Overriding `method` and `action` is not implemented.
        - [Issue #10](https://github.com/gost-dom/browser/issues/10)
    - Reset behaviour is not implemented (either as a call to `.reset()` or a
    - Missing validation. [Issue #30](https://github.com/gost-dom/browser/issues/30)
    reset button.
- `<template>` exposes its HTML children as children of the `content` attribute.

#### User interaction

The list of supported user interaction is yet quite limited

- `HTMLElement.click()` - implemented on `dom.Element`, due to some weirdness in
  the specs.

### URL API

- Doesn't support object URLs

### XHR API

- `XMLHttpRequest`
    - Only text response is supported.
    - Not all events are emitted, e.g. progress on file upload, but file upload
      isn't supported at all, so ...
- `FormData`
    - Doesn't support file/blob objects

### Console

- Console functions exist, but they are just the ones V8 provides. Logging can
be directed back to Go code, see [Getting started](./Getting-started.md)

### History

- Setting attributes on `window.history` has no effect, i.e., it doesn't navigate.
