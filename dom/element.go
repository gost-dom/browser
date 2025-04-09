package dom

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	. "github.com/gost-dom/browser/internal/dom"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// TODO: In the DOM, this is a `NamedNodeMap`. Is that useful in Go?
type Attributes []Attr

func (attrs Attributes) Length() int {
	return len(attrs)
}

// An Element in the document. Can be either an [HTMLElement] or an [XMLElement]
type Element interface {
	ElementContainer
	ClassList() DOMTokenList
	HasAttribute(name string) bool
	GetAttribute(name string) (string, bool)
	SetAttribute(name string, value string)
	RemoveAttribute(name string)
	GetAttributeNode(string) Attr
	SetAttributeNode(Attr) (Attr, error)
	RemoveAttributeNode(Attr) (Attr, error)
	Attributes() NamedNodeMap
	InsertAdjacentHTML(position string, text string) error
	OuterHTML() string
	InnerHTML() string
	SetInnerHTML(string) error
	TagName() string
	Matches(string) (bool, error)
	// unexported
	getAttributes() Attributes
	getSelfElement() Element
}

type element struct {
	node
	ParentNode
	tagName          string
	namespace        string
	attributes       Attributes
	selfElement      Element
	selfRenderer     Renderer
	childrenRenderer ChildrenRenderer
	// We might want a "prototype" as a value, rather than a Go type, as new types
	// can be created at runtime. But if so, we probably want them on the node
	// type.
}

func NewElement(tagName string, ownerDocument Document) Element {
	node := newNode(ownerDocument)
	result := &element{
		node: node,
		// ParentNode: newParentNode(node),
		tagName:    tagName,
		attributes: Attributes(nil),
	}
	result.ParentNode = newParentNode(&result.node)
	result.SetSelf(result)
	return result
}

func (e *element) SetSelf(n Node) {
	if self, ok := n.(Element); ok {
		e.selfElement = self
	} else {
		panic("Setting a non-element as element self")
	}
	if self, ok := n.(Renderer); ok {
		e.selfRenderer = self
	} else {
		panic("Setting a non-renderer as element self")
	}
	if self, ok := n.(ChildrenRenderer); ok {
		e.childrenRenderer = self
	} else {
		panic("Setting a non-child-renderer as element self")
	}
	e.node.SetSelf(n)
}

func (e *element) NodeName() string {
	return e.selfElement.TagName()
}

func (e *element) TagName() string {
	return strings.ToLower(e.tagName)
}

func (e *element) ClassList() DOMTokenList {
	return NewClassList(e)
}

func (e *element) OuterHTML() string {
	writer := &strings.Builder{}
	e.selfRenderer.Render(writer)
	return writer.String()
}

func (e *element) InnerHTML() string {
	writer := &strings.Builder{}
	e.renderChildren(writer)
	return writer.String()
}

func (e *element) SetInnerHTML(html string) error {
	fragment, err := e.nodeDocument().parseFragment(strings.NewReader(html))
	if err == nil {
		err = e.ReplaceChildren(fragment)
	}
	return err
}

func (e *element) HasAttribute(name string) bool {
	for _, a := range e.attributes {
		if a.Name() == name {
			return true
		}
	}
	return false
}

func (e *element) GetAttribute(name string) (string, bool) {
	if a := e.GetAttributeNode(name); a != nil {
		return a.Value(), true
	} else {
		return "", false
	}
}

func (e *element) RemoveAttribute(name string) {
	if a := e.GetAttributeNode(name); a != nil {
		e.RemoveAttributeNode(a)
	}
}

func (e *element) GetAttributeNode(name string) Attr {
	for _, a := range e.attributes {
		if a.Name() == name && a.NamespaceURI() == e.namespace {
			return a
		}
	}
	return nil
}

func (e *element) SetAttributeNode(node Attr) (Attr, error) {
	if node.Parent() != nil {
		return nil, newDomError("Attribute already in use")
	}
	for i, a := range e.attributes {
		if a.Name() == node.Name() && a.NamespaceURI() == node.NamespaceURI() {
			e.attributes[i] = node
			return a, nil
		}
	}
	e.attributes = append(e.attributes, node)
	return nil, nil
}

func (e *element) RemoveAttributeNode(node Attr) (Attr, error) {
	for i, a := range e.attributes {
		if a == node {
			e.attributes = slices.Delete(e.attributes, i, i+1)
			node.setParent(nil)
			return node, nil
		}
	}
	return nil, newDomErrorCode("Node was not found", domErrorNotFound)
}

