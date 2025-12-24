// This file is generated. Do not edit.

package htmlinterfaces

type WindowOrWorkerGlobalScope interface {
	SetTimeout(TimerHandler) int
	SetTimeoutTimeout(TimerHandler, int, ...any) int
	ClearTimeout(int)
	SetInterval(TimerHandler) int
	SetIntervalTimeout(TimerHandler, int, ...any) int
	ClearInterval(int)
	QueueMicrotask(VoidFunction)
}
