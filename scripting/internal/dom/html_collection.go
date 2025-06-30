package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *HTMLCollection[T]) CustomInitializer(class js.Class[T]) {
	iterator := js.NewIterator(
		func(ctx js.Scope[T], instance dom.Element) (js.Value[T], error) {
			return codec.EncodeEntityScoped(ctx, instance)
		})
	iterator.InstallPrototype(class)

	class.CreateIndexedHandler(
		js.WithIndexedGetterCallback(
			func(info js.CallbackScope[T], index int) (js.Value[T], error) {
				instance, err := js.As[dom.HTMLCollection](info.Instance())
				if err != nil {
					return nil, err
				}
				if item := instance.Item(index); item != nil {
					return codec.EncodeEntity(info, item)
				}
				return nil, nil
			},
		),
		js.WithLengthCallback(
			func(cbCtx js.CallbackScope[T]) (int, error) {
				instance, err := js.As[dom.HTMLCollection](cbCtx.Instance())
				if err != nil {
					return 0, err
				}
				return instance.Length(), nil
			},
		),
	)
}
