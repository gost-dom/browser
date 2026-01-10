package js

import "fmt"

type ErrorHandler[T any] interface {
	HandleError(Scope[T], error)
}

type ErrorHandlerFunc[T any] func(Scope[T], error)

func (f ErrorHandlerFunc[T]) HandleError(s Scope[T], err error) {
	f(s, err)
}

type ScriptEngine[T any] interface {
	CreateClass(name string, Parent Class[T], cb CallbackFunc[T]) Class[T]
	Class(name string) (Class[T], bool)
	CreateGlobalObject(name string) GlobalObject[T]
	SetUnhandledPromiseRejectionHandler(ErrorHandler[T])
	ConfigureGlobalScope(name string, Parent Class[T]) Class[T]

	// InstallPolyfill is an escape hatch for when creating the right DOM
	// environment was easier to do in JS - or usable polyfills were found.
	InstallPolyfill(script, src string)
}

// A callback function that returns an "Illegal constructor" TypeError in
// JavaScript. To be used for all classes that cannot be constructed by client
// code directly.
func IllegalConstructor[T any](ctx CallbackContext[T]) (Value[T], error) {
	return nil, ctx.NewTypeError("Illegal constructor")
}

// MustGetClass finds a previously registered class. Returns nil if className is
// empty. Panics if className is non-empty and not previously registered.
func MustGetClass[T any](e ScriptEngine[T], className string) Class[T] {
	if className == "" {
		return nil
	}
	res, ok := e.Class(className)
	if !ok {
		panic(fmt.Sprintf("gost-dom/scripting: %s: class not registered", className))
	}
	return res
}
