package js

type ScriptEngine[T any] interface {
	CreateClass(name string, Parent Constructor[T], cb FunctionCallback[T]) Constructor[T]
}
