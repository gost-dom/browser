// This file is generated. Do not edit.

package xhr

import (
	"errors"
	html "github.com/gost-dom/browser/internal/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type XMLHttpRequest[T any] struct{}

func NewXMLHttpRequest[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequest[T] {
	return &XMLHttpRequest[T]{}
}

func (wrapper XMLHttpRequest[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequest[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("open", w.open)
	jsClass.CreatePrototypeMethod("setRequestHeader", w.setRequestHeader)
	jsClass.CreatePrototypeMethod("send", w.send)
	jsClass.CreatePrototypeMethod("abort", w.abort)
	jsClass.CreatePrototypeMethod("getResponseHeader", w.getResponseHeader)
	jsClass.CreatePrototypeMethod("getAllResponseHeaders", w.getAllResponseHeaders)
	jsClass.CreatePrototypeMethod("overrideMimeType", w.overrideMimeType)
	jsClass.CreatePrototypeAttribute("readyState", w.readyState, nil)
	jsClass.CreatePrototypeAttribute("timeout", w.timeout, w.setTimeout)
	jsClass.CreatePrototypeAttribute("withCredentials", w.withCredentials, w.setWithCredentials)
	jsClass.CreatePrototypeAttribute("upload", w.upload, nil)
	jsClass.CreatePrototypeAttribute("responseURL", w.responseURL, nil)
	jsClass.CreatePrototypeAttribute("status", w.status, nil)
	jsClass.CreatePrototypeAttribute("statusText", w.statusText, nil)
	jsClass.CreatePrototypeAttribute("responseType", w.responseType, w.setResponseType)
	jsClass.CreatePrototypeAttribute("response", w.response, nil)
	jsClass.CreatePrototypeAttribute("responseText", w.responseText, nil)
	jsClass.CreatePrototypeAttribute("responseXML", w.responseXML, nil)
}

func (w XMLHttpRequest[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w XMLHttpRequest[T]) setRequestHeader(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setRequestHeader")
	instance, errInst := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetRequestHeader(name, value)
	return nil, nil
}

func (w XMLHttpRequest[T]) send(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.send")
	instance, errInst := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	body, found, errArg := js.ConsumeOptionalArg(cbCtx, "body", w.decodeDocument, w.decodeXMLHttpRequestBodyInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		errCall := instance.SendBody(body)
		return nil, errCall
	}
	errCall := instance.Send()
	return nil, errCall
}

func (w XMLHttpRequest[T]) abort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.abort")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Abort()
	return nil, errCall
}

func (w XMLHttpRequest[T]) getResponseHeader(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getResponseHeader")
	instance, errInst := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetResponseHeader(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w XMLHttpRequest[T]) getAllResponseHeaders(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getAllResponseHeaders")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, errCall := instance.GetAllResponseHeaders()
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequest[T]) overrideMimeType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.overrideMimeType")
	instance, errInst := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	mime, errArg1 := js.ConsumeArgument(cbCtx, "mime", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.OverrideMimeType(mime)
	return nil, errCall
}

func (w XMLHttpRequest[T]) readyState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.readyState")
	return nil, errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w XMLHttpRequest[T]) timeout(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.timeout")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return codec.EncodeInt(cbCtx, result)
}

func (w XMLHttpRequest[T]) setTimeout(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setTimeout")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTimeout(val)
	return nil, nil
}

func (w XMLHttpRequest[T]) withCredentials(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.withCredentials")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w XMLHttpRequest[T]) setWithCredentials(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setWithCredentials")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetWithCredentials(val)
	return nil, nil
}

func (w XMLHttpRequest[T]) responseURL(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseURL")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequest[T]) status(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.status")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return codec.EncodeInt(cbCtx, result)
}

func (w XMLHttpRequest[T]) statusText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.statusText")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequest[T]) responseType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseType")
	return nil, errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w XMLHttpRequest[T]) setResponseType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setResponseType")
	return nil, errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w XMLHttpRequest[T]) response(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.response")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return w.toAny(cbCtx, result)
}

func (w XMLHttpRequest[T]) responseText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseText")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequest[T]) responseXML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseXML")
	return nil, errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
