package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	"github.com/gost-dom/v8go"
)

func (w htmlOrSVGElementV8Wrapper) focus(cbCtx *argumentHelper) (*v8go.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.focus")
	instance, err := abstraction.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Focus()
	return nil, nil
}
