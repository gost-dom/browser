package dom

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"

	intdom "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"golang.org/x/net/html"
)

func getNode(n *intdom.Node) Node {
	if n == nil {
		return nil
	}
	if res, ok := entity.ComponentType[Node](n); ok {
		return res
	}
	return nil
}

type NodeType = intdom.NodeType

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

	self() Node
	setParent(Node)
	setOwnerDocument(owner Document)
	assertCanAddNode(Node) error
	cloneChildren() []Node
	createHtmlNode() *html.Node
	nodeDocument() Document
	notify(ChangeEvent)
	cloneNode(Document, bool) Node
	revision() int // Internal optimizations
	ptr() *node
}

type node struct {
	event.EventTarget
	*intdom.Node
	// revision of the node is incremented on any change. Used by
	// LiveHtmlCollection to check if a note has been changed.
	rev       int
	document  Document
	observers []observer
}

func newNode(ownerDocument Document) node {
	return node{
		Node:        intdom.NewNode(),
		EventTarget: event.NewEventTarget(),
		document:    ownerDocument,
	}
}

func (n *node) CloneNode(deep bool) Node {
	return n.self().cloneNode(n.OwnerDocument(), deep)
}

func (n *node) cloneChildren() []Node {
	nodes := make([]Node, len(n.Children))
	for i, n := range n.Children {
		nodes[i] = getNode(n).CloneNode(true)
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
	_, err := n.self().InsertBefore(node, nil)
	return node, err
}

func (n *node) NodeValue() (string, bool) { return "", false }
func (n *node) SetNodeValue(string)       {}

func (n *node) InsertBefore(newChild Node, referenceNode Node) (Node, error) {
	if err := n.assertCanAddNode(newChild); err != nil {
		return nil, err
	}
	err := n.insertBefore(newChild, referenceNode)
	return newChild, err
}

func (n *node) ChildNodes() NodeList {
	if res, ok := entity.ComponentType[NodeList](n); ok {
		return res
	}

	res := newDynamicNodeList(&(n.Node.Children))
	entity.SetComponentType(n, res)
	return res
}

func (n *node) Observe(observer observer) Closer {
	if slices.Contains(n.observers, observer) {
		panic("Observer already added to this node")
	}
	n.observers = append(n.observers, observer)
	return observerCloser{n, observer}
}

func (n *node) IsEqualNode(other Node) bool { return n.isEqualNode(other) }

func (n *node) isEqualNode(other Node) bool {
	if n.self().NodeType() != other.NodeType() {
		return false
	}
	l := len(n.Children)
	cmp := other.ChildNodes()
	if l != cmp.Length() {
		return false
	}
	for i := range l {
		if !getNode(n.Children[i]).IsEqualNode(cmp.Item(i)) {
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
	if n.Parent == nil {
		return n.self()
	} else {
		return getNode(n.Parent).GetRootNode(options...)
	}
}

func (n *node) Contains(node Node) bool {
	if n.self() == node {
		return true
	}
	for _, c := range n.Children {
		if c == node.ptr().Node || getNode(c).Contains(node) {
			return true
		}
	}
	return false
}

func (n *node) ParentNode() Node { return getNode(n.Parent) }

func (n *node) ParentElement() Element {
	r, _ := n.ParentNode().(Element)
	return r
}

func (n *node) setOwnerDocument(owner Document) {
	n.document = owner
	for _, n := range n.Children {
		getNode(n).setOwnerDocument(owner)
	}
}
func (n *node) setParent(parent Node) {
	if n.Parent != nil {
		getNode(n.Parent).RemoveChild(n.self())
	}
	if parent != nil {
		parentOwner := parent.nodeDocument()
		if n.document != parentOwner {
			n.setOwnerDocument(parentOwner)
		}
		n.Parent = parent.ptr().Node
	} else {
		n.Parent = nil
	}
	n.SetParentTarget(parent)
}

func (n *node) Connected() {
	if p := n.self().ParentNode(); p != nil {
		p.Connected()
	}
}

func (n *node) IsConnected() (result bool) {
	if n.Parent != nil {
		result = getNode(n.Parent).IsConnected()
	}
	return
}

func (n *node) IsSameNode(other Node) (result bool) {
	return n.self() == other
}

func (n *node) NodeName() string {
	return "#node"
}

func (n *node) RemoveChild(node Node) (Node, error) {
	idx := slices.Index(n.Children, node.ptr().Node)
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
	parentType := n.self().NodeType()
	childType := newNode.NodeType()
	if !parentType.CanHaveChildren() {
		return newDomErrorCode(
			fmt.Sprintf("May not add children to node type %s", parentType),
			hierarchy_request_err,
		)
	}
	if !childType.CanBeAChild() {
		return newDomErrorCode(
			fmt.Sprintf("May not add an node type %s as a child", childType), hierarchy_request_err,
		)
	}
	if newNode.Contains(n.self()) {
		return newDomError("May not add a parent as a child")
	}
	if childType == intdom.NodeTypeText && parentType == intdom.NodeTypeDocument {
		return newDomErrorCode(
			"Text nodes may not be direct descendants of a document", hierarchy_request_err,
		)
	}
	if childType == intdom.NodeTypeDocumentType && parentType != intdom.NodeTypeDocument {
		return newDomError("Document type may only be a parent of Document")
	}
	if doc, isDoc := n.self().(Document); isDoc {
		if doc.ChildElementCount() > 0 {
			return newDomErrorCode(
				"Document can have only one child element",
				hierarchy_request_err,
			)
		}
		if fragment, isFragment := newNode.(DocumentFragment); isFragment {
			if fragment.ChildElementCount() > 0 {
				return newDomErrorCode(
					"Document can have only one child element", hierarchy_request_err)
			}
			for n := range fragment.ChildNodes().All() {
				if n.NodeType() == intdom.NodeTypeText {
					return newDomErrorCode(
						"Text nodes may not be direct descendants of a document",
						hierarchy_request_err,
					)
				}
			}
		}
	}
	return nil
}

func (n *node) childElements() []Element {
	res := make([]Element, 0, len(n.Children))
	for _, n := range n.Children {
		if e, ok := getNode(n).(Element); ok {
			res = append(res, e)
		}
	}
	return res
}

func (n *node) insertBefore(node Node, referenceNode Node) error {
	var index int
	if referenceNode == nil {
		index = len(n.Children)
	} else {
		index = slices.Index(n.Children, referenceNode.ptr().Node)
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
	children := slices.Clone(n.Children)
	start := index
	end := index + count
	if start > 0 {
		prevSibling = getNode(children[start-1])
	}
	if end < len(children) {
		nextSibling = getNode(children[end])
	}

	var searchNode *intdom.Node
	if node != nil {
		searchNode = node.ptr().Node
	}
	currentIdx := slices.Index(children, searchNode)
	removedNodes := slices.Clone(children[start:end])
	children = slices.Replace(children, start, end, newNodes...)
	sameParent := node != nil && node.ParentNode() == n.self()
	if sameParent {
		if currentIdx == -1 {
			panic("replaceNodes: bad state - node has no index in it's parent collection")
		}
		if currentIdx >= end {
			currentIdx = currentIdx - count + 1
		}
		children = slices.Delete(children, currentIdx, currentIdx+1)
	}
	n.Children = children

	if !sameParent {
		for _, node := range removedNodes {
			getNode(node).setParent(nil)
		}
		for _, node := range newNodes {
			getNode(node).setParent(n.self())
			if getNode(node).IsConnected() {
				getNode(node).Connected()
			}
		}
	}

	n.notify(ChangeEvent{
		Target:          n.self(),
		Type:            ChangeEventChildList,
		PreviousSibling: prevSibling,
		NextSibling:     nextSibling,
		AddedNodes:      newStaticNodeList(newNodes),
		RemovedNodes:    newStaticNodeList(removedNodes),
	})
	return nil
}

func (n *node) ReplaceChild(node, child Node) (Node, error) {
	nodeType := n.self().NodeType()
	if nodeType != intdom.NodeTypeDocument &&
		nodeType != intdom.NodeTypeDocumentFragment &&
		nodeType != intdom.NodeTypeElement {
		return nil, newDomError("HierarchyRequestError")
	}
	if child.ParentNode() != n.self() {
		return nil, newDomError("NotFoundError")
	}
	for i, c := range n.Children {
		if getNode(c) == child {
			err := n.replaceNodes(i, 1, node)
			return getNode(c), err
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
	if _, isDoc := n.self().(Document); isDoc {
		return nil
	}
	return n.nodeDocument()
}

func (n *node) nodeDocument() Document {
	return n.document
}

func (n *node) FirstChild() Node {
	if len(n.Children) == 0 {
		return nil
	}
	return getNode(n.Children[0])
}

func (n *node) LastChild() Node {
	l := len(n.Children)
	if l == 0 {
		return nil
	}
	return getNode(n.Children[l-1])
}

func (n *node) NextSibling() Node {
	parent := n.ParentNode()
	if parent == nil {
		return nil
	}
	children := parent.ptr().Children
	idx := slices.IndexFunc(
		children,
		func(c *intdom.Node) bool { return n.IsSameNode(getNode(c)) },
	) + 1
	if idx == -1 {
		panic("We should exist in our parent's collection")
	}
	if idx >= len(children) {
		return nil
	}
	return getNode(children[idx])
}

func (n *node) PreviousSibling() Node {
	parent := n.ParentNode()
	if parent == nil {
		return nil
	}
	children := parent.ptr().Children
	idx := slices.IndexFunc(children, func(c *intdom.Node) bool { return n.IsSameNode(getNode(c)) })
	if idx == -1 {
		panic("We should exist in our parent's collection")
	}
	if idx == 0 {
		return nil
	}
	return getNode(children[idx-1])
}

func (n *node) ptr() *node { return n }

func (n *node) SetSelf(node Node) {
	entity.SetComponentType(n, node)
	event.SetEventTargetSelf(node)
	if doc, ok := node.(Document); ok {
		n.document = doc
	}
}

func (n *node) self() Node {
	res, _ := entity.ComponentType[Node](n)
	return res
}

func (n *node) SetTextContent(val string) {
	for x := n.FirstChild(); x != nil; x = n.FirstChild() {
		n.RemoveChild(x)
	}
	n.AppendChild(NewText(val, n.OwnerDocument()))
}

func (n *node) TextContent() string {
	b := &strings.Builder{}
	for _, node := range n.Children {
		b.WriteString(getNode(node).TextContent())
	}
	return b.String()
}

func (n *node) renderChildren(builder *strings.Builder) {
	childRenderer, ok := entity.ComponentType[intdom.ChildrenRenderer](n)
	if !ok {
		childRenderer, ok = n.self().(intdom.ChildrenRenderer)
	}
	if ok {
		childRenderer.RenderChildren(builder)
	}
}

func (n *node) RenderChildren(builder *strings.Builder) {
	for _, child := range n.Children {
		if renderer, ok := getNode(child).(intdom.Renderer); ok {
			renderer.Render(builder)
		}
	}
}

func (n *node) String() string { return n.self().NodeName() }

func (n *node) revision() int { return n.rev }

func (n *node) notify(event ChangeEvent) {
	n.rev++
	for _, o := range n.observers {
		o.Process(event)
	}
	if n.Parent != nil {
		getNode(n.Parent).notify(event)
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
	for _, child := range n.ptr().Children {
		htmlNode.AppendChild(nodeIterator{getNode(child)}.toHtmlNode(m))
	}
	return htmlNode
}
