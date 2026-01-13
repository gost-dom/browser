// This file is generated. Do not edit.

package htmlinterfaces

import (
	clock "github.com/gost-dom/browser/internal/clock"
	"time"
)

type WindowOrWorkerGlobalScope interface {
	SetTimeout(TimerHandler, time.Duration) clock.TaskHandle
	ClearTimeout(clock.TaskHandle)
	SetInterval(TimerHandler, time.Duration) clock.TaskHandle
	ClearInterval(clock.TaskHandle)
	QueueMicrotask(VoidFunction)
}
