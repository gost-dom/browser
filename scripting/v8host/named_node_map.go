package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w namedNodeMapV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
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
			return encodeEntity(cbCtx, item)
		},
	)
}
