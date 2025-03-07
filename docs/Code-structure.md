# Code structure

This is probably not the most obvious code base to get into, and some
unidiomatic choices were made. 

Hopefully this can help potential contributors to get a better understanding of
the code, and the unidiomatic choices - perhaps event come with better, more
idiomatic solutions to the problems.

```sh
browser/ # Main entry point, 
    internal/
        clock/ # Controls the passing of simulated time
        code-gen/ # Code generator. Not documented here!
        constants/ # String constants, like 
        dom/ # Hmmm 
        domslices/ # ... 
        entity/ # Generates objects with JS-friendly unique IDs
        html/ # XMLHTTPRequest implementation
        http/ # Helper for handling HTTP requests
        interfaces/ # Interfaces generated from webref specs
            url-interfaces/ # Interfaces for the URL spec.
        log/ # Logging functionality
        test/ # Code used for testing, 
            htmx-app/ # A test application hosting HTMX, exercised by tests
            html-app-main/ # Helps run the test app in a _real_ browser.
            scripttests/ # Common test suite for JavaScript implementations
        testing/ # Other test code, that makes you wonder, why there is 'test' and 'testing' as two separate folders ...
    dom/ # Code dom functionality. Node elements, and 
        event/ # EventTarget and Event types
    html/ # Window, HTMLDocument, and specific HTML element implementations
    logger/ # Provides a way for client code of receiving log messages
    scripting/ # Impementing of scripting engines
        v8host/ # Scripting support using v8 (and CGo)
        gojahost/ # INCOMPLETE scripting support using Goja, a pure Go engine
    testing/ # Provides usable test helpers for users of the library
        gomega-matchers/ # Matchers useful for users of the Gomega library
    url/ # Implementation of URL behaviour
```

## Structure loosely follows Web APIs

The web standards defines a series of named APIs. The code attempts to follow this to the extend possible.

- `browser/url` corresponds to the [URL API](https://developer.mozilla.org/en-US/docs/Web/API/URL)
- `browser/dom` corresponds to the [DOM API](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model)
- `browser/html` corresponds to the [HTML DOM API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API)

Following this structure strictly isn't without challenges. An `HTMLFormElement`
dispatches `formdata` events, referencing `FormData` objects. `FormData` is
defined in the [XMLHttpRequest API](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest_API), yet
the `FormData` constructor accepts an `HTMLFormElement` as an optional
constructor argument, causing a circular reference. This `FormData` currently
exists in the `html` package.

## Heavy use of internal packages

In order to keep as much freedom for refactoring without breaking compatibility,
code is largely placed in internal packages until it's known if it provided
value to test code.

For example, the XMLHttpRequest can be constructed from and used from
JavaScript, but it doesn't seem to have any value to provide to Go code. You Go
code can already make HTTP requests, so using the clunky XHR interface doesn't
seem to provide any value.

## Internal interfaces are for verification

Package `browser/internal/interfaces` contain interfaces generated from web IDL
specifications. The interfaces are not referenced in production code, but test
code for the `URL` type uses this interface. This effectively makes it a static
check, that the `URL` type conforms to the web IDL specifications.

