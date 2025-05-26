package v8host

import (
	"errors"
	"io"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

type xmlHttpRequestV8Wrapper struct {
	handleReffedObject[XmlHttpRequest]
}

func (xhr xmlHttpRequestV8Wrapper) decodeDocument(
	cbCtx jsCallbackContext,
	val *v8.Value,
) (io.Reader, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr xmlHttpRequestV8Wrapper) decodeXMLHttpRequestBodyInit(
	cbCtx jsCallbackContext,
	val *v8.Value,
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
	if !val.IsObject() {
		return nil, errors.New("Not supported yet")
	}
	obj := val.Object()
	if res, err := getWrappedInstance[*html.FormData](obj); err == nil {
		return res.GetReader(), nil
	} else {
		return nil, err
	}
}

func newXMLHttpRequestV8Wrapper(host *V8ScriptHost) xmlHttpRequestV8Wrapper {
	return xmlHttpRequestV8Wrapper{newHandleReffedObject[XmlHttpRequest](host)}
}

func (xhr xmlHttpRequestV8Wrapper) CreateInstance(
	cbCtx *argumentHelper,
) js.CallbackRVal {
	ctx := cbCtx.ScriptCtx()
	this := cbCtx.This()
	result := NewXmlHttpRequest(ctx.window, ctx.clock)
	result.SetCatchAllHandler(event.NewEventHandlerFunc(func(event *event.Event) error {
		prop := "on" + event.Type
		handler, err := this.Get(prop)
		if err == nil && handler.IsFunction() {
			v8Event, err := ctx.getInstanceForNode(event)
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

func (xhr xmlHttpRequestV8Wrapper) open(cbCtx *argumentHelper) js.CallbackRVal {
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

func (xhr xmlHttpRequestV8Wrapper) upload(cbCtx *argumentHelper) js.CallbackRVal {
	return cbCtx.ReturnWithValue(cbCtx.This().Value)
}
