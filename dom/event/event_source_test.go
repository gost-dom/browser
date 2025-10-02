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

type EventChan chan *event.Event

func TestEventsAreReceivedInOrder(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		src := event.EventSource{event.NewEventTarget()}
		const buf = 32
		c := src.Listen(ctx, "gost-event", event.BufSize(buf))

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
		synctest.Test(t, func(t *testing.T) {
			newCtx, cancel := context.WithCancel(t.Context())

			spy := &EventTargetSpy{EventTarget: target}
			src := event.EventSource{spy}
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
		synctest.Test(t, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			defer cancel()

			spy := &EventTargetSpy{EventTarget: target}
			src := &event.EventSource{spy}

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
