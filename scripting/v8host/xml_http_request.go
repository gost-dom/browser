package v8host

import (
	"errors"
	"io"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (xhr xmlHttpRequestV8Wrapper[T]) decodeDocument(
	_ js.CallbackContext[T],
	val js.Value[T],
) (io.Reader, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr xmlHttpRequestV8Wrapper[T]) decodeXMLHttpRequestBodyInit(
	_ js.CallbackContext[T],
	val js.Value[T],
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

func (xhr xmlHttpRequestV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
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
	return storeNewValue(result, cbCtx)
}

func (xhr xmlHttpRequestV8Wrapper[T]) open(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, errInstance := js.As[XmlHttpRequest](cbCtx.Instance())
	method, err0 := js.ConsumeArgument(cbCtx, "method", nil, codec.DecodeString)
	url, err1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if err := errors.Join(err0, err1, errInstance); err != nil {
		return nil, err
	}
	if async, found, err2 := js.ConsumeOptionalArg(cbCtx, "async", codec.DecodeBoolean); found {
		if err2 != nil {
			return nil, err2
		}
		instance.Open(method, url, RequestOptionAsync(async))
		return nil, nil
	}
	instance.Open(method, url)
	return nil, nil
}

func (xhr xmlHttpRequestV8Wrapper[T]) upload(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w xmlHttpRequestV8Wrapper[T]) toAny(
	cbCtx js.CallbackContext[T],
	val string,
) (js.Value[T], error) {
	return codec.EncodeString(cbCtx, val)
}
