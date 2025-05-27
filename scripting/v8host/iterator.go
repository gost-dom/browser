package v8host

import (
	"errors"
	"iter"

	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

type iterator[T any] struct {
	host           *V8ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	entityLookup   entityLookup[T]
}

type entityLookup[T any] func(value T, ctx *V8ScriptContext) (*v8.Value, error)

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
	iterator.ot.Set("next", v8.NewFunctionTemplateWithError(host.iso, iterator.next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(host.iso, iterator.newIterator),
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

func (i iterator[T]) newIteratorInstance(context *V8ScriptContext, items []T) (*v8.Value, error) {
	return i.newIteratorInstanceOfIterable(context, sliceIterable[T]{items})
}

func (i iterator[T]) newIteratorInstanceOfIterable(
	context *V8ScriptContext,
	items iterable[T],
) (*v8.Value, error) {
	seq := items.All()
	next, stop := iter.Pull(seq)

	iterator := &iteratorInstance[T]{
		items: items,
		next:  next,
		stop:  stop,
	}
	res, err := i.ot.NewInstance(context.v8ctx)
	if err == nil {
		return context.cacheNode(newV8Object(context.host.iso, res), iterator)
	}
	return res.Value, err
}

func (i iterator[T]) next(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iteratorInstance[T]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	next := instance.next
	stop := instance.stop
	if item, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(ctx.v8ctx)
	} else {
		value, err1 := i.entityLookup(item, ctx)
		result, err2 := i.createNotDoneIteratorResult(ctx.v8ctx, value)
		return result, errors.Join(err1, err2)
	}
}

func (i iterator[T]) newIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iteratorInstance[T]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	return i.newIteratorInstanceOfIterable(ctx, instance.items)
}

func (i iterator[T]) createDoneIteratorResult(ctx *v8.Context) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return result.Value, nil
}

func (i iterator[T]) createNotDoneIteratorResult(
	ctx *v8.Context,
	value interface{},
) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", false)
	result.Set("value", value)
	return result.Value, nil
}
func (i iterator[T]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := v8.NewFunctionTemplateWithError(iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := i.host.mustGetContext(info.Context())
			instance, err := getWrappedInstance[iterable[T]](info.This())
			if err != nil {
				return nil, err
			}
			return i.newIteratorInstanceOfIterable(ctx, instance)
		})
	prototypeTempl := ft.PrototypeTemplate()
	prototypeTempl.Set("entries", getEntries)
	prototypeTempl.SetSymbol(v8.SymbolIterator(iso), getEntries)
	// iterator for keys/values
	// keys := newIterator(i.host, func(v T, ctx *V8ScriptContext) (*v8.Value, error) {
	// 	return v8.NewValue(iso, v)
	// })
	// values := newIterator(i.host, func(v T, ctx *V8ScriptContext) (*v8.Value, error) {
	// 	return v8.NewValue(iso, v)
	// })
	// prototypeTempl.Set("keys", v8.NewFunctionTemplateWithError(iso,
	// 	func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	// 		ctx := i.host.mustGetContext(info.Context())
	// 		instance, err := getWrappedInstance[iterable[T]](info.This())
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return keys.newIteratorInstanceOfIterable(ctx, Keys[K, V]{instance})
	// 	},
	// ),
	// )
	// prototypeTempl.Set("values", v8.NewFunctionTemplateWithError(iso,
	// 	func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	// 		ctx := i.host.mustGetContext(info.Context())
	// 		instance, err := getWrappedInstance[iterable[T]](info.This())
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return values.newIteratorInstanceOfIterable(ctx, iterValues[K, V]{instance})
	// 	},
	// ),
	// )
}

/* -------- iterator2 -------- */

type iterator2[K, V any] struct {
	host           *V8ScriptHost
	ot             *v8.ObjectTemplate
	resultTemplate *v8.ObjectTemplate
	entityLookup   iteratorValueLookup2[K, V]
}

