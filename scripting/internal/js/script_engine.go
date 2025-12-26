package js

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
	CreateFunction(name string, cb CallbackFunc[T])
	SetUnhandledPromiseRejectionHandler(ErrorHandler[T])
	ConfigureGlobalScope(name string, Parent Class[T]) Class[T]

	// InstallPolyfill is an escape hatch for when creating the right DOM
	// environment was easier to do in JS - or usable polyfills were found.
	InstallPolyfill(script, src string)
}

// ScriptEngineFactory constructs ScriptEngine instances. Client code can
// register a Configurator to define values in global JavaScript scope.
type ScriptEngineFactory[T any] interface {
	AddConfigurator(Configurer[T])
}

func IllegalConstructor[T any](ctx CallbackContext[T]) (Value[T], error) {
	return nil, ctx.NewTypeError("Illegal Constructor")
}
