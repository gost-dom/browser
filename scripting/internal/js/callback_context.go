package js

import "errors"

var ErrMissingArgument = errors.New("missing argument")
var ErrNoInternalValue = errors.New("object does not have an internal instance")

// CallbackContext represents the execution context of a JavaScript function
// or handler callback. For example
//
// - Calling a native function
// - Getting or setting an accessor property backed by a native function
// - Named or indexed handler callbacks / interceptors.
type CallbackContext interface {
	// ConsumeRequiredArg pulls argument from the list of required arguments. If
	// there are no more arguments to consume, an error wrapping ErrMissingArgument error will be
	// returned. (e.g., if you call the method 3
	// times, but only two arguments were passed). The actual error will contain
	// the name of the argument in the error message.
	//
	// Client code does not need to deal with the actual error, and should
	// normally just return it as is.
	//  The script engine has the responsibility to translate the error to the
	//  proper JavaScript error. If client code wishes to create a new error
	//  type it must wrap the original error if the default JavaScript Error
	//  should be returned.
	//
	// The function is intended for parsing required arguments in the specs.
	ConsumeRequiredArg(name string) (Value, error)

	// ConsumeRestArgs returns all remaining arguments as a slice of values.
	//
	// This is intended for implementing functions with variadic arguments,
	// e.g., [Element.append]
	//
	// [Element.append]: https://developer.mozilla.org/en-US/docs/Web/API/Element/append
	ConsumeRestArgs() []Value

	// InternalInstance returns the Go value that is wrapped by "this", with
	// "this" referring the the JavaScript value of "this". If the object does
	// not contain an internal Go value an [ErrNoInternalValue] error is
	// returned.
	InternalInstance() (any, error)

	ReturnWithValue(Value) CallbackRVal
	ReturnWithError(error) CallbackRVal

	// ValueFactory returns a "factory" that can be used to produce JavaScript
	// values.
	ValueFactory() ValueFactory
}

type Value interface {
	AsString() string
}

type CallbackRVal any

type FunctionCallback func(CallbackContext) CallbackRVal

// ValueFactory allows creating JavaScript values from Go values
type ValueFactory interface {
	Null() Value
	String(string) Value
}
