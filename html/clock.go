package html

import (
	"time"

	"github.com/gost-dom/browser/internal/clock"
)

type clockWrapper struct {
	clock.Clock
}

func (w clockWrapper) Advance(d time.Duration) error {
	return clock.Advance(w.Clock, d)
}

func (w clockWrapper) RunAll() error { return clock.RunAll(w.Clock) }
