package v8host

import (
	"fmt"
	"iter"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type v8Iterator struct {
	host           *V8ScriptHost
	ot             *v8go.ObjectTemplate
	resultTemplate *v8go.ObjectTemplate
}

func newV8Iterator(host *V8ScriptHost) v8Iterator {
	// TODO: once we have weak handles in v8, we can release the iterator when it
	// goes out of scope.
	iso := host.iso
	iterator := v8Iterator{
		host,
		v8go.NewObjectTemplate(host.iso),
		v8go.NewObjectTemplate(host.iso),
	}
	iterator.ot.Set("next", wrapV8Callback(host, iterator.next))
	iterator.ot.SetSymbol(
		v8go.SymbolIterator(iso),
		wrapV8Callback(host, iterator.cloneIterator),
	)
	iterator.ot.SetInternalFieldCount(1)
	return iterator
}

func (i v8Iterator) cloneIterator(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*jsIteratorInstance](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return i.newIterator(cbCtx, instance.items), nil
}

type jsIteratorInstance struct {
	entity.Entity
	items iter.Seq2[jsValue, error]
	next  func() (jsValue, error, bool)
	stop  func()
}

func (i v8Iterator) newIterator(
	cbCtx jsCallbackContext,
	items iter.Seq2[jsValue, error],
) jsValue {
	seq := items
	next, stop := iter.Pull2(seq)

	iterator := &jsIteratorInstance{
		items: items,
		next:  next,
		stop:  stop,
	}
	res, err := i.ot.NewInstance(cbCtx.v8ctx())
	if err != nil {
		panic(fmt.Sprintf("Could not create iterator instance. %s", constants.BUG_ISSUE_URL))
	}
	obj := newV8Object(cbCtx.iso(), res)
	obj.SetNativeValue(iterator)
	return obj
}

func (i v8Iterator) next(cbCtx jsCallbackContext) (jsValue, error) {
	instance, ok := (cbCtx.This().NativeValue()).(*jsIteratorInstance)
	if !ok {
		return cbCtx.ReturnWithTypeError("Not an iterator instance")
	}
	next := instance.next
	stop := instance.stop
	if item, err, ok := next(); !ok {
		stop()
		return i.createDoneIteratorResult(cbCtx.v8ctx())
	} else {
		if err != nil {
			return nil, err
		}
		return i.createNotDoneIteratorResult(cbCtx.v8ctx(), item)
	}
}

func (i v8Iterator) createDoneIteratorResult(ctx *v8go.Context) (jsValue, error) {
	result, err := i.resultTemplate.NewInstance(ctx)
	if err != nil {
		return nil, err
	}
	result.Set("done", true)
	return newV8Object(ctx.Isolate(), result), nil
}

func (i v8Iterator) createNotDoneIteratorResult(
	ctx *v8go.Context,
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

/* -------- iterator[T] -------- */
