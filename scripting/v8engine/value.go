package v8engine

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/internal/monads/result"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type jsTypeParam = *v8Value
type jsValue = js.Value[*v8Value]
type jsFunction = js.Function[*v8Value]
type jsObject = js.Object[*v8Value]
type jsError = js.Error[*v8Value]

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

type v8Value struct {
	ctx   *V8ScriptContext
	Value *v8go.Value
}

func (v *v8Value) iso() *v8go.Isolate   { return v.ctx.iso() }
func (v *v8Value) global() *v8go.Object { return v.ctx.v8ctx.Global() }

func (v *v8Value) Self() *v8Value { return v }

// newV8Value creates a v8Value wrapping a v8go value. This is safe to use for
// for mapping values that can be nil. A nil input value will be treated as
// undefined.
func newV8Value(ctx *V8ScriptContext, v *v8go.Value) *v8Value {
	return &v8Value{ctx, v}
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

func (v v8Value) IsUndefined() bool { return v.Value == nil || v.Value.IsUndefined() }
func (v v8Value) IsNull() bool      { return v.Value.IsNull() }
func (v v8Value) IsBoolean() bool   { return v.Value.IsBoolean() }
func (v v8Value) IsString() bool    { return v.Value.IsString() }
func (v v8Value) IsSymbol() bool    { return v.Value.IsSymbol() }
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
		return newV8Object(v.ctx, o), true
	}
	return nil, false
}

func callV8Function(f *v8go.Function, arg0 *v8go.Value, arg *v8go.Value) (*v8go.Value, error) {
	return f.Call(arg0, arg)
}

func asSlice(v any) ([]any, error) {
	if res, ok := v.([]any); ok {
		return res, nil
	}
	return nil, errors.New("value not a go slice")
}

func mapSlice[T, U any](s []T, fn func(T) (U, error)) ([]U, error) {
	var err error
	r := make([]U, len(s))
	for i, e := range s {
		if r[i], err = fn(e); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func asString(v any) (string, error) {
	if s, ok := v.(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("not a string: %v", v)
}

func (o *v8Object) Keys() ([]string, error) {
	// v8go doesn't support retrieving all keys for an object, so this
	// evaluates "Object.keys" to get the JavaScript function that retrieves the
	// necessary data, and then convert the value into a slice of strings.
	global := o.global().Value
	objectKeys := result.Bind(
		result.New(o.ctx.runScript("Object.keys")),
		(*v8go.Value).AsFunction,
	)
	keysV8Value := result.Bind2(objectKeys, callV8Function, global, o.Value)
	keysAsAny := result.Bind(keysV8Value, v8ValueToGoValue)
	keysAsSlice := result.Bind(keysAsAny, asSlice)
	return result.Bind1(keysAsSlice, mapSlice, asString).Unwrap()
}

type v8GlobalObject struct {
	host *V8ScriptHost
	tmpl *v8go.ObjectTemplate
}

func newV8GlobalObject(host *V8ScriptHost, tmpl *v8go.ObjectTemplate) v8GlobalObject {
	return v8GlobalObject{host, tmpl}
}

func (o v8GlobalObject) CreateFunction(name string, cb js.CallbackFunc[jsTypeParam]) {
	o.tmpl.Set(name, wrapV8Callback(o.host, cb.WithLog("", name)))
}

type v8Class struct {
	host  *V8ScriptHost
	ft    *v8go.FunctionTemplate
	proto *v8go.ObjectTemplate
	inst  *v8go.ObjectTemplate

	name string
}

func newV8Class(host *V8ScriptHost, name string, ft *v8go.FunctionTemplate) v8Class {
	return v8Class{host, ft, ft.PrototypeTemplate(), ft.InstanceTemplate(), name}
}

func (c v8Class) CreateIteratorMethod(cb js.CallbackFunc[jsTypeParam]) {
	v8cb := wrapV8Callback(c.host, cb.WithLog(c.name, "Symbol.iterator"))
	it := v8go.SymbolIterator(c.host.iso)
	c.proto.SetSymbol(it, v8cb, v8go.ReadOnly)
}
func (c v8Class) CreateOperation(name string, cb js.CallbackFunc[jsTypeParam]) {
	v8cb := wrapV8Callback(c.host, cb.WithLog(c.name, name))
	c.proto.Set(name, v8cb, v8go.ReadOnly)
}

func (c v8Class) CreateAttribute(
	name string,
	getter js.CallbackFunc[jsTypeParam],
	setter js.CallbackFunc[jsTypeParam],
) {
	v8Getter := wrapV8Callback(c.host, getter.WithLog(c.name, fmt.Sprintf("%s get", name)))
	v8Setter := wrapV8Callback(c.host, setter.WithLog(c.name, fmt.Sprintf("%s set", name)))
	c.proto.SetAccessorProperty(name, v8Getter, v8Setter, v8go.None)
}

func (c v8Class) CreateInstanceAttribute(
	name string,
	getter js.CallbackFunc[jsTypeParam],
	setter js.CallbackFunc[jsTypeParam],
) {
	v8Getter := wrapV8Callback(c.host, getter.WithLog(c.name, fmt.Sprintf("%s get", name)))
	v8Setter := wrapV8Callback(c.host, setter.WithLog(c.name, fmt.Sprintf("%s set", name)))
	c.inst.SetAccessorProperty(name, v8Getter, v8Setter, v8go.None)
}

func (c v8Class) CreateIndexedHandler(opts ...js.IndexedHandlerOption[jsTypeParam]) {
	var oo js.IndexedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&oo)
	}
	c.inst.SetIndexedHandler(func(info *v8go.FunctionCallbackInfo) (*v8go.Value, error) {
		res, err := oo.Getter(newV8CallbackScope(c.host, info), int(info.Index()))
		return toV8Value(res), err
	})
}

func (c v8Class) CreateNamedHandler(opts ...js.NamedHandlerOption[jsTypeParam]) {
	var oo js.NamedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&oo)
	}
	c.inst.SetNamedHandler(v8HandlerWrapper{c.host, oo})
}

