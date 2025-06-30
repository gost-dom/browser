package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w NamedNodeMap[T]) CustomInitializer(class js.Class[T]) {
	iterator := js.NewIterator(
		func(ctx js.Scope[T], instance dom.Attr) (js.Value[T], error) {
			return codec.EncodeEntityScoped(ctx, instance)
		})
	iterator.InstallPrototype(class)
	class.CreateIndexedHandler(
		js.WithIndexedGetterCallback(func(cbCtx js.CallbackScope[T], key int) (js.Value[T], error) {
			instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			item := instance.Item(key)
			if item == nil {
				return nil, nil
			}
			return codec.EncodeEntity(cbCtx, item)
		}),
		js.WithLengthCallback(func(cbCtx js.CallbackScope[T]) (int, error) {
			instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err != nil {
				return 0, err
			}
			return instance.Length(), nil
		}),
	)
}
