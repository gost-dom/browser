// This file is generated. Do not edit.

package htmlinterfaces

import (
	event "github.com/gost-dom/browser/dom/event"
	entity "github.com/gost-dom/browser/internal/entity"
)

type MessagePort interface {
	entity.Components
	event.EventTarget

	Onclose() event.EventHandler
	SetOnclose(event.EventHandler)
	Start()
	Close()
}
