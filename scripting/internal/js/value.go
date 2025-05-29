package js

type Value[T any] interface {
	Self() T

	String() string
	Int32() int32
	Uint32() uint32
	Boolean() bool

	IsUndefined() bool
	IsNull() bool
	IsBoolean() bool
	IsString() bool
	IsObject() bool
	IsFunction() bool

	AsFunction() (Function[T], bool)
	AsObject() (Object[T], bool)

	StrictEquals(Value[T]) bool
}

type Function[T any] interface {
	Value[T]

	Call(this Object[T], args ...Value[T]) (Value[T], error)
}

type Object[T any] interface {
	Value[T]
	NativeValue() any
	SetNativeValue(any)
	Get(name string) (Value[T], error)
}

// Constructor represents a JavaScript "class" that wraps a Go object.
//
// While a constructor IS a function in JavaScript, this abstraction has two
// separate represenatation as they have two completely different roles. You
// cannot "call" a constructor, doing so at runtime will result in a TypeError,
// they can only be constructed using the JavaScript new operator. On the Go
// side, a new intance can be created using [Constructor.NewInstance], passing
// the object that should be wrapped.
type Constructor[T any] interface {
	NewInstance(cbCtx CallbackContext[T], nativeValue any) (Object[T], error)
	CreatePrototypeMethod(name string, cb FunctionCallback[T])
}

type ScriptEngineInitializer[T any] = func(ScriptEngine[T])
