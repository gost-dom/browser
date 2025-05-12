package v8host

import (
	"fmt"
	"io"
	"runtime/debug"
	"strings"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/url"
	"github.com/gost-dom/v8go"

	v8 "github.com/gost-dom/v8go"
)

type V8ScriptContext struct {
	htmxLoaded bool
	host       *V8ScriptHost
	v8ctx      *v8.Context
	window     html.Window
	v8nodes    map[entity.ObjectId]jsValue
	disposers  []js.Disposable
	clock      *clock.Clock
	disposed   bool
	global     jsObject
	resolver   moduleResolver
}

func (c *V8ScriptContext) iso() *v8.Isolate    { return c.host.iso }
func (c *V8ScriptContext) Window() html.Window { return c.window }

func (h *V8ScriptHost) getContext(v8ctx *v8.Context) (*V8ScriptContext, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	ctx, ok := h.contexts[v8ctx]
	return ctx, ok
}

func (h *V8ScriptHost) mustGetContext(v8ctx *v8.Context) *V8ScriptContext {
	if ctx, ok := h.getContext(v8ctx); ok {
		return ctx
	}
	panic("Unknown v8 context!!\n" + string(debug.Stack()))
}

func (c *V8ScriptContext) cacheEntity(obj jsObject, node entity.ObjectIder) {
	c.v8nodes[node.ObjectId()] = obj
}

func (c *V8ScriptContext) GetValue(entity entity.ObjectIder) (jsValue, bool) {
	res, ok := c.v8nodes[entity.ObjectId()]
	return res, ok
}

func (c *V8ScriptContext) SetValue(entity entity.ObjectIder, value jsValue) {
	c.v8nodes[entity.ObjectId()] = value
}

func (context *V8ScriptContext) initializeGlobals() error {
	win := context.window
	context.global = newV8Object(context, context.v8ctx.Global())
	context.global.SetNativeValue(win)
	{
		// Install window.location as a "data property".
		//
		// A better solution could have been an accessor property that
		// automatically retrieves a cached JS object for a Go value. This could
		// have been configured on the template, and be consistent with how code
		// gen works.
		//
		// Currently, Window.Location() returns a new value every time,
		// preventing that solution, as the location is a [SameValue] object -
		// and verified by tests. Furthermore, in JS, document.location must
		// return the same object as window.location (document.location looks up
		// window.location)
		l, err := context.Constructor("Location").NewInstance(win.Location())
		if err != nil {
			return err
		}
		if err := context.global.Set("location", l); err != nil {
			return fmt.Errorf("error installing location: %v\n%s", err, constants.BUG_ISSUE_URL)
		}
	}
	for _, s := range context.host.scripts {
		script := s[0]
		src := s[1]
		_, err := context.v8ctx.RunScript(script, src)
		context.clock.Tick()
		if err != nil {
			return fmt.Errorf("v8host: install globals (%s): %w", src, err)
		}
	}

	return nil
}

// Constructor returns the V8 FunctionTemplate with the specified name.
// Panics if the name is not one registered as a constructor. The name should
// not originate from client code, only from this library, so it should be
// guaranteed that this function is only called with valid values.
func (c *V8ScriptContext) Constructor(name string) js.Constructor[jsTypeParam] {
	return v8Constructable{c, c.getConstructor(name)}
}

func (c *V8ScriptContext) getConstructor(name string) v8Class {
	prototype, ok := c.host.globals.namedGlobals[name]
	if !ok {
		panic(fmt.Sprintf("Unrecognised constructor name: %s. %s", name, constants.BUG_ISSUE_URL))
	}
	return prototype
}

func (ctx *V8ScriptContext) Clock() html.Clock { return ctx.clock }

func (host *V8ScriptHost) addContext(ctx *V8ScriptContext) {
	host.mu.Lock()
	defer host.mu.Unlock()
	host.contexts[ctx.v8ctx] = ctx
}

