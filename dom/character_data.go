package dom

import (
	"strings"
	"unicode/utf8"

	intdom "github.com/gost-dom/browser/internal/dom"

	"golang.org/x/net/html"
)

/* -------- CharacterData -------- */

// CharacterData is a "base type" for [Text], [Comment], and
// [CDataSection], and [ProcessingInstruction].
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
type CharacterData interface {
	Node
	ChildNode
	Data() string
	SetData(string)
	Length() int
}

type characterData struct {
	node
	childNode
	data string
}

func newCharacterData(text string, ownerDocument Document, typ NodeType) *characterData {
	res := &characterData{node: newNode(ownerDocument, typ), data: text}
	res.childNode = childNode{&res.node}
	return res
}

func (n *characterData) NodeValue() (string, bool) { return n.Data(), true }
func (n *characterData) SetNodeValue(val string)   { n.SetData(val) }

func (n *characterData) Data() string {
	return n.data
}

func (d *characterData) SetData(data string) {
	oldValue := d.data
	d.data = data
	d.notify(ChangeEvent{
		Target:   d.self(),
		Type:     ChangeEventCData,
		OldValue: oldValue,
	})
}

func (n *characterData) Length() int {
	return utf8.RuneCountInString(n.data)
}

func (d *characterData) cloneChildren() []Node { return nil }

func (d *characterData) IsEqualNode(other Node) bool {
	if od, ok := other.(CharacterData); ok {
		return od.Data() == d.data && d.node.isEqualNode(other)
	}
	return false
}

/* -------- Comment -------- */

type Comment interface {
	CharacterData
}

type comment struct {
	*characterData
}

func NewComment(text string, ownerDocument Document) Comment {
	result := &comment{newCharacterData(text, ownerDocument, intdom.NodeTypeComment)}
	result.SetSelf(result)
	return result
}

func (n *comment) Render(builder *strings.Builder) {
	builder.WriteString("<!--")
	builder.WriteString(n.Data())
	builder.WriteString("-->")
}

func (n *comment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.CommentNode,
		Data: n.Data(),
	}
}

func (c *comment) cloneNode(doc Document, _ bool) Node {
	return NewComment(c.TextContent(), doc)
}

/* -------- Text -------- */

type Text interface {
	CharacterData
}

type textNode struct {
	*characterData
}

func NewText(text string, ownerDocument Document) Text {
	result := &textNode{newCharacterData(text, ownerDocument, intdom.NodeTypeText)}
	result.SetSelf(result)
	return result
}

func (n *textNode) cloneNode(doc Document, _ bool) Node {
	return NewText(n.characterData.data, doc)
}

func (n *textNode) Render(builder *strings.Builder) {
	builder.WriteString(n.Data())
}

func (n *textNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: n.Data(),
	}
}

func (n *textNode) NodeName() string    { return "#text" }
func (n *textNode) TextContent() string { return n.Data() }

/* -------- ProcessingInstruction -------- */

// ProcessingInstruction represents an XML processing instruction.
//
// Deprecated: This only exists to support Web Platform Tests. It may be
// removed.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/ProcessingInstruction
type ProcessingInstruction interface {
	CharacterData
}

type processingInstruction struct {
	*characterData
	target string
}

// NewProcessingInstruction creates a new ProcessingInstruction.
//
// Deprecated: This only exists to support Web Platform Tests
func NewProcessingInstruction(target, data string, ownerDocument Document) ProcessingInstruction {
	result := &processingInstruction{
		newCharacterData(data, ownerDocument, intdom.NodeTypeProcessingInstruction),
		target,
	}
	result.SetSelf(result)
	return result
}

func (n *processingInstruction) cloneNode(doc Document, _ bool) Node {
	return NewProcessingInstruction(n.target, n.characterData.data, doc)
}

func (n *processingInstruction) Render(builder *strings.Builder) {
	builder.WriteString("<?")
	builder.WriteString(n.target)
	builder.WriteString(" ")
	builder.WriteString(n.Data())
	builder.WriteString("?>")
}

func (n *processingInstruction) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.RawNode,
		Data: n.Data(),
	}
}

func (n *processingInstruction) NodeName() string { return n.target }

func (n *processingInstruction) TextContent() string { return "" }
