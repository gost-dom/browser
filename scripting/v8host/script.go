package v8host

import (
	"errors"
	"fmt"

	"github.com/gost-dom/v8go"
)

type V8Script struct {
	ctx    *V8ScriptContext
	script string
}

func (s V8Script) Run() error {
	_, err := s.ctx.runScript(s.script)
	return err
}

func (s V8Script) Eval() (any, error) {

	result, err := s.ctx.runScript(s.script)
	if err == nil {
		return v8ValueToGoValue(result)
	}
	return nil, err
}

func v8ValueToGoValue(result *v8go.Value) (any, error) {
	if o, err := result.AsObject(); err == nil {
		if c := o.InternalFieldCount(); c > 0 {
			f := o.GetInternalField(0)
			if f.IsExternal() {
				return f.ExternalHandle().Value(), nil
			}
		}
	}
	if result == nil {
		return nil, nil
	}
	if result.IsBoolean() {
		return result.Boolean(), nil
	}
	if result.IsInt32() {
		return result.Int32(), nil
	}
	if result.IsString() {
		return result.String(), nil
	}
	if result.IsNull() {
		return nil, nil
	}
	if result.IsUndefined() {
		return nil, nil
	}
	if o, err := result.AsObject(); err == nil {
		if c := o.InternalFieldCount(); c > 0 {
			f := o.GetInternalField(0)
			if f.IsExternal() {
				return f.ExternalHandle().Value(), nil
			}
		}
	}
	if result.IsArray() {
		obj, _ := result.AsObject()
		length, err := obj.Get("length")
		l := length.Uint32()
		errs := make([]error, l+1)
		errs[0] = err
		result := make([]any, l)
		for i := uint32(0); i < l; i++ {
			val, err := obj.GetIdx(i)
			if err == nil {
				result[i], err = v8ValueToGoValue(val)
			}
			errs[i+1] = err
		}
		return result, errors.Join(errs...)
	}
	return nil, fmt.Errorf("Value not yet supported: %v", *result)
}
