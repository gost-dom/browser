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
}

func (s *ClockTestSuite) TestNewClock() {
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	s.Assert().Equal(feb1st2025_noon, c.Time.Unix())
}

func (s *ClockTestSuite) TestAdvance() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))

	c.AddSafeTask(clock.Relative(101*time.Millisecond), func() { addLog("A") })
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { addLog("B") })
	s.Assert().NoError(c.Advance(10 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+10, c.Time.UnixMilli())

	s.Assert().NoError(c.Advance(90 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
	s.Assert().Equal([]string{"B"}, logs)

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
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(clock.Relative(200*time.Millisecond), func() { addLog("B") })
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { addLog("A") })
	c.AddSafeTask(clock.Relative(300*time.Millisecond), func() { addLog("C") })
	err := c.RunAll()
	s.Assert().NoError(err)
	s.Assert().Equal([]string{"A", "B", "C"}, logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestErrorsAreReturned() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(200*time.Millisecond), func() error {
		addLog("B")
		return errors.New("Error B")
	})
	c.AddSafeTask(clock.Relative(100*time.Millisecond), func() { addLog("A") })
	c.AddSafeTask(clock.Relative(300*time.Millisecond), func() { addLog("C") })

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "B", "C"}, logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestImmediatesAreExecutedBeforeScheduledTasks() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.OfIsoString("2025-02-01T12:00:00Z"))
	c.AddSafeTask(clock.Relative(1*time.Millisecond), func() {
		addLog("A")
		c.AddSafeMicrotask(func() {
			addLog("A2")
			c.AddSafeMicrotask(func() { addLog("A2A") })
		})
		c.AddSafeMicrotask(func() { addLog("A3") })
	})
	c.AddTask(
		clock.Relative(1*time.Millisecond),
		func() error { addLog("B"); return errors.New("Error B") },
	)

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "A2", "A3", "A2A", "B"}, logs)
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
