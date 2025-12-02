package js

import (
	"errors"
	"fmt"
	"iter"
)

var ErrNotIterable = errors.New("gost-dom/scripting: value not iterable")

func errNotIterable(msg string) error {
	return fmt.Errorf("%w: %s", ErrNotIterable, msg)
}

type Object[T any] interface {
	Value[T]
	NativeValue() any
	SetNativeValue(any)
	Keys() ([]string, error)
	Get(name string) (Value[T], error)
	Set(name string, val Value[T]) error
	Iterator() (Value[T], error)
}

type objectGetter[T any] interface {
	objectGet(Object[T]) (Value[T], error)
}

type StringKey[T any] string

func (s StringKey[T]) objectGet(o Object[T]) (Value[T], error) {
	return o.Get(string(s))
}

type IteratorKey[T any] struct{}

func (IteratorKey[T]) objectGet(o Object[T]) (Value[T], error) {
	return o.Iterator()
}

func ObjectGetFunctionx[T any](
	o Object[T],
	getter objectGetter[T],
) (res Function[T], ok bool, err error) {
	val, err := getter.objectGet(o)
	if err == nil {
		res, ok = val.AsFunction()
	}
	return
}
func ObjectGetFunction[T any](o Object[T], name string) (res Function[T], ok bool, err error) {
	val, err := o.Get(name)
	if err == nil {
		res, ok = val.AsFunction()
	}
	return
}

func ObjectGetIterator[T any](o Object[T]) (res Function[T], ok bool, err error) {
	return ObjectGetFunctionx(o, IteratorKey[T]{})
}

func ObjectEntries[T any](ctx CallbackScope[T], o Object[T]) (Value[T], error) {
	entries, err := ctx.Eval("Object.entries", "")
	if err != nil {
		return nil, err
	}
	fn, ok := entries.AsFunction()
	if !ok {
		return nil, errors.New("Object.entries is not a function")
	}
	return fn.Call(ctx.GlobalThis(), o)
}

// iterate returns a seq.Iter2 exposing a JavaScript iterable as a Go iterator.
// It will return an ErrNotIterable error if the JavaScript value is not an
// object implementing the [Iterable] protocol. An error is returned if
// obtaining the [Iterator] itself resulted in an error. The returned Seq will
// return yield an error value if the JavaScript iterator throws an error during
// iteration.
//
// [Iterable]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterable_protocol
// [Iterator]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterator_protocol
func Iterate[T any](v Value[T]) (iter.Seq2[Value[T], error], error) {
	obj, ok := v.AsObject()
	if !ok {
		return nil, ErrNotIterable
	}
	symIter, ok, err := ObjectGetIterator(obj)
	if err == nil && !ok {
		err = ErrNotIterable
	}
	if err != nil {
		return nil, err
	}
	iterVal, err := symIter.Call(obj)
	if err != nil {
		return nil, err
	}
	iter, ok := iterVal.AsObject()
	if !ok {
		return nil, ErrNotIterable
	}
	next, ok, err := ObjectGetFunction(iter, "next")
	if err == nil && !ok {
		err = ErrNotIterable
	}
	if err != nil {
		return nil, errNotIterable("next is not a function")
	}
	return func(yield func(Value[T], error) bool) {
		for {
			result, err := next.Call(iter)
			if err != nil {
				yield(nil, err)
				return
			}
			resultObj, ok := result.AsObject()
			if !ok {
				break
			}
			done, err := resultObj.Get("done")
			if err != nil {
				yield(nil, err)
				return
			}
			if done.Boolean() {
				return
			}
			val, err := resultObj.Get("value")
			if err != nil || !yield(val, err) {
				return
			}
		}
	}, nil
}
