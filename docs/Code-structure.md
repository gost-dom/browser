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
        constants/ # String constants, e.g., BUG issue URLs for error messages
        dom/ # Parts of the DOM API not yet available to client code.
        entity/ # Generates objects with JS-friendly unique IDs
        html/ # XMLHTTPRequest implementation
        gosthttp/ # Helper for handling HTTP requests
        interfaces/ # Interfaces generated from webref specs
            url-interfaces/ # Interfaces for the URL spec.
        log/ # Logging functionality
        test/ # Test code.
            integration/ # Contains a simple go web app for testing
            scripttests/ # Common test suite for JavaScript implementations
        testing/ # Test helpers
    dom/ # Core dom functionality. Nodes, elements, attributes, document, etc.
        event/ # EventTarget and Event types
    html/ # Window, HTMLDocument, and specific HTML element implementations
    logger/ # Provides a way for client code of receiving log messages
    scripting/ # Impementing of scripting engines
        internal/
            js/    # Defines the interface for a script engine
            codec/ # Conversion between Go and JS values
            dom/   # Implementation of the DOM API
            html/  # Implementation of the HTML DOM API
        v8host/ # Scripting support using v8 (and CGo)
        gojahost/ # scripting support using Goja, a pure Go engine
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
new code is largely placed in internal packages until it's known if it provides
value to test code; and the public API seems stable.

For example, the XMLHttpRequest can be constructed from and used from
JavaScript, but it doesn't seem to have any value to provide to Go code. Go code
can already make HTTP requests, so using the clunky XHR interface doesn't seem
to provide any value.

## Internal interfaces are for verification

Package `browser/internal/interfaces` contain interfaces generated from web IDL
specifications. The interfaces are not referenced in production code, but test
code for the `URL` type uses this interface. This effectively makes it a static
check, that the `URL` type conforms to the web IDL specifications.
