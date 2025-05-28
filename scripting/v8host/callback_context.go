package v8host

import (
	"fmt"
	"runtime/debug"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/v8go"
)

type v8ValueFactory struct{ host *V8ScriptHost }

func (f v8ValueFactory) iso() *v8go.Isolate { return f.host.iso }
func (f v8ValueFactory) Null() jsValue      { return f.toVal(v8go.Null(f.iso())) }

func (f v8ValueFactory) String(s string) jsValue {
	return f.mustVal(v8go.NewValue(f.iso(), s))
}

// mustVal is just a simple helper to crete Value wrappers on top of v8go values
// where construction is assumed to succeed
func (f v8ValueFactory) mustVal(val *v8go.Value, err error) jsValue {
	if err != nil {
		panic(
			fmt.Sprintf(
				"v8 value conversion. value construction was assumed to succeed but failed: %v. %s",
				err,
				constants.BUG_ISSUE_URL,
			),
		)
	}
	return &v8Value{f.iso(), val}
}
func (f v8ValueFactory) toVal(val *v8go.Value) jsValue {
	return &v8Value{f.iso(), val}
}

type internalCallback func(*argumentHelper) (jsValue, error)

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
			result, err := callback(cbCtx)
			val := assertV8Value(result)
			return val.v8Value(), err
		},
	)
}

func wrapV8CallbackFn(
	host *V8ScriptHost,
	callback internalCallback,
) v8go.FunctionCallbackWithError {
	return func(info *v8go.FunctionCallbackInfo) (res *v8go.Value, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("PANIC in callback: %v\n%s", r, debug.Stack())
			}
		}()
		cbCtx := newArgumentHelper(host, info)
		result, err := callback(cbCtx)
		val := assertV8Value(result)
		return val.v8Value(), err
	}
}

/* -------- Decoders -------- */

func decodeInt32(cbCtx jsCallbackContext, val jsValue) (int32, error) {
	return val.Int32(), nil
}
