package xhr

import (
	"errors"
	"io"

	"github.com/gost-dom/browser/dom/event"
	xhrint "github.com/gost-dom/browser/internal/html/xhr"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func decodeDocument[T any](
	_ js.Scope[T],
	val js.Value[T],
) (io.Reader, error) {
	if js.IsNullish(val) {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func decodeXMLHttpRequestBodyInit[T any](
	s js.Scope[T],
	val js.Value[T],
) (io.Reader, error) {
	return codec.DecodeRequestBody(s, val)
}

func CreateXMLHttpRequest[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := codec.GetWindow(cbCtx)
	if err != nil {
		return nil, err
	}
	this := cbCtx.This()
	result := xhrint.NewXmlHttpRequest(win, cbCtx.Clock())
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
	return codec.EncodeConstructedValue(cbCtx, result)
}

func XMLHttpRequest_open[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, errInstance := js.As[xhrint.XmlHttpRequest](cbCtx.Instance())
	method, err0 := js.ConsumeArgument(cbCtx, "method", nil, codec.DecodeString)
	url, err1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if err := errors.Join(err0, err1, errInstance); err != nil {
		return nil, err
	}
	if async, found, err2 := js.ConsumeOptionalArg(cbCtx, "async", codec.DecodeBoolean); found {
		if err2 != nil {
			return nil, err2
		}
		instance.Open(method, url, xhrint.RequestOptionAsync(async))
		return nil, nil
	}
	instance.Open(method, url)
	return nil, nil
}

func XMLHttpRequest_upload[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func encodeAny[T any](s js.Scope[T], val string) (js.Value[T], error) {
	return codec.EncodeString(s, val)
}

func decodeXMLHttpRequestResponseType[T any](
	s js.Scope[T], val js.Value[T],
) (string, error) {
	return codec.DecodeString(s, val)
}

func encodeXMLHttpRequestResponseType[T any](
	ctx js.CallbackContext[T], val string) (js.Value[T], error) {
	return codec.EncodeString(ctx, val)
}
