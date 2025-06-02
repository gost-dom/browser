package js

// AssertObject asserts that an argument is of an object type. A TypeError is
// returned if the value is not an object
func AssertObjectArg[T any](cbCtx CallbackContext[T], v Value[T]) (Object[T], error) {
	if obj, ok := v.AsObject(); ok {
		return obj, nil
	}
	return nil, cbCtx.ValueFactory().NewTypeError("Value must be an object")
}

func DecodeString[T any](cbCtx CallbackContext[T], val Value[T]) (string, error) {
	return val.String(), nil
}
