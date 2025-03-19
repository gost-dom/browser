// This file is generated. Do not edit.

package dominterfaces

import dom "github.com/gost-dom/browser/dom"

type MutationObserver interface {
	Observe(dom.Node, ...func(*MutationObserverInit)) error
	Disconnect()
	TakeRecords() []MutationRecord
}
