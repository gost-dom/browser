package event_test

import (
	"context"
	"fmt"
	"testing"
	"testing/synctest"
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/stretchr/testify/assert"
)

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

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	src := EventSource{event.NewEventTarget()}
	const buf = 32
	c := src.Listen(ctx, "gost-event", buf)

	for i := range buf << 1 {
		// Dispatch twice as many events as the buffer size
		src.DispatchEvent(event.New("gost-event", i))
	}

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
		fmt.Printf("Compare %t\n", e == nil)
		assert.Equal(t, i, e.Data, "Event data at index: %d", i)
	}
}

// func TestEventStreamIsClosedOnCancel(t *testing.T) {
// }

func TestEventStreamSource(t *testing.T) {
	t.Parallel()

	target := event.NewEventTarget()

	t.Run("RemoveEventListener on cancel", func(t *testing.T) {
		synctest.Run(func() {
			newCtx, cancel := context.WithCancel(t.Context())

			spy := &EventTargetSpy{EventTarget: target}
			src := EventSource{spy}
			events := src.Listen(newCtx, "gost-event", 1)

			assert.Equal(t, 1, spy.addCallCount)
			assert.Equal(t, 0, spy.removeCallCount)

			cancel()
			synctest.Wait()

			assert.Equal(t, 1, spy.addCallCount)
			assert.Equal(t, 1, spy.removeCallCount)

			var closed bool
			select {
			case _, ok := <-events:
				closed = !ok
			default:
			}
			assert.True(t, closed, "channel closed after cancel")
		})
	})

	t.Run("Events are delivered in order", func(t *testing.T) {
		synctest.Run(func() {
			ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			defer cancel()

			spy := &EventTargetSpy{EventTarget: target}
			src := &EventSource{spy}

			events := src.Listen(ctx, "gost-event", 1)

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

// Listen returns a read channel of all events of type t until ctx is cancelled.
// If ctx is nil, the event stream will continue indefinitely. The channel will
// have a buffer size of buf.
//
// Ordering of events is guaranteed only when events are dispatched from the
// same goroutine, and the channel buffer is not full.
func (s EventSource) Listen(ctx context.Context, t string, buf int) <-chan *event.Event {
	c := make(chan *event.Event, 32)
	handler := event.NewEventHandlerFunc(func(e *event.Event) error {
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