type v8HandlerWrapper struct {
	host      *V8ScriptHost
	callbacks js.NamedHandlerCallbacks[jsTypeParam]
}

func (w v8HandlerWrapper) NamedPropertyGet(
	property *v8go.Value,
	info v8go.PropertyCallbackInfo,
) (*v8go.Value, error) {
	w.host.Logger().Debug("NamedPropertyGet", "property", property)
	if w.callbacks.Getter == nil {
		return nil, v8go.NotIntercepted
	}
	ctx := w.host.mustGetContext(info.Context())
	result, err := w.callbacks.Getter(newV8CallbackScope(w.host, info), newV8Value(ctx, property))
	if err == nil && result != nil {
		return result.Self().v8Value(), nil
	}
	return nil, w.convertErr(err)
}

func (w v8HandlerWrapper) NamedPropertySet(
	property *v8go.Value,
	value *v8go.Value,
	info v8go.PropertyCallbackInfo,
) error {
	w.host.Logger().Debug("NamedPropertySet", "property", property, "value", value)
	if w.callbacks.Setter == nil {
		return v8go.NotIntercepted
	}
	ctx := w.host.mustGetContext(info.Context())
	err := w.callbacks.Setter(newV8CallbackScope(w.host, info),
		newV8Value(ctx, property),
		newV8Value(ctx, value),
	)
	return w.convertErr(err)
}

func (w v8HandlerWrapper) NamedPropertyDelete(
	property *v8go.Value,
	info v8go.PropertyCallbackInfo,
) (success bool, err error) {
	w.host.Logger().Debug("NamedPropertyDelete", "property", property)
	if w.callbacks.Deleter == nil {
		return false, v8go.NotIntercepted
	}
	ctx := w.host.mustGetContext(info.Context())
	success, err = w.callbacks.Deleter(newV8CallbackScope(w.host, info), newV8Value(ctx, property))
	return success, w.convertErr(err)
}

func (w v8HandlerWrapper) NamedPropertyEnumerator(
	info v8go.PropertyCallbackInfo,
) (names []*v8go.Value, err error) {
	w.host.Logger().Debug("NamedPropertyEnumerator")
	if w.callbacks.Enumerator == nil {
		return nil, v8go.NotIntercepted
	}
	scope := newV8CallbackScope(w.host, info)
	result, err := w.callbacks.Enumerator(scope)
	if err == nil {
		res := make([]*v8go.Value, len(result))
		for i, r := range result {
			res[i] = r.Self().v8Value()
		}
		return res, nil
	}
	return nil, w.convertErr(err)
}

func (w v8HandlerWrapper) convertErr(err error) error {
	if err == js.NotIntercepted {
		return v8go.NotIntercepted
	}
	return err
}
