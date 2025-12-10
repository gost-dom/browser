package sobekengine

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom"
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

func newScriptError(s scope, err error) js.Error[jsTypeParam] {
	ctx := s.scriptContext
	if errors.Is(err, dom.ErrDom) {
		domException, constructErr := s.Constructor("DOMException").NewInstance(err)
		if constructErr == nil {
			return scriptError{domException}
		} else {
			err = errors.Join(err, fmt.Errorf("constructing DOMException: %w", constructErr))
		}
	}
	return scriptError{newObject(ctx, ctx.vm.NewGoError(err))}
}
