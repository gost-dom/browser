package html

import (
	"time"

	"github.com/gost-dom/browser/internal/clock"
)

type windowOrWorkerGlobalScope struct {
	clock *clock.Clock
}

func (w windowOrWorkerGlobalScope) SetTimeout(
	task clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	return w.clock.SetTimeout(func() error { task(); return nil }, delay)
}

func (w windowOrWorkerGlobalScope) ClearTimeout(handle clock.TaskHandle) {
	w.clock.Cancel(handle)
}

func (w windowOrWorkerGlobalScope) SetInterval(
	task clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	return w.clock.SetInterval(task, delay)
}

func (w windowOrWorkerGlobalScope) ClearInterval(handle clock.TaskHandle) {
	w.clock.Cancel(handle)
}

func (w windowOrWorkerGlobalScope) QueueMicrotask(task clock.TaskCallback) {
	w.clock.AddMicrotask(task)
}
