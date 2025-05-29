package v8host

import (
	"errors"
	"iter"

	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type iterator[T any] struct {
	host           *V8ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	entityLookup   entityLookup[T]
}

type entityLookup[T any] func(ctx jsCallbackContext, value T) (jsValue, error)

func newIterator[T any](host *V8ScriptHost, entityLookup entityLookup[T]) iterator[T] {
	iso := host.iso
	// TODO: once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator[T]{
		host,
		v8.NewObjectTemplate(host.iso),
		v8.NewObjectTemplate(host.iso),
		entityLookup,
	}
	iterator.ot.Set("next", wrapV8Callback(host, iterator.next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		wrapV8Callback(host, iterator.newIterator),
	)
	iterator.ot.SetInternalFieldCount(1)
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

func (i iterator[T]) newIteratorInstanceOfIterable(
	cbCtx jsCallbackContext,
	items iterable[T],
) (jsValue, error) {
	seq := items.All()
	next, stop := iter.Pull(seq)

	iterator := &iteratorInstance[T]{
		items: items,
		next:  next,
		stop:  stop,
	}
	res, err := i.ot.NewInstance(cbCtx.v8ctx())
	if err == nil {
		obj := newV8Object(cbCtx.iso(), res)
		obj.SetNativeValue(iterator)
		return obj, nil
	}
	return newV8Object(cbCtx.iso(), res), err
}

func (i iterator[T]) next(cbCtx jsCallbackContext) (jsValue, error) {
	instance, ok := (cbCtx.This().NativeValue()).(*iteratorInstance[T])
	if !ok {
		return cbCtx.ReturnWithTypeError("Not an iterator instance")
	}
	next := instance.next
	stop := instance.stop
	if item, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(cbCtx.v8ctx())
	} else {
		value, err := i.entityLookup(cbCtx, item)
		if err != nil {
			return nil, err
		}
		return i.createNotDoneIteratorResult(cbCtx.v8ctx(), value)
	}
}

func (i iterator[T]) newIterator(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*iteratorInstance[T]](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return i.newIteratorInstanceOfIterable(cbCtx, instance.items)
}

func (i iterator[T]) createDoneIteratorResult(ctx *v8.Context) (jsValue, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return newV8Object(ctx.Isolate(), result), nil
}

func (i iterator[T]) createNotDoneIteratorResult(
	ctx *v8.Context,
	value jsValue,
) (jsValue, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", false)
	result.Set("value", value.Self().v8Value())
	return newV8Object(ctx.Isolate(), result), nil
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
	host           *V8ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	keyLookup      entityLookup[K]
	valueLookup    entityLookup[V]
}

func newIterator2[K, V any](
	host *V8ScriptHost,
	keyLookup entityLookup[K],
	valueLookup entityLookup[V],
) iterator2[K, V] {
	iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator2[K, V]{
		host,
		v8.NewObjectTemplate(host.iso),
		v8.NewObjectTemplate(host.iso),
		keyLookup, valueLookup,
	}
	iterator.ot.Set("next", wrapV8Callback(host, iterator.next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		wrapV8Callback(host, iterator.newIterator),
	)
	iterator.ot.SetInternalFieldCount(1)
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

func (i iterator2[K, V]) newIteratorInstanceOfIterable(
	cbCtx jsCallbackContext,
	items iterable2[K, V],
) (jsValue, error) {
	seq := items.All()
	next, stop := iter.Pull2(seq)

	iterator := &iterator2Instance[K, V]{
		items: items,
		next:  next,
		stop:  stop,
	}
	res, err := i.ot.NewInstance(cbCtx.v8ctx())
	if err == nil {
		obj := newV8Object(cbCtx.iso(), res)
		obj.SetNativeValue(iterator)
		return obj, nil
	}
	return newV8Object(cbCtx.iso(), res), err
}

func (i iterator2[K, V]) next(cbCtx jsCallbackContext) (jsValue, error) {
	instance, ok := (cbCtx.This().NativeValue()).(*iterator2Instance[K, V])
	if !ok {
		return cbCtx.ReturnWithTypeError("Not an iterator instance")
	}
	next := instance.next
	stop := instance.stop
	if k, v, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(cbCtx)
	} else {
		val1, err1 := i.keyLookup(cbCtx, k)
		val2, err2 := i.valueLookup(cbCtx, v)
		result, err3 := i.createNotDoneIteratorResult(cbCtx, val1, val2)
		return result, errors.Join(err1, err2, err3)
	}
}

func (i iterator2[K, V]) newIterator(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*iterator2Instance[K, V]](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return cbCtx.ReturnWithJSValueErr(i.newIteratorInstanceOfIterable(cbCtx, instance.items))
}

func (i iterator2[K, V]) createDoneIteratorResult(cbCtx jsCallbackContext) (jsValue, error) {
	result, err := i.resultTemplate.NewInstance(cbCtx.v8ctx())
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return newV8Object(cbCtx.iso(), result), nil
}

func (i iterator2[K, V]) createNotDoneIteratorResult(
	cbCtx jsCallbackContext,
	key, value jsValue,
) (jsValue, error) {
	result, err := i.resultTemplate.NewInstance(cbCtx.v8ctx())
	if err != nil {
		return nil, err
	}
	pair, err := toArray(
		cbCtx.v8ctx(),
		assertV8Value(key).v8Value(),
		assertV8Value(value).v8Value(),
	)
	if err != nil {
		result.Release()
		return nil, err
	}
	result.Set("done", false)
	result.Set("value", pair)
	return newV8Value(cbCtx.iso(), result.Value), nil
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
				return cbCtx.ReturnWithError(err)
			}
			return keys.newIteratorInstanceOfIterable(cbCtx, Keys[K, V]{instance})
		}),
	)
	prototypeTempl.Set("values",
		wrapV8Callback(i.host, func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[iterable2[K, V]](cbCtx.Instance())
			if err != nil {
				return cbCtx.ReturnWithError(err)
			}
			return values.newIteratorInstanceOfIterable(cbCtx, iterValues[K, V]{instance})
		}),
	)
}
