// This file is generated. Do not edit.

package dominterfaces

import event "github.com/gost-dom/browser/dom/event"

type AbortSignal interface {
	event.EventTarget
	Aborted() bool
	Reason() any
	Onabort() event.EventHandler
	SetOnabort(event.EventHandler)
	ThrowIfAborted() error
}
