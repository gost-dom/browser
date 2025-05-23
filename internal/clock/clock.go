// Package clock provides a simulated time for Gost-DOM.
package clock

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/gost-dom/browser/internal/dom/mutation"
)

type TaskHandle uint32

// A TaskCallback is the callback registered for a [futureTask].
type TaskCallback func() error

// A SafeTaskCallback is the callback that registered for a [futureTask] that
// can't generate an err.r
type SafeTaskCallback func()

func (t SafeTaskCallback) toTask() TaskCallback { return func() error { t(); return nil } }

// A futureTask represents a task to run when simulated time advances past the
// specified Time value.
type futureTask struct {
	time   time.Time
	task   TaskCallback
	handle TaskHandle
	repeat bool
	delay  time.Duration
}

// Clock simulates passing of time, as well as potential future tasks. Simulated
// Time can be advanced using [Clock.Advance] or [Clock.RunAll]. The zero value
// for Clock represents Unix time 0, and will not panic.
//
// Advancing time will perform the following
// - Run microtasks.
// - Flush any registered "Flushers"
// - Run timeout/interval callbacks
//
// Microtasks are JavaScript callbacks that should execute when a script has
// completed, but before returning to the event loop. The script engine _may_
// run microtasks, possibly making microtask management unnecessary in this
// space.
//
// Flushers are operations that would run when returning to the event loop.
// Events generated by the MutationObserver are dispatched here.
//
// Advance and RunAll will return an error if any of the executed tasks return
// an error; Advance and RunAll panics if the task list isn't reducing.
//
// Advancing the clock panics if the task queue isn't decreasing, i.e, tasks are
// adding new tasks. This will happen if a JavaScript call to setInterval isn't
// cleared before calling [RunAll], or a setInterval callback that continuously
// adds a new immediate callback, or a microtasks creating a new microtask.
//
// This is designed to panic rather than return an error, as an error represents
// an error generated by the code being tested. You may choose to ignore it, or
// even assert on it's precense when testing error conditions. But a task list
// that doesn't decrease is an error, that the developer should be notified of;
// which is why the design is a panic.
//
// The behaviour is configurable by the MaxLoopWithoutDecrement value.
type Clock struct {
	Time time.Time
	mutation.FlusherSet

	// Sets the number of times a task is allowed to run without seeing a
	// reduction in task list size. I.e., the task list doesn't need to be
	// emptied after the specified no of tasks executed, but the smallest
	// observed size of the queue must have decreased. The counter is reset
	// every time a new minimum of remaining number of tasks is noticed
	MaxLoopWithoutDecrement int
	tasks                   []futureTask
	// The tasks in the microtask queue must execute before the next task in the
	// task queue, even if next tasks is set to execute now. Microtasks are
	// really a property of the JavaScript language and should be eventually be
	// carried out by the JavaScript host.
	microtasks []TaskCallback
	nextHandle TaskHandle
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
func New(options ...NewClockOption) *Clock {
	c := &Clock{
		MaxLoopWithoutDecrement: 100,
	}
	for _, o := range options {
		o(c)
	}
	if c.Time.IsZero() {
		c.Time = time.Now()
	}
	return c
}

// runMicrotasksAndFlush runs first microtasks, e.g., tasks added using
// `window.queueMicrotask`.
func (c *Clock) runMicrotasksAndFlush() []error {
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
			if count > c.MaxLoopWithoutDecrement {
				panic("Clock: Size of pending microtasks isn't decreasing. Are tasks adding new tasks?")
			}
		}
	}
	c.FlusherSet.Flush()
	return errs
}

func (c *Clock) runWhile(predicate func() bool) []error {
	var errs []error
	minLength := len(c.tasks)
	count := 0
	errs = append(errs, c.runMicrotasksAndFlush()...)
	for predicate() {
		task := c.tasks[0]
		c.tasks = c.tasks[1:]
		c.Time = task.time
		if err := task.task(); err != nil {
			errs = append(errs, err)
		}
		if task.repeat {
			task.time = task.time.Add(task.delay)
			c.insertTask(task)
		}
		errs = append(errs, c.runMicrotasksAndFlush()...)
		newLength := len(c.tasks)
		if newLength < minLength {
			minLength = newLength
			count = 0
		} else {
			count++
			if count > c.MaxLoopWithoutDecrement {
				panic("Clock: Size of pending tasks isn't decreasing. Are tasks adding new tasks?")
			}
		}
	}
	return errs
}

// Advances the clock by the specified amount of time. Any new tasks being
// registered while running will be executed; if they are scheduled _before_ the
// timeout. When returning, the clock time will be the current time + the
// duration.
//
// Returns an error if any of the added tasks generate an error. Panics if the
// task list doesn't decrease in size. See [Clock] documentation for more info.
func (c *Clock) Advance(d time.Duration) error {
	endTime := c.Time.Add(d)
	errs := c.runMicrotasksAndFlush()
	errs = append(errs, c.runWhile(func() bool {
		return len(c.tasks) > 0 && !c.tasks[0].time.After(endTime)
	})...)
	c.Time = endTime
	return errors.Join(errs...)
}

