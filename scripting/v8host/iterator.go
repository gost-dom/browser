package v8host

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type iterator[T any] struct {
	entityLookup entityLookup[T]
}

type entityLookup[T any] func(ctx jsCallbackContext, value T) (jsValue, error)

func newIterator[T any](entityLookup entityLookup[T]) iterator[T] {
	return iterator[T]{entityLookup}
}

type iterable[T any] interface {
	All() iter.Seq[T]
}
type sliceIter[T any] interface {
	All() []T
}

func (i iterator[T]) newIteratorOfSlice(cbCtx jsCallbackContext, items []T) (jsValue, error) {
	return i.newIterator(cbCtx, slices.Values(items))
}

func (i iterator[T]) mapItems(
	cbCtx jsCallbackContext,
	items iter.Seq[T],
) iter.Seq2[jsValue, error] {
	return func(yield func(jsValue, error) bool) {
		for item := range items {
			if !yield(i.entityLookup(cbCtx, item)) {
				return
			}
		}
	}
}

func (i iterator[T]) newIterator(cbCtx jsCallbackContext, items iter.Seq[T]) (jsValue, error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items)), nil
}

func (i iterator[T]) installPrototype(class jsClass) {
	class.CreatePrototypeMethod("entries", i.entries)
	class.CreateIteratorMethod(i.entries)
}

func (i iterator[T]) entries(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err1 := js.As[iterable[T]](cbCtx.Instance())
	if err1 == nil {
		return i.newIterator(cbCtx, instance.All())
	}
	sliceIter, err2 := js.As[sliceIter[T]](cbCtx.Instance())
	if err2 == nil {
		return i.newIteratorOfSlice(cbCtx, sliceIter.All())
	}
	return nil, fmt.Errorf("iterator.getEntries: %w", errors.Join(err1, err2))
}

/* -------- iterator2 -------- */

type iterator2[K, V any] struct {
	keyLookup   entityLookup[K]
	valueLookup entityLookup[V]
}

func newIterator2[K, V any](
	keyLookup entityLookup[K],
	valueLookup entityLookup[V],
) iterator2[K, V] {
	// iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator2[K, V]{keyLookup, valueLookup}
	return iterator
}

type iterable2[K, V any] interface {
	All() iter.Seq2[K, V]
}

func (i iterator2[K, V]) mapItems(
	cbCtx jsCallbackContext,
	items iter.Seq2[K, V]) iter.Seq2[jsValue, error] {
	return func(yield func(jsValue, error) bool) {
		for k, v := range items {
			kk, err1 := i.keyLookup(cbCtx, k)
			vv, err2 := i.valueLookup(cbCtx, v)
			err := errors.Join(err1, err2)
			res := cbCtx.ValueFactory().NewArray(kk, vv) // Safe to call on nil jsValues
			if !yield(res, err) {
				return
			}
		}
	}
}

func (i iterator2[K, V]) newIterator(
	cbCtx jsCallbackContext,
	items iterable2[K, V],
) (jsValue, error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items.All())), nil
}

func (i iterator2[K, V]) installPrototype(cls js.Class[jsTypeParam]) {
	getEntries := func(cbCtx jsCallbackContext) (jsValue, error) {
		instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return i.newIterator(cbCtx, instance)
	}
	cls.CreatePrototypeMethod("entries", getEntries)
	cls.CreateIteratorMethod(getEntries)
	keys := newIterator(i.keyLookup)
	values := newIterator(i.valueLookup)
	cls.CreatePrototypeMethod("keys", func(cbCtx jsCallbackContext) (jsValue, error) {
		instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return keys.newIterator(cbCtx, pairKeys(instance.All()))
	})
	cls.CreatePrototypeMethod("values", func(cbCtx jsCallbackContext) (jsValue, error) {
		instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return values.newIterator(cbCtx, pairValues(instance.All()))
	})
}

// pairKeys returns a sequences of the keys in a sequence of key/value pairs
func pairKeys[K, V any](i iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range i {
			if !yield(k) {
				return
			}
		}
	}
}

// pairValues returns a sequences of the values in a sequence of key/value pairs
func pairValues[K, V any](i iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range i {
			if !yield(v) {
				return
			}
		}
	}
}
