package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *NodeList[T]) CustomInitializer(class js.Class[T]) {
	nodeListIterator := js.NewIterator(
		func(ctx js.Scope[T], instance dom.Node) (js.Value[T], error) {
			return codec.EncodeEntityScoped(ctx, instance)
		})
	nodeListIterator.InstallPrototype(class)

	class.CreateIndexedHandler(
		js.WithGetterCallback(
			func(info js.CallbackScope[T], index int) (js.Value[T], error) {
				instance := info.This().NativeValue()
				if nodemap, ok := instance.(dom.NodeList); ok {
					item := nodemap.Item(index)
					if item == nil {
						return nil, nil
					}
					return codec.EncodeEntity(info, item)
				}
				return nil, info.ValueFactory().NewTypeError("dunno")
			},
		))
}
