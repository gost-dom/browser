package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

// scriptError embeds a JavaScript value but adds method Error() permitting the
// value to be treated as an error in Go.
type scriptError struct {
	js.Value[jsTypeParam]
	err error
}

func (e scriptError) Unwrap() error { return e.err }

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
