package v8engine

import (
	"runtime/cgo"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/v8go"
)

type v8Object struct {
	v8Value
	Object *v8go.Object
	handle cgo.Handle
}

// newV8Object returns a jsObject wrapping o, a v8go *Object value. The function
// returns nil when o is nil.
func newV8Object(ctx *V8ScriptContext, o *v8go.Object) jsObject {
	if o == nil {
		panic("gost-dom/v8engine: newV8Object called with nil. " + constants.BUG_ISSUE_URL)
	}
	return &v8Object{v8Value{ctx, o.Value}, o, 0}
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
	ext := v8go.NewValueExternalHandle(o.iso(), o.handle)
	defer ext.Release()
	o.Object.SetInternalField(0, ext)
}

func (o *v8Object) Dispose() {
	if o.handle != 0 {
		o.handle.Delete()
		o.handle = 0
	}
}

func (o *v8Object) Iterator() (jsValue, error) {
	res, err := o.Object.GetSymbol(v8go.SymbolIterator(o.iso()))
	if err != nil {
		return nil, err
	}
	return newV8Value(o.ctx, res), nil
}

func (o *v8Object) Get(name string) (jsValue, error) {
	res, err := o.Object.Get(name)
	if err != nil {
		return nil, err
	}
	return newV8Value(o.ctx, res), nil
}

func (o *v8Object) Set(name string, val jsValue) error {
	return o.Object.Set(name, val.Self().v8Value())
}
