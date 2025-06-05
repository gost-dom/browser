package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w namedNodeMapV8Wrapper) CustomInitializer(class jsClass) {
	class.CreateIndexedHandler(
		func(cbCtx js.GetterCallbackContext[jsTypeParam, int]) (jsValue, error) {
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
