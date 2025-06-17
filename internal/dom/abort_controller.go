package dom

import "github.com/gost-dom/browser/dom/event"

type AbortController struct{}

func (c AbortController) Abort(any)           {}
func (c AbortController) Signal() AbortSignal { return AbortSignal{} }

type AbortSignal struct {
	event.EventTarget
}

func (s AbortSignal) Aborted() bool                 { return false }
func (s AbortSignal) Reason() any                   { return nil }
func (s AbortSignal) Onabort() event.EventHandler   { return nil }
func (s AbortSignal) SetOnabort(event.EventHandler) {}
func (s AbortSignal) ThrowIfAborted() error         { return nil }
