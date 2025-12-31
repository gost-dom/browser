package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func NamedNodeMapCustomInitializer[T any](class js.Class[T]) {
	iterator := js.NewIterator(
		func(s js.Scope[T], instance dom.Attr) (js.Value[T], error) {
			return codec.EncodeEntity(s, instance)
		})
	iterator.InstallPrototype(class)
	class.CreateIndexedHandler(
		js.WithIndexedGetterCallback(func(s js.CallbackScope[T], key int) (js.Value[T], error) {
			instance, err := js.As[dom.NamedNodeMap](s.Instance())
			if err != nil {
				return nil, err
			}
			item := instance.Item(key)
			if item == nil {
				return nil, nil
			}
			return codec.EncodeEntity(s, item)
		}),
		js.WithLengthCallback(func(s js.CallbackScope[T]) (int, error) {
			instance, err := js.As[dom.NamedNodeMap](s.Instance())
			if err != nil {
				return 0, err
			}
			return instance.Length(), nil
		}),
	)
}
