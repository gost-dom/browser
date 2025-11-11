package html

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w HTMLOrSVGElement[T]) focus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.focus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err == nil {
		instance.Focus()
	}
	return nil, err
}

func (w HTMLOrSVGElement[T]) InstallPrototype(jsClass js.Class[T]) {
	w.installPrototype(jsClass)
}