func (e *element) getAttributes() Attributes {
	return e.attributes
}

func (e *element) getSelfElement() Element {
	if r := e.selfElement; r != nil {
		return r
	}
	panic(
		"Calling method on an element which isn't an element. Did a custom type forget to call 'setSelf()'?",
	)
}

func (e *element) Attributes() NamedNodeMap {
	return &namedNodeMap{ownerElement: e}
}

func (e *element) SetAttribute(name string, value string) {
	if a := e.GetAttributeNode(name); a != nil {
		a.SetValue(value)
	} else {
		a := newAttr(name, value, e.OwnerDocument())
		a.setParent(e.selfElement)
		e.attributes = append(e.attributes, a)
	}
}

func (e *element) createHtmlNode() *html.Node {
	tag := strings.ToLower(e.tagName)
	attrs := make([]html.Attribute, len(e.attributes))
	for i, a := range e.attributes {
		attrs[i] = a.htmlAttr()
	}
	return &html.Node{
		Type:      html.ElementNode,
		Data:      tag,
		DataAtom:  atom.Lookup([]byte(tag)),
		Namespace: e.namespace,
		Attr:      attrs,
	}
}

func (n *element) InsertAdjacentHTML(position string, text string) error {
	var (
		parent    Node
		reference Node
	)
	switch position {
	case "beforebegin":
		parent = n.Parent()
		reference = n.getSelf()
	case "afterbegin":
		parent = n
		reference = n.ChildNodes().Item(0)
	case "beforeend":
		parent = n
		reference = nil
	case "afterend":
		parent = n.Parent()
		reference = n.NextSibling()
	default:
		return errors.New("Invalid position")
	}
	fragment, err := n.nodeDocument().parseFragment(strings.NewReader(text))
	if err == nil {
		_, err = parent.InsertBefore(fragment, reference)
	}
	return err
}

func (n *element) NodeType() NodeType { return NodeTypeElement }

func (e *element) Render(writer *strings.Builder) {
	renderElement(e, writer)
}

func renderElement(e *element, writer *strings.Builder) {
	tagName := strings.ToLower(e.TagName())
	writer.WriteRune('<')
	writer.WriteString(tagName)
	for a := range e.Attributes().All() {
		writer.WriteRune(' ')
		writer.WriteString(a.Name())
		writer.WriteString("=\"")
		writer.WriteString(a.Value())
		writer.WriteString("\"")
	}
	writer.WriteRune('>')
	e.childrenRenderer.RenderChildren(writer)
	writer.WriteString("</")
	writer.WriteString(tagName)
	writer.WriteRune('>')
}

var tagNameRegExp = regexp.MustCompile("(?m:^[a-zA-Z]+$)")
var attributeRegExp = regexp.MustCompile("(?m:^[[]([a-zA-Z-]+)[]]$)")
var tagNameAndAttribute = regexp.MustCompile(`(?m:^([a-zA-Z]+)+[[]([a-zA-Z-]+)="([a-zA-Z-]+)"[]]$)`)

// Element.Matches returns true if the current element matches the specified CSS
// selectors; accepting a comma-separated list of selectors with any leading and
// trailing whitespace trimmed. Returns an error if the patterns is not
// supported (or invalid)
func (e *element) Matches(pattern string) (res bool, err error) {
	dummy := e.OwnerDocument().CreateElement("div")
	clone := e.self.CloneNode(true)
	dummy.Append(clone)
	el, err := dummy.QuerySelectorAll(pattern)
	if err == nil {
		for _, e := range el.All() {
			if e == clone {
				return true, nil
			}
		}
	}
	return false, err
}

func (e *element) String() string {
	childLen := e.ChildNodes().Length()

	id, found := e.GetAttribute("id")
	if found {
		id = "id='" + id + "'"
	}
	return fmt.Sprintf("<%s %s(child count=%d) />", e.tagName, id, childLen)
}

func (e *element) CloneNode(deep bool) Node {
	doc := e.OwnerDocument()
	tag := e.selfElement.TagName()
	res := doc.CreateElement(tag)
	for a := range e.Attributes().All() {
		res.SetAttributeNode(a.CloneNode(deep).(Attr))
	}
	if deep {
		res.Append(e.cloneChildren()...)
	}
	return res
}
