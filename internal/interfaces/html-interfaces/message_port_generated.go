// This file is generated. Do not edit.

package htmlinterfaces

import event "github.com/gost-dom/browser/dom/event"

type MessagePort interface {
	event.EventTarget

	Onclose() event.EventHandler
	SetOnclose(event.EventHandler)
	Start()
	Close()
}
