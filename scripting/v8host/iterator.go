package v8host

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type iterator[T, U any] struct {
	entityLookup entityLookup[T, U]
}

type entityLookup[T, U any] func(ctx js.CallbackScope[U], value T) (js.Value[U], error)

func newIterator[T, U any](entityLookup entityLookup[T, U]) iterator[T, U] {
	return iterator[T, U]{entityLookup}
}

type iterable[T any] interface {
	All() iter.Seq[T]
}
type sliceIter[T any] interface {
	All() []T
}

func (i iterator[T, U]) newIteratorOfSlice(
	cbCtx js.CallbackContext[U],
	items []T,
) (js.Value[U], error) {
	return i.newIterator(cbCtx, slices.Values(items))
}

func (i iterator[T, U]) mapItems(
	cbCtx js.CallbackContext[U],
	items iter.Seq[T],
) iter.Seq2[js.Value[U], error] {
	return func(yield func(js.Value[U], error) bool) {
		for item := range items {
			if !yield(i.entityLookup(cbCtx, item)) {
				return
			}
		}
	}
}

func (i iterator[T, U]) newIterator(
	cbCtx js.CallbackContext[U],
	items iter.Seq[T],
) (js.Value[U], error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items)), nil
}

func (i iterator[T, U]) installPrototype(class js.Class[U]) {
	class.CreatePrototypeMethod("entries", i.entries)
	class.CreateIteratorMethod(i.entries)
}

func (i iterator[T, U]) entries(cbCtx js.CallbackContext[U]) (js.Value[U], error) {
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

type iterator2[K, V, U any] struct {
	keyLookup   entityLookup[K, U]
	valueLookup entityLookup[V, U]
}

func newIterator2[K, V, U any](
	keyLookup entityLookup[K, U],
	valueLookup entityLookup[V, U],
) iterator2[K, V, U] {
	// iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator2[K, V, U]{keyLookup, valueLookup}
	return iterator
}

type iterable2[K, V any] interface {
	All() iter.Seq2[K, V]
}

func (i iterator2[K, V, U]) mapItems(
	cbCtx js.CallbackContext[U],
	items iter.Seq2[K, V]) iter.Seq2[js.Value[U], error] {
	return func(yield func(js.Value[U], error) bool) {
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

func (i iterator2[K, V, U]) newIterator(
	cbCtx js.CallbackContext[U],
	items iterable2[K, V],
) (js.Value[U], error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items.All())), nil
}

func (i iterator2[K, V, U]) installPrototype(cls js.Class[U]) {
	getEntries := func(cbCtx js.CallbackContext[U]) (js.Value[U], error) {
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
	cls.CreatePrototypeMethod("keys", func(cbCtx js.CallbackContext[U]) (js.Value[U], error) {
		instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return keys.newIterator(cbCtx, pairKeys(instance.All()))
	})
	cls.CreatePrototypeMethod("values", func(cbCtx js.CallbackContext[U]) (js.Value[U], error) {
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
