package dom

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"golang.org/x/net/html"
)

type NodeType int

const (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDataSection          NodeType = 4
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
)

// canHaveChildren returns true for note types that allow child nodes
func (t NodeType) canHaveChildren() bool {
	switch t {
	case NodeTypeElement:
		return true
	case NodeTypeDocument:
		return true
	case NodeTypeDocumentFragment:
		return true
	default:
		return false
	}
}

// isCharacterDataNode returns true the 4 node types that are [Characterdata]
// nodes.
//
// [Characterdata]: https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
func (t NodeType) isCharacterDataNode() bool {
	switch t {
	case NodeTypeText:
		return true
	case NodeTypeCDataSection:
		return true
	case NodeTypeComment:
		return true
	case NodeTypeProcessingInstruction:
		return true
	default:
		return false
	}
}

// canBeAChild returns true for node types that are allowed as children. Node,
// special rules may apply, such as a document node can only contain one child
// element.
func (t NodeType) canBeAChild() bool {
	if t.isCharacterDataNode() {
		return true
	}
	switch t {
	case NodeTypeDocumentFragment:
		return true
	case NodeTypeDocumentType:
		return true
	case NodeTypeElement:
		return true
	default:
		return false
	}
}

// String returns name of the node type. For invalid values, a string
// representation of the integer value is returned.
func (t NodeType) String() string {
	switch t {
	case NodeTypeElement:
		return "Element"
	case NodeTypeAttribute:
		return "Attribute"
	case NodeTypeText:
		return "Text"
	case NodeTypeCDataSection:
		return "CDataSection"
	case NodeTypeProcessingInstruction:
		return "ProcessingInstruction"
	case NodeTypeComment:
		return "Comment"
	case NodeTypeDocument:
		return "Document"
	case NodeTypeDocumentType:
		return "DocumentType"
	case NodeTypeDocumentFragment:
		return "DocumentFragment"
	default:
		return strconv.Itoa(int(t))
	}
}

type GetRootNodeOptions bool

type Closer interface {
	Close()
}

type ChangeEvent struct {
	// The original target of the change.
	Target       Node
	AddedNodes   NodeList
	RemovedNodes NodeList
}

type observer interface {
	Process(ChangeEvent)
}

type Node interface {
	entity.ObjectIder
	event.EventTarget
	Logger() log.Logger
	AppendChild(node Node) (Node, error)
	GetRootNode(options ...GetRootNodeOptions) Node
	ChildNodes() NodeList
	CloneNode(deep bool) Node
	IsConnected() bool
	// IsSameNode shouldn't be used and may be removed in a future version.
	IsSameNode(Node) bool
	Contains(node Node) bool
	InsertBefore(newNode Node, referenceNode Node) (Node, error)
	NodeName() string
	NodeType() NodeType
	OwnerDocument() Document
	Parent() Node
	ParentElement() Element
	RemoveChild(node Node) (Node, error)
	NextSibling() Node
	PreviousSibling() Node
	FirstChild() Node
	TextContent() string
	SetTextContent(value string)
	Connected()
	// SetSelf must be called when creating instances of structs embedding a Node.
	//
	// If this is not called, the specialised type, which is itself a Node, will
	// not be returned from functions that should have returned it, e.g., through
	// ChildNodes. Only the embedded Node will be returned, and any specialised
	// behaviour, including HTML output, will not work.
	//
	// This function is a workaround to solve a fundamental problem. The DOM
	// specifies a model that is fundamentally object-oriented, with sub-classes
	// overriding behaviour in super-classes. This is not a behaviour that Go has.
	SetSelf(node Node)

	Observe(observer) Closer

	getSelf() Node
	setParent(Node)
	setOwnerDocument(owner Document)
	nodes() []Node
	assertCanAddNode(Node) error
	cloneChildren() []Node
	createHtmlNode() *html.Node
	nodeDocument() Document
	notify(ChangeEvent)
}

type node struct {
	event.EventTarget
	entity.Entity
	self       Node
	childNodes NodeList
	parent     Node
	document   Document
	observers  []observer
}

func newNode(ownerDocument Document) node {
	return node{
		EventTarget: event.NewEventTarget(),
		childNodes:  &nodeList{},
		document:    ownerDocument,
	}
}

func newNodePtr(ownerDocument Document) *node {
	n := newNode(ownerDocument)
	return &n
}

func (n *node) cloneChildren() []Node {
	children := n.ChildNodes().All()
	nodes := make([]Node, len(children))
	for i, n := range children {
		nodes[i] = n.CloneNode(true)
	}
	return nodes
}

// AppendChild adds node to the end of the list of the current node's child
// nodes. The appended child is returned. If node is not a valid child of the
// parent, a [DOMError] is returned. The [MDN docs for appendChild] lists the
// error conditions.
//
// [MDN docs for appendChild]: https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (n *node) AppendChild(node Node) (Node, error) {
	log.Debug(n.Logger(), "Node.AppendChild", "target", n.String(), "child", node.NodeName())
	_, err := n.self.InsertBefore(node, nil)
	return node, err
}