func (ctx *V8ScriptContext) Close() {
	if ctx.disposed {
		panic("Context already disposed")
	}
	ctx.disposed = true

	ctx.host.inspector.ContextDestroyed(ctx.v8ctx)
	log.Debug(ctx.host.logger, "ScriptContext: Dispose")
	for _, dispose := range ctx.disposers {
		dispose.Dispose()
	}
	ctx.host.deleteContext(ctx)
	ctx.v8ctx.Close()
}

func (ctx *V8ScriptContext) addDisposer(disposer js.Disposable) {
	ctx.disposers = append(ctx.disposers, disposer)
}

func (ctx *V8ScriptContext) runScript(script string) (res *v8.Value, err error) {
	res, err = ctx.v8ctx.RunScript(script, "")
	ctx.clock.Tick()
	return
}

func (ctx *V8ScriptContext) Run(script string) error {
	return ctx.compile(script).Run()
}

func (ctx *V8ScriptContext) Eval(script string) (any, error) {
	return ctx.compile(script).Eval()
}

func (ctx *V8ScriptContext) compile(script string) html.Script {
	return V8Script{ctx, script}
}

func (ctx *V8ScriptContext) Compile(script string) (html.Script, error) {
	return ctx.compile(script), nil
}

func (ctx *V8ScriptContext) DownloadScript(url string) (html.Script, error) {
	script, err := ctx.resolver.download(url)
	if err != nil {
		return nil, err
	}
	return ctx.Compile(script)
}

func (ctx *V8ScriptContext) DownloadModule(url string) (html.Script, error) {
	module, err := ctx.resolver.downloadAndCompile(url)
	if err = module.InstantiateModule(ctx.v8ctx, &ctx.resolver); err != nil {
		return nil, fmt.Errorf("gost: v8host: module instantiation: %w", err)
	}
	return V8Module{ctx, module, ctx.host.logger, url}, nil
}

type resolvedModule struct {
	scriptID int
	location string
	module   *v8go.Module
}

type moduleResolver struct {
	host    *V8ScriptHost
	modules []resolvedModule
}

func (r *moduleResolver) download(url string) (string, error) {
	resp, err := r.host.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("gost: v8host: download errors: %w", err)
	}
	defer resp.Body.Close()
	var buf strings.Builder
	io.Copy(&buf, resp.Body)
	script := buf.String()

	if resp.StatusCode != 200 {
		err := fmt.Errorf(
			"gost: v8host: ScriptContext: bad status code: %d, downloading %s",
			resp.StatusCode,
			url,
		)
		r.host.logger.Error("Script download error", "err", err, "body", script)
		return "", err
	}
	return script, nil
}

func (r *moduleResolver) cached(url string) *v8go.Module {
	for _, m := range r.modules {
		if m.location == url {
			return m.module
		}
	}
	return nil
}

func (r *moduleResolver) downloadAndCompile(url string) (*v8go.Module, error) {
	if cached := r.cached(url); cached != nil {
		return cached, nil
	}

	script, err := r.download(url)
	if err != nil {
		return nil, err
	}
	module, err := v8.CompileModule(r.host.iso, script, url)
	if err != nil {
		return nil, fmt.Errorf("gost: v8host: module compilation: %w", err)
	}
	r.modules = append(r.modules, resolvedModule{module.ScriptID(), url, module})
	return module, nil
}

func (r *moduleResolver) get(scriptID int) (resolvedModule, bool) {
	for _, m := range r.modules {
		if m.scriptID == scriptID {
			return m, true
		}
	}
	return resolvedModule{}, false
}

func (r *moduleResolver) ResolveModule(
	v8ctx *v8go.Context,
	spec string,
	attr v8go.ImportAttributes,
	ref *v8go.Module,
) (*v8go.Module, error) {
	refModule, found := r.get(ref.ScriptID())
	if !found {
		return nil, fmt.Errorf(
			"gost: referrer not cached. This is a bug in Gost-DOM. Please file an issue at: %s",
			constants.BUG_ISSUE_URL,
		)
	}
	url := url.ParseURLBase(spec, refModule.location).Href()
	return r.downloadAndCompile(url)
}
