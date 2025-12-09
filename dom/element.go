package dom

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gost-dom/browser/internal/constants"
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
	ChildNode
	NonDocumentTypeChildNode
	ClassList() DOMTokenList
	Closest(string) (Element, error)
	HasAttribute(name string) bool
	GetAttribute(name string) (string, bool)
	SetAttribute(name string, value string)
	RemoveAttribute(name string)
	GetAttributeNode(string) Attr
	SetAttributeNode(Attr) (Attr, error)
	GetElementsByTagName(string) NodeList
	RemoveAttributeNode(Attr) (Attr, error)
	Attributes() NamedNodeMap
	InsertAdjacentElement(position string, element Element) (Element, error)
	InsertAdjacentHTML(position string, text string) error
	InsertAdjacentText(position string, text string) error
	OuterHTML() string
	SetOuterHTML(string) error
	InnerHTML() string
	SetInnerHTML(string) error
	TagName() string
	Namespace() string
	Matches(string) (bool, error)
	ID() string
	SetID(string)

	// unexported

	getAttributes() Attributes
	getSelfElement() Element
}

type element struct {
	childNode
	ElementParent
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
	result := &element{
		childNode:  newChildNode(ownerDocument),
		tagName:    tagName,
		attributes: Attributes(nil),
	}
	result.ElementParent = newParentNode(&result.node)
	result.SetSelf(result)
	return result
}

func (e *element) IsEqualNode(n Node) bool {
	if !e.isEqualNode(n) {
		return false
	}
	other, ok := n.(Element)
	if !ok {
		return false
	}
	if len(e.attributes) != other.Attributes().Length() {
		return false
	}
	if e.tagName != other.TagName() || e.namespace != other.Namespace() {
		return false
	}
	for _, a := range e.attributes {
		if v, ok := other.GetAttribute(a.Name()); !ok || v != a.Value() {
			return false
		}
	}
	return true
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

func (e *element) Namespace() string { return e.namespace }

func (e *element) ID() string {
	id, _ := e.GetAttribute("id")
	return id
}

func (e *element) SetID(val string) {
	e.SetAttribute("id", val)
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

func (e *element) SetOuterHTML(html string) error {
	parent := e.ParentNode()
	if parent == nil {
		return nil
	}
	if parent.NodeType() == NodeTypeDocument {
		return newDomError("NoModificationAllowed")
	}
	if parent.NodeType() == NodeTypeDocumentFragment {
		return fmt.Errorf(
			"SetOuterHTML not yet supported when parent is a fragment. %s",
			constants.BUG_ISSUE_DETAILS,
		)
	}
	fragment, err := e.nodeDocument().parseFragment(strings.NewReader(html))
	if err == nil {
		_, err = parent.ReplaceChild(fragment, e.getSelf())
	}
	return err
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
	if node.ParentNode() != nil {
		return nil, newDomError("Attribute already in use")
	}
	for i, a := range e.attributes {
		if a.Name() == node.Name() && a.NamespaceURI() == node.NamespaceURI() {
			e.attributes[i] = node
			e.notify(e.attributeChangedEvent(a, a.Value()))
			return a, nil
		}
	}
	node.setParent(e.selfElement)

	e.notify(e.attributeChangedEvent(node, ""))
	e.attributes = append(e.attributes, node)
	return nil, nil
}

func (e *element) attributeChangedEvent(attr Attr, oldVal string) ChangeEvent {
	return ChangeEvent{
		Target:   e.self,
		Type:     ChangeEventAttributes,
		Attr:     attr,
		OldValue: oldVal,
	}
}

func (e *element) RemoveAttributeNode(node Attr) (Attr, error) {
	for i, a := range e.attributes {
		if a == node {
			e.attributes = slices.Delete(e.attributes, i, i+1)
			node.setParent(nil)
			e.notify(e.attributeChangedEvent(a, a.Value()))
			return node, nil
		}
	}
	return nil, newDomErrorCode("Node was not found", not_found_err)
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
		prevVal := a.Value()
		a.SetValue(value)
		e.notify(e.attributeChangedEvent(a, prevVal))
	} else {
		e.SetAttributeNode(newAttr(name, value, e.OwnerDocument()))
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

func (n *element) insertAdjacentNode(position string, node Node) error {
	var (
		parent    Node
		reference Node
	)
	switch position {
	case "beforebegin":
		parent = n.ParentNode()
		reference = n.getSelf()
	case "afterbegin":
		parent = n
		reference = n.ChildNodes().Item(0)
	case "beforeend":
		parent = n
		reference = nil
	case "afterend":
		parent = n.ParentNode()
		reference = n.NextSibling()
	default:
		return errors.New("Invalid position")
	}
	_, err := parent.InsertBefore(node, reference)
	return err
}

func (n *element) InsertAdjacentElement(position string, element Element) (res Element, err error) {
	err = n.insertAdjacentNode(position, element)
	if err == nil {
		res = element
	}
	return
}

func (n *element) InsertAdjacentText(position string, text string) (err error) {
	node := n.OwnerDocument().CreateTextNode(text)
	return n.insertAdjacentNode(position, node)
}

func (n *element) InsertAdjacentHTML(position string, text string) error {
	fragment, err := n.nodeDocument().parseFragment(strings.NewReader(text))
	if err == nil {
		err = n.insertAdjacentNode(position, fragment)
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
	// This less-than-obvious implementation is due to the fact that Gost-DOM
	// uses a library for CSS selectors, but that library doesn't support the
	// "Matches" function.
	dummy := e.OwnerDocument().CreateElement("div")
	clone := e.self.CloneNode(true)
	dummy.Append(clone)
	el, err := dummy.QuerySelectorAll(pattern)
	if err == nil {
		return slices.Contains(el.All(), clone), nil
	}
	return false, err
}

func (e *element) Closest(pattern string) (Element, error) {
	ok, err := e.Matches(pattern)
	if ok {
		return e.getSelfElement(), nil
	}
	if err != nil {
		return nil, err
	}
	parent := e.ParentElement()
	if parent == nil {
		return nil, nil
	}
	return parent.Closest(pattern)
}

func (e *element) String() string {
	childLen := e.ChildNodes().Length()

	id, found := e.GetAttribute("id")
	if found {
		id = "id='" + id + "'"
	}
	return fmt.Sprintf("<%s %s(child count=%d) />", e.tagName, id, childLen)
}

func (e *element) cloneNode(doc Document, deep bool) Node {
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

func (e *element) NextElementSibling() Element {
	var n Node = e
	for {
		n = n.NextSibling()
		if n == nil {
			return nil
		}
		if res, ok := n.(Element); ok {
			return res
		}
	}
}

func (e *element) PreviousElementSibling() Element {
	var n Node = e
	for {
		n = n.PreviousSibling()
		if n == nil {
			return nil
		}
		if res, ok := n.(Element); ok {
			return res
		}
	}
}

func (n *element) GetElementsByTagName(qualifiedName string) NodeList {
	res, err := n.QuerySelectorAll(qualifiedName)
	if err != nil {
		panic(fmt.Sprintf("element.GetElementsByTagName: %v", err))
	}
	return res
}
