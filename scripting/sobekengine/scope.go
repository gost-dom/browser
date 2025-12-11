package sobekengine

import (
	"errors"
	"fmt"
	"iter"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type scope struct {
	*scriptContext
	global jsObject
}

func newScope(ctx *scriptContext) scope {
	return scope{
		ctx, newObject(ctx, ctx.vm.GlobalObject()),
	}
}

func (s scope) Window() html.Window  { return s.window }
func (s scope) GlobalThis() jsObject { return s.global }

func (s scope) Clock() *clock.Clock { return s.clock }
func (s scope) Constructor(name string) js.Constructor[jsTypeParam] {
	if class, ok := s.classes[name]; ok {
		return class
	}
	return nil
}

func (s scope) GetValue(e entity.ObjectIder) (js.Value[jsTypeParam], bool) {
	v, ok := s.cachedNodes[e.ObjectId()]
	return newValue(s.scriptContext, v), ok
}
func (s scope) SetValue(e entity.ObjectIder, v js.Value[jsTypeParam]) {
	s.cachedNodes[e.ObjectId()] = v.Self().value
}

func (f scope) JSONParse(s string) (js.Value[jsTypeParam], error) {
	parse, err := f.vm.RunString("JSON.parse")
	if err != nil {
		return nil, err
	}
	fn, ok := sobek.AssertFunction(parse)
	if !ok {
		return nil, errors.New("sobek: JSON.parse not in global scope")
	}
	res, err := fn(f.vm.GlobalObject(), f.vm.ToValue(s))
	return newValue(f.scriptContext, res), err
}

func (f scope) JSONStringify(v js.Value[jsTypeParam]) string {
	if o := v.Self().value.ToObject(f.vm); o != nil {
		b, err := o.MarshalJSON()
		if err != nil {
			panic(fmt.Sprintf("gost-dom/sobekhost: JSONStringify: %v", err))
		}
		return string(b)
	}
	panic(fmt.Sprintf("gost-dom/sobekhost: JSONStringify only supports objects. Got: %v", v))
}

func (f scope) NewArray(v ...js.Value[jsTypeParam]) js.Value[jsTypeParam] {
	arr := make([]any, len(v))
	for i, val := range v {
		arr[i] = unwrapValue(val)
	}
	return newObject(f.scriptContext, f.vm.NewArray(arr...))
}

func (f scope) NewBoolean(v bool) js.Value[jsTypeParam] {
	return newValue(f.scriptContext, f.vm.ToValue(v))
}

func (f scope) Undefined() js.Value[jsTypeParam] {
	return newValue(f.scriptContext, sobek.Undefined())
}

func (f scope) Null() js.Value[jsTypeParam] {
	return newValue(f.scriptContext, sobek.Null())
}

func (f scope) NewUint32(v uint32) js.Value[jsTypeParam] {
	return newValue(f.scriptContext, f.vm.ToValue(v))
}

func (f scope) NewInt32(v int32) js.Value[jsTypeParam] {
	return newValue(f.scriptContext, f.vm.ToValue(v))
}

func (f scope) NewInt64(v int64) js.Value[jsTypeParam] {
	return newValue(f.scriptContext, f.vm.ToValue(v))
}

func (f scope) NewString(v string) js.Value[jsTypeParam] {
	return newValue(f.scriptContext, f.vm.ToValue(v))
}

// NewTypeError implements [js.ValueFactory].
func (c scope) NewTypeError(v string) js.Error[jsTypeParam] {
	sobekErrVal := c.vm.NewTypeError(v)
	return scriptError{
		Value: newValue(c.scriptContext, sobekErrVal),
	}
}

func (c scope) NewPromise() js.Promise[jsTypeParam] { return newPromise(c.scriptContext) }

func (c scope) NewObject() jsObject {
	return newObject(c.scriptContext, c.vm.NewObject())
}

func (c scope) NewUint8Array(data []byte) js.Value[jsTypeParam] {
	vm := c.scriptContext.vm
	arrayBuf := vm.NewArrayBuffer(data)
	fVal, err := vm.RunScript("gost-dom/sobekhost:NewUint8Array", "Uint8Array")
	if err != nil {
		panic(fmt.Sprintf("gost-dom/sobekhost: Uint8Array: %v", err))
	}
	ctor, ok := sobek.AssertConstructor(fVal)
	if !ok {
		panic("gost-dom/sobekhost: Uint8Array is not a constructor")
	}
	value, err := ctor(nil, vm.ToValue(arrayBuf))
	if err != nil {
		panic(fmt.Sprintf("gost-dom/sobekhost: Uint8Array call: %v", err))
	}
	return newValue(c.scriptContext, value)
}

func (c scope) NewError(err error) js.Error[jsTypeParam] {
	ctx := c.scriptContext
	return c.NewValueError(newObject(ctx, ctx.vm.NewGoError(err)), err)
}

func (c scope) NewValueError(val jsValue, err error) jsError {
	return scriptError{val, err}
}

func (f scope) NewIterator(
	items iter.Seq2[js.Value[jsTypeParam], error],
) js.Value[jsTypeParam] {
	next, stop := iter.Pull2(items)
	iter := &iterator{next: next, stop: stop}
	jsIterator := f.vm.NewObject()
	obj := newObject(f.scriptContext, jsIterator)
	obj.SetNativeValue(iter)

	var nextJs js.CallbackFunc[jsTypeParam] = func(cbCtx js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
		instance, ok := (cbCtx.This().NativeValue()).(*iterator)
		if !ok {
			return cbCtx.ReturnWithTypeError("Not an iterator instance")
		}
		res := f.vm.NewObject()
		item, err, ok := instance.next()
		res.Set("done", f.vm.ToValue(!ok))
		if !ok {
			instance.stop()
		} else {
			if err == nil {
				res.Set("value", item.Self().value)
			}
		}
		return newObject(f.scriptContext, res), err
	}

	var iteratorJS js.CallbackFunc[jsTypeParam] = func(cbCtx js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
		return f.NewIterator(items), nil
	}

	jsIterator.Set("next", wrapJSCallback(f.scriptContext, nextJs.WithLog("Iterator", "next")))
	jsIterator.SetSymbol(
		sobek.SymIterator,
		wrapJSCallback(f.scriptContext, iteratorJS.WithLog("Iterator", "Symbol.iterator")),
	)

	return obj
}
