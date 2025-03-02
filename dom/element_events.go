package dom

type ElementEvents interface {
	Click() bool
}

type elementEvents struct {
	target EventTarget
}

func (n *elementEvents) Click() bool {
	return n.target.DispatchEvent(
		NewPointerEvent("click", EventCancelable(true), EventBubbles(true)),
	)
}
