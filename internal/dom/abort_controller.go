package dom

import "github.com/gost-dom/browser/dom/event"

const EventTypeAbort = "abort"

type AbortController struct {
	signal AbortSignal
}

func NewAbortController() *AbortController {
	return &AbortController{signal: newAbortSignal()}
}

func (c *AbortController) Signal() *AbortSignal { return &c.signal }

func (c *AbortController) Abort(reason any) {
	c.signal.aborted = true
	c.signal.reason = reason
	c.signal.DispatchEvent(&event.Event{Type: EventTypeAbort})
}

type AbortSignal struct {
	event.EventTarget
	aborted bool
	reason  any
}

func newAbortSignal() AbortSignal {
	return AbortSignal{EventTarget: event.NewEventTarget()}
}

func (s AbortSignal) Aborted() bool                 { return s.aborted }
func (s AbortSignal) Reason() any                   { return s.reason }
func (s AbortSignal) Onabort() event.EventHandler   { return nil }
func (s AbortSignal) SetOnabort(event.EventHandler) {}
func (s AbortSignal) ThrowIfAborted() error         { return nil }
