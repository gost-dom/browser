package js

import "fmt"

type Initializer[T any] interface {
	Constructor(CallbackContext[T]) (Value[T], error)
	Initialize(Class[T])
}

type InitializerFactory[T any, U Initializer[T]] = func(ScriptEngine[T]) U

func RegisterClass[T any, U Initializer[T], V InitializerFactory[T, U]](
	e ScriptEngine[T], className, superClassName string, constructorFactory V,
) {
	var superClass Class[T]
	if superClassName != "" {
		var ok bool
		if superClass, ok = e.Class(superClassName); !ok {
			msg := fmt.Sprintf(
				"gost-dom/js: RegisterClass: %s: not registered", superClassName,
			)
			panic(msg)
		}
	}
	wrapper := constructorFactory(e)
	class := e.CreateClass(className, superClass, wrapper.Constructor)
	wrapper.Initialize(class)
}
