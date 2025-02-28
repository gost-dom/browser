package uievents

import "github.com/gost-dom/browser/dom"

type UIEvent struct {
	dom.Event
}

func NewUIEvent(type_ string) UIEvent {
	return UIEvent{dom.NewEvent(type_)}
}
