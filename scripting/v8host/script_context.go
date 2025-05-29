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

	v8 "github.com/gost-dom/v8go"
)

type V8ScriptContext struct {
	htmxLoaded bool
	host       *V8ScriptHost
	v8ctx      *v8.Context
	window     html.Window
	v8nodes    map[entity.ObjectId]jsValue
	eventLoop  *eventLoop
	disposers  []disposable
	clock      *clock.Clock
	disposed   bool
	global     jsObject
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
	panic("Unknown v8 context!!\n" + string(debug.Stack()))
}

func (c *V8ScriptContext) cacheEntity(obj jsObject, node entity.ObjectIder) {
	c.v8nodes[node.ObjectId()] = obj
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

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func (c *V8ScriptContext) getJSInstance(
	node entity.ObjectIder,
) (jsValue, error) {
	iso := c.host.iso
	if node == nil {
		return newV8Value(iso, v8.Null(iso)), nil
	}
	objectId := node.ObjectId()
	if cached, ok := c.v8nodes[objectId]; ok {
		return cached, nil
	}

	prototypeName := lookupJSPrototype(node)
	prototype := c.getConstructor(prototypeName)
	value, err := prototype.NewInstance(c, node)
	if err == nil {
		c.cacheEntity(value, node)
	}
	return value, err
}

// getConstructor returns the V8 FunctionTemplate with the specified name.
// Panics if the name is not one registered as a constructor. The name should
// not originate from client code, only from this library, so it should be
// guaranteed that this function is only called with valid values.
func (c *V8ScriptContext) getConstructor(name string) v8Constructor {
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
