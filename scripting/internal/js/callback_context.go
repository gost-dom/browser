package js

import (
	"errors"
	"iter"
	"log/slog"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
)

var ErrMissingArgument = errors.New("missing argument")
var ErrNoInternalValue = errors.New("object does not have an internal instance")

var NotIntercepted = errors.New("Not intercepted")

// disposable represents a resource that needs cleanup when a context is closed.
// E.g., cgo handles that need to be released.
type Disposable interface{ Dispose() }

// Scope provides access to the JavaScript execution context not coupled to a
// specific function callback. CallbackScope is safe use outside the scope of a
// function callback.
//
// One case for storing the callback scope is when JavaScript code passes a
// function as a callback argument, which Go code need to call at a later point
// in time. E.g., when adding event listeners.
type Scope[T any] interface {
	ValueFactory[T]

	Window() html.Window
	GlobalThis() Object[T]
	Clock() *clock.Clock
	GetValue(entity.ObjectIder) (Value[T], bool)
	SetValue(entity.ObjectIder, Value[T])
	Constructor(string) Constructor[T]
}

type ArgumentConsumer[T any] interface {
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
}

type CallbackScope[T any] interface {
	Scope[T]

	// Instance returns the Go value that is wrapped by "this", with "this"
	// referring the the JavaScript value of "this". If the object does not
	// contain an internal Go value an [ErrNoInternalValue] error is returned.
	Instance() (any, error)

	Logger() *slog.Logger
	This() Object[T]
}

// CallbackContext represents the execution context of a JavaScript function
// or handler callback. For example
//
// - Calling a native function or constructor
// - Getting or setting an accessor property backed by a native function
//
// An instance of CallbackContext can only be used inside the actual function
// callback.
type CallbackContext[T any] interface {
	ArgumentConsumer[T]
	CallbackScope[T]

	// ReturnWithValue(Value[T]) (Value[T], error)
	// ReturnWithError(error) (Value[T], error)
	ReturnWithTypeError(msg string) (Value[T], error)
}

type FunctionCallback[T any] func(CallbackContext[T]) (Value[T], error)

type HandlerGetterCallback[T, U any] func(scope CallbackScope[T], key U) (Value[T], error)
type HandlerSetterCallback[T, U any] func(scope CallbackScope[T], key U, value Value[T]) error
type HandlerDeleterCallback[T, U any] func(scope CallbackScope[T], key U) (bool, error)
type HandlerEnumeratorCallback[T, U any] func(CallbackScope[T]) ([]U, error)

type IndexLenCallback[T any] func(CallbackScope[T]) (int, error)

type HandlerCallbacks[Tjs, Tkey any] struct {
	Getter     HandlerGetterCallback[Tjs, Tkey]
	Setter     HandlerSetterCallback[Tjs, Tkey]
	Deleter    HandlerDeleterCallback[Tjs, Tkey]
	Enumerator HandlerEnumeratorCallback[Tjs, Tkey]
}

type NamedHandlerCallbacks[T any] = HandlerCallbacks[T, Value[T]]

type IndexedHandlerCallbacks[T any] struct {
	Getter HandlerGetterCallback[T, int]
	Len    IndexLenCallback[T]
}

type HandlerOption[T, U any] = func(*HandlerCallbacks[T, U])
type IndexedHandlerOption[T any] = func(*IndexedHandlerCallbacks[T])
type NamedHandlerOption[T any] = func(*HandlerCallbacks[T, Value[T]])

func WithLengthCallback[T any](cb IndexLenCallback[T]) IndexedHandlerOption[T] {
	return func(opt *IndexedHandlerCallbacks[T]) { opt.Len = cb }
}

func WithIndexedGetterCallback[T any](cb HandlerGetterCallback[T, int]) IndexedHandlerOption[T] {
	return func(opt *IndexedHandlerCallbacks[T]) { opt.Getter = cb }
}

func WithGetterCallback[T, U any](cb HandlerGetterCallback[T, U]) HandlerOption[T, U] {
	return func(opt *HandlerCallbacks[T, U]) { opt.Getter = cb }
}

func WithSetterCallback[T, U any](cb HandlerSetterCallback[T, U]) HandlerOption[T, U] {
	return func(opt *HandlerCallbacks[T, U]) { opt.Setter = cb }
}

func WithDeleterCallback[T, U any](cb HandlerDeleterCallback[T, U]) HandlerOption[T, U] {
	return func(opt *HandlerCallbacks[T, U]) { opt.Deleter = cb }
}

func WithEnumeratorCallback[T, U any](cb HandlerEnumeratorCallback[T, U]) HandlerOption[T, U] {
	return func(opt *HandlerCallbacks[T, U]) { opt.Enumerator = cb }
}

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
