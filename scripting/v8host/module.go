package v8host

import (
	"fmt"

	"github.com/gost-dom/v8go"
)

// V8Module represents a compiled ECMAScript module
type V8Module struct {
	ctx    *V8ScriptContext
	module *v8go.Module
}

func (mod V8Module) Run() error {
	_, err := mod.module.Evaluate(mod.ctx.v8ctx)
	if err != nil {
		err = fmt.Errorf("gost: v8host: run module: %w", err)
	}
	return err
}

func (mod V8Module) Eval() (any, error) {
	_, err := mod.module.Evaluate(mod.ctx.v8ctx)
	if err != nil {
		err = fmt.Errorf("gost: v8host: evaluate module: %w", err)
	}
	return nil, err
}