// Tick runs all tasks scheduled for immediate execution. This is synonymous
// with calling Advance(0).
func (c *Clock) Tick() error { return c.Advance(0) }

// Cancel removes the task that have been added using [Clock.SetTimeout] or
// [Clock.SetInterval]. This corresponds to either [clearTimeout] or
// [clearInterval] in the browser, which by specification can be used
// interchangably; but shouldn't for clarity.
//
// [clearTimeout]: https://developer.mozilla.org/en-US/docs/Web/API/Window/clearTimeout
// [clearInterval]: https://developer.mozilla.org/en-US/docs/Web/API/Window/clearInterval
func (c *Clock) Cancel(handle TaskHandle) {
	idx := slices.IndexFunc(
		c.tasks,
		func(t futureTask) bool { return t.handle == handle },
	)
	if idx >= 0 {
		c.tasks = slices.Delete(c.tasks, idx, idx+1)
	}
}

// AddMicrotask adds a task to the "microtask queue".
//
// This shouldn't be called from Go code. Microtasks are a property of
// JavaScript execution and should really be carried out by the javascript
// engine.
func (c *Clock) AddMicrotask(task TaskCallback) {
	c.microtasks = append(c.microtasks, task)
}

// AddSafeMicrotask is a version of AddMicrotask, where the caller can guarantee
// the task doesn't generate an error.
func (c *Clock) AddSafeMicrotask(task SafeTaskCallback) { c.AddMicrotask(task.toTask()) }

func (c *Clock) generateHandle() TaskHandle {
	c.nextHandle++
	return c.nextHandle
}

func (c *Clock) insertTask(future futureTask) TaskHandle {
	if future.handle == 0 {
		future.handle = c.generateHandle()
	}
	idx := slices.IndexFunc(c.tasks, func(t futureTask) bool { return t.time.After(future.time) })
	if idx >= 0 {
		c.tasks = slices.Insert(c.tasks, idx, future)
	} else {
		c.tasks = append(c.tasks, future)
	}
	return future.handle
}

// SetInterval corresponds to the browser's [setInterval] function. Panics if
// the delay is negative.
//
// [SetInterval]: https://developer.mozilla.org/en-US/docs/Web/API/Window/setInterval
func (c *Clock) SetInterval(task SafeTaskCallback, delay time.Duration) TaskHandle {
	if delay < 0 {
		panic(fmt.Sprintf("Clock.SetInterval: negative delay: %d", delay))
	}
	return c.insertTask(
		futureTask{
			time:   c.Time.Add(delay),
			task:   task.toTask(),
			repeat: true,
			delay:  delay,
		})
}

// SetInterval corresponds to the browser's [setTimeout] function. Panics if the
// delay is negative.
//
// [SetTimeout]: https://developer.mozilla.org/en-US/docs/Web/API/Window/setTimeout
func (c *Clock) SetTimeout(task TaskCallback, delay time.Duration) TaskHandle {
	if delay < 0 {
		panic(fmt.Sprintf("Clock.SetTimeout: negative delay: %d", delay))
	}
	return c.insertTask(futureTask{
		time: c.Time.Add(delay),
		task: task,
	})
}

// Schedules a task to run at a specified time in the future. Panics if the time
// is in the past.
func (c *Clock) AddSafeTask(task SafeTaskCallback, delay time.Duration) TaskHandle {
	return c.SetTimeout(task.toTask(), delay)
}

// Keeps running as long as there are tasks in the task queue. New tasks
// appended while running will also run. When returning, the current time will
// be the time of the last executed task.
//
// Returns an error if any of the added tasks generate an error. Panics if the
// task list doesn't decrease in size. See [Clock] documentation for more info.
func (c *Clock) RunAll() error {
	errs := c.runWhile(func() bool { return len(c.tasks) > 0 })

	return errors.Join(errs...)
}

/* -------- Options -------- */

// NewClockOption are used to initialize a new [Clock]
type NewClockOption func(*Clock)

// Initializes the clock's simulated time from a concrete [time.Time] value.
func OfTime(t time.Time) NewClockOption {
	return func(c *Clock) {
		c.Time = t
	}
}

// Initializes the clock's simulated time based on an RFC3339 time string.
// Panics if the string is not valid.
//
// This is intended for the use case where the time is a constant in a test
// case, and as such will either fail or succeed consistently. For variable
// input, the caller should parse the time and use [OfTime] instead.
func OfIsoString(iso string) NewClockOption {
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		panic(fmt.Sprintf("clock.IsoTime: error parsing string - %v", err))
	}
	return OfTime(t)
}
