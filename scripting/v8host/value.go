package v8host

import (
	"runtime/cgo"

	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type jsTypeParam = *v8Value
type jsValue = js.Value[*v8Value]
type jsFunction = js.Function[*v8Value]
type jsObject = js.Object[*v8Value]
type jsConstructor = v8Constructor

func toV8Value(v jsValue) *v8go.Value {
	if v == nil {
		return nil
	}
	return v.Self().v8Value()
}

func assertV8Object(v jsObject) *v8Object {
	if r, ok := v.(*v8Object); ok {
		return r
	}
	panic("Expected a V8 Object")
}

//*/

type v8Value struct {
	iso   *v8go.Isolate
	Value *v8go.Value
}

func (v *v8Value) Self() *v8Value { return v }

// newV8Value creates a v8Value wrapping a v8go value. This is safe to use for
// for mapping values that can be nil. If the v8go value is nil, this will
// return nil.
func newV8Value(iso *v8go.Isolate, v *v8go.Value) jsValue {
	if v == nil {
		return nil
	}
	return &v8Value{iso, v}
}

func (v *v8Value) v8Value() *v8go.Value {
	if v == nil {
		return nil
	}
	return v.Value
}

func (v v8Value) String() string { return v.Value.String() }
func (v v8Value) Int32() int32   { return v.Value.Int32() }
func (v v8Value) Uint32() uint32 { return v.Value.Uint32() }
func (v v8Value) Boolean() bool  { return v.Value.Boolean() }

func (v v8Value) IsUndefined() bool { return v.Value.IsUndefined() }
func (v v8Value) IsNull() bool      { return v.Value.IsNull() }
func (v v8Value) IsBoolean() bool   { return v.Value.IsBoolean() }
func (v v8Value) IsString() bool    { return v.Value.IsString() }
func (v v8Value) IsObject() bool    { return v.Value.IsObject() }
func (v v8Value) IsFunction() bool  { return v.Value.IsFunction() }

func (v v8Value) StrictEquals(
	other jsValue,
) bool {
	return v.Value.StrictEquals(toV8Value(other))
}

func (v v8Value) AsFunction() (jsFunction, bool) {
	f, err := v.Value.AsFunction()
	if err == nil {
		return &v8Function{v, f}, true
	}
	return nil, false
}

func (v v8Value) AsObject() (jsObject, bool) {
	o, err := v.Value.AsObject()
	if err == nil {
		return newV8Object(v.iso, o), true
	}
	return nil, false
}

type v8Function struct {
	v8Value
	v8fn *v8go.Function
}

func (f v8Function) Call(this jsObject, args ...jsValue) (jsValue, error) {
	v8Args := make([]v8go.Valuer, len(args))
	for i, a := range args {
		v8Args[i] = toV8Value(a)
	}
	var res jsValue
	v, err := f.v8fn.Call(assertV8Object(this).Object, v8Args...)
	if err == nil {
		res = &v8Value{f.iso, v}
	}
	return res, err
}

/* -------- v8Object -------- */

type v8Object struct {
	v8Value
	Object *v8go.Object
	handle cgo.Handle
}

// newV8Object returns a jsObject wrapping o, a v8go *Object value. The function
// returns nil when o is nil.
func newV8Object(iso *v8go.Isolate, o *v8go.Object) jsObject {
	if o == nil {
		return nil
	}
	return &v8Object{v8Value{iso, o.Value}, o, 0}
}

// NativeValue returns the native Go value if any that this JS object is
// wrapping. I.e., for a JS HTMLFormElement, this will return the Go
// HTMLFormElement implementation. Returns nil when no native value is being
// wrapped by this object.
func (o *v8Object) NativeValue() any {
	if o.Object.InternalFieldCount() == 0 {
		return nil
	}
	internal := o.Object.GetInternalField(0)
	defer internal.Release()

	if !internal.IsExternal() {
		return nil
	}

	return internal.ExternalHandle().Value()
}

func (o *v8Object) SetNativeValue(v any) {
	if o.handle != 0 {
		o.handle.Delete()
	}
	o.handle = cgo.NewHandle(v)
	ext := v8go.NewValueExternalHandle(o.iso, o.handle)
	defer ext.Release()
	o.Object.SetInternalField(0, ext)
}

func (o *v8Object) Dispose() {
	if o.handle != 0 {
		o.handle.Delete()
		o.handle = 0
	}
}

func (o *v8Object) Get(name string) (jsValue, error) {
	res, err := o.Object.Get(name)
	if err != nil {
		return nil, err
	}
	return &v8Value{o.iso, res}, nil
}

type v8Constructor struct {
	host *V8ScriptHost
	ft   *v8go.FunctionTemplate
}

func newV8Constructor(host *V8ScriptHost, ft *v8go.FunctionTemplate) jsConstructor {
	return v8Constructor{host, ft}
}

func (c v8Constructor) NewInstance(
	ctx /* TODO: jsCallbackContext */ *V8ScriptContext,
	nativeValue any,
) (jsObject, error) {
	val, err := c.ft.InstanceTemplate().NewInstance(ctx.v8ctx)
	obj := newV8Object(c.host.iso, val).(*v8Object)
	if err == nil {
		obj.SetNativeValue(nativeValue)
		ctx.addDisposer(obj)
	}
	return obj, err
}

func (c v8Constructor) CreatePrototypeMethod(name string, cb internalCallback) {
	v8cb := wrapV8Callback(c.host, cb)
	c.ft.PrototypeTemplate().Set(name, v8cb, v8go.ReadOnly)
}
