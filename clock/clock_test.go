package clock_test

import (
	"errors"
	"testing"
	"time"

	"github.com/gost-dom/browser/clock"
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

	c.AddSafeTask(clock.Relative(101*time.Millisecond), func() { s.log("A") })
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { s.log("B") })
	s.Assert().NoError(c.Advance(10 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+10, c.Time.UnixMilli())

	s.Assert().NoError(c.Advance(90 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
	s.Assert().Equal([]string{"B"}, s.logs)
}

func (s *ClockTestSuite) TestCancelTask() {
	c := clock.New()
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { s.log("A") })
	handle := c.AddSafeTask(clock.Relative(200*time.Millisecond), func() { s.log("B") })
	c.AddSafeTask(clock.Relative(300*time.Millisecond), func() { s.log("C") })

	c.Advance(150 * time.Millisecond)
	c.Cancel(handle)
	c.Advance(150 * time.Millisecond)

	s.Assert().Equal([]string{"A", "C"}, s.logs)
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
	c.AddSafeTask(clock.Relative(100*time.Millisecond), task)
	c.RunAll()
	s.Assert().Equal(1, runCount)
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestOrderedExecution() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(clock.Relative(200*time.Millisecond), func() { s.log("B") })
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { s.log("A") })
	c.AddSafeTask(clock.Relative(300*time.Millisecond), func() { s.log("C") })
	err := c.RunAll()
	s.Assert().NoError(err)
	s.Assert().Equal([]string{"A", "B", "C"}, s.logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestErrorsAreReturned() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(200*time.Millisecond), func() error {
		s.log("B")
		return errors.New("Error B")
	})
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { s.log("A") })
	c.AddSafeTask(clock.Relative(300*time.Millisecond), func() { s.log("C") })

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "B", "C"}, s.logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestImmediatesAreExecutedBeforeScheduledTasks() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(clock.Relative(1*time.Millisecond), func() {
		s.log("A")
		c.AddSafeMicrotask(func() {
			s.log("A2")
			c.AddSafeMicrotask(func() { s.log("A2A") })
		})
		c.AddSafeMicrotask(func() { s.log("A3") })
	})
	c.AddTask(
		clock.Relative(1*time.Millisecond),
		func() error { s.log("B"); return errors.New("Error B") },
	)

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "A2", "A3", "A2A", "B"}, s.logs)
}

func (s *ClockTestSuite) TestImmediatesPanicWhenListDoesntReduce() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() { c.AddSafeMicrotask(task) }

	c.AddSafeTask(clock.Relative(1*time.Millisecond), task)

	s.Assert().Panics(func() { c.RunAll() })
}

func (s *ClockTestSuite) TestImmediatesPropagateErrors() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(clock.Relative(1*time.Millisecond), func() {
		c.AddMicrotask(func() error { return errors.New("Microtask error") })
	})

	err := c.RunAll()
	s.Assert().Error(err, "Microtask error")
}

func (s *ClockTestSuite) TestRepeatingTasksGeneratePanicOnRunAll() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() {
		c.AddSafeTask(clock.Relative(100*time.Millisecond), task)
	}
	c.AddSafeTask(clock.Relative(100*time.Millisecond), task)

	s.Assert().Panics(func() { c.RunAll() })
}

func (s *ClockTestSuite) TestSingleRepeatingTask() {
	c := clock.New()
	var task clock.SafeTaskCallback
	task = func() {
		c.AddSafeTask(clock.Relative(1*time.Millisecond), func() {})
	}
	c.AddSafeTask(clock.Relative(1*time.Millisecond), task)

	s.Assert().NotPanics(func() { c.Advance(100 * time.Millisecond) })
}

func (s *ClockTestSuite) TestRepeatingTasksGeneratePanicOnRunAdvance() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	var task clock.SafeTaskCallback
	task = func() { c.AddSafeTask(clock.Relative(1*time.Millisecond), task) }
	c.AddSafeTask(clock.Relative(1*time.Millisecond), task)

	s.Assert().Panics(func() { c.Advance(1000 * time.Millisecond) })
}

func TestClock(t *testing.T) {
	suite.Run(t, new(ClockTestSuite))
}
