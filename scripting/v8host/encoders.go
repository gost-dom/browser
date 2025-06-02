package v8host

import (
	"github.com/gost-dom/browser/internal/entity"
)

func encodeUint32(cbCtx jsCallbackContext, val uint32) (jsValue, error) {
	return cbCtx.ValueFactory().NewUint32(val), nil
}

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func encodeEntity(cbCtx jsCallbackContext, e entity.ObjectIder) (jsValue, error) {
	fact := cbCtx.ValueFactory()
	scope := cbCtx.Scope()

	if e == nil {
		return fact.Null(), nil
	}

	if cached, ok := scope.GetValue(e); ok {
		return cached, nil
	}

	prototypeName := lookupJSPrototype(e)
	prototype := cbCtx.Scope().Constructor(prototypeName)
	value, err := prototype.NewInstance(cbCtx.Scope(), e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err
}
