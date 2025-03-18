// This file is generated. Do not edit.

package dominterfaces

import dom "github.com/gost-dom/browser/dom"

type MutationObserver interface {
	Observe(dom.Node)
	ObserveOptions(dom.Node, MutationObserverInit)
	Disconnect()
	TakeRecords() []MutationRecord
}
