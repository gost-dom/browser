package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w Element[T]) CustomInitializer(jsClass js.Class[T]) {
	jsClass.CreateAttribute("style", w.style, nil)
}

func (w Element[T]) style(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err == nil {
		var ok bool
		if res, ok = entity.Component[js.Value[T]](instance, "style"); !ok {
			res = cbCtx.NewObject()
			entity.SetComponent(instance, "style", res)
		}
	}
	return
}

func (e Element[T]) classList(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()
	tokenList := cbCtx.Constructor("DOMTokenList")
	return tokenList.NewInstance(cl)
}

func encodeNamedNodeMap[T any](
	cbCtx js.CallbackContext[T],
	n dom.NamedNodeMap,
) (js.Value[T], error) {
	return codec.EncodeEntity(cbCtx, n)
}

func decodeElement[T any](s js.Scope[T], v js.Value[T]) (dom.Element, error) {
	return codec.DecodeAs[dom.Element](s, v)
}
