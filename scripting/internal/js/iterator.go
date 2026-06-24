package js

import (
	"errors"
	"fmt"
	"iter"
	"slices"

	"github.com/gost-dom/browser/internal/log"
)

// The type used for the index, when iterating value iterators
type index = int32

// ValueIterator implements the iterator protocol for a Go iter.Seq[E]. Type
// parameter T is the type parameter for the script engine. The field Resolver
// is a function that is used to generate a JavaScript value for an element of
// type E.
type ValueIterator[E, T any] struct {
	Resolver ValueResolver[E, T]
}

type ValueResolver[T, U any] func(s Scope[U], value T) (Value[U], error)

// InstallValueIterator creates prototype operations the web IDL value iterables
// should implement, including Symbol.iterator and entries. This requires that
// instances must wrap a value providing method All() returning an iter.Seq[T].
//
// See also: https://webidl.spec.whatwg.org/#idl-iterable
func InstallValueIterator[T, U any](class Class[U], entityLookup ValueResolver[T, U]) {
	ValueIterator[T, U]{entityLookup}.InstallPrototype(class)
}

// InstallPairIterator creates prototype operations the web IDL pair iterators
// should implement, including Symbol.iterator and entries. This requires that
// instances must wrap a value providing method All() returning an iter.Seq2[K,V].
//
// See also: https://webidl.spec.whatwg.org/#idl-iterable
func InstallPairIterator[K, V, U any](
	class Class[U],
	keyLookup ValueResolver[K, U],
	valueLookup ValueResolver[V, U],
) {
	PairIterator[K, V, U]{keyLookup, valueLookup}.InstallPrototype(class)
}

func newIterator[T, U any](entityLookup ValueResolver[T, U]) ValueIterator[T, U] {
	return ValueIterator[T, U]{entityLookup}
}

// valueIterable is the interface for a type that can implement the iterable
// interface for a value iterator.
type valueIterable[T any] interface {
	All() iter.Seq[T]
}

// sliceIter is a variation of [valueIterable] that works on a slice of data
// rather that obtaining a [seq.Iterator]
type sliceIter[T any] interface {
	All() []T
}

