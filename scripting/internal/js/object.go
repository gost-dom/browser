package js

type Object[T any] interface {
	Value[T]
	NativeValue() any
	SetNativeValue(any)
	Keys() ([]string, error)
	Get(name string) (Value[T], error)
	Set(name string, val Value[T]) error
	Iterator() (Value[T], error)
}

type objectGetter[T any] interface {
	objectGet(Object[T]) (Value[T], error)
}

type StringKey[T any] string

func (s StringKey[T]) objectGet(o Object[T]) (Value[T], error) {
	return o.Get(string(s))
}

type IteratorKey[T any] struct{}

func (IteratorKey[T]) objectGet(o Object[T]) (Value[T], error) {
	return o.Iterator()
}

func ObjectGetFunctionx[T any](
	o Object[T],
	getter objectGetter[T],
) (res Function[T], ok bool, err error) {
	val, err := getter.objectGet(o)
	if err == nil {
		res, ok = val.AsFunction()
	}
	return
}
func ObjectGetFunction[T any](o Object[T], name string) (res Function[T], ok bool, err error) {
	val, err := o.Get(name)
	if err == nil {
		res, ok = val.AsFunction()
	}
	return
}

func ObjectGetIterator[T any](o Object[T]) (res Function[T], ok bool, err error) {
	return ObjectGetFunctionx(o, IteratorKey[T]{})
}
