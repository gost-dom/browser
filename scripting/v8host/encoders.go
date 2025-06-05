package v8host

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeUint32(cbCtx jsCallbackContext, val uint32) (jsValue, error) {
	return cbCtx.ValueFactory().NewUint32(val), nil
}

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func encodeEntity[T any](cbCtx js.CallbackScope[T], e entity.ObjectIder) (js.Value[T], error) {
	return codec.EncodeEntity[T](cbCtx, e)
}
