package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"

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
			return nil, v8.NewTypeError(iso, "Not an instance of Attr")
		}
	}
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp("name", dom.Attr.Name)
	proto.CreateReadWriteProp("value", dom.Attr.Value, dom.Attr.SetValue)
	return builder.constructor
}

func (w namedNodeMapV8Wrapper) CustomInitializer(ft *v8go.FunctionTemplate) {
	ft.InstanceTemplate().SetIndexedHandler(
		// NOTE: This is the prototype index handler implementation.
		wrapV8IndexedGetterCallbackFn(
			w.scriptHost,
			func(cbCtx js.GetterCallbackContext[jsTypeParam, int]) (jsValue, error) {
				instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
				if err != nil {
					return nil, err
				}
				index := int(cbCtx.Key())
				item := instance.Item(index)
				if item == nil {
					return nil, nil
				}
				return encodeEntity(cbCtx, item)
			},
		))
}
