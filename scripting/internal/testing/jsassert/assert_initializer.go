package jsassert

import (
	_ "embed"

	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

//go:embed assertions.js
var assertions string

func Configure[T any](e js.ScriptEngine[T]) {
	gost := e.CreateGlobalObject("gost")
	gost.CreateFunction("error", func(ctx js.CallbackContext[T]) (js.Value[T], error) {
		if win, err := codec.GetWindow(ctx); err == nil {
			if t, ok := htmltest.GetTestingT(win); ok {
				t.Helper()
				msg, err := js.ConsumeArgument(ctx, "message", nil, codec.DecodeString)
				if err != nil {
					msg = "(missing message)"
				}
				t.Error(msg)
				return nil, nil
			}
		}
		return nil, ctx.NewTypeError("error")
	})
	e.InstallPolyfill(assertions, "gost-dom/assertions")
}
