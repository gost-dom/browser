// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	html1 "github.com/gost-dom/browser/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("FormData", "", createFormDataPrototype)
}

type formDataV8Wrapper struct {
	handleReffedObject[*html.FormData, jsTypeParam]
}

func newFormDataV8Wrapper(scriptHost *V8ScriptHost) *formDataV8Wrapper {
	return &formDataV8Wrapper{newHandleReffedObject[*html.FormData](scriptHost)}
}

func createFormDataPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newFormDataV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	wrapper.CustomInitializer(constructor)
	return constructor
}

func (w formDataV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("delete", w.delete)
	jsClass.CreatePrototypeMethod("get", w.get)
	jsClass.CreatePrototypeMethod("getAll", w.getAll)
	jsClass.CreatePrototypeMethod("has", w.has)
	jsClass.CreatePrototypeMethod("set", w.set)
}

func (w formDataV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.constructor")
	form, found, errArg := consumeOptionalArg(cbCtx, "form", w.decodeHTMLFormElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceForm(cbCtx, form)
	}
	submitter, found, errArg := consumeOptionalArg(cbCtx, "submitter", w.decodeHTMLElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceFormSubmitter(cbCtx, form, submitter)
	}
	return w.CreateInstance(cbCtx)
}

func (w formDataV8Wrapper) append(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.append")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w formDataV8Wrapper) delete(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.delete")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Delete(name)
	return nil, nil
}

func (w formDataV8Wrapper) get(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.get")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Get(name)
	return w.toFormDataEntryValue(cbCtx, result)
}

func (w formDataV8Wrapper) getAll(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.getAll")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceFormDataEntryValue(cbCtx, result)
}

func (w formDataV8Wrapper) has(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.has")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Has(name)
	return w.toBoolean(cbCtx, result)
}

func (w formDataV8Wrapper) set(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: FormData.set")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}

func init() {
	registerJSClass("XMLHttpRequest", "XMLHttpRequestEventTarget", createXMLHttpRequestPrototype)
}

func createXMLHttpRequestPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newXMLHttpRequestV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return constructor
}

func (w xmlHttpRequestV8Wrapper) installPrototype(jsClass v8Class) {
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

func (w xmlHttpRequestV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.constructor")
	return w.CreateInstance(cbCtx)
}

func (w xmlHttpRequestV8Wrapper) setRequestHeader(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setRequestHeader")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetRequestHeader(name, value)
	return nil, nil
}

func (w xmlHttpRequestV8Wrapper) send(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.send")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	body, found, errArg := consumeOptionalArg(cbCtx, "body", w.decodeDocument, w.decodeXMLHttpRequestBodyInit)
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

func (w xmlHttpRequestV8Wrapper) abort(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.abort")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Abort()
	return nil, errCall
}

func (w xmlHttpRequestV8Wrapper) getResponseHeader(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getResponseHeader")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetResponseHeader(name)
	return w.toNillableString_(cbCtx, result, hasValue)
}

func (w xmlHttpRequestV8Wrapper) getAllResponseHeaders(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.getAllResponseHeaders")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result, errCall := instance.GetAllResponseHeaders()
	if errCall != nil {
		return nil, errCall
	}
	return w.toString_(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) overrideMimeType(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.overrideMimeType")
	instance, errInst := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	mime, errArg1 := consumeArgument(cbCtx, "mime", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.OverrideMimeType(mime)
	return nil, errCall
}

func (w xmlHttpRequestV8Wrapper) readyState(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.readyState")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w xmlHttpRequestV8Wrapper) timeout(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.timeout")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Timeout()
	return w.toUnsignedLong(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) setTimeout(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setTimeout")
	instance, err0 := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeUnsignedLong)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTimeout(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w xmlHttpRequestV8Wrapper) withCredentials(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.withCredentials")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.WithCredentials()
	return w.toBoolean(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) setWithCredentials(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setWithCredentials")
	instance, err0 := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetWithCredentials(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w xmlHttpRequestV8Wrapper) responseURL(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseURL")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ResponseURL()
	return w.toString_(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) status(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.status")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Status()
	return w.toUnsignedShort(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) statusText(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.statusText")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.StatusText()
	return w.toString_(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) responseType(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseType")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w xmlHttpRequestV8Wrapper) setResponseType(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.setResponseType")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w xmlHttpRequestV8Wrapper) response(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.response")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Response()
	return w.toAny(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) responseText(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseText")
	instance, err := js.As[html1.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ResponseText()
	return w.toString_(cbCtx, result)
}

func (w xmlHttpRequestV8Wrapper) responseXML(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequest.responseXML")
	return cbCtx.ReturnWithError(errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
