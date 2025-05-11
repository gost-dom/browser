package v8host

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/cgo"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
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
		fmt.Println("ERROR!", n)
		panic(fmt.Sprintf("Cannot lookup node: %V", n))
	}
}

func (c *V8ScriptContext) getInstanceForNodeByName(
	constructor string,
	node entity.ObjectIder,
) (*v8.Value, error) {
	iso := c.host.iso
	if node == nil {
		return v8.Null(iso), nil
	}
	prototype, ok := c.host.globals.namedGlobals[constructor]
	if !ok {
		panic("Bad constructor name")
	}
	value, err := prototype.InstanceTemplate().NewInstance(c.v8ctx)
	if err == nil {
		objectId := node.ObjectId()
		if cached, ok := c.v8nodes[objectId]; ok {
			return cached, nil
		}
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
	log.Debug(ctx.host.logger,
		"ScriptContext: Dispose")
	for _, dispose := range ctx.disposers {
		dispose.dispose()
	}
	delete(ctx.host.contexts, ctx.v8ctx)
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

func (ctx *V8ScriptContext) compile(script string) html.Script {
	return V8Script{ctx, script}
}

func (ctx *V8ScriptContext) Compile(script string) (html.Script, error) {
	return ctx.compile(script), nil
}

func (ctx *V8ScriptContext) DownloadScript(url string) (html.Script, error) {
	resp, err := ctx.host.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(resp.Body)
	script := string(buf.Bytes())

	if resp.StatusCode != 200 {
		err := fmt.Errorf(
			"v8host: ScriptContext: bad status code: %d, downloading %s",
			resp.StatusCode,
			url,
		)
		ctx.host.logger.Error("Script download error", "err", err, "body", script)
		return nil, err
	}
	return ctx.Compile(script)
}

func (ctx *V8ScriptContext) DownloadModule(url string) (html.Script, error) {
	return nil, errors.New("v8: esm not yet supported")
}
