// This file is generated. Do not edit.

package htmlinterfaces

type History interface {
	Length() int
	State() any
	Go(int) error
	Back() error
	Forward() error
	PushState(any, string) error
	PushStateUrl(any, string, string) error
	ReplaceState(any, string) error
	ReplaceStateUrl(any, string, string) error
}
