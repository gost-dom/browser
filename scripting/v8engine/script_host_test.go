package v8engine

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestScriptHostDocumentScriptLoading(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	win := browsertest.InitWindow(t, defaultEngine)
	err := win.LoadHTML(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
	Expect(err).ToNot(HaveOccurred())
	Expect(
		win.Eval("window.sut"),
	).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
}

func TestBasics(t *testing.T) {
	scripttests.RunBasicSuite(t, assertEngine)
}

type dummyContext struct {
	*entity.Entity
	ctx context.Context
}

func (c dummyContext) Context() context.Context { return c.ctx }
func (c dummyContext) HTTPClient() http.Client  { return *http.DefaultClient }
func (c dummyContext) LocationHREF() string     { return "http://example.com" }
func (c dummyContext) Logger() *slog.Logger     { return nil }

func TestClone(t *testing.T) {
	type T = jsTypeParam
	type Global = entity.Entity

	e := newEngine(js.ConfigurerFunc[jsTypeParam](func(e js.ScriptEngine[jsTypeParam]) {
		global := e.ConfigureGlobalScope("Global", nil)
		global.CreateOperation("store", func(ctx js.CallbackContext[T]) (js.Value[T], error) {
			t.Log("Store called")
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
			fmt.Println("Cloning")

			res, err := js.Clone(val, ctx)
			if err != nil {
				fmt.Printf("Clone err: %v\n", err)
			}
			fmt.Println("Cloned")
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
