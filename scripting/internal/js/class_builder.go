package js

import "fmt"

// ClassBuilder provides a simplified way of registering classes in global
// scope. Method [CreateGlobals] initializes a script engine with the classes
// registered using [Register]. This handles the correct order of registration,
// creating superclasses before subclasses.
type ClassBuilder[T any] struct {
	e ScriptEngine[T]
}

func NewClassBuilder[T any](e ScriptEngine[T]) ClassBuilder[T] { return ClassBuilder[T]{e} }

type Initializer[T any] interface {
	Constructor(CallbackContext[T]) (Value[T], error)
	Initialize(Class[T])
}

type InitializerFactory[T any, U Initializer[T]] = func(ScriptEngine[T]) U

func RegisterClass[T any, U Initializer[T], V InitializerFactory[T, U]](
	reg ClassBuilder[T],
	className, superClassName string,
	constructorFactory V) {
	var superClassConstructor Class[T]
	if superClassName != "" {
		var ok bool
		if superClassConstructor, ok = reg.e.Class(superClassName); !ok {
			msg := fmt.Sprintf(
				"gost-dom/js: createGlobals: %s: not registered", superClassName,
			)
			panic(msg)
		}
	}
	wrapper := constructorFactory(reg.e)
	class := reg.e.CreateClass(className, superClassConstructor, wrapper.Constructor)
	wrapper.Initialize(class)
}
