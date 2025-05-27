package v8host

import (
	"github.com/gost-dom/browser/dom"

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
	proto.CreateFunction(
		"item",
		func(instance dom.NamedNodeMap, info *argumentHelper) (*v8.Value, error) {
			idx, err := info.consumeInt32()
			item := instance.Item(int(idx))
			if item != nil && err == nil {
				val, err := info.ScriptCtx().getJSInstance(item)
				return val, err
			}
			return v8.Null(iso), err
		},
	)
	instance := builder.NewInstanceBuilder()
	instance.proto.SetIndexedHandler(func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		ctx := host.mustGetContext(info.Context())
		instance, ok := ctx.getCachedNode(info.This())
		nodemap, ok_2 := instance.(dom.NamedNodeMap)
		if ok && ok_2 {
			index := int(info.Index())
			item := nodemap.Item(index)
			if item == nil {
				return v8.Undefined(iso), nil
			}
			return ctx.getJSInstance(item)
		}
		return nil, v8.NewTypeError(iso, "dunno")
	})

	return builder.constructor
}
