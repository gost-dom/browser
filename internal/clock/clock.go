package clock

import (
	"context"
	"errors"
	"time"
)

type processEventOptions struct {
	keepCurrentTime bool
}

type ProcessEventOption func(*processEventOptions)

func KeepCurrentTime() ProcessEventOption {
	return func(o *processEventOptions) {
		o.keepCurrentTime = true
	}
}

type Clock interface {
	runMicrotasks() error
	QueueMicrotask(TaskCallback)
	QueueMacrotask(TaskCallback) TaskHandle
	SetTimeout(TaskCallback, time.Duration) TaskHandle
	SetInterval(TaskCallback, time.Duration) TaskHandle
	Cancel(TaskHandle)
	ProcessEvents(context.Context, ...ProcessEventOption) error
	ProcessEventsWhile(context.Context, func() bool, ...ProcessEventOption) error
	BeginEvent() *EventLoopCallback
	Close()
	Time() time.Time
	SetTime(time.Time)
	runWhile(predicate func() bool) []error
	length() int
	enter()
	exit() error
	peek() (futureTask, bool)
}

// Creates a new clock. If the options don't set a specific time, the clock is
// initialised with the current system time as the initial simulated wall clock
// time.
//
// If tests depend on an actual time, e.g., verifying a local time displaed in
// the user interface, then test code should pass a concrete starting time;
// letting the test execution be independent of the running environment.
//
// The option should only be left out if the test only needs to verify behaviour
// due to passing of time. E.g., testing throttling/debouncing/timeouts.
func New(options ...NewClockOption) Clock {
	c := &singleClock{
		MaxLoopWithoutDecrement: 100,
	}
	for _, o := range options {
		o(c)
	}
	if c.time.IsZero() {
		c.time = time.Now()
	}
	return c
}

// Advances the clock by the specified amount of time. Any new tasks being
// registered while running will be executed; if they are scheduled _before_ the
// timeout. When returning, the clock time will be the current time + the
// duration.
//
// Returns an error if any of the added tasks generate an error. Panics if the
// task list doesn't decrease in size. See [Clock] documentation for more info.
func Advance(c Clock, d time.Duration) error {
	endTime := c.Time().Add(d)
	errs := []error{c.runMicrotasks()}
	errs = append(errs, c.runWhile(func() bool {
		task, ok := c.peek()
		return ok && !task.time.After(endTime)
	})...)
	c.SetTime(endTime)
	return errors.Join(errs...)
}

// Tick runs all tasks scheduled for immediate execution. This is synonymous
// with calling Advance(0).
func Tick(c Clock) error {
	return Advance(c, 0)
}

// Keeps running as long as there are tasks in the task queue. New tasks
// appended while running will also run. When returning, the current time will
// be the time of the last executed task.
//
// Returns an error if any of the added tasks generate an error. Panics if the
// task list doesn't decrease in size. See [Clock] documentation for more info.
func RunAll(c Clock) error {
	errs := c.runWhile(func() bool { return c.length() > 0 })

	return errors.Join(errs...)
}

// Do wraps a task call and runs microtasks when it has completed. Nested tasks
// will not trigger microtasks; only when the original root task completes.
func Do(c Clock, f func() error) (err error) {
	c.enter()
	defer func() {
		err = errors.Join(err, c.exit())
	}()

	return f()
}
