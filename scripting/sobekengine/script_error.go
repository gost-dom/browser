package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

// scriptError embeds a JavaScript value but adds method Error() permitting the
// value to be treated as an error in Go.
type scriptError struct {
	js.Value[jsTypeParam]
}

func (e scriptError) Error() string {
	if obj, ok := e.AsObject(); ok {
		s, err := obj.Get("message")
		if err != nil {
			return err.Error()
		}
		if s.IsUndefined() {
			return obj.String()
		} else {
			return s.String()
		}
	}
	return "undefined"
}

func newScriptError(ctx *scriptContext, err error) js.Error[jsTypeParam] {
	return scriptError{newObject(ctx, ctx.vm.NewGoError(err))}
}
