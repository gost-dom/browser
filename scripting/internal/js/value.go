package js

import (
	"errors"
	"fmt"
)

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
	Number() float64
	Int32() int32
	Uint32() uint32
	Boolean() bool

	IsUndefined() bool
	IsNull() bool
	IsSymbol() bool
	IsString() bool
	IsNumber() bool
	IsObject() bool
	IsBoolean() bool
	IsFunction() bool

	AsFunction() (Function[T], bool)
	AsObject() (Object[T], bool)

	StrictEquals(Value[T]) bool
}

type Function[T any] interface {
	Value[T]

	Call(this Object[T], args ...Value[T]) (Value[T], error)
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
	CreateOperation(name string, cb CallbackFunc[T])
	CreateIteratorMethod(cb CallbackFunc[T])
	CreateAttribute(
		name string,
		getter CallbackFunc[T],
		setter CallbackFunc[T],
		opts ...PropertyOption,
	)
	CreateIndexedHandler(getter ...IndexedHandlerOption[T])
	CreateNamedHandler(opts ...NamedHandlerOption[T])
}

// GlobalObject represents an object that will be present in global scope. The
// JavaScript console object is an example of a global object.
type GlobalObject[T any] interface {
	CreateFunction(name string, cb CallbackFunc[T])
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

type Configurer[T any] interface{ Configure(ScriptEngine[T]) }
type ConfigurerFunc[T any] func(ScriptEngine[T])

func (f ConfigurerFunc[T]) Configure(e ScriptEngine[T]) { f(e) }

// Promise represents a JavaScript promise that is controlled from Go-code.
type Promise[T any] interface {
	Value[T]
	Resolve(Value[T])

	// Reject rejects the promise with an Error instance, representing a Go
	// error value. This is based on two assumptions
	//
	//  - You always want to reject with Error values.
	//  - The cause will be an instance of a Go error value.
	//
	// So while you can reject with any value in JavaScript, it is best practice
	// to only use instances of the Error class. This implementation assumes
	// that all API implementations follow that practice.
	Reject(Value[T])
}

// IsNullish returns whether a JavaScript value is null or undefined.
func IsNullish[T any](v Value[T]) bool { return v == nil || v.IsNull() || v.IsUndefined() }

func IsBoolean[T any](v Value[T]) bool { return v != nil && v.IsBoolean() }

func AsFunction[T any](v Value[T]) (Function[T], bool) {
	if IsNullish(v) {
		return nil, false
	}
	return v.AsFunction()
}

func Clone[T any](v Value[T], s Scope[T]) (Value[T], error) {
	var objects [][2]Value[T]
	return clone(v, s, &objects)
}

func clone[T any](v Value[T], s Scope[T], objects *[][2]Value[T]) (Value[T], error) {
	switch {
	case v.IsNull():
		return s.Null(), nil
	case v.IsUndefined():
		return s.Undefined(), nil
	case v.IsString():
		return s.NewString(v.String()), nil
	case v.IsNumber():
		return s.NewNumber(v.Number()), nil
	case v.IsBoolean():
		return s.NewBoolean(v.Boolean()), nil
	case v.IsFunction():
		//TODO: Use correct error
		return nil, errors.New("Serialize function")
	}
	if o, ok := v.AsObject(); ok {
		return cloneObject(o, s, objects)
	}
	return nil, fmt.Errorf("Unable to clone value: %v", v)
}

func cloneObject[T any](o Object[T], s Scope[T], knownObjects *[][2]Value[T]) (Value[T], error) {
	for _, pair := range *knownObjects {
		known := pair[0]
		res := pair[1]
		if o.StrictEquals(known) {
			return res, nil
		}
	}
	res := s.NewObject()
	*knownObjects = append(*knownObjects, [2]Value[T]{o, res})
	keys, err := o.Keys()
	if err != nil {
		return nil, err
	}
	for _, k := range keys {
		oldV, err := o.Get(k)
		if err != nil {
			return nil, err
		}
		newV, err := clone(oldV, s, knownObjects)
		if err != nil {
			return nil, err
		}
		res.Set(k, newV)
	}
	return res, nil
}
