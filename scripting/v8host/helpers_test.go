package v8host_test

import (
	"fmt"

	"github.com/gost-dom/browser/html"
	. "github.com/onsi/ginkgo/v2"
)

type TestScriptContext struct {
	html.ScriptContext
	ignoreUnhandledErrors bool
	window                html.Window
}

func (c TestScriptContext) Window() html.Window {
	return c.window
}

func (c TestScriptContext) MustRunTestScript(script string) any {
	result, err := c.Eval(script)
	if err != nil {
		panic(
			fmt.Sprintf(
				"Script error. Script src:\n-----\n%s\n-----\nError:\n%s",
				script,
				err.Error(),
			),
		)
	}
	return result
}

type CreateHook func(ctx *TestScriptContext)

var IgnoreUnhandledErrors CreateHook = func(ctx *TestScriptContext) {
	ctx.ignoreUnhandledErrors = true
}

// NewTextContext loads HTML into a browser for a single Ginkgo test. It
// installs the proper Ginkgo cleanup handler.
func NewTestContext(hooks ...CreateHook) TestScriptContext {
	ctx := TestScriptContext{}
	window := html.NewWindow(html.WindowOptions{
		// ScriptEngineFactory: (*Wrapper)(host),
	})
	ctx.window = window
	ctx.ScriptContext = host.NewContext(window)
	DeferCleanup(ctx.Close)
	for _, hook := range hooks {
		hook(&ctx)
	}
	return ctx
}

// InitializeContextWithEmptyHtml is useful when multiple tests need has the
// same initial HTML. The html will be parsed by a normal HTML parser, which
// automatically wraps content in <html> and <body> if those are missing. So you
// So passing `<div>foo</div>` will be the same as
// `<html><body><div>foo</div></body></html>`.
//
// Example:
//
//	Describe("Tests with shared setup", func () {
//		ctx := InitializeContextWithEmptyHtml(
//			"<body><div>Hello, world!</div></body>");
//
//		 It("Should should find Hello, world! in first div", func () { /*...*/ }
//		It("Should should have one child of body", func () { /*...*/ }
//	})
func InitializeContext(hooks ...CreateHook) *TestScriptContext {
	ctx := TestScriptContext{}

	BeforeEach(func() {
		window := html.NewWindow(html.WindowOptions{
			ScriptHost: host,
		})
		ctx.window = window
		ctx.ScriptContext = window.ScriptContext()
		for _, hook := range hooks {
			hook(&ctx)
		}
		// ctx.ScriptContext = window.
		DeferCleanup(func() { window.Close() })
	})

	return &ctx
}

func LoadHTML(html string) CreateHook {
	return func(ctx *TestScriptContext) {
		ctx.Window().LoadHTML(html)
	}
}

func InitializeContextWithEmptyHtml() *TestScriptContext {
	return InitializeContext(LoadHTML("<html></html>"))
}
