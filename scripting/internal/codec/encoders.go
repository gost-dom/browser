package codec

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func EncodeEntity[T any](cbCtx js.CallbackScope[T], e entity.ObjectIder) (js.Value[T], error) {
	fact := cbCtx.ValueFactory()
	scope := cbCtx.Scope()

	if e == nil {
		return fact.Null(), nil
	}

	if cached, ok := scope.GetValue(e); ok {
		return cached, nil
	}

	prototypeName := LookupJSPrototype(e)
	prototype := cbCtx.Scope().Constructor(prototypeName)
	value, err := prototype.NewInstance(cbCtx.Scope(), e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err
}
