package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w NamedNodeMap[T]) CustomInitializer(class js.Class[T]) {
	class.CreateIndexedHandler(
		func(cbCtx js.GetterCallbackContext[T, int]) (js.Value[T], error) {
			instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
			if err != nil {
				return nil, err
			}
			index := int(cbCtx.Key())
			item := instance.Item(index)
			if item == nil {
				return nil, nil
			}
			return codec.EncodeEntity(cbCtx, item)
		},
	)
}
