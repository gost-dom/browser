package v8engine

import (
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

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
	opts ...js.PropertyOption,
) {
	o := js.InitOpts(opts...)
	v8Getter := wrapV8Callback(c.host, getter.WithLog(c.name, fmt.Sprintf("%s get", name)))
	v8Setter := wrapV8Callback(c.host, setter.WithLog(c.name, fmt.Sprintf("%s set", name)))
	if o.InstanceMember {
		c.inst.SetAccessorProperty(name, v8Getter, v8Setter, v8go.None)
	} else {
		c.proto.SetAccessorProperty(name, v8Getter, v8Setter, v8go.None)
	}
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
