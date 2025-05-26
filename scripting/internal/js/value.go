package js

type Value interface {
	String() string
	Int32() int32

	AsFunction() (Function, bool)
	AsObject() (Object, bool)
}

type Function interface {
	Value

	Call(this Object, args ...Value) (Value, error)
}

type Object interface {
	Value
	NativeValue() any
	Get(name string) (Value, error)
}
