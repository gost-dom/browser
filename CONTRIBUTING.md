# Contributor's Guide

Contributors are welcome, but there is very poor documentation as of now. The
rest of this document aims to bring a high-level overview of code structure.

I would encourage you to [join the
discussions](https://github.com/orgs/gost-dom/discussions), also as a user, in
order to help prioritise the future direction.

There are a few issues marked as "good first issue", as they seem isolated
enough to be easy to start on. But they aren't the highest priority.

## Code structure

This is probably not the most obvious code base to get into, and some unidiomatic choices were made. 

Hopefully this can help potential contributors to get a better understanding of the code, and the unidiomatic choices - perhaps event come with better, more idiomatic solutions to the problems.

```sh
browser/ # Main entry point, 
    internal/
        clock/ # Controls the passing of simulated time
        code-gen/ # Code generator. Not documented here!
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

### Structure loosely follows Web APIs

The web standards defines a series of named APIs. The code attempts to follow this to the extend possible.

- `browser/url` corresponds to the [URL API](https://developer.mozilla.org/en-US/docs/Web/API/URL)
- `browser/dom` corresponds to the [DOM API](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model)
- `browser/html` corresponds to the [HTML DOM API](https://developer.mozilla.org/en-US/docs/Web/API/HTML_DOM_API)

Following this structure strictly isn't without challenges. An `HTMLFormElement` dispatches `formdata` events, referencing `FormData` objects. `FormData` is defined in the [XMLHttpRequest API](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest_API), yet the `FormData` constructor accepts an `HTMLFormElement` as an optional constructor argument, causing a circular reference. This `FormData` currently exists in the `html` package.

### Heavy use of internal packages

In order to keep as much freedom for refactoring without breaking compatibility, code is largely placed in internal packages until it's known if it provided value to test code.

For example, the XMLHttpRequest can be constructed from and used from JavaScript, but it doesn't seem to have any value to provide to Go code. You Go code can already make HTTP requests, so using the clunky XHR interface doesn't seem to provide any value.

### Internal interfaces are for verification

Package `browser/internal/interfaces` contain interfaces generated from web IDL specifications. The interfaces are not referenced in production code, but test code for the `URL` type uses this interface. This effectively makes it a static check, that the `URL` type conforms to the web IDL specifications.

## Unidiomatic Go

Idiomatic Go uses small interfaces, and this leads to many elegant solutions.

But the problem domain itself, implementing browser behaviour **and implementing the standards** does place a significant constraint on the types of solutions.

As a result, there are many types, and specifically interfaces, that are bigger than desired, particularly with the DOM tree itself.

Occasionally, [new insights](https://github.com/orgs/gost-dom/discussions/50) lead to smaller more elegant solutions, and hopefully, this will happen with the DOM tree over time as well, where I'm not entirely happy with the solution.

## The DOM Tree

Implementing the actual DOM tree in Go presents a problem. How to implement the DOM specification, which has a very object-oriented interface, in Go, which doesn't support "function overloading on base classes".

So code structure is an attempt at a solution to this problem. 

```go
type Node interface {
	AppendChild(node Node) (Node, error)
	Parent() Node
 	FirstChild() Node
	TextContent() string
	SetTextContent(string)
}
```

Often, but not always the `Node` is an `Element`.

```Go
type Element interface {
	Node
	GetAttribute(name string) (string, bool)
	SetAttribute(name string, value string)
}
```

The actual implementations have similar relationships

```Go
type node struct {
    children []Node
    parent   Node
}

type element struct {
    node
    attributes []Attr
}
```

Obviously methods on the `node` struct implement the `Node` interface

So why is the interface necessary? If a function like `Parent()`  returns a `node` struct, that couldn't be treated as an element, if it was embedded inside an element. While an embedder receives all the methods of the embedded object; and may "override" some, the embedded object doesn't "know" anything about being embedded.

Only by returning a `Node` interface, does it provide for the actual return value to be a `Document`, `Doctype`, `Element`, an `HTMLElement`, or any of the special HTML elements, like `HTMLButtonElement`, etc.

### The self problem!

When an element is appended using `AppendChild` it's parent is updated to the new parent. An unexported method is added to the `Node` interface

```Go
type Node interface {
    // ...
    setParent(Node)
}

func (n *node) AppendChild(child Node) { // Return values ignored in example
    n.children = append(n.children, child)
    child.setParent(n)
}
```

This doesn't work!!!

When calling `AppendChild` on a specialised `Element`, when the method implementation on `Node` passes itself to `setParent`, it is the embedded node, not the `Element`. So calling `Parent()` on the child wouldn't return an `Element` any longer, just a `Node`.

#### The solution

The solution is that a `node` has a `self` value.

```Go
type node struct {
    // ...
    self Node
}
```

Any specialised element type **must** pass a reference to _its outer self_ as a function call to `SetSelf` after being constructed. For example:

```Go
func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{
		HTMLElement: NewHTMLElement("a", ownerDoc),
		URL:         nil,
	}
	result.SetSelf(result)
	return result
}
```

Failure to do this will cause the system to not work correctly

#### Not happy about it

I am certainly not happy about this solution. Objectively the fact that every element _must_ do something special in order to work correctly.

But it just feels wrong.

#### Maybe an improvement?

Actually, while writing this, I realised that Elements can't be created directly, only through a call to `document.createElement`. So at least the possibility of not calling `SetSelf` can be mitigated, as there's a single point in code that fix this for all elements.

But it still doesn't feel right.

#### A strategy pattern?

Maybe a strategy pattern had been a better solution, but each specialised HTML element has special methods, e.g., for reading/writing IDL attribute. E.g., the `HTMLFormElement` has the `method` and `action` IDL attributes, that need to exist on the Go implementation of that object as well. So all the different implementations _don't_ have the same interface, they merely share a common subset.

## Script hosts

The `html` package defines an interface for a script host, so the actual elements are decoupled from how they are presented to client side scripts.

The two wrappers for v8 and Goja (Goja isn't fully working yet) each proves a way for JavaScript objects.

## Testing

All code must be covered by tests

But the package is a mix of different paradigms, so what to use?

**Don't use Ginkgo**

Existing Ginkgo tests will eventually be migrated to non-Ginko tests. Currently the replacement is largely based on [Testify suites](https://pkg.go.dev/github.com/stretchr/testify/suite) that integrate with Go's subtest functionality. Contributor feedback will affect the future of this.

Gingko was used historically, but this was a mistake, and it's a barrier for new contributors to join.[^1]

### Assertions

You are free to write assertions in any way you like.

Gomega has been used historically, and provides expressive assertions through custom matchers, but the true benefit isn't achieved until you take your time to write the custom matchers.

---

[^1]: When Ginkgo was first conceived, provided some benefits. But these have diminished with advancements in Go's own testing capabilities.

