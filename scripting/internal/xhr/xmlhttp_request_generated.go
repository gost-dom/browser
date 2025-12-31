// This file is generated. Do not edit.

package xhr

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	xhr "github.com/gost-dom/browser/internal/html/xhr"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type XMLHttpRequest[T any] struct{}

func NewXMLHttpRequest[T any](scriptHost js.ScriptEngine[T]) XMLHttpRequest[T] {
	return XMLHttpRequest[T]{}
}

func (wrapper XMLHttpRequest[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequest[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("open", XMLHttpRequest_open)
	jsClass.CreateOperation("setRequestHeader", XMLHttpRequest_setRequestHeader)
	jsClass.CreateOperation("send", XMLHttpRequest_send)
	jsClass.CreateOperation("abort", XMLHttpRequest_abort)
	jsClass.CreateOperation("getResponseHeader", XMLHttpRequest_getResponseHeader)
	jsClass.CreateOperation("getAllResponseHeaders", XMLHttpRequest_getAllResponseHeaders)
	jsClass.CreateOperation("overrideMimeType", XMLHttpRequest_overrideMimeType)
	jsClass.CreateAttribute("readyState", XMLHttpRequest_readyState, nil)
	jsClass.CreateAttribute("timeout", XMLHttpRequest_timeout, XMLHttpRequest_setTimeout)
	jsClass.CreateAttribute("withCredentials", XMLHttpRequest_withCredentials, XMLHttpRequest_setWithCredentials)
	jsClass.CreateAttribute("upload", XMLHttpRequest_upload, nil)
	jsClass.CreateAttribute("responseURL", XMLHttpRequest_responseURL, nil)
	jsClass.CreateAttribute("status", XMLHttpRequest_status, nil)
	jsClass.CreateAttribute("statusText", XMLHttpRequest_statusText, nil)
	jsClass.CreateAttribute("responseType", XMLHttpRequest_responseType, XMLHttpRequest_setResponseType)
	jsClass.CreateAttribute("response", XMLHttpRequest_response, nil)
	jsClass.CreateAttribute("responseText", XMLHttpRequest_responseText, nil)
	jsClass.CreateAttribute("responseXML", XMLHttpRequest_responseXML, nil)
}

func XMLHttpRequestConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return CreateXMLHttpRequest(cbCtx)
}

func XMLHttpRequest_setRequestHeader[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeByteString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetRequestHeader(name, value)
	return nil, nil
}

func XMLHttpRequest_send[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	body, errArg1 := js.ConsumeArgument(cbCtx, "body", codec.ZeroValue, decodeDocument, decodeXMLHttpRequestBodyInit)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Send(body)
	return nil, errCall
}

func XMLHttpRequest_abort[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Abort()
	return nil, errCall
}

func XMLHttpRequest_getResponseHeader[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetResponseHeader(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func XMLHttpRequest_getAllResponseHeaders[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, errCall := instance.GetAllResponseHeaders()
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeString(cbCtx, result)
}

func XMLHttpRequest_overrideMimeType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
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

func XMLHttpRequest_readyState[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "XMLHttpRequest.XMLHttpRequest_readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func XMLHttpRequest_timeout[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return codec.EncodeInt(cbCtx, result)
}

func XMLHttpRequest_setTimeout[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTimeout(val)
	return nil, nil
}

func XMLHttpRequest_withCredentials[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return codec.EncodeBoolean(cbCtx, result)
}

func XMLHttpRequest_setWithCredentials[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetWithCredentials(val)
	return nil, nil
}

func XMLHttpRequest_responseURL[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return codec.EncodeString(cbCtx, result)
}

func XMLHttpRequest_status[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return codec.EncodeInt(cbCtx, result)
}

func XMLHttpRequest_statusText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return codec.EncodeString(cbCtx, result)
}

func XMLHttpRequest_responseType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseType()
	return encodeXMLHttpRequestResponseType(cbCtx, result)
}

func XMLHttpRequest_setResponseType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, decodeXMLHttpRequestResponseType)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetResponseType(val)
	return nil, nil
}

func XMLHttpRequest_response[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return encodeAny(cbCtx, result)
}

func XMLHttpRequest_responseText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[xhr.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return codec.EncodeString(cbCtx, result)
}

func XMLHttpRequest_responseXML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "XMLHttpRequest.XMLHttpRequest_responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
