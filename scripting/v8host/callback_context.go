package v8host

import (
	"fmt"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	"github.com/gost-dom/v8go"
)

func runV8FunctionCallback(
	host *V8ScriptHost,
	info *v8go.FunctionCallbackInfo,
	cb abstraction.FunctionCallback,
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
) abstraction.CallbackContext {
	ctx := host.mustGetContext(info.Context())
	return &v8CallbackContext{
		host: host,
		ctx:  ctx,
		info: info,
		args: info.Args(),
	}
}

func (c *v8CallbackContext) iso() *v8go.Isolate { return c.host.iso }

func (c *v8CallbackContext) ConsumeRequiredArg(name string) (abstraction.Value, error) {
	if c.argsConsumed >= len(c.args) {
		return nil, fmt.Errorf("%w: %w",
			abstraction.ErrMissingArgument, v8go.NewTypeError(c.iso(), "missing argument: "+name))
	}
	arg := c.args[c.argsConsumed]
	c.argsConsumed++
	return v8Value{arg}, nil
}

func (c *v8CallbackContext) ConsumeOptionalArg() (abstraction.Value, bool) {
	// TODO: Implement
	panic("NOT IMPLEMENTED")
}

func (c *v8CallbackContext) ConsumeRestArgs() []abstraction.Value {
	// TODO: Implement
	panic("NOT IMPLEMENTED")
}

func (c *v8CallbackContext) InternalInstance() (any, error) {
	if c.info.This().InternalFieldCount() == 0 {
		return nil, abstraction.ErrNoInternalValue
	}
	field := c.info.This().GetInternalField(0)
	handle := field.ExternalHandle()
	return handle.Value(), nil
}

func (v *v8CallbackContext) ReturnWithValue(val abstraction.Value) abstraction.CallbackRVal {
	return v8CallbackRVal{rtnVal: val.(v8Value).Value}
}

func (v *v8CallbackContext) ReturnWithError(err error) abstraction.CallbackRVal {
	return v8CallbackRVal{err: err}
}

func (c *v8CallbackContext) ValueFactory() abstraction.ValueFactory {
	return v8ValueFactory{c.host}
}

type v8CallbackRVal struct {
	rtnVal *v8go.Value
	err    error
}

type v8Value struct{ *v8go.Value }

func (v v8Value) AsString() string { return v.Value.String() }

type v8ValueFactory struct{ host *V8ScriptHost }

func (f v8ValueFactory) iso() *v8go.Isolate      { return f.host.iso }
func (f v8ValueFactory) Null() abstraction.Value { return f.toVal(v8go.Null(f.iso())) }

func (f v8ValueFactory) String(s string) abstraction.Value {
	return f.mustVal(v8go.NewValue(f.iso(), s))
}

// mustVal is just a simple helper to crete Value wrappers on top of v8go values
// where construction is assumed to succeed
func (f v8ValueFactory) mustVal(val *v8go.Value, err error) abstraction.Value {
	if err != nil {
		panic(
			fmt.Sprintf(
				"v8 value conversion. value construction was assumed to succeed but failed: %v. %s",
				err,
				constants.BUG_USSUE_URL,
			),
		)
	}
	return v8Value{val}
}
func (f v8ValueFactory) toVal(val *v8go.Value) abstraction.Value {
	return v8Value{val}
}

func As[T any](val any, err error) (rtnVal T, rtnErr error) {
	if err != nil {
		rtnErr = err
		return
	}
	var ok bool
	if rtnVal, ok = val.(T); !ok {
		rtnErr = fmt.Errorf("value %+v is not assignable to requested type %T", val, rtnVal)
	}
	return
}
