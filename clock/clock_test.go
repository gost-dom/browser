package clock_test

import (
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
	c := clock.New(clock.IsoTime("2025-02-01T12:00:00Z"))
	c.Advance(100 * time.Millisecond)
	s.Assert().Equal(feb1st2025_noon_milli+100, c.Time.UnixMilli())
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
	c.RunAll()
	s.Assert().Equal([]string{"B", "A", "C"}, logs)
	s.Assert().Equal(feb1st2025_noon_milli+300, c.Time.UnixMilli())
}

func TestClock(t *testing.T) {
	suite.Run(t, new(ClockTestSuite))
}
