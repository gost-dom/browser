// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/internal/html"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("XMLHttpRequest", "XMLHttpRequestEventTarget", createXMLHttpRequestPrototype)
}

func createXMLHttpRequestPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newXMLHttpRequestV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w xmlHttpRequestV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("open", wrapV8Callback(w.scriptHost, w.open))
	prototypeTmpl.Set("setRequestHeader", wrapV8Callback(w.scriptHost, w.setRequestHeader))
	prototypeTmpl.Set("send", wrapV8Callback(w.scriptHost, w.send))
	prototypeTmpl.Set("abort", wrapV8Callback(w.scriptHost, w.abort))
	prototypeTmpl.Set("getResponseHeader", wrapV8Callback(w.scriptHost, w.getResponseHeader))
	prototypeTmpl.Set("getAllResponseHeaders", wrapV8Callback(w.scriptHost, w.getAllResponseHeaders))
	prototypeTmpl.Set("overrideMimeType", wrapV8Callback(w.scriptHost, w.overrideMimeType))

	prototypeTmpl.SetAccessorProperty("readyState",
		wrapV8Callback(w.scriptHost, w.readyState),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("timeout",
		wrapV8Callback(w.scriptHost, w.timeout),
		wrapV8Callback(w.scriptHost, w.setTimeout),
		v8.None)
	prototypeTmpl.SetAccessorProperty("withCredentials",
		wrapV8Callback(w.scriptHost, w.withCredentials),
		wrapV8Callback(w.scriptHost, w.setWithCredentials),
		v8.None)
	prototypeTmpl.SetAccessorProperty("upload",
		wrapV8Callback(w.scriptHost, w.upload),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseURL",
		wrapV8Callback(w.scriptHost, w.responseURL),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		wrapV8Callback(w.scriptHost, w.status),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusText",
		wrapV8Callback(w.scriptHost, w.statusText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseType",
		wrapV8Callback(w.scriptHost, w.responseType),
		wrapV8Callback(w.scriptHost, w.setResponseType),
		v8.None)
	prototypeTmpl.SetAccessorProperty("response",
		wrapV8Callback(w.scriptHost, w.response),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseText",
		wrapV8Callback(w.scriptHost, w.responseText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseXML",
		wrapV8Callback(w.scriptHost, w.responseXML),
		nil,
		v8.None)
}

func (w xmlHttpRequestV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return w.CreateInstance(cbCtx.Context(), info.This())
}

func (w xmlHttpRequestV8Wrapper) setRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.setRequestHeader")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
	}
	return nil, errors.New("XMLHttpRequest.setRequestHeader: Missing arguments")
}

func (w xmlHttpRequestV8Wrapper) send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.send")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	body, err1 := consumeArgument(cbCtx, "body", zeroValue, w.decodeDocument, w.decodeXMLHttpRequestBodyInit)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.SendBody(body)
		return nil, callErr
	}
	if err0 != nil {
		return nil, err0
	}
	callErr := instance.Send()
	return nil, callErr
}

func (w xmlHttpRequestV8Wrapper) abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.abort")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	callErr := instance.Abort()
	return nil, callErr
}

func (w xmlHttpRequestV8Wrapper) getResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.getResponseHeader")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetResponseHeader(name)
		return w.toNullableString_(cbCtx.Context(), result)
	}
	return nil, errors.New("XMLHttpRequest.getResponseHeader: Missing arguments")
}

func (w xmlHttpRequestV8Wrapper) getAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.getAllResponseHeaders")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, callErr := instance.GetAllResponseHeaders()
	if callErr != nil {
		return nil, callErr
	} else {
		return w.toString_(cbCtx.Context(), result)
	}
}

func (w xmlHttpRequestV8Wrapper) overrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.overrideMimeType")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	mime, err1 := consumeArgument(cbCtx, "mime", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.OverrideMimeType(mime)
		return nil, callErr
	}
	return nil, errors.New("XMLHttpRequest.overrideMimeType: Missing arguments")
}

func (w xmlHttpRequestV8Wrapper) readyState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.readyState")
	return nil, errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) timeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.timeout")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return w.toUnsignedLong(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) setTimeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.setTimeout")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.Context(), info, w.decodeUnsignedLong)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTimeout(val)
	return nil, nil
}

func (w xmlHttpRequestV8Wrapper) withCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.withCredentials")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return w.toBoolean(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) setWithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.setWithCredentials")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.Context(), info, w.decodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetWithCredentials(val)
	return nil, nil
}

func (w xmlHttpRequestV8Wrapper) responseURL(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.responseURL")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return w.toString_(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.status")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return w.toUnsignedShort(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) statusText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.statusText")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return w.toString_(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) responseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.responseType")
	return nil, errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) setResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.setResponseType")
	return nil, errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) response(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.response")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return w.toAny(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) responseText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.responseText")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return w.toString_(cbCtx.Context(), result)
}

func (w xmlHttpRequestV8Wrapper) responseXML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: XMLHttpRequest.responseXML")
	return nil, errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
