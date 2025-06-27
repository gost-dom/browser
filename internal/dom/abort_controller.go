package dom

import (
	"context"
	"fmt"

	"github.com/gost-dom/browser/dom/event"
)

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

// AbortContext will listen to abort events from an [event.EventTarget]. The
// return value is a child context of ctx which will be cancelled if a, abort
// event is dispatched before the parent context cancels.
//
// If the context is cancelled due to an abort event, the abort reason can be
// used as the cancel cause, which can be read using [context.Cause]. If the
// cause is not an error type, an [ErrReason] will be returned.
func AbortContext(ctx context.Context, signal event.EventTarget) context.Context {
	abortEvents := event.NewEventSource(signal).Listen(ctx, EventTypeAbort, event.BufSize(1))
	ctx, cancel := context.WithCancelCause(ctx)
	go func() {
		select {
		case e := <-abortEvents:
			err, ok := e.Data.(error)
			if !ok {
				err = ErrReason{Reason: e.Data}
			}
			cancel(err)
		case <-ctx.Done():
			cancel(nil)
		}
	}()
	return ctx
}

// ErrReason wraps the [AbortSignal.Abort] reason when the reason is not an
// error type. This is used when integrating with context cancellation causes
// that require the cause to be of an error type, but in JavaScript, any value
// is a valid cause.
type ErrReason struct{ Reason any }

func (err ErrReason) Error() string {
	return fmt.Sprintf("aborted: reason: %v", err.Reason)
}
