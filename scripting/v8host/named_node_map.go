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

func (w namedNodeMapV8Wrapper) CustomInitialiser(ft *v8go.FunctionTemplate) {
	ft.InstanceTemplate().SetIndexedHandler(
		// NOTE: This is the prototype index handler implementation.
		wrapV8CallbackFn(w.scriptHost, func(cbCtx jsCallbackContext) (jsValue, error) {
			instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err != nil {
				return cbCtx.ReturnWithError(err)
			}
			index := int(cbCtx.v8Info.Index())
			item := instance.Item(index)
			if item == nil {
				return cbCtx.ReturnWithValue(nil)
			}
			return encodeEntity(
				cbCtx,
				item,
			)
		}))
}
