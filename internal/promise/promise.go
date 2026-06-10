package promise

import (
	"io"
	"time"
)

type PromiseOptions struct {
	Delay time.Duration
}
type PromiseOption func(*PromiseOptions)

func WithDelay(d time.Duration) PromiseOption {
	return func(o *PromiseOptions) { o.Delay = d }
}

// Result represents the outcome of a promise. If the promise is rejected, Err
// will contain a non-nil value. If Err is nil, the promise was fulfilled, the
// result contained in field Value.
type Result[T any] struct {
	Value T
	Err   error
}

type Promise[T any] struct {
	PromiseOptions
	C chan Result[T]
}

func (p Promise[T]) Close()              { close(p.C) }
func (p Promise[T]) Resolve(v T)         { p.C <- Result[T]{Value: v} }
func (p Promise[T]) Send(v T, err error) { p.C <- Result[T]{Value: v, Err: err} }

// Reject rejects the promise,
func (p Promise[T]) Reject(err error) { p.C <- Result[T]{Err: err} }

// New returns a promise. If function f is non-nil, it will be executed in a
// separate gorouting, and the promise will eventually settle with the result of
// the function. If f returns a non-nil error, the promise will reject,
// otherwise it will resolve with a value of type T.
func New[T any](f func() (T, error), opts ...PromiseOption) Promise[T] {
	p := Promise[T]{C: make(chan Result[T], 1)}
	for _, o := range opts {
		o(&p.PromiseOptions)
	}
	if f != nil {
		go func() { p.Send(f()) }()
	}
	return p
}

func ReadAll(reader io.Reader) Promise[[]byte] {
	return New(func() ([]byte, error) { return io.ReadAll(reader) })
}
