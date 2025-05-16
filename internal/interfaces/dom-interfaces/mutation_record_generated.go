// This file is generated. Do not edit.

package dominterfaces

import dom "github.com/gost-dom/browser/dom"

type MutationRecord struct {
	Type               string
	Target             dom.Node
	AddedNodes         dom.NodeList
	RemovedNodes       dom.NodeList
	PreviousSibling    dom.Node
	NextSibling        dom.Node
	AttributeName      *string
	AttributeNamespace *string
	OldValue           *string
}