func withIndex[T any](s iter.Seq[T]) iter.Seq2[index, T] {
	return func(yield func(index, T) bool) {
		var idx index = 0
		for v := range s {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

// entry encodes a single "entry" in "entries". For the ValueIterator, the
// "key", i.e. index, is ignored.
func (i ValueIterator[T, U]) entry(scope Scope[U], _ index, v T) (Value[U], error) {
	return i.encodeValue(scope, v)
}

func (i ValueIterator[T, U]) mapItems(
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
func (i ValueIterator[T, U]) NewIterator(s Scope[U], items iter.Seq[T]) (Value[U], error) {
	return s.NewIterator(i.mapItems(s, items)), nil
}

func (i ValueIterator[T, U]) encodeKey(s Scope[U], idx index) (Value[U], error) {
	return s.NewInt32(idx), nil
}

// Method InstallPrototype creates the following prototype methods:
//
// - Symbol iterator - implementing the iterable protocol
// - "entries" - which all web API implement
func (i ValueIterator[T, U]) InstallPrototype(class Class[U]) {
	fe := iterableOperations[index, T, U]{
		i,
	}
	class.CreateOperation("entries", fe.entries)
	class.CreateOperation("forEach", fe.forEach)
	class.CreateIteratorMethod(i.entries)
}

func (i ValueIterator[T, U]) encodeValue(s Scope[U], v T) (Value[U], error) {
	return i.Resolver(s, v)
}

// seq2 implements iterableSource
func (i ValueIterator[T, U]) seq2(cbCtx CallbackContext[U]) (iter.Seq2[index, T], error) {
	idx, err := i.seq(cbCtx)
	if err != nil {
		return nil, err
	}
	return withIndex(idx), nil
}

// seq creates a new [iter.Seq] for iterating the collection from start.
func (i ValueIterator[T, U]) seq(cbCtx CallbackContext[U]) (iter.Seq[T], error) {
	instance, err1 := As[valueIterable[T]](cbCtx.Instance())
	if err1 == nil {
		return instance.All(), nil
	}
	sliceIter, err2 := As[sliceIter[T]](cbCtx.Instance())
	if err2 == nil {
		return slices.Values(sliceIter.All()), nil
	}
	return nil, err1
}

func (i ValueIterator[T, U]) entries(cbCtx CallbackContext[U]) (res Value[U], err error) {
	defer cbCtx.Logger().
		Debug("JS Function call: Iterator.entries", ThisLogAttr(cbCtx), LogAttr("retVal", res), log.ErrAttr(err))
	instance, err := i.seq(cbCtx)
	if err != nil {
		return nil, fmt.Errorf("iterator.getEntries: %w", err)
	}
	return i.NewIterator(cbCtx, instance)
}

/* -------- PairIterator -------- */

// PairIterator is like [ValueIterator], but implements a pair iterator over an
// iter.Seq2[K,V] value.
type PairIterator[K, V, U any] struct {
	keyLookup   ValueResolver[K, U]
	valueLookup ValueResolver[V, U]
}

type pairIterable[K, V any] interface {
	All() iter.Seq2[K, V]
}

// iterableSource supports a common implementation for both value-, and
// pair collections.
type iterableSource[K, V, T any] interface {
	// seq2 returns a new iter.seq2 for iterating the collection. For
	// value collections, the iterated key is the index.
	seq2(cbCtx CallbackContext[T]) (iter.Seq2[K, V], error)
	// entry creates the iterated value for the entries iterator
	entry(Scope[T], K, V) (Value[T], error)
	// encodeKey encodes a JavaScript value of the iterated key/index in forEach callbacks
	encodeKey(Scope[T], K) (Value[T], error)
	// encodeKey encodes a JavaScript value of the iterated value in forEach callbacks
	encodeValue(Scope[T], V) (Value[T], error)
}

// iterableOperations provides common implementations for iterable methods
// present on both value- and pair iterators.
//
// Differences in the two types are captured in iterableSource, notably the
// entries iterator returns a single value for value iterators, and a key/value
// array for pair iterators.
type iterableOperations[K, V, T any] struct {
	iterableSource[K, V, T]
}

func (e iterableOperations[K, V, U]) entries(cbCtx CallbackContext[U]) (res Value[U], err error) {
	cbCtx.Logger().Debug("JS Function call: Iterator2.entries", ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().
			Debug("JS Function call: Iterator2.entries", ThisLogAttr(cbCtx), LogAttr("retVal", res), log.ErrAttr(err))
	}()
	items, err := e.seq2(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewIterator(e.mapItems(cbCtx, items)), nil
}

func (e iterableOperations[K, V, U]) mapItems(
	cbCtx CallbackContext[U],
	items iter.Seq2[K, V]) iter.Seq2[Value[U], error] {
	return func(yield func(Value[U], error) bool) {
		for k, v := range items {
			res, err := e.entry(cbCtx, k, v)
			if !yield(res, err) {
				return
			}
		}
	}
}

func (e iterableOperations[K, V, U]) forEach(cbCtx CallbackContext[U]) (res Value[U], err error) {
	defer cbCtx.Logger().
		Debug("JS Function call: Iterator.entries", ThisLogAttr(cbCtx), LogAttr("retVal", res), log.ErrAttr(err))
	instance, err1 := e.seq2(cbCtx)
	if err1 != nil {
		return nil, err
	}
	cb, ok := cbCtx.ConsumeArg()
	if !ok {
		return nil, cbCtx.NewTypeError("no argument passed to forEach")
	}
	fn, ok := cb.AsFunction()
	if !ok {
		return nil, cbCtx.NewTypeError("callback not a function")
	}
	for k, v := range instance {
		key, err := e.encodeKey(cbCtx, k)
		if err != nil {
			return nil, err
		}
		val, err := e.encodeValue(cbCtx, v)
		if err != nil {
			return nil, err
		}
		if _, err := fn.Call(cbCtx.This(), val, key); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// entry implements iterableSource, returning an array with key and value
func (i PairIterator[K, V, U]) entry(cbCtx Scope[U], k K, v V) (Value[U], error) {
	kk, err1 := i.keyLookup(cbCtx, k)
	vv, err2 := i.valueLookup(cbCtx, v)
	err := errors.Join(err1, err2)
	res := cbCtx.NewArray(kk, vv) // Safe to call on nil jsValues
	return res, err
}

// InstallPrototype creates the following prototype methods on cls.
//
// - Symbol Iterator
// - "entries" - Returns the same iterator as Symbol Iterator
// - "keys" - Returns an iterator over all keys
// - "values" - Returns an iterator over all values
func (i PairIterator[K, V, U]) InstallPrototype(cls Class[U]) {
	fe := iterableOperations[K, V, U]{
		i,
	}
	cls.CreateOperation("entries", fe.entries)
	cls.CreateOperation("forEach", fe.forEach)
	cls.CreateIteratorMethod(fe.entries)
	keys := newIterator(i.keyLookup)
	values := newIterator(i.valueLookup)
	cls.CreateOperation("keys", func(cbCtx CallbackContext[U]) (Value[U], error) {
		cbCtx.Logger().Debug("JS Function call: Iterator2.keys", ThisLogAttr(cbCtx))
		instance, err := As[pairIterable[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return keys.NewIterator(cbCtx, pairKeys(instance.All()))
	})
	cls.CreateOperation("values", func(cbCtx CallbackContext[U]) (Value[U], error) {
		cbCtx.Logger().Debug("JS Function call: Iterator2.values", ThisLogAttr(cbCtx))
		instance, err := As[pairIterable[K, V]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return values.NewIterator(cbCtx, pairValues(instance.All()))
	})
}

func (i PairIterator[K, V, U]) encodeKey(scope Scope[U], key K) (Value[U], error) {
	return i.keyLookup(scope, key)
}

func (i PairIterator[K, V, U]) encodeValue(scope Scope[U], value V) (Value[U], error) {
	return i.valueLookup(scope, value)
}

// seq2 implements iterableSource
func (i PairIterator[K, V, U]) seq2(cbCtx CallbackContext[U]) (res iter.Seq2[K, V], err error) {
	var it pairIterable[K, V]
	if it, err = As[pairIterable[K, V]](cbCtx.Instance()); err == nil {
		res = it.All()
	}
	return
}

// pairKeys returns a sequences of the keys in a sequence of key/value pairs.
func pairKeys[K, V any](i iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range i {
			if !yield(k) {
				return
			}
		}
	}
}

// pairValues returns a sequences of the values in a sequence of key/value
// pairs.
func pairValues[K, V any](i iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range i {
			if !yield(v) {
				return
			}
		}
	}
}
