package v8host

import (
	"errors"
	"iter"
	"slices"

	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type iterator[T any] struct {
	host         *V8ScriptHost
	entityLookup entityLookup[T]
}

type entityLookup[T any] func(ctx jsCallbackContext, value T) (jsValue, error)

func newIterator[T any](host *V8ScriptHost, entityLookup entityLookup[T]) iterator[T] {
	return iterator[T]{host, entityLookup}
}

type iterable[T any] interface {
	All() iter.Seq[T]
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

func (i iterator[T]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
		instance, err := js.As[iterable[T]](cbCtx.Instance())
		if err != nil {
			return nil, err
		}
		return i.newIterator(cbCtx, instance.All())
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
	}
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

func (i iterator2[K, V]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := wrapV8Callback(i.host,
		func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			return i.newIterator(cbCtx, instance)
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
			return keys.newIterator(cbCtx, pairKeys(instance.All()))
		}),
	)
	prototypeTempl.Set("values",
		wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			return values.newIterator(cbCtx, pairValues(instance.All()))
		}),
	)
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
