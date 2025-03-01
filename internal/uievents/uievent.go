package uievents

import "github.com/gost-dom/browser/dom"

type UIEvent struct {
	dom.Event
}

func NewUIEvent(type_ string) UIEvent {
	return UIEvent{dom.NewEvent(type_)}
}

type PointerEvent = UIEvent
type MouseEvent = UIEvent

var NewMouseEvent = NewUIEvent
var NewPointerEvent = NewUIEvent