func (n *node) InsertBefore(newChild Node, referenceNode Node) (Node, error) {
	if err := n.assertCanAddNode(newChild); err != nil {
		return nil, err
	}
	if fragment, ok := newChild.(DocumentFragment); ok {
		for fragment.ChildNodes().Length() > 0 {
			if _, err := n.InsertBefore(fragment.ChildNodes().Item(0), referenceNode); err != nil {
				return nil, err
			}
		}
		return fragment, nil
	}
	result, err := n.insertBefore(newChild, referenceNode)
	if err == nil {
		newChild.setParent(n.self)
	}
	if newChild.IsConnected() {
		newChild.Connected()
	}
	return result, err
}

func (n *node) ChildNodes() NodeList { return n.childNodes }

func (n *node) Observe(observer observer) Closer {
	if slices.Contains(n.observers, observer) {
		panic("Observer already added to this node")
	}
	n.observers = append(n.observers, observer)
	return observerCloser{n, observer}
}

func (n *node) removeObserver(o observer) {
	n.observers = slices.DeleteFunc(n.observers, func(x observer) bool { return o == x })
}

type observerCloser struct {
	n *node
	o observer
}

func (c observerCloser) Close() {
	c.n.removeObserver(c.o)
}

// func (n *node) CloneNode(deep bool) Node { return nil }

func (n *node) GetRootNode(options ...GetRootNodeOptions) Node {
	if len(options) > 1 {
		log.Warn(n.Logger(), "Node.GetRootNode: composed not yet implemented")
	}
	if n.parent == nil {
		return n.self
	} else {
		return n.parent.GetRootNode(options...)
	}
}

func (n *node) Contains(node Node) bool {
	for _, c := range n.ChildNodes().All() {
		if c == node || c.Contains(node) {
			return true
		}
	}
	return false
}

func (n *node) Parent() Node { return n.parent }

func (n *node) ParentElement() Element {
	r, _ := n.Parent().(Element)
	return r
}

func (n *node) setOwnerDocument(owner Document) {
	n.document = owner
	for _, n := range n.ChildNodes().All() {
		n.setOwnerDocument(owner)
	}
}
func (n *node) setParent(parent Node) {
	if parent != nil {
		parentOwner := parent.nodeDocument()
		if n.document != parentOwner {
			n.setOwnerDocument(parentOwner)
		}
	}
	n.parent = parent
	n.SetParentTarget(parent)
}

func (n *node) Connected() {
	if p := n.getSelf().Parent(); p != nil {
		p.Connected()
	}
}

func (n *node) IsConnected() (result bool) {
	if n.parent != nil {
		result = n.parent.IsConnected()
	}
	return
}

func (n *node) IsSameNode(other Node) (result bool) {
	return n.getSelf() == other
}

func (n *node) NodeName() string {
	return "#node"
}

// removeNodeFromParent removes the node from the current parent, _if_ it has
// one. Does nothing for disconnected nodes.
func removeNodeFromParent(node Node) {
	parent := node.Parent()
	if parent != nil {
		parent.RemoveChild(node)
	}
}

func (n *node) RemoveChild(node Node) (Node, error) {
	idx := slices.Index(n.childNodes.All(), node)
	if idx == -1 {
		return nil, newDomError(
			"Node.removeChild: The node to be removed is not a child of this node",
		)
	}
	n.childNodes.setNodes(slices.Delete(n.childNodes.All(), idx, idx+1))
	n.notify(n.removedNodeEvent(node))
	return node, nil
}

// assertCanAddNode verifies that the node can be added as a child. The function
// returns the corresponding [DOMError] that should be returned from the
// relevant function. If the node is a valid new child in the current state of
// the node, the return value is nill
//
// This is a separate function for the purpose of checking all arguments to
// [Element.Append] before adding, to avoid a partial update if the last
// argument was invalid.
func (n *node) assertCanAddNode(newNode Node) error {
	parentType := n.getSelf().NodeType()
	childType := newNode.NodeType()
	if !parentType.canHaveChildren() {
		return newDomError(
			fmt.Sprintf("May not add children to node type %s", parentType),
		)
	}
	if !childType.canBeAChild() {
		return newDomError(
			fmt.Sprintf("May not add an node type %s as a child", childType),
		)
	}
	if newNode.Contains(n.getSelf()) {
		return newDomError("May not add a parent as a child")
	}
	if childType == NodeTypeText && parentType == NodeTypeDocument {
		return newDomError("Text nodes may not be direct descendants of a document")
	}
	if childType == NodeTypeDocumentType && parentType != NodeTypeDocument {
		return newDomError("Document type may only be a parent of Document")
	}
	if doc, isDoc := n.getSelf().(Document); isDoc {
		if doc.ChildElementCount() > 0 {
			return newDomError("Document can have only one child element")
		}
		if fragment, isFragment := newNode.(DocumentFragment); isFragment {
			if fragment.ChildElementCount() > 0 {
				return newDomError("Document can have only one child element")
			}
			for _, n := range fragment.ChildNodes().All() {
				if n.NodeType() == NodeTypeText {
					return newDomError("Text nodes may not be direct descendants of a document")
				}
			}
		}
	}
	return nil
}

