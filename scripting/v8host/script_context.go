package v8host

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting"
	"github.com/gost-dom/browser/scripting/internal/js"

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

func lookupJSPrototype(entity entity.ObjectIder) string {
	switch n := entity.(type) {
	case *event.Event:
		switch n.Data.(type) {
		case event.CustomEventInit:
			return "CustomEvent"
		case uievents.PointerEventInit:
			return "PointerEvent"
		case uievents.MouseEventInit:
			return "MouseEvent"
		case uievents.UIEventInit:
			return "UIEvent"
		default:
			return "Event"
		}
	case dom.Element:
		if constructor, ok := scripting.HtmlElements[strings.ToLower(n.TagName())]; ok {
			return constructor
		} else {
			return "Element"
		}
	case html.HTMLDocument:
		return "HTMLDocument"
	case dom.Document:
		return "Document"
	case dom.DocumentFragment:
		return "DocumentFragment"
	case dom.NamedNodeMap:
		return "NamedNodeMap"
	case dom.Attr:
		return "Attr"
	case dom.NodeList:
		return "NodeList"
	case dom.Node:
		return "Node"
	default:
		panic(fmt.Sprintf("Cannot lookup node: %v", n))
	}
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
		scope := v8Scope{context.host, context}
		l, err := scope.Constructor("Location").NewInstance(win.Location())
		if err != nil {
			return err
		}
		context.Constructor("Location")
		if err := context.global.Set("location", l); err != nil {
			return fmt.Errorf("error installing location: %v\n%s", err, constants.BUG_ISSUE_URL)
		}
	}
	err := installPolyfills(context)
	if err != nil {
		return fmt.Errorf(
			"Error installing polyfills: %v. Should not be possible on a passing build of Gost-DOM.\n%s",
			err,
			constants.BUG_ISSUE_URL,
		)
	}

	return nil
}

// Constructor returns the V8 FunctionTemplate with the specified name.
// Panics if the name is not one registered as a constructor. The name should
// not originate from client code, only from this library, so it should be
// guaranteed that this function is only called with valid values.
func (c *V8ScriptContext) Constructor(name string) v8Class {
	return c.getConstructor(name)
}

func (c *V8ScriptContext) getConstructor(name string) v8Class {
	prototype, ok := c.host.globals.namedGlobals[name]
	if !ok {
		panic(fmt.Sprintf("Unrecognised constructor name: %s. %s", name, constants.BUG_ISSUE_URL))
	}
	return prototype
}

func (c *V8ScriptContext) getCachedNode(this *v8.Object) (entity.ObjectIder, bool) {
	h := this.GetInternalField(0).ExternalHandle()
	r, ok := h.Value().(entity.ObjectIder)
	return r, ok
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

func (ctx *V8ScriptContext) compile(script string) V8Script {
	return V8Script{ctx, script}
}

func (ctx *V8ScriptContext) Compile(script string) (V8Script, error) {
	return ctx.compile(script), nil
}
