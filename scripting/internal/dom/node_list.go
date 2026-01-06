package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeNode[T any](s js.Scope[T], n dom.Node) (js.Value[T], error) {
	return codec.EncodeEntity(s, n)
}

func NodeListCustomInitializer[T any](class js.Class[T]) {
	class.CreateIndexedHandler(
		js.WithIndexedGetterCallback(
			func(info js.CallbackScope[T], index int) (js.Value[T], error) {
				instance, err := js.As[dom.NodeList](info.Instance())
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
			func(s js.CallbackScope[T]) (int, error) {
				instance, err := js.As[dom.NodeList](s.Instance())
				if err != nil {
					return 0, err
				}
				return instance.Length(), nil
			},
		),
	)
}
