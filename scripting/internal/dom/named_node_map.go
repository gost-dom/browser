package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w NamedNodeMap[T]) CustomInitializer(class js.Class[T]) {
	class.CreateIndexedHandler(
		js.WithGetterCallback(func(cbCtx js.CallbackScope[T], key int) (js.Value[T], error) {
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
	)
}
