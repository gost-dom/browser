package clock

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Task func() error

type FutureTask struct {
	Time time.Time
	Task Task
}

type Clock struct {
	Time time.Time
	// Sets the number of times the function allows running without seeing a
	// reduction in task list size. I.e., the task list doesn't need to be
	// emptied after the specified no of tasks executed, but the list of pending
	// tasks must at least have been lower. The counter is reset every time a
	// new minimum of remaining number of tasks is noticed
	MaxLoopWithoutDecrement int
	tasks                   []FutureTask
	microtasks              []Task
}

func New(options ...NewClockOption) *Clock {
	c := &Clock{
		MaxLoopWithoutDecrement: 100,
	}
	for _, o := range options {
		o(c)
	}
	return c
}

func (c *Clock) runMicrotasks() []error {
	var errs []error
	for len(c.microtasks) > 0 {
		t := c.microtasks[0]
		c.microtasks = c.microtasks[1:]
		if err := t(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (c *Clock) runWhile(predicate func() bool) []error {
	var errs []error
	minLength := len(c.tasks)
	count := 0
	for predicate() {
		task := c.tasks[0]
		c.tasks = c.tasks[1:]
		c.Time = task.Time
		if err := task.Task(); err != nil {
			errs = append(errs, err)
		}
		errs = append(errs, c.runMicrotasks()...)
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

func (c *Clock) Advance(d time.Duration) error {
	endTime := c.Time.Add(d)
	errs := c.runWhile(func() bool {
		return len(c.tasks) > 0 && !c.tasks[0].Time.After(endTime)
	})
	c.Time = endTime
	return errors.Join(errs...)
}

func (c *Clock) AddMicrotask(task Task) {
	c.microtasks = append(c.microtasks, task)
}

func (c *Clock) AddTask(when FutureTimeSpec, task Task) {
	taskTime := when(c.Time)
	future := FutureTask{
		Time: taskTime,
		Task: task,
	}
	idx := slices.IndexFunc(c.tasks, func(t FutureTask) bool { return t.Time.After(taskTime) })
	if idx >= 0 {
		c.tasks = slices.Insert(c.tasks, idx, future)
	} else {
		c.tasks = append(c.tasks, future)
	}
}

// Keeps running as long as there are tasks in the task queue. New tasks
// appended while running will also run. Panics if the task list doesn't
// decrease in size.
//
// A note about panic. If you call this, and have an active setInterval, this
// could go into an infinite loop. It accomodate this, the implementation
// monitors the task list size. If the task list size doesn't decrease after a
// default of 100 iterations (configuratble on [Clock.MaxLoopWithoutDecrement]),
// it will panic.
//
// To test code that uses a setInterval, you can call [Clock.Advance] instead.
//
// Why panic, not error? This is to be executed from a test, that _may_ or _may
// not_ care about the error itself, i.e., the test may not assert success. This
// case however, is an issue in the test case that the developer should be aware
// of.
func (c *Clock) RunAll() error {
	errs := c.runWhile(func() bool { return len(c.tasks) > 0 })

	return errors.Join(errs...)
}

/* -------- Options -------- */

type NewClockOption func(*Clock)

// Initializes the clock's time from a std [time.Time] value
func WithTime(t time.Time) NewClockOption {
	return func(c *Clock) {
		c.Time = t
	}
}

// Initializes the clock's time based on an RFC3339 time string, using the as
// implemented by the [time.RFC3339] format. Panics if the string is not valid.
//
// This is intended for the use case where the time is a constant in a test
// case, and as such will either fail or succeed consistently. For variable
// input, the caller should parse the time and use [WithTime] instead.
func IsoTime(iso string) NewClockOption {
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		panic(fmt.Errorf("clock.IsoTime: error parsing string - %w", err))
	}
	return WithTime(t)
}

/* -------- DelaySpecifier -------- */

// Specifies a future time.
type FutureTimeSpec func(current time.Time) time.Time

func Relative(d time.Duration) FutureTimeSpec {
	if d < 0 {
		panic(fmt.Errorf("clock.Relative: Negative relative time - %d", d))
	}
	return func(t time.Time) time.Time {
		return t.Add(d)
	}
}
