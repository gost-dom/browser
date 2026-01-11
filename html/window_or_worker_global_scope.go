package html

import (
	"time"

	"github.com/gost-dom/browser/internal/clock"
)

type windowOrWorkerGlobalScope struct {
	clock *clock.Clock
}

func (w windowOrWorkerGlobalScope) SetTimeout(
	task clock.SafeTaskCallback,
	delay time.Duration,
) int {
	return int(w.clock.SetTimeout(func() error { task(); return nil }, delay))
}

func (w windowOrWorkerGlobalScope) ClearTimeout(handle int) {
	w.clock.Cancel(clock.TaskHandle(handle))
}

func (w windowOrWorkerGlobalScope) SetInterval(
	task clock.SafeTaskCallback,
	delay time.Duration,
) int {
	return int(w.clock.SetInterval(task, delay))
}

func (w windowOrWorkerGlobalScope) ClearInterval(handle int) {
	w.clock.Cancel(clock.TaskHandle(handle))
}

func (w windowOrWorkerGlobalScope) QueueMicrotask(task clock.SafeTaskCallback) {
	w.clock.AddMicrotask(func() error { task(); return nil })
}
