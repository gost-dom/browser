package clock_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	"github.com/stretchr/testify/suite"
)

const feb1st2025_noon int64 = 1738411200
const feb1st2025_noon_milli int64 = feb1st2025_noon * 1000

type ClockTestSuite struct {
	suite.Suite
	logs []string
}

func (s *ClockTestSuite) SetupTest() {
	s.logs = nil
}

func (s *ClockTestSuite) log(l string) {
	s.logs = append(s.logs, l)
}

func (s *ClockTestSuite) TestNewClock() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	s.Assert().Equal(feb1st2025_noon, c.Time.Unix())
}

func (s *ClockTestSuite) TestAdvance() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))

	c.AddSafeTask(func() { s.log("A") }, 101*time.Millisecond)
	c.AddSafeTask(func() { s.log("B") }, 100*time.Millisecond)
	s.Assert().NoError(c.Advance(10 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+10, c.Time.UnixMilli())

	s.Assert().NoError(c.Advance(90 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
	s.Assert().Equal([]string{"B"}, s.logs)
}

func (s *ClockTestSuite) TestCancelTask() {
	c := clock.New()
	c.AddSafeTask(func() { s.log("A") }, 100*time.Millisecond)
	c.AddSafeTask(func() { s.log("B1") }, 200*time.Millisecond)
	handle := c.AddSafeTask(func() { s.log("B2") }, 200*time.Millisecond)
	c.AddSafeTask(func() { s.log("C") }, 300*time.Millisecond)

	c.Advance(150 * time.Millisecond)
	c.Cancel(handle)
	c.Advance(150 * time.Millisecond)

	s.Assert().Equal([]string{"A", "B1", "C"}, s.logs)
}

func (s *ClockTestSuite) TestInitialTimeIsRelevant() {
	// Without any defaults, the clock would start at UTS 0. This can be
	// problematic converting to a local time on the Western Hemisphere, as that
	// could result in a negative time.

	// A sensible default is to use the current system time as the default
	// starting time.
	currentTime := time.Now()
	c := clock.New()

	// This is pretty vague specification, but the test is not too specific
	// about what the time should be, to allow some deliberate jitter in the
	// clock. Just make sure we're in the right ballpark of a sensible default.
	s.Assert().True(c.Time.After(currentTime.Add(-time.Second)))
	s.Assert().True(c.Time.Before(currentTime.Add(time.Second)))
}

func (s *ClockTestSuite) TestRunAll() {
	var runCount int
	task := func() { runCount++ }
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(task, 100*time.Millisecond)
	c.RunAll()
	s.Assert().Equal(1, runCount)
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestOrderedExecution() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(func() { s.log("B") }, 200*time.Millisecond)
	c.AddSafeTask(func() { s.log("A") }, 100*time.Millisecond)
	c.AddSafeTask(func() { s.log("C") }, 300*time.Millisecond)
	err := c.RunAll()
	s.Assert().NoError(err)
	s.Assert().Equal([]string{"A", "B", "C"}, s.logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestErrorsAreReturned() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.SetTimeout(func() error { s.log("B"); return errors.New("Error B") }, 200*time.Millisecond)
	c.AddSafeTask(func() { s.log("A") }, 100*time.Millisecond)
	c.AddSafeTask(func() { s.log("C") }, 300*time.Millisecond)

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "B", "C"}, s.logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestImmediatesAreExecutedBeforeScheduledTasks() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(
		func() {
			s.log("A")
			c.AddSafeMicrotask(func() {
				s.log("A2")
				c.AddSafeMicrotask(func() { s.log("A2A") })
			})
			c.AddSafeMicrotask(func() { s.log("A3") })
		},
		1*time.Millisecond,
	)
	c.SetTimeout(
		func() error { s.log("B"); return errors.New("Error B") },
		1*time.Millisecond,
	)

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "A2", "A3", "A2A", "B"}, s.logs)
}

func (s *ClockTestSuite) TestTick() {
	c := clock.New()
	c.AddSafeTask(
		func() { s.log("Task") },
		0,
	)
	c.AddSafeMicrotask(func() { s.log("Microtask") })
	c.Tick()
	s.Assert().Equal([]string{"Microtask", "Task"}, s.logs)
}

func (s *ClockTestSuite) TestImmediatesPanicWhenListDoesntReduce() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() { c.AddSafeMicrotask(task) }

	c.AddSafeTask(task, 1*time.Millisecond)

	s.Assert().Panics(func() { c.RunAll() })
}

func (s *ClockTestSuite) TestImmediatesPropagateErrors() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(func() {
		c.AddMicrotask(func() error { return errors.New("Microtask error") })
	}, 1*time.Millisecond)

	err := c.RunAll()
	s.Assert().Error(err, "Microtask error")
}

func (s *ClockTestSuite) TestRepeatingTasksGeneratePanicOnRunAll() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() {
		c.AddSafeTask(task, 100*time.Millisecond)
	}
	c.AddSafeTask(task, 100*time.Millisecond)

	s.Assert().Panics(func() { c.RunAll() })
}

func (s *ClockTestSuite) TestSingleRepeatingTask() {
	c := clock.New()
	var task clock.SafeTaskCallback
	task = func() {
		c.AddSafeTask(func() {}, 1*time.Millisecond)
	}
	c.AddSafeTask(task, 1*time.Millisecond)

	s.Assert().NotPanics(func() { c.Advance(100 * time.Millisecond) })
}

func (s *ClockTestSuite) TestRepeatingTasksGeneratePanicOnRunAdvance() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() { c.AddSafeTask(task, 1*time.Millisecond) }
	c.AddSafeTask(task, 1*time.Millisecond)

	s.Assert().Panics(func() { c.Advance(1000 * time.Millisecond) })
}

func (s *ClockTestSuite) TestProcessEvents() {
	var count int

	ctx, cancel := context.WithTimeout(s.T().Context(), time.Millisecond)
	defer cancel()

	c := clock.New()
	c.BeginEvent()
	go func() {
		c.AddSafeEvent(func() {
			count++
		})
	}()

	s.Assert().Equal(0, count, "count before processEvents")
	c.ProcessEvents(ctx)
	s.Assert().Equal(1, count, "count before processEvents")
}

func (s *ClockTestSuite) TestProcessEventsUntil() {
	var count int

	ctx, cancel := context.WithTimeout(s.T().Context(), time.Millisecond)
	defer cancel()

	c := clock.New()
	c.BeginEvent()
	go func() {
		c.AddSafeEvent(func() { count++ })
	}()
	c.BeginEvent()
	go func() {
		c.AddSafeEvent(func() { count++ })
	}()

	s.Assert().Equal(0, count, "count before ProcessEventsWhile")
	c.ProcessEventsWhile(ctx, func() bool { return count == 0 })
	s.Assert().Equal(1, count, "count after ProcessEventsWhile")
	c.ProcessEvents(ctx)
	s.Assert().Equal(2, count, "count before ProcessEvents")
}

func TestClock(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ClockTestSuite))
}
