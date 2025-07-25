package event

import "context"

// DefaultBuf is the default buffer size for event channels created by
// [EventSource] when buffer size not specified explicitly. Buffer size affects
// event ordering guarantees.
const DefaultBuf = 16

// EventSource embeds an [EventTarget] and provides events in a channel,
// simplifying Go code consuming events.
type EventSource struct {
	EventTarget
}

func NewEventSource(tgt EventTarget) EventSource { return EventSource{tgt} }

type eventSourceOptions struct {
	buf int
}

type EventSourceOption func(*eventSourceOptions)

// BufSize is an option to [EventSource.Listen], indicating the buffer size of
// the event channel returned.
func BufSize(buf int) EventSourceOption {
	return func(o *eventSourceOptions) { o.buf = buf }
}

// Listen adds an event listener for events of type t and returns a channel of
// events containing all the events. Cancelling context ctx will remove the
// event listener and close the channel. If no context is passed, the event
// listener will never be removed.
//
// Ordering of events is guaranteed when the channel buffer is not full and all
// events are dispatched from the same goroutine. The channel buffer size is
// controlled with the [BufSize] option. Default value is [DefaultBuf].
func (s EventSource) Listen(
	ctx context.Context,
	t string,
	opts ...EventSourceOption,
) <-chan *Event {
	if ctx == nil {
		panic("gost-dom/event: EventSource.Listen: ctx is nil")
	}
	opt := eventSourceOptions{buf: DefaultBuf}
	for _, o := range opts {
		o(&opt)
	}
	c := make(chan *Event, opt.buf)
	handler := NewEventHandlerFunc(func(e *Event) error {
		// It is assumed that all events are dispatched from the same
		// goroutine. If the buffer permits, send as a blocking call,
		// ensuring ordering of events.
		//
		// If the channel is not ready to accept messages, send in a new
		// goroutine to avoid blocking the EventTarget.
		//
		// TODO: Reconcider if the channel buffer should be a hard limit,
		// resulting in an error if the channel is not ready to accept new
		// messages
		select {
		case c <- e:
		default:
			go func() {
				select {
				case c <- e:
				case <-ctx.Done():
					// This is technically not needed because:
					//
					// - The function doesn't block
					// - The event handler is removed when the context cancels.
					//
					// But review tools complain about possible goroutine leaks - so
					// this just adds a line to fix that issue
				}
			}()
		}
		return nil
	})

	s.EventTarget.AddEventListener(t, handler)

	go func() {
		<-ctx.Done()
		s.EventTarget.RemoveEventListener(t, handler)
		close(c)
	}()

	return c
}
