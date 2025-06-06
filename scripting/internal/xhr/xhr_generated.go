// This file is generated. Do not edit.

package xhr

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	html1 "github.com/gost-dom/browser/internal/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type FormDataV8Wrapper[T any] struct{}

func NewFormDataV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *FormDataV8Wrapper[T] {
	return &FormDataV8Wrapper[T]{}
}

func (wrapper FormDataV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w FormDataV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("delete", w.delete)
	jsClass.CreatePrototypeMethod("get", w.get)
	jsClass.CreatePrototypeMethod("getAll", w.getAll)
	jsClass.CreatePrototypeMethod("has", w.has)
	jsClass.CreatePrototypeMethod("set", w.set)
}

func (w FormDataV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.Constructor")
	form, found, errArg := js.ConsumeOptionalArg(cbCtx, "form", w.decodeHTMLFormElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceForm(cbCtx, form)
	}
	submitter, found, errArg := js.ConsumeOptionalArg(cbCtx, "submitter", codec.DecodeHTMLElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceFormSubmitter(cbCtx, form, submitter)
	}
	return w.CreateInstance(cbCtx)
}

func (w FormDataV8Wrapper[T]) append(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.append")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w FormDataV8Wrapper[T]) delete(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.delete")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Delete(name)
	return nil, nil
}

func (w FormDataV8Wrapper[T]) get(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.get")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Get(name)
	return w.toFormDataEntryValue(cbCtx, result)
}

func (w FormDataV8Wrapper[T]) getAll(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.getAll")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceFormDataEntryValue(cbCtx, result)
}

func (w FormDataV8Wrapper[T]) has(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.has")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Has(name)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w FormDataV8Wrapper[T]) set(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.set")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}

type XMLHttpRequestV8Wrapper[T any] struct{}

func NewXMLHttpRequestV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequestV8Wrapper[T] {
	return &XMLHttpRequestV8Wrapper[T]{}
}

func (wrapper XMLHttpRequestV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequestV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w XMLHttpRequestV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w XMLHttpRequestV8Wrapper[T]) setRequestHeader(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setRequestHeader")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
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

func (w XMLHttpRequestV8Wrapper[T]) send(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.send")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
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

func (w XMLHttpRequestV8Wrapper[T]) abort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.abort")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Abort()
	return nil, errCall
}

func (w XMLHttpRequestV8Wrapper[T]) getResponseHeader(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getResponseHeader")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetResponseHeader(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w XMLHttpRequestV8Wrapper[T]) getAllResponseHeaders(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getAllResponseHeaders")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result, errCall := instance.GetAllResponseHeaders()
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) overrideMimeType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.overrideMimeType")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	mime, errArg1 := js.ConsumeArgument(cbCtx, "mime", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.OverrideMimeType(mime)
	return nil, errCall
}

func (w XMLHttpRequestV8Wrapper[T]) readyState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.readyState")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w XMLHttpRequestV8Wrapper[T]) timeout(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.timeout")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Timeout()
	return codec.EncodeInt(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) setTimeout(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setTimeout")
	instance, err0 := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTimeout(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w XMLHttpRequestV8Wrapper[T]) withCredentials(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.withCredentials")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.WithCredentials()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) setWithCredentials(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setWithCredentials")
	instance, err0 := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetWithCredentials(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w XMLHttpRequestV8Wrapper[T]) responseURL(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseURL")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ResponseURL()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) status(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.status")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Status()
	return codec.EncodeInt(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) statusText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.statusText")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.StatusText()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) responseType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseType")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w XMLHttpRequestV8Wrapper[T]) setResponseType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setResponseType")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w XMLHttpRequestV8Wrapper[T]) response(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.response")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Response()
	return w.toAny(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) responseText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseText")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ResponseText()
	return codec.EncodeString(cbCtx, result)
}

func (w XMLHttpRequestV8Wrapper[T]) responseXML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseXML")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

type XMLHttpRequestEventTargetV8Wrapper[T any] struct{}

func NewXMLHttpRequestEventTargetV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequestEventTargetV8Wrapper[T] {
	return &XMLHttpRequestEventTargetV8Wrapper[T]{}
}

func (wrapper XMLHttpRequestEventTargetV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequestEventTargetV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {}

func (w XMLHttpRequestEventTargetV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequestEventTarget.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
