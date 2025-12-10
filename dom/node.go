package dom

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
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

type ChangeEventType string

const (
	ChangeEventChildList  ChangeEventType = "childList"
	ChangeEventAttributes ChangeEventType = "attributes"
	ChangeEventCData      ChangeEventType = "characterData"
)

type ChangeEvent struct {
	// The original target of the change.
	Target          Node
	Attr            Attr
	Type            ChangeEventType
	AddedNodes      NodeList
	RemovedNodes    NodeList
	OldValue        string
	PreviousSibling Node
	NextSibling     Node
}

type observer interface {
	Process(ChangeEvent)
}

type Node interface {
	entity.ObjectIder
	entity.Components
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
	NodeValue() (string, bool)
	SetNodeValue(string)
	OwnerDocument() Document
	ParentNode() Node
	ParentElement() Element
	RemoveChild(node Node) (Node, error)
	ReplaceChild(node, child Node) (Node, error)
	NextSibling() Node
	PreviousSibling() Node
	FirstChild() Node
	LastChild() Node
	TextContent() string
	SetTextContent(value string)
	Connected()
	IsEqualNode(Node) bool
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
	cloneNode(Document, bool) Node
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

func (n *node) CloneNode(deep bool) Node {
	return n.self.cloneNode(n.OwnerDocument(), deep)
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
	n.Logger().Debug("Node.AppendChild", "target", n.String(), "child", node.NodeName())
	_, err := n.self.InsertBefore(node, nil)
	return node, err
}

func (n *node) NodeValue() (string, bool) { return "", false }
func (n *node) SetNodeValue(string)       {}

func (n *node) InsertBefore(newChild Node, referenceNode Node) (Node, error) {
	if err := n.assertCanAddNode(newChild); err != nil {
		return nil, err
	}
	if fragment, ok := newChild.(DocumentFragment); ok {
		return fragment, n.insertBefore(fragment, referenceNode)
	}
	err := n.insertBefore(newChild, referenceNode)
	return newChild, err
}

func (n *node) ChildNodes() NodeList { return n.childNodes }

func (n *node) Observe(observer observer) Closer {
	if slices.Contains(n.observers, observer) {
		panic("Observer already added to this node")
	}
	n.observers = append(n.observers, observer)
	return observerCloser{n, observer}
}

func (n *node) IsEqualNode(other Node) bool { return n.isEqualNode(other) }

func (n *node) isEqualNode(other Node) bool {
	if n.self.NodeType() != other.NodeType() {
		return false
	}
	if n.childNodes.Length() != other.ChildNodes().Length() {
		return false
	}
	for i := 0; i < n.childNodes.Length(); i++ {
		if !n.childNodes.Item(i).IsEqualNode(other.ChildNodes().Item(i)) {
			return false
		}
	}
	return true
}

func (n *node) removeObserver(o observer) {
	n.observers = slices.DeleteFunc(n.observers, func(x observer) bool { return o == x })
}

func (n *node) GetRootNode(options ...GetRootNodeOptions) Node {
	if len(options) > 1 {
		n.Logger().Warn("Node.GetRootNode: composed not yet implemented")
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

func (n *node) ParentNode() Node { return n.parent }

func (n *node) ParentElement() Element {
	r, _ := n.ParentNode().(Element)
	return r
}

func (n *node) setOwnerDocument(owner Document) {
	n.document = owner
	for _, n := range n.ChildNodes().All() {
		n.setOwnerDocument(owner)
	}
}
func (n *node) setParent(parent Node) {
	if n.parent != nil {
		n.parent.RemoveChild(n.getSelf())
	}
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
	if p := n.getSelf().ParentNode(); p != nil {
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

func (n *node) RemoveChild(node Node) (Node, error) {
	idx := slices.Index(n.childNodes.All(), node)
	if idx == -1 {
		return nil, newDomError(
			"Node.removeChild: The node to be removed is not a child of this node",
		)
	}
	return node, n.replaceNodes(idx, 1, nil)
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
	if newNode == nil {
		return nil
	}
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

func (n *node) insertBefore(node Node, referenceNode Node) error {
	var index int
	if referenceNode == nil {
		index = n.childNodes.Length()
	} else {
		index = slices.Index(n.childNodes.All(), referenceNode)
		if index == -1 {
			return errors.New("Reference node not found")
		}
	}
	return n.replaceNodes(index, 0, node)
}

// replaceNodes removes count nodes from index, and inserts the content of node
// at the position.
//
// replaceNodes does not validate the node, and should only be called by an
// exported function that
//
// Which new nodes depend on the node. If node is nil,
// the specified nodes will be removed, and no new nodes inserted. If node is a
// [DocumentFragment], the children of the fragment is inserted. Otherwise, a
// single node is inserted
//
// If count is zero, the new node is inserted before the node at the specified
// index if index < len(n.childNodes.Length()), otherwise it is appended after
// the last element.
//
// replaceNodes panics if index < 0 or index + count > len(n.childNodes.Length()).
func (n *node) replaceNodes(index, count int, node Node) error {
	if err := n.assertCanAddNode(node); err != nil {
		return err
	}

	var (
		prevSibling Node
		nextSibling Node
	)
	newNodes := expandNode(node)
	children := slices.Clone(n.ChildNodes().All())
	if index > 0 {
		prevSibling = children[index-1]
	}
	if index+count < len(children) {
		nextSibling = children[index+count]
	}

	removedNodes := slices.Clone(children[index : index+count])
	children = slices.Replace(children, index, index+count, newNodes...)
	n.childNodes.setNodes(children)

	for _, node := range removedNodes {
		node.setParent(nil)
	}
	for _, node := range newNodes {
		node.setParent(n.self)
		if node.IsConnected() {
			node.Connected()
		}
	}

	n.notify(ChangeEvent{
		Target:          n.self,
		Type:            ChangeEventChildList,
		PreviousSibling: prevSibling,
		NextSibling:     nextSibling,
		AddedNodes:      &nodeList{nodes: newNodes},
		RemovedNodes:    &nodeList{nodes: removedNodes},
	})
	return nil

}

func (n *node) ReplaceChild(node, child Node) (Node, error) {

	_, isDocument := n.self.(Document)
	_, isDocumentFragment := n.self.(DocumentFragment)
	_, isElement := n.self.(Element)
	if !isDocument && !isDocumentFragment && !isElement {
		return nil, newDomError("HierarchyRequestError")
	}
	if child.ParentNode() != n.self {
		return nil, newDomError("NotFoundError")
	}
	for i, c := range n.childNodes.All() {
		if c == child {
			err := n.replaceNodes(i, 1, node)
			return c, err
		}
	}
	panic(
		fmt.Sprintf(
			"gost-dom/dom: ReplaceChild: child not found in child list\n. %s",
			constants.BUG_ISSUE_URL,
		),
	)
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

func (n *node) LastChild() Node {
	l := n.childNodes.Length()
	if l == 0 {
		return nil
	}
	return n.childNodes.Item(l - 1)
}

func (n *node) NextSibling() Node {
	children := n.ParentNode().nodes()
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
	parent := n.ParentNode()
	if parent == nil {
		return nil
	}
	children := parent.nodes()
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

func (n *node) Logger() (res *slog.Logger) {
	if docLogger, ok := n.document.(log.LogSource); ok {
		res = docLogger.Logger()
	}
	if res == nil {
		res = log.Default()
	}
	return res
}

/* -------- observerCloser -------- */

type observerCloser struct {
	n *node
	o observer
}

func (c observerCloser) Close() {
	c.n.removeObserver(c.o)
}

/* -------- nodeIterator -------- */

type nodeIterator struct{ Node }

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
