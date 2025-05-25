package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"

	v8 "github.com/gost-dom/v8go"
)

type elementV8Wrapper struct {
	handleReffedObject[dom.Element]
	parentNode               *parentNodeV8Wrapper
	nonDocumentTypeChildNode *nonDocumentTypeChildNodeV8Wrapper
}

func newElementV8Wrapper(host *V8ScriptHost) *elementV8Wrapper {
	return &elementV8Wrapper{
		newHandleReffedObject[dom.Element](host),
		newParentNodeV8Wrapper(host),
		newNonDocumentTypeChildNodeV8Wrapper(host),
	}
}

func (e *elementV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	iso := e.scriptHost.iso
	prototype := constructor.PrototypeTemplate()
	prototype.Set(
		"insertAdjacentHTML",
		v8.NewFunctionTemplateWithError(iso, e.insertAdjacentHTML),
	)
	prototype.SetAccessorProperty(
		"outerHTML",
		v8.NewFunctionTemplateWithError(iso, e.outerHTML),
		nil,
		v8.None,
	)
	prototype.SetAccessorProperty(
		"textContent",
		nil,
		v8.NewFunctionTemplateWithError(iso, e.setTextContent),
		v8.None,
	)
}

func (e *elementV8Wrapper) insertAdjacentHTML(
	info *v8.FunctionCallbackInfo,
) (val *v8.Value, err error) {
	iso := e.scriptHost.iso
	arg := newArgumentHelper(e.scriptHost, info)
	element, e0 := e.getInstance(info)
	position, e1 := arg.consumeString()
	html, e2 := arg.consumeString()
	err = errors.Join(e0, e1, e2)
	if err == nil {
		element.InsertAdjacentHTML(position, html)
		val, err = v8.NewValue(iso, element.OuterHTML())
	}
	return
}

func (e *elementV8Wrapper) outerHTML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	if i, err := e.getInstance(info); err == nil {
		return v8.NewValue(e.scriptHost.iso, i.OuterHTML())
	} else {
		return nil, err
	}
}

func (w *elementV8Wrapper) setTextContent(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	e, err := w.getInstance(info)
	if err == nil {
		e.SetTextContent(info.Args()[0].String())
	}
	return nil, err
}

func (e elementV8Wrapper) classList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	tokenList := e.scriptHost.globals.namedGlobals["DOMTokenList"]
	ctx := e.scriptHost.mustGetContext(info.Context())
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	res, err := tokenList.InstanceTemplate().NewInstance(ctx.v8ctx)
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()

	storeObjectHandleInV8Instance(cl, ctx, res)
	return res.Value, nil
}

func (e *elementV8Wrapper) toNamedNodeMap(
	ctx *V8ScriptContext,
	n dom.NamedNodeMap,
) (*v8.Value, error) {
	return ctx.getInstanceForNodeByName("NamedNodeMap", n)
}

func (w elementV8Wrapper) getAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	helper := newArgumentHelper(w.scriptHost, info)
	element, e0 := w.getInstance(info)
	name, e1 := helper.consumeString()
	err := errors.Join(e0, e1)
	if err != nil {
		return nil, err
	}
	if r, ok := element.GetAttribute(name); ok {
		return v8.NewValue(w.scriptHost.iso, r)
	} else {
		return v8.Null(w.scriptHost.iso), nil
	}
}
