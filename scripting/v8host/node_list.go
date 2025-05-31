package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func (w *nodeListV8Wrapper) CustomInitialiser(ft *v8.FunctionTemplate) {
	host := w.scriptHost
	iso := w.iso()
	prototype := ft.PrototypeTemplate()
	nodeListIterator := newIterator(host,
		func(ctx jsCallbackContext, instance dom.Node) (jsValue, error) {
			return ctx.ScriptCtx().getJSInstance(instance)
		},
	)
	prototype.SetSymbol(v8.SymbolIterator(iso),
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			nodeList, err := js.As[dom.NodeList](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			return nodeListIterator.newIteratorInstance(cbCtx, nodeList.All())
		}))

	instanceTemplate := ft.InstanceTemplate()
	instanceTemplate.SetIndexedHandler(
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			instance, ok := ctx.getCachedNode(info.This())
			nodemap, ok_2 := instance.(dom.NodeList)
			if ok && ok_2 {
				index := int(info.Index())
				item := nodemap.Item(index)
				if item == nil {
					return v8.Undefined(iso), nil
				}
				v, err := ctx.getJSInstance(item)
				return assertV8Value(v).v8Value(), err
			}
			return nil, v8.NewTypeError(iso, "dunno")
		},
	)
}
