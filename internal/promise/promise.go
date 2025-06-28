package promise

import (
	"fmt"
	"io"
)

// Result represents the outcome of a promise. If the promise is rejected, Err
// will contain a non-nil value. If Err is nil, the promise was fulfilled, the
// result contained in field Value.
type Result[T any] struct {
	Value T
	Err   error
}

type Promise[T any] chan Result[T]

func (p Promise[T]) Close()              { close(p) }
func (p Promise[T]) Resolve(v T)         { p <- Result[T]{Value: v} }
func (p Promise[T]) Send(v T, err error) { p <- Result[T]{Value: v, Err: err} }

// Reject rejects the promise,
func (p Promise[T]) Reject(err error) { p <- Result[T]{Err: err} }

// New returns a promise. If function f is non-nil, it will be executed in a
// separate gorouting, and the promise will eventually settle with the result of
// the function. If f returns a non-nil error, the promise will reject,
// otherwise it will resolve with a value of type T.
func New[T any](f func() (T, error)) Promise[T] {
	p := make(Promise[T], 1)
	if f != nil {
		go func() { p.Send(f()) }()
	}
	return p
}

func ReadAll(reader io.Reader) Promise[[]byte] {
	return New(func() ([]byte, error) { return io.ReadAll(reader) })
}

// ErrAny wraps any value as a valid go [error] value. While errors originating
// from Go code will always be instances of error, in JavaScript, any value can
// be an error.
//
// When an error is generated in JavaScript code, and not representable directly
// as an error in Go, ErrAny will represent the value.
//
// TODO: This isn't specific to promises, but there's not really another package
// that's a good fit ATM, and I don't want to create a new package just to have
// this type. Consider moving in the future
type ErrAny struct{ Reason any }

func (err ErrAny) Error() string {
	return fmt.Sprintf("aborted: reason: %v", err.Reason)
}

func (err ErrAny) As(target any) bool {
	if errAny, ok := target.(*ErrAny); ok {
		*errAny = err
		return true
	}
	return false
}
