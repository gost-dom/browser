package v8host

import (
	"fmt"
	"runtime/debug"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

func runV8FunctionCallback(
	host *V8ScriptHost,
	info *v8go.FunctionCallbackInfo,
	cb js.FunctionCallback,
) (*v8go.Value, error) {
	ctx := newCallbackContext(host, info)
	rtn := cb(ctx).(v8CallbackRVal)
	return rtn.rtnVal, rtn.err
}

type v8CallbackContext struct {
	host         *V8ScriptHost
	ctx          *V8ScriptContext
	info         *v8go.FunctionCallbackInfo
	args         []*v8go.Value
	argsConsumed int
}

func newCallbackContext(
	host *V8ScriptHost,
	info *v8go.FunctionCallbackInfo,
) js.CallbackContext {
	ctx := host.mustGetContext(info.Context())
	return &v8CallbackContext{
		host: host,
		ctx:  ctx,
		info: info,
		args: info.Args(),
	}
}

func (c *v8CallbackContext) iso() *v8go.Isolate { return c.host.iso }

func (c *v8CallbackContext) ConsumeRequiredArg(name string) (js.Value, error) {
	if c.argsConsumed >= len(c.args) {
		return nil, fmt.Errorf("%w: %w",
			js.ErrMissingArgument, v8go.NewTypeError(c.iso(), "missing argument: "+name))
	}
	arg := c.args[c.argsConsumed]
	c.argsConsumed++
	return v8Value{arg}, nil
}

func (c *v8CallbackContext) ConsumeOptionalArg() (js.Value, bool) {
	// TODO: Implement
	panic("NOT IMPLEMENTED")
}

func (c *v8CallbackContext) ConsumeRestArgs() []js.Value {
	// TODO: Implement
	panic("NOT IMPLEMENTED")
}

func (c *v8CallbackContext) InternalInstance() (any, error) {
	if c.info.This().InternalFieldCount() == 0 {
		return nil, js.ErrNoInternalValue
	}
	field := c.info.This().GetInternalField(0)
	handle := field.ExternalHandle()
	return handle.Value(), nil
}

func (v *v8CallbackContext) ReturnWithValue(val js.Value) js.CallbackRVal {
	return v8CallbackRVal{rtnVal: val.(v8Value).Value}
}

func (v *v8CallbackContext) ReturnWithError(err error) js.CallbackRVal {
	return v8CallbackRVal{err: err}
}

func (c *v8CallbackContext) ValueFactory() js.ValueFactory {
	return v8ValueFactory{c.host}
}

type v8CallbackRVal struct {
	rtnVal *v8go.Value
	err    error
}

type v8Value struct{ *v8go.Value }

func (v v8Value) AsString() string { return v.Value.String() }

type v8ValueFactory struct{ host *V8ScriptHost }

func (f v8ValueFactory) iso() *v8go.Isolate { return f.host.iso }
func (f v8ValueFactory) Null() js.Value     { return f.toVal(v8go.Null(f.iso())) }

func (f v8ValueFactory) String(s string) js.Value {
	return f.mustVal(v8go.NewValue(f.iso(), s))
}

// mustVal is just a simple helper to crete Value wrappers on top of v8go values
// where construction is assumed to succeed
func (f v8ValueFactory) mustVal(val *v8go.Value, err error) js.Value {
	if err != nil {
		panic(
			fmt.Sprintf(
				"v8 value conversion. value construction was assumed to succeed but failed: %v. %s",
				err,
				constants.BUG_ISSUE_URL,
			),
		)
	}
	return v8Value{val}
}
func (f v8ValueFactory) toVal(val *v8go.Value) js.Value {
	return v8Value{val}
}

type internalCallback func(*argumentHelper) js.CallbackRVal

func wrapV8Callback(
	host *V8ScriptHost,
	callback internalCallback,
) *v8go.FunctionTemplate {
	return v8go.NewFunctionTemplateWithError(
		host.iso,
		func(info *v8go.FunctionCallbackInfo) (res *v8go.Value, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("PANIC in callback: %v\n%s", r, debug.Stack())
				}
			}()
			cbCtx := newArgumentHelper(host, info)
			result := callback(cbCtx).(v8CallbackRVal)
			return result.rtnVal, result.err
		},
	)
}
