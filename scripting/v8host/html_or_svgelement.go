package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w htmlOrSVGElementV8Wrapper) focus(cbCtx *v8CallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.focus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err == nil {
		instance.Focus()
	}
	return nil, err
}
