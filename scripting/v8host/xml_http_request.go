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
	handleReffedObject[XmlHttpRequest, jsTypeParam]
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
	cbCtx jsCallbackContext,
) (jsValue, error) {
	this := cbCtx.This()
	result := NewXmlHttpRequest(cbCtx.Scope().Window(), cbCtx.Scope().Clock())
	result.SetCatchAllHandler(event.NewEventHandlerFunc(func(event *event.Event) error {
		prop := "on" + event.Type
		handler, err := this.Get(prop)
		if err == nil && handler.IsFunction() {
			ev, err := encodeEntity(cbCtx, event)
			if err == nil {
				f, _ := handler.AsFunction()
				f.Call(this, ev)
			}
		}
		return nil
	}))
	return xhr.store(result, cbCtx)
}

func (xhr xmlHttpRequestV8Wrapper) open(cbCtx jsCallbackContext) (jsValue, error) {
	instance, errInstance := js.As[XmlHttpRequest](cbCtx.Instance())
	method, err0 := consumeArgument(cbCtx, "method", nil, xhr.decodeString)
	url, err1 := consumeArgument(cbCtx, "url", nil, xhr.decodeString)
	if err := errors.Join(err0, err1, errInstance); err != nil {
		return nil, err
	}
	if async, found, err2 := consumeOptionalArg(cbCtx, "async", xhr.decodeBoolean); found {
		if err2 != nil {
			return nil, err2
		}
		instance.Open(method, url, RequestOptionAsync(async))
		return nil, nil
	}
	instance.Open(method, url)
	return nil, nil
}

func (xhr xmlHttpRequestV8Wrapper) upload(cbCtx jsCallbackContext) (jsValue, error) {
	return cbCtx.This(), nil
}
