package gojahost

import (
	"fmt"
	"iter"

	"github.com/dop251/goja"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaValueFactory struct {
	*GojaContext
}

func newGojaValueFactory(c *GojaContext) js.ValueFactory[jsTypeParam] {
	return gojaValueFactory{c}
}

func (f gojaValueFactory) JSONParse(s string) (js.Value[jsTypeParam], error) {
	o := f.vm.NewObject()
	err := o.UnmarshalJSON([]byte(s))
	return newGojaObject(f.GojaContext, o), err
}

func (f gojaValueFactory) JSONStringify(v js.Value[jsTypeParam]) string {
	if o := v.Self().value.ToObject(f.vm); o != nil {
		b, err := o.MarshalJSON()
		if err == nil {
			panic(fmt.Sprintf("Goja JSON marshalling failed: %v", err))
		}
		return string(b)
	}
	panic("Goja only support JSON for objects")
}

func (f gojaValueFactory) NewArray(v ...js.Value[jsTypeParam]) js.Value[jsTypeParam] {
	arr := make([]any, len(v))
	for i, val := range v {
		arr[i] = toGojaValue(val)
	}
	return newGojaObject(f.GojaContext, f.vm.NewArray(arr...))
}

func (f gojaValueFactory) NewBoolean(v bool) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) Null() js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, goja.Null())
}

func (f gojaValueFactory) NewUint32(v uint32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewInt32(v int32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewInt64(v int64) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewString(v string) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewTypeError(v string) error {
	panic(f.vm.NewTypeError(v))
}

func (f gojaValueFactory) NewIterator(
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
