package v8engine

import (
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

var _ jsClass = &v8Class{}
var _ jsClass = &v8GlobalClass{}

type v8Class struct {
	host           *V8ScriptHost
	ft             *v8go.FunctionTemplate
	proto          *v8go.ObjectTemplate
	inst           *v8go.ObjectTemplate
	parent         *v8Class
	instanceAtttrs []attribute

	name string
}

func newV8Class(
	host *V8ScriptHost,
	name string,
	cb js.CallbackFunc[jsTypeParam],
	parent *v8Class,
) *v8Class {
	ft := wrapV8Callback(host, cb.WithLog(name, "Constructor"))
	if parent != nil {
		ft.Inherit(parent.ft)
	}
	result := v8Class{
		host:   host,
		ft:     ft,
		proto:  ft.PrototypeTemplate(),
		inst:   ft.InstanceTemplate(),
		name:   name,
		parent: parent,
	}
	result.inst.SetInternalFieldCount(1)
	for parent != nil {
		// V8 docs says that inherited classes _does_ get the instance
		// attributes of parent classes. But ...
		for _, attr := range parent.instanceAtttrs {
			v8Getter := wrapV8Callback(host, attr.getter.WithLog(name, fmt.Sprintf("%s get", name)))
			v8Setter := wrapV8Callback(host, attr.setter.WithLog(name, fmt.Sprintf("%s set", name)))
			result.inst.SetAccessorProperty(attr.name, v8Getter, v8Setter, v8go.None)
		}
		parent = parent.parent
	}
	return &result
}

type attribute struct {
	name   string
	getter js.CallbackFunc[jsTypeParam]
	setter js.CallbackFunc[jsTypeParam]
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

func (c *v8Class) CreateAttribute(
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
		c.instanceAtttrs = append(c.instanceAtttrs, attribute{
			name:   name,
			getter: getter,
			setter: setter,
		})
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

/* -------- v8GlobalClass -------- */

type v8GlobalClass struct {
	v8Class *v8Class
}

func newV8GlobalClass(
	host *V8ScriptHost,
	name string,
	cb js.CallbackFunc[jsTypeParam],
	parent *v8Class,
) *v8GlobalClass {
	return &v8GlobalClass{newV8Class(host, name, cb, parent)}
}

func (c v8GlobalClass) CreateOperation(name string, cb js.CallbackFunc[jsTypeParam]) {
	v8cb := wrapV8Callback(c.v8Class.host, cb.WithLog(c.v8Class.name, name))
	c.v8Class.host.windowTemplate.Set(name, v8cb, v8go.ReadOnly)
}

func (c v8GlobalClass) CreateAttribute(
	name string,
	getter js.CallbackFunc[jsTypeParam],
	setter js.CallbackFunc[jsTypeParam],
	opts ...js.PropertyOption,
) {
	host := c.v8Class.host
	className := c.v8Class.name
	v8Getter := wrapV8Callback(host, getter.WithLog(className, fmt.Sprintf("%s get", name)))
	v8Setter := wrapV8Callback(host, setter.WithLog(className, fmt.Sprintf("%s set", name)))
	host.windowTemplate.SetAccessorProperty(name, v8Getter, v8Setter, v8go.None)
}

func (c v8GlobalClass) CreateIndexedHandler(opts ...js.IndexedHandlerOption[jsTypeParam]) {
	c.v8Class.CreateIndexedHandler(opts...)
}

func (c v8GlobalClass) CreateNamedHandler(opts ...js.NamedHandlerOption[jsTypeParam]) {
	c.v8Class.CreateNamedHandler(opts...)
}

func (c v8GlobalClass) CreateIteratorMethod(cb js.CallbackFunc[jsTypeParam]) {
	c.v8Class.CreateIteratorMethod(cb)
}
