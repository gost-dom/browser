package scripttests

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/testing/htmltest"
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
				t.Errorf("Clone error: %v", err)
			}
			return res, err
		})
	}))

	global1 := new(Global)
	global2 := new(Global)

	c1 := htmltest.NewScriptContextHelper(
		t,
		e.NewHost(html.ScriptEngineOptions{}).NewContext(dummyContext{global1, t.Context()}),
	)
	c2 := htmltest.NewScriptContextHelper(
		t,
		e.NewHost(html.ScriptEngineOptions{}).NewContext(dummyContext{global2, t.Context()}),
	)

	assert.NoError(t, c1.Run(`
		const b = {
			id: "b",
		}
		const arr = [1,2,3]
		const recursiveArray = [1,2,3]
		recursiveArray.push(recursiveArray)
		const a = {
			stringVal: "hello",
			numberVal: 42.5,
			trueVal: true,
			falseVal: false,
			b1: b,
			b2: b,
			arr1: arr,
			arr2: arr,
			recursiveArray,
		}
		globalThis.store(a)
	`))
	val, ok := entity.ComponentType[js.Value[T]](global1)
	assert.True(t, ok)
	entity.SetComponentType(global2, val)

	assert.NoError(t, c2.Run("const cloned = globalThis.get()"))
	assert.Equal(t, "hello", c2.MustEval("cloned.stringVal"))
	assert.Equal(t, 42.5, c2.MustEval("cloned.numberVal"))
	assert.True(t, c2.MustEval("cloned.trueVal").(bool))
	assert.False(t, c2.MustEval("cloned.falseVal").(bool))
	assert.True(t, c2.MustEval("cloned.b1 === cloned.b2").(bool))
	assert.Equal(t, "1,2,3", c2.MustEval("cloned.arr1.join(',')"))
}
