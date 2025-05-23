// This file is generated. Do not edit.

package htmlinterfaces

type History interface {
	Length() int
	State() HistoryState
	Go(int) error
	Back() error
	Forward() error
	PushState(HistoryState, string) error
	ReplaceState(HistoryState, string) error
}
