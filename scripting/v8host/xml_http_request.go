package v8host

import (
	"errors"
	"io"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type xmlHttpRequestV8Wrapper struct {
	handleReffedObject[XmlHttpRequest]
}

func (xhr xmlHttpRequestV8Wrapper) decodeDocument(
	cbCtx jsCallbackContext,
	val jsValue,
) (io.Reader, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr xmlHttpRequestV8Wrapper) decodeXMLHttpRequestBodyInit(
	cbCtx jsCallbackContext,
	val jsValue,
) (io.Reader, error) {
	if val == nil {
		return nil, nil
	}
	if val.IsUndefined() || val.IsNull() {
		return nil, nil
	}
	if val.IsString() {
		return strings.NewReader(val.String()), nil
	}
	if obj, ok := val.AsObject(); ok {
		if res, ok := obj.NativeValue().(*html.FormData); ok {
			return res.GetReader(), nil
		}
	}
	return nil, errors.New("XMLHTTPRequest only accepts FormData body yet")
}

func newXMLHttpRequestV8Wrapper(host *V8ScriptHost) xmlHttpRequestV8Wrapper {
	return xmlHttpRequestV8Wrapper{newHandleReffedObject[XmlHttpRequest](host)}
}

func (xhr xmlHttpRequestV8Wrapper) CreateInstance(
	cbCtx *v8CallbackContext,
) (jsValue, error) {
	ctx := cbCtx.ScriptCtx()
	this := cbCtx.This()
	result := NewXmlHttpRequest(ctx.window, ctx.clock)
	result.SetCatchAllHandler(event.NewEventHandlerFunc(func(event *event.Event) error {
		prop := "on" + event.Type
		handler, err := this.Get(prop)
		if err == nil && handler.IsFunction() {
			v8Event, err := ctx.getJSInstance(event)
			if err == nil {
				f, _ := handler.AsFunction()
				f.Call(this, v8Event)
			}
		}
		return nil
	}))
	xhr.store(result, ctx, this)
	return cbCtx.ReturnWithValue(nil)
}

func (xhr xmlHttpRequestV8Wrapper) open(cbCtx *v8CallbackContext) (jsValue, error) {
	method, err0 := consumeArgument(cbCtx, "method", nil, xhr.decodeString)
	url, err1 := consumeArgument(cbCtx, "url", nil, xhr.decodeString)
	async, err2 := consumeArgument(cbCtx, "async", nil, xhr.decodeBoolean)
	instance, errInstance := js.As[XmlHttpRequest](cbCtx.Instance())
	if cbCtx.noOfReadArguments > 2 {
		if err := errors.Join(err0, err1, err2, errInstance); err != nil {
			return cbCtx.ReturnWithError(err)
		}
		instance.Open(method, url, RequestOptionAsync(async))
		return cbCtx.ReturnWithValue(nil)
	}
	if cbCtx.noOfReadArguments < 2 {
		cbCtx.ReturnWithError(errors.New("Not enough arguments"))
	}
	if err := errors.Join(err0, err1, errInstance); err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Open(method, url)
	return cbCtx.ReturnWithValue(nil)
}

func (xhr xmlHttpRequestV8Wrapper) upload(cbCtx *v8CallbackContext) (jsValue, error) {
	return cbCtx.ReturnWithJSValue(cbCtx.This())
}
