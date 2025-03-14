package v8host

import (
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/v8go"
)

func (w htmlOrSVGElementV8Wrapper) focus(info *v8go.FunctionCallbackInfo) (*v8go.Value, error) {
	log.Debug("V8 Function call: HTMLOrSVGElement.focus")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.Focus()
	return nil, nil
}