type iteratorValueLookup2[K, V any] func(k K, v V, ctx *V8ScriptContext) (*v8.Value, *v8.Value, error)

func newIterator2[K, V any](
	host *V8ScriptHost,
	entityLookup iteratorValueLookup2[K, V],
) iterator2[K, V] {
	iso := host.iso
	// TODO, once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iterator := iterator2[K, V]{
		host,
		v8.NewObjectTemplate(host.iso),
		v8.NewObjectTemplate(host.iso),
		entityLookup,
	}
	iterator.ot.Set("next", v8.NewFunctionTemplateWithError(host.iso, iterator.next))
	iterator.ot.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(host.iso, iterator.newIterator),
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
	context *V8ScriptContext,
	items iterable2[K, V],
) (*v8.Value, error) {
	seq := items.All()
	next, stop := iter.Pull2(seq)

	iterator := &iterator2Instance[K, V]{
		items: items,
		next:  next,
		stop:  stop,
	}
	res, err := i.ot.NewInstance(context.v8ctx)
	if err == nil {
		return context.cacheNode(newV8Object(context.host.iso, res), iterator)
	}
	return res.Value, err
}

func (i iterator2[K, V]) next(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iterator2Instance[K, V]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	next := instance.next
	stop := instance.stop
	if k, v, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(ctx.v8ctx)
	} else {
		val1, val2, err1 := i.entityLookup(k, v, ctx)
		result, err2 := i.createNotDoneIteratorResult(ctx.v8ctx, val1, val2)
		return result, errors.Join(err1, err2)
	}
}

func (i iterator2[K, V]) newIterator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := i.host.mustGetContext(info.Context())
	instance, err := getInstanceFromThis[*iterator2Instance[K, V]](ctx, info.This())
	if err != nil {
		return nil, err
	}
	return i.newIteratorInstanceOfIterable(ctx, instance.items)
}

func (i iterator2[K, V]) createDoneIteratorResult(ctx *v8.Context) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return result.Value, nil
}

func (i iterator2[K, V]) createNotDoneIteratorResult(
	ctx *v8.Context,
	key, value *v8go.Value,
) (*v8.Value, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	pair, err := toArray(ctx, key, value)
	if err != nil {
		result.Release()
		return nil, err
	}
	result.Set("done", false)
	result.Set("value", pair)
	return result.Value, nil
}

func (i iterator2[K, V]) installPrototype(ft *v8.FunctionTemplate) {
	iso := i.host.iso
	getEntries := v8.NewFunctionTemplateWithError(iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := i.host.mustGetContext(info.Context())
			instance, err := getWrappedInstance[iterable2[K, V]](info.This())
			if err != nil {
				return nil, err
			}
			return i.newIteratorInstanceOfIterable(ctx, instance)
		})
	prototypeTempl := ft.PrototypeTemplate()
	prototypeTempl.Set("entries", getEntries)
	prototypeTempl.SetSymbol(v8.SymbolIterator(iso), getEntries)
	// iterator for keys/values
	keys := newIterator(i.host, func(k K, ctx *V8ScriptContext) (*v8.Value, error) {
		return v8.NewValue(iso, k)
	})
	values := newIterator(i.host, func(v V, ctx *V8ScriptContext) (*v8.Value, error) {
		return v8.NewValue(iso, v)
	})
	prototypeTempl.Set("keys", v8.NewFunctionTemplateWithError(iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := i.host.mustGetContext(info.Context())
			instance, err := getWrappedInstance[iterable2[K, V]](info.This())
			if err != nil {
				return nil, err
			}
			return keys.newIteratorInstanceOfIterable(ctx, Keys[K, V]{instance})
		},
	),
	)
	prototypeTempl.Set("values", v8.NewFunctionTemplateWithError(iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := i.host.mustGetContext(info.Context())
			instance, err := getWrappedInstance[iterable2[K, V]](info.This())
			if err != nil {
				return nil, err
			}
			return values.newIteratorInstanceOfIterable(ctx, iterValues[K, V]{instance})
		},
	),
	)
}
