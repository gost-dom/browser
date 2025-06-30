package gojahost

import (
	"errors"
	"fmt"
	"iter"

	"github.com/dop251/goja"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaScope struct {
	*GojaContext
	global js.Object[jsTypeParam]
}

func newGojaScope(ctx *GojaContext) gojaScope {
	return gojaScope{
		ctx, newGojaObject(ctx, ctx.vm.GlobalObject()),
	}
}

func (s gojaScope) Window() html.Window                { return s.window }
func (s gojaScope) GlobalThis() js.Object[jsTypeParam] { return s.global }

func (s gojaScope) Clock() *clock.Clock { return s.clock }
func (s gojaScope) Constructor(name string) js.Constructor[jsTypeParam] {
	if class, ok := s.classes[name]; ok {
		return class
	}
	return nil
}

func (s gojaScope) GetValue(e entity.ObjectIder) (js.Value[jsTypeParam], bool) {
	v, ok := s.cachedNodes[e.ObjectId()]
	return newGojaValue(s.GojaContext, v), ok
}
func (s gojaScope) SetValue(e entity.ObjectIder, v js.Value[jsTypeParam]) {
	s.cachedNodes[e.ObjectId()] = v.Self().value
}

type gojaConstructor struct {
	ctx       *GojaContext
	class     *gojaClass
	prototype *goja.Object
}

func (f gojaScope) JSONParse(s string) (js.Value[jsTypeParam], error) {
	parse, err := f.vm.RunString("JSON.parse")
	if err != nil {
		return nil, err
	}
	fn, ok := goja.AssertFunction(parse)
	if !ok {
		return nil, errors.New("Goja error, retrieving JSON.parse")
	}
	res, err := fn(f.vm.GlobalObject(), f.vm.ToValue(s))
	return newGojaValue(f.GojaContext, res), err
}

func (f gojaScope) JSONStringify(v js.Value[jsTypeParam]) string {
	if o := v.Self().value.ToObject(f.vm); o != nil {
		b, err := o.MarshalJSON()
		if err == nil {
			panic(fmt.Sprintf("Goja JSON marshalling failed: %v", err))
		}
		return string(b)
	}
	panic("Goja only support JSON for objects")
}

func (f gojaScope) NewArray(v ...js.Value[jsTypeParam]) js.Value[jsTypeParam] {
	arr := make([]any, len(v))
	for i, val := range v {
		arr[i] = toGojaValue(val)
	}
	return newGojaObject(f.GojaContext, f.vm.NewArray(arr...))
}

func (f gojaScope) NewBoolean(v bool) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaScope) Undefined() js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, goja.Undefined())
}

func (f gojaScope) Null() js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, goja.Null())
}

func (f gojaScope) NewUint32(v uint32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaScope) NewInt32(v int32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaScope) NewInt64(v int64) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaScope) NewString(v string) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaScope) NewTypeError(v string) error {
	panic(f.vm.NewTypeError(v))
}

func (c gojaScope) NewPromise() js.Promise[jsTypeParam] { return newGojaPromise(c.GojaContext) }

func (c gojaScope) NewObject() js.Object[jsTypeParam] {
	return newGojaObject(c.GojaContext, c.vm.NewObject())
}

func (c gojaScope) NewError(err error) js.Error[jsTypeParam] {
	return newGojaError(c.GojaContext, err)
}

func (f gojaScope) NewIterator(
	items iter.Seq2[js.Value[jsTypeParam], error],
) js.Value[jsTypeParam] {
	next, stop := iter.Pull2(items)
	iter := &gojaIteratorInstance{next: next, stop: stop}
	gojaObj := f.vm.NewObject()
	obj := newGojaObject(f.GojaContext, gojaObj)
	obj.SetNativeValue(iter)

	gojaObj.Set(
		"next",
		wrapJSCallback(
			f.GojaContext,
			func(cbCtx js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
				instance, ok := (cbCtx.This().NativeValue()).(*gojaIteratorInstance)
				if !ok {
					return cbCtx.ReturnWithTypeError("Not an iterator instance")
				}
				res := f.vm.NewObject()
				item, err, ok := instance.next()
				res.Set("done", instance.vm.ToValue(!ok))
				if !ok {
					instance.stop()
				} else {
					if err == nil {
						res.Set("value", item.Self().value)
					}
				}
				return newGojaObject(f.GojaContext, res), err
			},
		),
	)
	gojaObj.SetSymbol(
		goja.SymIterator,
		wrapJSCallback(
			f.GojaContext,
			func(cbCtx js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
				return f.NewIterator(items), nil
			},
		),
	)

	return obj
}
