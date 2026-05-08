// This file is generated. Do not edit.

package dominterfaces

import (
	"fmt"
	dom "github.com/gost-dom/browser/dom"
)

type Range interface {
	AbstractRange
	fmt.Stringer
	CommonAncestorContainer() dom.Node
	SetStart(dom.Node, int)
	SetEnd(dom.Node, int)
	SetStartBefore(dom.Node)
	SetStartAfter(dom.Node)
	SetEndBefore(dom.Node)
	SetEndAfter(dom.Node)
	Collapse(bool)
	SelectNode(dom.Node)
	SelectNodeContents(dom.Node)
	CompareBoundaryPoints(int, Range) int
	DeleteContents()
	ExtractContents() dom.DocumentFragment
	CloneContents() dom.DocumentFragment
	InsertNode(dom.Node)
	SurroundContents(dom.Node)
	CloneRange() Range
	Detach()
	IsPointInRange(dom.Node, int) bool
	ComparePoint(dom.Node, int) int
	IntersectsNode(dom.Node) bool
}
