package js

import (
	"errors"
	"fmt"
)

// AssertObject asserts that an argument is of an object type. A TypeError is
// returned if the value is not an object
func AssertObjectArg[T any](cbCtx CallbackContext[T], v Value[T]) (Object[T], error) {
	if obj, ok := v.AsObject(); ok {
		return obj, nil
	}
	return nil, cbCtx.NewTypeError("Value must be an object")
}

// ParseSetterArg parses a single argument and is intended for attribute
// setters, where exactly one argument must be passed by v8.
func ParseSetterArg[T, U any](
	ctx CallbackContext[T],
	parsers ...func(CallbackContext[T], Value[T]) (U, error),
) (result U, err error) {
	arg, ok := ctx.ConsumeArg()
	if !ok {
		err = errors.New("parseSetterArg: expected one argument. got none")
	}

	errs := make([]error, len(parsers))
	for i, parser := range parsers {
		result, errs[i] = parser(ctx, arg)
		if errs[i] == nil {
			return
		}
	}
	err = fmt.Errorf("parseSetterArg: conversion errors: %w", errors.Join(errs...))
	return
}
