package v8host

import v8 "github.com/gost-dom/v8go"

func toArray(ctx *v8.Context, values ...*v8.Value) (*v8.Value, error) {
	// Total hack, v8go doesn't expose Array values, so we polyfill the engine
	var err error
	if arrayOf, err := ctx.RunScript("Array.of", "gost-polyfills-array"); err == nil {
		if fn, err := arrayOf.AsFunction(); err == nil {
			args := make([]v8.Valuer, len(values))
			for i, v := range values {
				args[i] = v
			}
			return fn.Call(ctx.Global(), args...)
		}
	}
	return nil, err
}
