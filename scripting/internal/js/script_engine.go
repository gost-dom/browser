package js

type ScriptEngine[T any] interface {
	CreateClass(name string, Parent Class[T], cb CallbackFunc[T]) Class[T]
	CreateGlobalObject(name string) GlobalObject[T]
	CreateFunction(name string, cb CallbackFunc[T])
	RunScript(script, src string)
}

// ScriptEngineFactory constructs ScriptEngine instances. Client code can
// register a Configurator to define values in global JavaScript scope.
type ScriptEngineFactory[T any] interface {
	AddConfigurator(Configurator[T])
}
