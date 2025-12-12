package html

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w HTMLElement[T]) CustomInitializer(jsClass js.Class[T]) {
	jsClass.CreateInstanceAttribute("style", w.style, nil)
}

func (w HTMLElement[T]) style(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err == nil {
		var ok bool
		if res, ok = entity.Component[js.Value[T]](instance, "style"); !ok {
			res = cbCtx.NewObject()
			entity.SetComponent(instance, "style", res)
		}
	}
	return
}
