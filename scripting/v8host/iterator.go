package v8host

import (
	"errors"
	"iter"

	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type iterator[T any] struct {
	host         *V8ScriptHost
	entityLookup entityLookup[T]
	jsIterator   v8Iterator
}

type entityLookup[T any] func(ctx jsCallbackContext, value T) (jsValue, error)

func newIterator[T any](host *V8ScriptHost, entityLookup entityLookup[T]) iterator[T] {
	iterator := iterator[T]{host, entityLookup, newV8Iterator(host)}
	return iterator
}

type iterable[T any] interface {
	All() iter.Seq[T]
}

type iteratorInstance[T any] struct {
	entity.Entity
	items iterable[T]
	next  func() (T, bool)
	stop  func()
}

func seqOfSlice[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

type sliceIterable[T any] struct {
	items []T
}

func (i sliceIterable[T]) All() iter.Seq[T] {
	return seqOfSlice(i.items)
}

func (i iterator[T]) newIteratorInstance(cbCtx jsCallbackContext, items []T) (jsValue, error) {
	return i.newIteratorInstanceOfIterable(cbCtx, sliceIterable[T]{items})
}

func (i iterator[T]) mapItems(
	cbCtx jsCallbackContext,
	items iterable[T],
) iter.Seq2[jsValue, error] {
	return func(yield func(jsValue, error) bool) {
		for item := range items.All() {
			if !yield(i.entityLookup(cbCtx, item)) {
				return
			}
		}
	}
}

func (i iterator[T]) newIteratorInstanceOfIterable(
	cbCtx jsCallbackContext,
	items iterable[T],
) (jsValue, error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items)), nil
}

func (i iterator[T]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
		instance, err := js.As[iterable[T]](cbCtx.Instance())
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		return i.newIteratorInstanceOfIterable(cbCtx, instance)
	})
	prototypeTempl := ft.PrototypeTemplate()
	prototypeTempl.Set("entries", getEntries)
	prototypeTempl.SetSymbol(v8.SymbolIterator(iso), getEntries)
}

/* -------- iterator2 -------- */

type iterator2[K, V any] struct {
	host        *V8ScriptHost
	keyLookup   entityLookup[K]
	valueLookup entityLookup[V]
	jsIterator  v8Iterator
}

func newIterator2[K, V any](
	host *V8ScriptHost,
	keyLookup entityLookup[K],
	valueLookup entityLookup[V],
) iterator2[K, V] {
	// iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator2[K, V]{
		host,
		keyLookup, valueLookup,
		newV8Iterator(host),
	}
	return iterator
}

type iterable2[K, V any] interface {
	All() iter.Seq2[K, V]
}

type iterator2Instance[K, V any] struct {
	entity.Entity
	items iterable2[K, V]
	next  func() (K, V, bool)
	stop  func()
}

func (i iterator2[K, V]) mapItems(
	cbCtx jsCallbackContext,
	items iter.Seq2[K, V]) iter.Seq2[jsValue, error] {
	return func(yield func(jsValue, error) bool) {
		for k, v := range items {
			var res jsValue
			kk, err1 := i.keyLookup(cbCtx, k)
			vv, err2 := i.valueLookup(cbCtx, v)
			err := errors.Join(err1, err2)
			if err == nil {
				res = cbCtx.ValueFactory().NewArray(kk, vv)
			}
			if !yield(res, err) {
				return
			}
		}
	}
}

func (i iterator2[K, V]) newIteratorInstanceOfIterable(
	cbCtx jsCallbackContext,
	items iterable2[K, V],
) (jsValue, error) {
	return cbCtx.ValueFactory().NewIterator(i.mapItems(cbCtx, items.All())), nil
}

func (i iterator2[K, V]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := wrapV8Callback(i.host,
		func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return cbCtx.ReturnWithError(err)
			}
			return cbCtx.ReturnWithJSValueErr(i.newIteratorInstanceOfIterable(cbCtx, instance))
		})
	prototypeTempl := ft.PrototypeTemplate()
	prototypeTempl.Set("entries", getEntries)
	prototypeTempl.SetSymbol(v8.SymbolIterator(iso), getEntries)
	keys := newIterator(i.host, i.keyLookup)
	values := newIterator(i.host, i.valueLookup)
	prototypeTempl.Set("keys",
		wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			return keys.newIteratorInstanceOfIterable(cbCtx, Keys[K, V]{instance})
		}),
	)
	prototypeTempl.Set("values",
		wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			return values.newIteratorInstanceOfIterable(cbCtx, iterValues[K, V]{instance})
		}),
	)
}
