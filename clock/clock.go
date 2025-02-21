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
	Time       time.Time
	tasks      []FutureTask
	microtasks []Task
}

func New(options ...NewClockOption) *Clock {
	c := new(Clock)
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

func (c *Clock) runTo(endTime time.Time) []error {
	var errs []error
	for len(c.tasks) > 0 && !c.tasks[0].Time.After(endTime) {
		task := c.tasks[0]
		c.tasks = c.tasks[1:]
		c.Time = task.Time
		if err := task.Task(); err != nil {
			errs = append(errs, err)
		}
		errs = append(errs, c.runMicrotasks()...)
	}
	c.Time = endTime
	return errs
}

func (c *Clock) Advance(d time.Duration) error {
	endTime := c.Time.Add(d)
	return errors.Join(c.runTo(endTime)...)
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

func (c *Clock) RunAll() error {
	var errs []error
	for len(c.tasks) > 0 {
		errs = append(errs, c.runTo(c.tasks[0].Time)...)
	}
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
