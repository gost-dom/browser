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
		func(instance dom.Node, ctx *V8ScriptContext) (*v8.Value, error) {
			v, err := ctx.getJSInstance(instance)
			return v.v8Value(), err
		},
	)
	prototype.SetSymbol(v8.SymbolIterator(iso),
		wrapV8Callback(host, func(cbCtx *argumentHelper) js.CallbackRVal {
			nodeList, err := js.As[dom.NodeList](cbCtx.Instance())
			if err != nil {
				return cbCtx.ReturnWithError(err)
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
				return v.v8Value(), err
			}
			return nil, v8.NewTypeError(iso, "dunno")
		},
	)
}
