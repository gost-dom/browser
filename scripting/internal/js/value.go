package js

type Value interface {
	String() string
	Int32() int32
	Uint32() uint32
	Boolean() bool

	IsUndefined() bool
	IsNull() bool
	IsBoolean() bool
	IsString() bool
	IsObject() bool
	IsFunction() bool

	AsFunction() (Function, bool)
	AsObject() (Object, bool)

	StrictEquals(Value) bool
}

type Function interface {
	Value

	Call(this Object, args ...Value) (Value, error)
}

type Object interface {
	Value
	NativeValue() any
	SetNativeValue(any)
	Get(name string) (Value, error)
}
