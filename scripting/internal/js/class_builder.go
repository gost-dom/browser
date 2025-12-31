package js

import "fmt"

type Initializer[T any] interface {
	Initialize(Class[T])
}

type InitializerFactory[T any, U Initializer[T]] = func(ScriptEngine[T]) U

// CreateClass creates a new JavaScript "class" with a constructor, implementing
// an IDL interface. If inherits is non-empty, the new class will inherit from
// the named class. If constructor is nil, [IllegalConstructor] will be used.
// Panics if a non-empty inherits argument doesn't match a previously created
// class.
//
// Note, web IDL standards use the term "inherits". JavaScript classes use the
// keyword "extends". The web IDL term is used here.
func CreateClass[T any](
	e ScriptEngine[T],
	className, inherits string,
	constructorCallback CallbackFunc[T],
) Class[T] {
	var superClass Class[T]
	if inherits != "" {
		var ok bool
		if superClass, ok = e.Class(inherits); !ok {
			msg := fmt.Sprintf(
				"gost-dom/js: RegisterClass: %s: not registered", inherits,
			)
			panic(msg)
		}
	}
	if constructorCallback == nil {
		constructorCallback = IllegalConstructor
	}
	return e.CreateClass(className, superClass, constructorCallback)
}
