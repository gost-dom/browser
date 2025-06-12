package html

import (
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
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

func (w HTMLOrSVGElement[T]) dataset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.focus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.Dataset())
}
