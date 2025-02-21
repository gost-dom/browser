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
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	s.Assert().Equal(feb1st2025_noon, c.Time.Unix())
}

func (s *ClockTestSuite) TestAdvance() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))

	c.AddTask(clock.Relative(101*time.Millisecond), func() error {
		addLog("A")
		return nil
	})
	c.AddTask(clock.Relative(100*time.Millisecond), func() error {
		addLog("B")
		return nil
	})

	s.Assert().NoError(c.Advance(10 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+10, c.Time.UnixMilli())

	s.Assert().NoError(c.Advance(90 * time.Millisecond))
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
	s.Assert().Equal([]string{"B"}, logs)

}

func (s *ClockTestSuite) TestRunAll() {
	var runCount int
	task := func() error { runCount++; return nil }
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(100*time.Millisecond), task)
	c.RunAll()
	s.Assert().Equal(1, runCount)
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestOrderedExecution() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(200*time.Millisecond), func() error {
		addLog("A")
		return nil
	})
	c.AddTask(clock.Relative(100*time.Millisecond), func() error {
		addLog("B")
		return nil
	})
	c.AddTask(clock.Relative(300*time.Millisecond), func() error {
		addLog("C")
		return nil
	})
	err := c.RunAll()
	s.Assert().NoError(err)
	s.Assert().Equal([]string{"B", "A", "C"}, logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestErrorsAreReturned() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(200*time.Millisecond), func() error {
		addLog("A")
		return nil
	})
	c.AddTask(clock.Relative(100*time.Millisecond), func() error {
		addLog("B")
		return errors.New("Error B")
	})
	c.AddTask(clock.Relative(300*time.Millisecond), func() error {
		addLog("C")
		return nil
	})

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"B", "A", "C"}, logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func (s *ClockTestSuite) TestImmediatesAreExecutedBeforeScheduledTasks() {
	var logs []string
	var addLog = func(s string) { logs = append(logs, s) }
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(1*time.Millisecond), func() error {
		addLog("A")
		c.AddMicrotask(func() error {
			addLog("A2")
			c.AddMicrotask(func() error {
				addLog("A2A")
				return nil
			})
			return nil
		})
		c.AddMicrotask(func() error {
			addLog("A3")
			return nil
		})
		return nil
	})
	c.AddTask(clock.Relative(1*time.Millisecond), func() error {
		addLog("B")
		return errors.New("Error B")
	})

	err := c.RunAll()
	s.Assert().Error(err)
	// Subsequent tasks should be performed
	s.Assert().Equal([]string{"A", "A2", "A3", "A2A", "B"}, logs)
}

func (s *ClockTestSuite) TestImmediatesPropagateErrors() {
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.AddTask(clock.Relative(1*time.Millisecond), func() error {
		c.AddMicrotask(func() error {
			return errors.New("Microtask error")
		})
		return nil
	})

	err := c.RunAll()
	s.Assert().Error(err, "Microtask error")

}

func TestClock(t *testing.T) {
	suite.Run(t, new(ClockTestSuite))
}
