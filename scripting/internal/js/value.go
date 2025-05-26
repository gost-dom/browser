package js

type Value interface {
	String() string
	Int32() int32

	AsFunction() Function
}

type Function interface {
	Value

	Call(this Value, args ...Value) (Value, error)
}
