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
	Keys() ([]string, error)
	Get(name string) (Value[T], error)
}

// Constructor represents a JavaScript "class" that wraps a Go object.
//
// While a constructor IS a function in JavaScript, this abstraction has two
// separate represenatation as they have two completely different roles. You
// cannot "call" a constructor, doing so at runtime will result in a TypeError,
// they can only be constructed using the JavaScript new operator.
type Constructor[T any] interface {
	CreatePrototypeMethod(name string, cb FunctionCallback[T])
}

type Constructable[T any] interface {
	NewInstance(cbCtx Scope[T], nativeValue any) (Object[T], error)
}

type Configurator[T any] interface{ Configure(ScriptEngine[T]) }
type ConfigurerFunc[T any] func(ScriptEngine[T])

func (f ConfigurerFunc[T]) Configure(e ScriptEngine[T]) { f(e) }

func Register[T any](fact ScriptEngineFactory[T], conf ConfigurerFunc[T]) {
	fact.Register(conf)
}
