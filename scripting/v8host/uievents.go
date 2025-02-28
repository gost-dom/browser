package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/v8go"
)

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
	e := uievents.NewUIEvent(type_)
	return w.store(e, ctx, this)
}

func (w uIEventV8Wrapper) decodeUIEventInit(
	ctx *V8ScriptContext,
	v *v8go.Value,
) (dom.EventOption, error) {
	return w.decodeEventInit(ctx, v)
}
