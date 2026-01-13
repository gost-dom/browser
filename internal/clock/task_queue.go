package clock

import (
	"errors"
	"log/slog"
)

// Implements queue intended for Microtasks. While the JavaScript engine itself
// has a microtask queue; not all engine implementations expose it to the
// embedded. In a browser environment, non-native JavaScript operations can
// result in microtasks.
//
//   - enqueueMicrotask function in global scope
//   - MutationObserver events.
//
// While enqueueMicrotask is a JavaScript function, it's not a native JavaScript
// function but one provided by the embedder; so the embedder needs to handle
// those microtasks.
type TaskQueue struct {
	microtasks []TaskCallback
}

// TODO: maxIterationsWithoutDecrement was originally configurable by the
// called. Make it so again.

// maxIterationsWithoutDecrement is the maximum no of iterations of queue
// processing allowed before seeing a decrease in queue size. This is to prevent
// an infinite loop when one task continuously enqueues the same task.
const maxIterationsWithoutDecrement = 100

func (c *TaskQueue) RunMicrotasks() error {
	var errs []error
	minLength := len(c.microtasks)
	count := 0

	for len(c.microtasks) > 0 {
		t := c.microtasks[0]
		c.microtasks = c.microtasks[1:]
		if err := t(); err != nil {
			errs = append(errs, err)
		}
		newLength := len(c.microtasks)
		if newLength < minLength {
			minLength = newLength
			count = 0
		} else {
			count++
			if count > maxIterationsWithoutDecrement {
				panic("Clock: Size of pending microtasks isn't decreasing. Are tasks adding new tasks?")
			}
		}
	}
	return errors.Join(errs...)
}

func (q *TaskQueue) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("noMicrotasks", len(q.microtasks)),
	)
}

func (c *TaskQueue) Enqueue(task TaskCallback) {
	c.microtasks = append(c.microtasks, task)
}
