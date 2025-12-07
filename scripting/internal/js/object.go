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
		res, ok = AsFunction(val)
	}
	return
}
func ObjectGetFunction[T any](o Object[T], name string) (res Function[T], ok bool, err error) {
	val, err := o.Get(name)
	if err == nil {
		res, ok = AsFunction(val)
	}
	return
}

func ObjectGetIterator[T any](o Object[T]) (res Function[T], ok bool, err error) {
	return ObjectGetFunctionx(o, IteratorKey[T]{})
}

func ObjectKeys[T any](ctx CallbackScope[T], o Object[T]) (Value[T], error) {
	entries, err := ctx.Eval("Reflect.ownKeys", "")
	if err != nil {
		return nil, err
	}
	fn, ok := entries.AsFunction()
	if !ok {
		return nil, errors.New("Reflect.ownKeys is not a function")
	}
	return fn.Call(ctx.GlobalThis(), o)
}

func ObjectOwnPropertyDescriptor[T any](
	ctx CallbackScope[T],
	o Object[T],
	p Value[T],
) (PropertyDescriptor[T], error) {
	entries, err := ctx.Eval("Object.getOwnPropertyDescriptor", "")
	if err != nil {
		return nil, err
	}
	fn, ok := entries.AsFunction()
	if !ok {
		return nil, errors.New("Object.getOwnPropertyDescriptor is not a function")
	}
	v, err := fn.Call(ctx.GlobalThis(), o, p)
	if err != nil {
		return nil, err
	}
	if v == nil || v.IsUndefined() {
		return nil, nil
	}
	obj, _ := v.AsObject()
	return propertyDescriptor[T]{obj}, nil
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
// The returned Seq will yield an error value if the value is not an Iterable,
// or JavaScript iterator throws an error during iteration.
//
// [Iterable]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterable_protocol
// [Iterator]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterator_protocol
func Iterate[T any](v Value[T]) iter.Seq2[Value[T], error] {
	return func(yield func(Value[T], error) bool) {
		obj, ok := v.AsObject()
		if !ok {
			yield(nil, ErrNotIterable)
			return
		}
		symIter, ok, err := ObjectGetIterator(obj)
		if err == nil && !ok {
			err = ErrNotIterable
		}
		if err != nil {
			yield(nil, err)
			return
		}
		iterVal, err := symIter.Call(obj)
		if err != nil {
			yield(nil, err)
			return
		}
		iter, ok := iterVal.AsObject()
		if !ok {
			yield(nil, ErrNotIterable)
			return
		}
		next, ok, err := ObjectGetFunction(iter, "next")
		if err == nil && !ok {
			err = ErrNotIterable
		}
		if err != nil {
			yield(nil, fmt.Errorf("%w: %w", errNotIterable("iterator.next: "), err))
			return
		}
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
	}
}

/* -------- PropertyDescriptor -------- */

type PropertyDescriptor[T any] interface {
	Object[T]
	Enumerable() bool
}

type propertyDescriptor[T any] struct {
	Object[T]
}

func (d propertyDescriptor[T]) Enumerable() bool {
	val, err := d.Get("enumerable")
	return err == nil && val.Boolean()
}

type PropertyDescriptorIter[T any] struct {
	PropertyDescriptor[T]
	Key Value[T]
}

// ObjectEnumerableOwnPropertyKeys iterates over all enumerable [[OwnPropertyKeys]]
func ObjectEnumerableOwnPropertyKeys[T any](
	scope CallbackScope[T],
	obj Object[T],
) iter.Seq2[Value[T], error] {
	return func(yield func(Value[T], error) bool) {
		keys, err := ObjectKeys(scope, obj)
		if err != nil {
			yield(nil, scope.NewTypeError(err.Error()))
			return
		}

		var (
			key Value[T]
			pd  PropertyDescriptor[T]
		)
		for key, err = range Iterate(keys) {
			if err != nil {
				break
			}
			if pd, err = ObjectOwnPropertyDescriptor(scope, obj, key); err != nil {
				break
			}
			if pd == nil {
				continue
			}
			if pd.Enumerable() && !yield(key, nil) {
				return
			}
		}
		if err != nil {
			yield(nil, err)
		}
	}
}
