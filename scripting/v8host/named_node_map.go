package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

func createAttr(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := newIllegalConstructorBuilder[dom.Attr](host)
	builder.instanceLookup = func(ctx *V8ScriptContext, this *v8.Object) (dom.Attr, error) {
		instance, ok := ctx.getCachedNode(this)
		if e, e_ok := instance.(dom.Attr); e_ok && ok {
			return e, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp("name", dom.Attr.Name)
	proto.CreateReadWriteProp("value", dom.Attr.Value, dom.Attr.SetValue)
	return builder.constructor
}

func createNamedNodeMap(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := newIllegalConstructorBuilder[dom.NamedNodeMap](host)
	builder.instanceLookup = func(ctx *V8ScriptContext, this *v8.Object) (dom.NamedNodeMap, error) {
		instance, ok := ctx.getCachedNode(this)
		if e, e_ok := instance.(dom.NamedNodeMap); e_ok && ok {
			return e, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp2(
		"length",
		func(instance dom.NamedNodeMap, ctx *V8ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, int32(instance.Length()))
		},
	)
	proto.proto.Set("item",
		wrapV8Callback(host, func(cbCtx *argumentHelper) js.CallbackRVal {
			idx, err0 := cbCtx.consumeInt32()
			instance, err1 := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err := errors.Join(err0, err1); err != nil {
				return cbCtx.ReturnWithError(err)
			}
			item := instance.Item(int(idx))
			if item != nil {
				return cbCtx.ReturnWithJSValueErr(cbCtx.ScriptCtx().getJSInstance(item))
			}
			return cbCtx.ReturnWithValue(v8.Null(iso))
		}),
		v8.ReadOnly,
	)
	instance := builder.NewInstanceBuilder()
	instance.proto.SetIndexedHandler(
		// NOTE: This is the prototype index handler implementation.
		wrapV8CallbackFn(host, func(cbCtx *argumentHelper) js.CallbackRVal {
			instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err != nil {
				return cbCtx.ReturnWithError(err)
			}
			index := int(cbCtx.v8Info.Index())
			item := instance.Item(index)
			if item == nil {
				return cbCtx.ReturnWithValue(nil)
			}
			return cbCtx.ReturnWithJSValueErr(cbCtx.ScriptCtx().getJSInstance(item))
		}))
	return builder.constructor
}