func (n *node) childElements() []Element {
	nodes := n.childNodes.All()
	res := make([]Element, 0, len(nodes))
	for _, n := range nodes {
		if e, ok := n.(Element); ok {
			res = append(res, e)
		}
	}
	return res
}

func (n *node) insertBefore(newNode Node, referenceNode Node) (Node, error) {
	if referenceNode == nil {
		n.childNodes.append(newNode)
	} else {
		i := slices.Index(n.childNodes.All(), referenceNode)
		if i == -1 {
			return nil, errors.New("Reference node not found")
		}
		n.childNodes.setNodes(slices.Insert(n.childNodes.All(), i, newNode))
	}
	removeNodeFromParent(newNode)
	n.notify(n.addedNodeEvent(newNode))
	return newNode, nil
}

type nodeIterator struct{ Node }

func toHtmlNode(node Node) *html.Node {
	return nodeIterator{node}.toHtmlNode(nil)
}
func toHtmlNodeAndMap(node Node) (*html.Node, map[*html.Node]Node) {
	m := make(map[*html.Node]Node)
	result := nodeIterator{node}.toHtmlNode(m)
	return result, m
}

func (n nodeIterator) toHtmlNode(m map[*html.Node]Node) *html.Node {
	htmlNode := n.Node.createHtmlNode()
	if m != nil {
		m[htmlNode] = n.Node
	}
	for _, child := range n.nodes() {
		htmlNode.AppendChild(nodeIterator{child}.toHtmlNode(m))
	}
	return htmlNode
}

func (n *node) OwnerDocument() Document {
	if _, isDoc := n.getSelf().(Document); isDoc {
		return nil
	}
	return n.nodeDocument()
}

func (n *node) nodeDocument() Document {
	return n.document
}

func (n *node) FirstChild() Node {
	if n.childNodes.Length() == 0 {
		return nil
	}
	return n.childNodes.Item(0)
}

func (n *node) NextSibling() Node {
	children := n.Parent().nodes()
	idx := slices.IndexFunc(
		children,
		func(child Node) bool { return n.ObjectId() == child.ObjectId() },
	) + 1
	if idx == 0 {
		panic("We should exist in our parent's collection")
	}
	if idx >= len(children) {
		return nil
	}
	return children[idx]
}
func (n *node) PreviousSibling() Node {
	children := n.Parent().nodes()
	idx := slices.IndexFunc(
		children,
		func(child Node) bool { return n.ObjectId() == child.ObjectId() },
	) - 1
	if idx == -2 {
		panic("We should exist in our parent's collection")
	}
	if idx < 0 {
		return nil
	}
	return children[idx]
}

func (n *node) nodes() []Node {
	return n.childNodes.All()
}

func (n *node) SetSelf(node Node) {
	n.self = node
	event.SetEventTargetSelf(node)
	if doc, ok := node.(Document); ok {
		n.document = doc
	}
}

func (n *node) getSelf() Node { return n.self }

func (n *node) SetTextContent(val string) {
	for x := n.FirstChild(); x != nil; x = n.FirstChild() {
		n.RemoveChild(x)
	}
	n.AppendChild(NewText(val, n.OwnerDocument()))
}

func (n *node) TextContent() string {
	b := &strings.Builder{}
	for _, node := range n.nodes() {
		b.WriteString(node.TextContent())
	}
	return b.String()
}

func (n *node) renderChildren(builder *strings.Builder) {
	if childRenderer, ok := n.self.(ChildrenRenderer); ok {
		childRenderer.RenderChildren(builder)
	}
}

func (n *node) RenderChildren(builder *strings.Builder) {
	for _, child := range n.childNodes.All() {
		if renderer, ok := child.(Renderer); ok {
			renderer.Render(builder)
		}
	}
}

func (n *node) String() string { return n.self.NodeName() }

func (n *node) notify(event ChangeEvent) {
	for _, o := range n.observers {
		o.Process(event)
	}
	if n.parent != nil {
		n.parent.notify(event)
	}
}

func (n *node) addedNodeEvent(newNode ...Node) ChangeEvent {
	return ChangeEvent{
		Target:     n.self,
		AddedNodes: &nodeList{nodes: newNode},
	}
}

func (n *node) removedNodeEvent(nodes ...Node) ChangeEvent {
	return ChangeEvent{
		Target:       n.self,
		RemovedNodes: &nodeList{nodes: nodes},
	}
}

func (n *node) Logger() *slog.Logger {
	if docLogger, ok := n.document.(log.LogSource); ok {
		return docLogger.Logger()
	}
	return nil
}
