package js

type Array[T any] interface {
	Value[T]
	Push(Value[T]) error
}
