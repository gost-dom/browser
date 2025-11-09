package xhr

import (
	"errors"
	"io"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	inthtml "github.com/gost-dom/browser/internal/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (xhr XMLHttpRequest[T]) decodeDocument(
	_ js.Scope[T],
	val js.Value[T],
) (io.Reader, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr XMLHttpRequest[T]) decodeXMLHttpRequestBodyInit(
	_ js.Scope[T],
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

func (xhr XMLHttpRequest[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	this := cbCtx.This()
	result := inthtml.NewXmlHttpRequest(cbCtx.Window(), cbCtx.Clock())
	result.SetCatchAllHandler(event.NewEventHandlerFunc(func(event *event.Event) error {
		prop := "on" + event.Type
		handler, err := this.Get(prop)
		if err == nil && handler.IsFunction() {
			ev, err := codec.EncodeEntity(cbCtx, event)
			if err == nil {
				f, _ := handler.AsFunction()
				f.Call(this, ev)
			}
		}
		return nil
	}))
	return codec.EncodeConstrucedValue(cbCtx, result)
}

func (xhr XMLHttpRequest[T]) open(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, errInstance := js.As[inthtml.XmlHttpRequest](cbCtx.Instance())
	method, err0 := js.ConsumeArgument(cbCtx, "method", nil, codec.DecodeString)
	url, err1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if err := errors.Join(err0, err1, errInstance); err != nil {
		return nil, err
	}
	if async, found, err2 := js.ConsumeOptionalArg(cbCtx, "async", codec.DecodeBoolean); found {
		if err2 != nil {
			return nil, err2
		}
		instance.Open(method, url, inthtml.RequestOptionAsync(async))
		return nil, nil
	}
	instance.Open(method, url)
	return nil, nil
}

func (xhr XMLHttpRequest[T]) upload(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w XMLHttpRequest[T]) toAny(s js.Scope[T], val string) (js.Value[T], error) {
	return codec.EncodeString(s, val)
}

func (xhr XMLHttpRequest[T]) decodeXMLHttpRequestResponseType(
	s js.Scope[T], val js.Value[T],
) (string, error) {
	return codec.DecodeString(s, val)
}

func (xhr XMLHttpRequest[T]) toXMLHttpRequestResponseType(
	ctx js.CallbackContext[T], val string) (js.Value[T], error) {
	return codec.EncodeString(ctx, val)
}
