// package result provides monadic result binding
//
// Use of code in this package is discouraged, as you should generally prefer
// handling errors.
//
// This exists as an experiment to simplify code containing long chains of
// functions that can produce errors. The only usage right now is retrieving all
// keys of a JavaScript object in V8, which isn't exposed in v8go, and a current
// workaround involves multiple calls to V8 each of which can produce an error,
// and the previous version had 7 levels if nested if-statements.
//
// Fixing the lack of support for Object.keys in v8go is a better solution to
// the same problem than using monadic bindings to go.
package result

type Result[T any] struct {
	Value T
	Err   error
}

func New[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Err: err}
}

func (r Result[T]) Unwrap() (T, error) { return r.Value, r.Err }

type BindFunc0[T, V any] = func(T) (V, error)
type BindFunc1[T, V, A1 any] = func(val T, arg A1) (V, error)
type BindFunc2[T, V, A1, A2 any] = func(val T, a1 A1, a2 A2) (V, error)

func Bind[T, U any](r Result[T], f BindFunc0[T, U]) Result[U] {
	if r.Err != nil {
		return Result[U]{Err: r.Err}
	}
	res, err := f(r.Value)
	return New(res, err)
}

func Bind1[T, U, A1 any](r Result[T], f BindFunc1[T, U, A1], arg A1) Result[U] {
	if r.Err != nil {
		return Result[U]{Err: r.Err}
	}
	return New(f(r.Value, arg))
}
func Bind2[T, U, A1, A2 any](r Result[T], f BindFunc2[T, U, A1, A2], a1 A1, a2 A2) Result[U] {
	if r.Err != nil {
		return Result[U]{Err: r.Err}
	}
	return New(f(r.Value, a1, a2))
}

func Unwrap[T any](r Result[T]) (T, error) { return r.Unwrap() }
