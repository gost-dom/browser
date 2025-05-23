package v8host

import (
	"fmt"
	"runtime/cgo"
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

	v8 "github.com/gost-dom/v8go"
)

type V8ScriptContext struct {
	htmxLoaded bool
	host       *V8ScriptHost
	v8ctx      *v8.Context
	window     html.Window
	v8nodes    map[entity.ObjectId]*v8.Value
	eventLoop  *eventLoop
	disposers  []disposable
	clock      *clock.Clock
	disposed   bool
}

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
	panic("Unknown v8 context")
}

func (c *V8ScriptContext) cacheNode(obj *v8.Object, node entity.ObjectIder) (*v8.Value, error) {
	val := obj.Value
	objectId := node.ObjectId()
	c.v8nodes[objectId] = val
	handle := cgo.NewHandle(node)
	c.addDisposer(handleDisposable(handle))
	internal := v8.NewValueExternalHandle(c.host.iso, handle)
	obj.SetInternalField(0, internal)
	return val, nil
}

func (c *V8ScriptContext) getInstanceForNode(
	node entity.ObjectIder,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	switch n := node.(type) {
	case *event.Event:
		switch n.Data.(type) {
		case event.CustomEventInit:
			return c.getInstanceForNodeByName("CustomEvent", n)
		case uievents.PointerEventInit:
			return c.getInstanceForNodeByName("PointerEvent", n)
		case uievents.MouseEventInit:
			return c.getInstanceForNodeByName("MouseEvent", n)
		case uievents.UIEventInit:
			return c.getInstanceForNodeByName("UIEvent", n)
		default:
			return c.getInstanceForNodeByName("Event", n)
		}
	case dom.Element:
		if constructor, ok := scripting.HtmlElements[strings.ToLower(n.TagName())]; ok {
			return c.getInstanceForNodeByName(constructor, n)
		}
		return c.getInstanceForNodeByName("Element", n)
	case html.HTMLDocument:
		return c.getInstanceForNodeByName("HTMLDocument", n)
	case dom.Document:
		return c.getInstanceForNodeByName("Document", n)
	case dom.DocumentFragment:
		return c.getInstanceForNodeByName("DocumentFragment", n)
	case dom.Node:
		return c.getInstanceForNodeByName("Node", n)
	case dom.Attr:
		return c.getInstanceForNodeByName("Attr", n)
	default:
		panic(fmt.Sprintf("Cannot lookup node: %v", n))
	}
}

// getConstructor returns the V8 FunctionTemplate with the specified name.
// Panics if the name is not one registered as a constructor. The name should
// not originate from client code, only from this library, so it should be
// guaranteed that this function is only called with valid values.
func (c *V8ScriptContext) getConstructor(name string) *v8.FunctionTemplate {
	prototype, ok := c.host.globals.namedGlobals[name]
	if !ok {
		panic(fmt.Sprintf("Unrecognised constructor name: %s. %s", name, constants.BUG_ISSUE_URL))
	}
	return prototype
}

func (c *V8ScriptContext) createJSInstanceForObjectOfType(
	constructor string,
	instance any,

) (*v8.Value, error) {
	iso := c.host.iso
	if instance == nil {
		return v8.Null(iso), nil
	}
	prototype := c.getConstructor(constructor)
	jsThis, err := prototype.InstanceTemplate().NewInstance(c.v8ctx)
	storeObjectHandleInV8Instance(instance, c, jsThis)
	return jsThis.Value, err
}

func (c *V8ScriptContext) getInstanceForNodeByName(
	constructor string,
	node entity.ObjectIder,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	prototype := c.getConstructor(constructor)
	objectId := node.ObjectId()
	if cached, ok := c.v8nodes[objectId]; ok {
		return cached, nil
	}
	value, err := prototype.InstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		return c.cacheNode(value, node)
	}
	return nil, err
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
		dispose.dispose()
	}
	ctx.host.deleteContext(ctx)
	ctx.v8ctx.Close()
}

func (ctx *V8ScriptContext) addDisposer(disposer disposable) {
	ctx.disposers = append(ctx.disposers, disposer)
}

func (ctx *V8ScriptContext) runScript(script string) (res *v8.Value, err error) {
	res, err = ctx.v8ctx.RunScript(script, "")
	ctx.eventLoop.tick()
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
