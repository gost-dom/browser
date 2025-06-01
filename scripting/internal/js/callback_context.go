package js

import (
	"errors"
	"iter"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
)

var ErrMissingArgument = errors.New("missing argument")
var ErrNoInternalValue = errors.New("object does not have an internal instance")

// disposable represents a resource that needs cleanup when a context is closed.
// E.g., cgo handles that need to be released.
type Disposable interface{ Dispose() }

type Scope[T any] interface {
	Window() html.Window
	GlobalThis() Object[T]
	Clock() *clock.Clock
	GetValue(entity.ObjectIder) (Value[T], bool)
	SetValue(entity.ObjectIder, Value[T])
	AddDisposable(Disposable)
}

// CallbackContext represents the execution context of a JavaScript function
// or handler callback. For example
//
// - Calling a native function
// - Getting or setting an accessor property backed by a native function
// - Named or indexed handler callbacks / interceptors.
type CallbackContext[T any] interface {
	// ConsumeArg pulls argument from the list of passed arguments. The return
	// value arg will contain the argument. If no argument is passed, or
	// the value is undefined, arg will be nil. Ok indicates whether there were
	// more arguments to consume,
	// the function returns nil. (e.g., if you call the method 3 times, but only
	// two arguments were passed).
	//
	// For most use cases, the client shouldn't care about the ok return value;
	// but treat the value as if undefined was passed. The primary use for ok is
	// when consuming the remaining arguments for a variadic argument list, e.g.
	// [Element.append]
	//
	// [Element.append]: https://developer.mozilla.org/en-US/docs/Web/API/Element/append
	ConsumeArg() (arg Value[T], ok bool)

	// ConsumeRestArgs returns all remaining arguments as a slice of values.
	//
	// This is intended for implementing functions with variadic arguments,
	// e.g., [Element.append]
	//
	// [Element.append]: https://developer.mozilla.org/en-US/docs/Web/API/Element/append
	// ConsumeRestArgs() []Value[T]

	// Instance returns the Go value that is wrapped by "this", with "this"
	// referring the the JavaScript value of "this". If the object does not
	// contain an internal Go value an [ErrNoInternalValue] error is returned.
	Instance() (any, error)

	ReturnWithValue(Value[T]) (Value[T], error)
	ReturnWithError(error) (Value[T], error)
	ReturnWithTypeError(msg string) (Value[T], error)

	// ValueFactory returns a "factory" that can be used to produce JavaScript
	// values.
	ValueFactory() ValueFactory[T]

	Scope() Scope[T]
}

type FunctionCallback[T any] func(CallbackContext[T]) (Value[T], error)

// ValueFactory allows creating JavaScript values from Go values
type ValueFactory[T any] interface {
	Null() Value[T]

	NewString(string) Value[T]
	NewBoolean(bool) Value[T]
	NewUint32(uint32) Value[T]
	NewInt32(int32) Value[T]
	NewInt64(int64) Value[T]

	// NewArray creates a JavaScript array containing the values. If any value
	// is nil, it will become undefined in the resulting array.
	NewArray(...Value[T]) Value[T]
	// NewIterator returns an object implementing the [Iterator protocol]
	//
	// [Iterator protocol]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterator_protocol
	NewIterator(iter.Seq2[Value[T], error]) Value[T]

	NewTypeError(msg string) error

	JSONStringify(val Value[T]) string
	JSONParse(val string) (Value[T], error)
}
