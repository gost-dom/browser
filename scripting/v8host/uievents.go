package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/v8go"
)

func (w uIEventV8Wrapper) decodeMouseEventInit(
	ctx *V8ScriptContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeUIEventInit(ctx, v)
}

func (w uIEventV8Wrapper) decodePointerEventInit(
	ctx *V8ScriptContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeMouseEventInit(ctx, v)
}

func (w uIEventV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8go.Object,
	type_ string,
) (*v8go.Value, error) {
	return w.CreateInstanceEventInitDict(ctx, this, type_)
}

func (w uIEventV8Wrapper) CreateInstanceEventInitDict(
	ctx *V8ScriptContext,
	this *v8go.Object,
	type_ string,
	options ...interface{}) (*v8go.Value, error) {
	e := dom.NewUIEvent(type_)
	return w.store(e, ctx, this)
}

func (w uIEventV8Wrapper) decodeUIEventInit(
	ctx *V8ScriptContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeEventInit(ctx, v)
}

type mouseEventV8Wrapper struct {
	uIEventV8Wrapper
}

type pointerEventV8Wrapper struct {
	mouseEventV8Wrapper
}

func newMouseEventV8Wrapper(host *V8ScriptHost) mouseEventV8Wrapper {
	return mouseEventV8Wrapper{*newUIEventV8Wrapper(host)}
}

func newPointerEventV8Wrapper(host *V8ScriptHost) pointerEventV8Wrapper {
	return pointerEventV8Wrapper{newMouseEventV8Wrapper(host)}
}
