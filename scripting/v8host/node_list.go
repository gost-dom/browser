package v8host

import (
	"github.com/gost-dom/browser/dom"
	v8 "github.com/gost-dom/v8go"
)

func (w *nodeListV8Wrapper) CustomInitializer(class jsClass) {
	host := w.scriptHost
	ft := class.(v8Class).ft
	nodeListIterator := newIterator(host,
		func(ctx jsCallbackContext, instance dom.Node) (jsValue, error) {
			return encodeEntity(ctx, instance)
		})
	nodeListIterator.installPrototype(ft)

	instanceTemplate := ft.InstanceTemplate()
	instanceTemplate.SetIndexedHandler(
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			iso := w.iso()
			ctx := host.mustGetContext(info.Context())
			cbCtx := newCallbackContext(host, info)
			instance, ok := ctx.getCachedNode(info.This())
			nodemap, ok_2 := instance.(dom.NodeList)
			if ok && ok_2 {
				index := int(info.Index())
				item := nodemap.Item(index)
				if item == nil {
					return v8.Undefined(iso), nil
				}
				v, err := encodeEntity(cbCtx, item)
				return toV8Value(v), err
			}
			return nil, v8.NewTypeError(iso, "dunno")
		},
	)
}
