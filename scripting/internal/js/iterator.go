package js

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/gost-dom/browser/internal/log"
)

// Iterator implements the iterator protocol for a Go iter.Seq[E]. Type
// parameter T is the type parameter for the script engine. The field Resolver
// is a function that is used to generate a JavaScript value for an element of
// type E.
type Iterator[E, T any] struct {
	Resolver ValueResolver[E, T]
}

type ValueResolver[T, U any] func(s Scope[U], value T) (Value[U], error)

func NewIterator[T, U any](entityLookup ValueResolver[T, U]) Iterator[T, U] {
	return Iterator[T, U]{entityLookup}
}

type iterable[T any] interface {
	All() iter.Seq[T]
}
type sliceIter[T any] interface {
	All() []T
}

func (i Iterator[T, U]) newIteratorOfSlice(s Scope[U], items []T) (Value[U], error) {
	return i.NewIterator(s, slices.Values(items))
}

func (i Iterator[T, U]) mapItems(
	scope Scope[U],
	items iter.Seq[T],
) iter.Seq2[Value[U], error] {
	return func(yield func(Value[U], error) bool) {
		for item := range items {
			if !yield(i.Resolver(scope, item)) {
				return
			}
		}
	}
}

// Method NewIterator on an iterator returns a new iterator, iterating from the
// beginning of the specified sequence.
func (i Iterator[T, U]) NewIterator(s Scope[U], items iter.Seq[T]) (Value[U], error) {
	return s.NewIterator(i.mapItems(s, items)), nil
}

// Method InstallPrototype creates the following prototype methods:
//
// - Symbol iterator - implementing the iterable protocol
// - "entries" - which all web API implement
func (i Iterator[T, U]) InstallPrototype(class Class[U]) {
	class.CreatePrototypeMethod("entries", i.entries)
	class.CreateIteratorMethod(i.entries)
}

func (i Iterator[T, U]) entries(cbCtx CallbackContext[U]) (res Value[U], err error) {
	defer cbCtx.Logger().
		Debug("JS Function call: Iterator.entries", ThisLogAttr(cbCtx), LogAttr("retVal", res), log.ErrAttr(err))

	instance, err1 := As[iterable[T]](cbCtx.Instance())
	if err1 == nil {
		return i.NewIterator(cbCtx, instance.All())
	}
	sliceIter, err2 := As[sliceIter[T]](cbCtx.Instance())
	if err2 == nil {
		return i.newIteratorOfSlice(cbCtx, sliceIter.All())
	}
	return nil, fmt.Errorf("iterator.getEntries: %w", errors.Join(err1, err2))
}

/* -------- iterator2 -------- */

// Iterator2 is like [Iterator], but implements a key/value iterator over an
// iter.Seq2[K,V] value.
type Iterator2[K, V, U any] struct {
	keyLookup   ValueResolver[K, U]
	valueLookup ValueResolver[V, U]
}

func NewIterator2[K, V, U any](
	keyLookup ValueResolver[K, U],
	valueLookup ValueResolver[V, U],
) Iterator2[K, V, U] {
	iterator := Iterator2[K, V, U]{keyLookup, valueLookup}
	return iterator
}

type iterable2[K, V any] interface {
	All() iter.Seq2[K, V]
}

func (i Iterator2[K, V, U]) mapItems(
	cbCtx CallbackContext[U],
	items iter.Seq2[K, V]) iter.Seq2[Value[U], error] {
	return func(yield func(Value[U], error) bool) {
		for k, v := range items {
			kk, err1 := i.keyLookup(cbCtx, k)
			vv, err2 := i.valueLookup(cbCtx, v)
			err := errors.Join(err1, err2)
			res := cbCtx.NewArray(kk, vv) // Safe to call on nil jsValues
			if !yield(res, err) {
				return
			}
		}
	}
}

func (i Iterator2[K, V, U]) newIterator(
	cbCtx CallbackContext[U],
	items iterable2[K, V],
) (Value[U], error) {
	return cbCtx.NewIterator(i.mapItems(cbCtx, items.All())), nil
}

// InstallPrototype creates the following prototype methods on cls.
//
// - Symbol Iterator
// - "entries" - Returns the same iterator as Symbol Iterator
// - "keys" - Returns an iterator over all keys
// - "values" - Returns an iterator over all values
func (i Iterator2[K, V, U]) InstallPrototype(cls Class[U]) {
	getEntries := func(cbCtx CallbackContext[U]) (res Value[U], err error) {
		cbCtx.Logger().Debug("JS Function call: Iterator2.entries", ThisLogAttr(cbCtx))
		defer func() {
			cbCtx.Logger().
				Debug("JS Function call: Iterator2.entries", ThisLogAttr(cbCtx), LogAttr("retVal", res), log.ErrAttr(err))
		}()
		instance, err := As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return i.newIterator(cbCtx, instance)
	}
	cls.CreatePrototypeMethod("entries", getEntries)
	cls.CreateIteratorMethod(getEntries)
	keys := NewIterator(i.keyLookup)
	values := NewIterator(i.valueLookup)
	cls.CreatePrototypeMethod("keys", func(cbCtx CallbackContext[U]) (Value[U], error) {
		cbCtx.Logger().Debug("JS Function call: Iterator2.keys", ThisLogAttr(cbCtx))
		instance, err := As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return keys.NewIterator(cbCtx, pairKeys(instance.All()))
	})
	cls.CreatePrototypeMethod("values", func(cbCtx CallbackContext[U]) (Value[U], error) {
		cbCtx.Logger().Debug("JS Function call: Iterator2.values", ThisLogAttr(cbCtx))
		instance, err := As[iterable2[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return values.NewIterator(cbCtx, pairValues(instance.All()))
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
