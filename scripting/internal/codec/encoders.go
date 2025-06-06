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
