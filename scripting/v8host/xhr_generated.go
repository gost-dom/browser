// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("XMLHttpRequest", "XMLHttpRequestEventTarget", createXMLHttpRequestPrototype)
}

func createXMLHttpRequestPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newXMLHttpRequestV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

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

func (w xmlHttpRequestV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w xmlHttpRequestV8Wrapper) setRequestHeader(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.setRequestHeader")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
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

func (w xmlHttpRequestV8Wrapper) send(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.send")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
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

func (w xmlHttpRequestV8Wrapper) abort(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.abort")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	callErr := instance.Abort()
	return nil, callErr
}

func (w xmlHttpRequestV8Wrapper) getResponseHeader(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.getResponseHeader")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetResponseHeader(name)
		return w.toNullableString_(cbCtx.ScriptCtx(), result)
	}
	return nil, errors.New("XMLHttpRequest.getResponseHeader: Missing arguments")
}

func (w xmlHttpRequestV8Wrapper) getAllResponseHeaders(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.getAllResponseHeaders")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, callErr := instance.GetAllResponseHeaders()
	if callErr != nil {
		return nil, callErr
	} else {
		return w.toString_(cbCtx.ScriptCtx(), result)
	}
}

func (w xmlHttpRequestV8Wrapper) overrideMimeType(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.overrideMimeType")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
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

func (w xmlHttpRequestV8Wrapper) readyState(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.readyState")
	return nil, errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) timeout(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.timeout")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return w.toUnsignedLong(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) setTimeout(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.setTimeout")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.ScriptCtx(), cbCtx, w.decodeUnsignedLong)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTimeout(val)
	return nil, nil
}

func (w xmlHttpRequestV8Wrapper) withCredentials(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.withCredentials")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return w.toBoolean(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) setWithCredentials(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.setWithCredentials")
	instance, err0 := js.As[html.XMLHttpRequest](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.ScriptCtx(), cbCtx, w.decodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetWithCredentials(val)
	return nil, nil
}

func (w xmlHttpRequestV8Wrapper) responseURL(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.responseURL")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) status(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.status")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return w.toUnsignedShort(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) statusText(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.statusText")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) responseType(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.responseType")
	return nil, errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) setResponseType(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.setResponseType")
	return nil, errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w xmlHttpRequestV8Wrapper) response(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.response")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return w.toAny(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) responseText(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.responseText")
	instance, err := js.As[html.XMLHttpRequest](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w xmlHttpRequestV8Wrapper) responseXML(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: XMLHttpRequest.responseXML")
	return nil, errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
