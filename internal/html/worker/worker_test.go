package worker_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/html/worker"
	"github.com/stretchr/testify/assert"
)

func TestWorkerEnqueue(t *testing.T) {
	c := clock.New()
	w := worker.New(c)
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()
	var called bool
	delay := make(chan struct{})
	item := func(worker.GlobalScope) error {
		<-delay
		called = true
		return nil
	}
	assert.NoError(t, w.Enqueue(item))
	assert.False(t, called)
	close(delay)
	assert.NoError(t, c.ProcessEvents(ctx))
	assert.True(t, called)
}

func TestWorkerSetTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := clock.New()
		w := worker.New(c)
		defer w.Close()
		var called bool
		w.Enqueue(func(s worker.GlobalScope) error {
			s.SetTimeout(clock.SafeTask(func() {
				called = true
			}), time.Millisecond*100)
			return nil
		})
		synctest.Wait() // Make sure setTimeout has been called

		// Wait 99 milliseconds. callback shouldn't have been called
		c.Advance(99 * time.Millisecond)
		synctest.Wait()
		assert.False(t, called)

		// Wait one more millisecond, callback should have been called.
		c.Advance(1 * time.Millisecond)
		synctest.Wait()
		assert.True(t, called, "callback should be called after one more millisecond")
	})
}
