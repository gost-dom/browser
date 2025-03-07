# DOM Nodes

## A tricky problem in the domain itself.

Implementing the actual DOM tree in Go presents a problem: How to implement the
DOM specification, which has a very object-oriented interface, in Go, which
doesn't support "function overloading on base classes".

Unfortunately, this does lead to very large interfaces, which is very
unidiomatic Go. Hopefully, new insights can over time lead to more idiomatic
solutions, while still being able to provide the full DOM interface where
necessary, e.g., for Script integration.

## The current state

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

So why is the interface necessary? If a function like `Parent()`  returns a
`node` struct, that couldn't be treated as an element, if it was embedded inside
an element. While an embedder receives all the methods of the embedded object;
and may "override" some, the embedded object doesn't "know" anything about being
embedded.

Only by returning a `Node` interface, does it provide for the actual return
value to be a `Document`, `Doctype`, `Element`, an `HTMLElement`, or any of the
special HTML elements, like `HTMLButtonElement`, etc.

## The self problem!

When an element is appended using `AppendChild` it's parent is updated to the
new parent. An unexported method is added to the `Node` interface

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

When calling `AppendChild` on a specialised `Element`, when the method
implementation on `Node` passes itself to `setParent`, it is the embedded node,
not the `Element`. So calling `Parent()` on the child wouldn't return an
`Element` any longer, just a `Node`.

### The solution

The solution is that a `node` has a `self` value.

```Go
type node struct {
    // ...
    self Node
}
```

Any specialised element type **must** pass a reference to _its outer self_ as a
function call to `SetSelf` after being constructed. For example:

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

### Not happy about it

I am certainly not happy about this solution. Objectively the fact that every
element _must_ do something special in order to work correctly.

But it just feels wrong.

### Maybe an improvement?

Actually, while writing this, I realised that Elements can't be created
directly, only through a call to `document.createElement`. So at least the
possibility of not calling `SetSelf` can be mitigated, as there's a single point
in code that fix this for all elements.

But it still doesn't feel right.

### A strategy pattern?

Maybe a strategy pattern had been a better solution, but each specialised HTML
element has special methods, e.g., for reading/writing IDL attribute. E.g., the
`HTMLFormElement` has the `method` and `action` IDL attributes, that need to
exist on the Go implementation of that object as well. So all the different
implementations _don't_ have the same interface, they merely share a common
subset.
