// This file is generated. Do not edit.

package htmlinterfaces

type WindowOrWorkerGlobalScope interface {
	IsSecureContext() bool
	CrossOriginIsolated() bool
	ReportError(any)
	Btoa(string) string
	Atob(string) string
	SetTimeout(TimerHandler) int
	SetTimeoutTimeout(TimerHandler, int, ...any) int
	ClearTimeout(int)
	SetInterval(TimerHandler) int
	SetIntervalTimeout(TimerHandler, int, ...any) int
	ClearInterval(int)
	QueueMicrotask(VoidFunction)
}
