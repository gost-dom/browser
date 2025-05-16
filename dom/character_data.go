package dom

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

/* -------- CharacterData -------- */

// CharacterData is a "base type" for [Text], [Comment], and
// [CDataSection], and [ProcessingInstruction].
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
type CharacterData interface {
	Node
	Data() string
	SetData(string)
	Length() int
}

type characterData struct {
	node
	data string
}

func (n *characterData) Data() string {
	return n.data
}

func (d *characterData) SetData(data string) {
	oldValue := d.data
	d.data = data
	d.notify(ChangeEvent{
		Target:   d.self,
		Type:     ChangeEventCData,
		OldValue: oldValue,
	})
}

func (n *characterData) Length() int {
	return utf8.RuneCountInString(n.data)
}

func (d *characterData) cloneChildren() []Node { return nil }

/* -------- Comment -------- */

type Comment interface {
	CharacterData
}

type comment struct {
	characterData
}

func NewComment(text string, ownerDocument Document) Comment {
	result := &comment{characterData{newNode(ownerDocument), text}}
	result.SetSelf(result)
	return result
}

func (n *comment) Render(builder *strings.Builder) {
	builder.WriteString("<!--")
	builder.WriteString(n.Data())
	builder.WriteString("-->")
}

func (n *comment) NodeType() NodeType {
	return NodeTypeComment
}

func (n *comment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.CommentNode,
		Data: n.Data(),
	}
}

func (c *comment) CloneNode(bool) Node {
	return NewComment(c.TextContent(), c.OwnerDocument())
}

/* -------- Text -------- */

type Text interface {
	CharacterData
}

type textNode struct {
	characterData
}

func NewText(text string, ownerDocument Document) Text {
	result := &textNode{characterData{newNode(ownerDocument), text}}
	result.SetSelf(result)
	return result
}

func (n *textNode) CloneNode(bool) Node {
	return NewText(n.characterData.data, n.OwnerDocument())
}

func (n *textNode) Render(builder *strings.Builder) {
	builder.WriteString(n.Data())
}

func (n *textNode) NodeType() NodeType { return NodeTypeText }

func (n *textNode) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.TextNode,
		Data: n.Data(),
	}
}

func (n *textNode) NodeName() string    { return "#text" }
func (n *textNode) TextContent() string { return n.Data() }
