package js

// Value represents a value in JavaScript. Referential equality cannot be used
// to check if to Value instances represent the same value in JavaScript. Use
// StrictEquals to check if two values are equal.
//
// A script engine can use the method Self to return internal values.
//
// The type parameter T is controlled by the actual script engine, permitting
// retriving the internal value used by the engine without type assertions, but
// more importantly, prevents bugs caused by misuse by client code. With the
// type parameter, only values produced by a script engine can be supplied to
// the engine, not any object that conforms to the interface.
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
	Set(name string, val Value[T]) error
}

// Class represents a JavaScript "class" that wraps a Go object.
//
// This package has two separate abstractions for a class serving two different
// roles. This abstractions serves the role of configuring the methods and
// attributes that exists on a class. To create an instance of a class, you use
// the [Constructor]
//
// This is independent of any actual execution context, so values in global
// scope can be declared before creating a JavaScript execution context.
type Class[T any] interface {
	CreatePrototypeMethod(name string, cb FunctionCallback[T])
	CreateIteratorMethod(cb FunctionCallback[T])
	CreatePrototypeAttribute(name string, getter FunctionCallback[T], setter FunctionCallback[T])
	CreateInstanceAttribute(name string, getter FunctionCallback[T], setter FunctionCallback[T])
	CreateIndexedHandler(getter HandlerGetterCallback[T, int])
}

// Constructor represents a JavaScript "class" that wraps a Go object.
//
// This package has two separate abstractions for a class serving two different
// roles. This abstraction is used to create instances of a class in a
// JavaScript execution context.
//
// The class must previously have been configured using the [Class] interface.
type Constructor[T any] interface {
	NewInstance(nativeValue any) (Object[T], error)
}

type Configurator[T any] interface{ Configure(ScriptEngine[T]) }
type ConfigurerFunc[T any] func(ScriptEngine[T])

func (f ConfigurerFunc[T]) Configure(e ScriptEngine[T]) { f(e) }

func AddConfigurator[T any](fact ScriptEngineFactory[T], conf ConfigurerFunc[T]) {
	fact.AddConfigurator(conf)
}
