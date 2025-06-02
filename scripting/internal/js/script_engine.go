package js

type ScriptEngine[T any] interface {
	CreateClass(name string, Parent Constructor[T], cb FunctionCallback[T]) Constructor[T]
}

// ScriptEngineFactory constructs ScriptEngine instances. Client code can
// register a Configurator to define values in global JavaScript scope.
type ScriptEngineFactory[T any] interface {
	Register(Configurator[T])
}
