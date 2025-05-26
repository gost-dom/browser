package v8host

import (
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

func (w uIEventV8Wrapper) decodeMouseEventInit(
	cbCtx jsCallbackContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeUIEventInit(cbCtx, v)
}

func (w uIEventV8Wrapper) decodePointerEventInit(
	cbCtx jsCallbackContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeMouseEventInit(cbCtx, v)
}

func (w uIEventV8Wrapper) CreateInstance(
	cbCtx *argumentHelper,
	type_ string,
) js.CallbackRVal {
	return w.CreateInstanceEventInitDict(cbCtx, type_)
}

func (w uIEventV8Wrapper) CreateInstanceEventInitDict(
	cbCtx *argumentHelper,
	type_ string,
	options ...interface{}) js.CallbackRVal {
	e := uievents.NewUIEvent(type_)
	return cbCtx.ReturnWithValueErr(w.store(e, cbCtx.ScriptCtx(), cbCtx.This()))
}

func (w uIEventV8Wrapper) decodeUIEventInit(
	cbCtx jsCallbackContext,
	v *v8go.Value,
) (eventInitWrapper, error) {
	return w.decodeEventInit(cbCtx, v)
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
