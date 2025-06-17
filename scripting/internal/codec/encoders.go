package codec

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func EncodeEntity[T any](cbCtx js.CallbackScope[T], e entity.ObjectIder) (js.Value[T], error) {
	return EncodeEntityScoped(cbCtx, e)
}

// TODO: Embed scope in CallbackScope, so only one function is necessary
func EncodeEntityScoped[T any](scope js.Scope[T], e entity.ObjectIder) (js.Value[T], error) {
	if e == nil {
		return scope.Null(), nil
	}

	if cached, ok := scope.GetValue(e); ok {
		return cached, nil
	}

	prototypeName := LookupJSPrototype(e)
	prototype := scope.Constructor(prototypeName)
	value, err := prototype.NewInstance(e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err
}

func EncodeBoolean[T any](cbCtx js.CallbackScope[T], b bool) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewBoolean(b), nil
}
func EncodeInt[T any](cbCtx js.CallbackScope[T], i int) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewInt32(int32(i)), nil
}

// TODO: Embed scope in CallbackScope, so only one function is necessary
func EncodeStringScoped[T any](cbCtx js.Scope[T], s string) (js.Value[T], error) {
	return cbCtx.NewString(s), nil
}

func EncodeString[T any](cbCtx js.CallbackScope[T], s string) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewString(s), nil
}

func EncodeNullableString[T any](
	cbCtx js.CallbackScope[T],
	s *string,
) (js.Value[T], error) {
	if s != nil {
		return cbCtx.ValueFactory().NewString(*s), nil
	}
	return EncodeNull(cbCtx)
}
func EncodeNillableString[T any](
	cbCtx js.CallbackScope[T],
	s string,
	hasValue bool,
) (js.Value[T], error) {
	if hasValue {
		return cbCtx.ValueFactory().NewString(s), nil
	}
	return EncodeNull(cbCtx)
}

func EncodeNull[T any](cbCtx js.CallbackScope[T]) (js.Value[T], error) {
	return cbCtx.ValueFactory().Null(), nil
}

// EncodeConstrucedValue is a simple helper for JS constructor callbacks to
// store the constructed Go value in the JavaScript object, and possibly cache
// it with the script context.
func EncodeConstrucedValue[T any](cbCtx js.CallbackScope[T], val any) (js.Value[T], error) {
	// TODO: Figure out if this function should survive
	cbCtx.This().SetNativeValue(val)
	if e, ok := val.(entity.ObjectIder); ok {
		cbCtx.SetValue(e, cbCtx.This())
	}
	return nil, nil
}
