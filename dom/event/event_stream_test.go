package event_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"
)

const DEFAULT_BUF = 16

type EventChan chan *event.Event

type EventStream struct {
	EventChan
}

func NewEventStream(tgt event.EventTarget, t string, ctx context.Context) EventChan {
	c := make(chan *event.Event)
	handler := event.NewEventHandlerFunc(func(e *event.Event) error {
		go func() { c <- e }()
		return nil
	})
	tgt.AddEventListener(t, handler)
	go func() {
		<-ctx.Done()
		tgt.RemoveEventListener(t, handler)
	}()

	return c
}

func TestEventsAreReceivedInOrder(t *testing.T) {
	t.Parallel()
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		src := EventSource{event.NewEventTarget()}
		const buf = 32
		c := src.Listen(ctx, "gost-event", BufSize(buf))

		for i := range buf << 1 {
			// Dispatch twice as many events as the buffer size
			src.DispatchEvent(event.New("gost-event", i))
		}
		synctest.Wait()

		// Verify the events in the scope of the buffer size
		var i int
		events := make([]*event.Event, buf)
		for e := range c {
			events[i] = e
			i++
			if i == buf {
				break
			}
		}
		for i := range buf {
			e := events[i]
			assert.Equal(t, i, e.Data, "Event data at index: %d", i)
		}
	})
}

func TestEventStreamSource(t *testing.T) {
	t.Parallel()

	target := event.NewEventTarget()

	t.Run("RemoveEventListener on cancel", func(t *testing.T) {
		synctest.Run(func() {
			newCtx, cancel := context.WithCancel(t.Context())

			spy := &EventTargetSpy{EventTarget: target}
			src := EventSource{spy}
			events := src.Listen(newCtx, "gost-event")

			assert.Equal(t, 1, spy.addCallCount)
			assert.Equal(t, 0, spy.removeCallCount)

			cancel()
			synctest.Wait()

			assert.Equal(t, 1, spy.addCallCount)
			assert.Equal(t, 1, spy.removeCallCount)

			Expect(t, events).To(BeClosed())
		})
	})

	t.Run("Events are delivered in order", func(t *testing.T) {
		synctest.Run(func() {
			ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			defer cancel()

			spy := &EventTargetSpy{EventTarget: target}
			src := &EventSource{spy}

			events := src.Listen(ctx, "gost-event")

			e := event.New("gost-event", nil)
			target.DispatchEvent(e)

			var got *event.Event

			select {
			case got = <-events:
			case <-ctx.Done():
				t.Fatal("timeout")
			}

			assert.Same(t, e, got)
		})
	})
}

type EventTargetSpy struct {
	event.EventTarget

	addCallCount    int
	removeCallCount int
}

func (s *EventTargetSpy) AddEventListener(
	t string,
	h event.EventHandler,
	opts ...event.EventListenerOption,
) {
	s.addCallCount++
	s.EventTarget.AddEventListener(t, h, opts...)
}

func (s *EventTargetSpy) RemoveEventListener(
	t string,
	h event.EventHandler,
	opts ...event.EventListenerOption,
) {
	s.removeCallCount++
	s.EventTarget.RemoveEventListener(t, h, opts...)
}

// EventSource embeds an [EventTarget] and provides events in a channel,
// simplifying Go code consuming events.
type EventSource struct {
	event.EventTarget
}

type eventSourceOptions struct {
	buf int
}

type EventSourceOption func(*eventSourceOptions)

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
// controlled with the [BufSize] option. Default value is [DEFAULT_BUF].
func (s EventSource) Listen(
	ctx context.Context,
	t string,
	opts ...EventSourceOption,
) <-chan *event.Event {
	opt := eventSourceOptions{buf: DEFAULT_BUF}
	for _, o := range opts {
		o(&opt)
	}
	c := make(chan *event.Event, opt.buf)
	handler := event.NewEventHandlerFunc(func(e *event.Event) error {
		// It is assumed that all events are dispatched from the same
		// goroutine. If the buffer permits, send as a blocking call,
		// ensuring ordering of events.
		//
		// If the channel is not ready to accept messages, send in a new
		// goroutine to avoid blocking the EventTarget
		select {
		case c <- e:
		default:
			go func() { c <- e }()
		}
		return nil
	})

	s.EventTarget.AddEventListener(t, handler)

	if ctx != nil {
		// If no context is provided, the event listener is never removed.
		go func() {
			<-ctx.Done()
			s.EventTarget.RemoveEventListener(t, handler)
			close(c)
		}()
	}

	return c
}
