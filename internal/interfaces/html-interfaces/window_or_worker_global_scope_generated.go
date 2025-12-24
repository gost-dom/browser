// This file is generated. Do not edit.

package htmlinterfaces

import "time"

type WindowOrWorkerGlobalScope interface {
	SetTimeout(TimerHandler, time.Duration) int
	ClearTimeout(int)
	SetInterval(TimerHandler, time.Duration) int
	ClearInterval(int)
	QueueMicrotask(VoidFunction)
}
