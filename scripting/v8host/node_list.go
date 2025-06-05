package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *nodeListV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
	nodeListIterator := newIterator(
		func(ctx js.CallbackContext[T], instance dom.Node) (js.Value[T], error) {
			return encodeEntity(ctx, instance)
		})
	nodeListIterator.installPrototype(class)

	class.CreateIndexedHandler(
		func(info js.GetterCallbackContext[T, int]) (js.Value[T], error) {
			instance := info.This().NativeValue()
			if nodemap, ok := instance.(dom.NodeList); ok {
				index := int(info.Key())
				item := nodemap.Item(index)
				if item == nil {
					return nil, nil
				}
				return codec.EncodeEntity(info, item)
			}
			return nil, info.ValueFactory().NewTypeError("dunno")
		},
	)
}
