package scripttests

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/stretchr/testify/assert"
)

// RunScriptEngineSuites runs test suites of the script engine without a
// predefined global scope. In contrast to [RunSuites] that expect an engine
// configured for the global scope in the Window realm.
func RunScriptEngineSuites[T any](t *testing.T, f ScriptEngineFactory[T]) {
	type Global = entity.Entity

	e := f(js.ConfigurerFunc[T](func(e js.ScriptEngine[T]) {
		global := e.ConfigureGlobalScope("Global", nil)
		global.CreateOperation("store", func(ctx js.CallbackContext[T]) (js.Value[T], error) {
			v, ok := ctx.ConsumeArg()
			if !ok {
				return nil, ctx.NewTypeError("Missing argument")
			}
			c, err := js.As[entity.Components](ctx.GlobalThis().NativeValue(), nil)
			if err != nil {
				return nil, err
			}
			entity.SetComponentType(c, v)
			return nil, nil
		})

		global.CreateOperation("get", func(ctx js.CallbackContext[T]) (js.Value[T], error) {
			c, err := js.As[entity.Components](ctx.GlobalThis().NativeValue(), nil)
			if err != nil {
				return nil, err
			}
			val, ok := entity.ComponentType[js.Value[T]](c)
			if !ok {
				return nil, fmt.Errorf("Value missing")
			}

			res, err := js.Clone(val, ctx)
			if err != nil {
				fmt.Printf("Clone err: %v\n", err)
			}
			return res, err
		})
	}))

	global1 := new(Global)
	global2 := new(Global)

	c1 := e.NewHost(html.ScriptEngineOptions{}).NewContext(dummyContext{global1, t.Context()})
	c2 := e.NewHost(html.ScriptEngineOptions{}).NewContext(dummyContext{global2, t.Context()})

	assert.NoError(t, c1.Run(`
		const a = {
			foo: "hello"
		}
		globalThis.store(a)
	`))
	val, ok := entity.ComponentType[js.Value[T]](global1)
	assert.True(t, ok)
	entity.SetComponentType(global2, val)

	res, err := c2.Eval(`
		globalThis.get().foo
	`)
	assert.NoError(t, err)
	assert.Equal(t, "hello", res)
}
