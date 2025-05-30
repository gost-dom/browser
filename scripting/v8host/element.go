package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

type elementV8Wrapper struct {
	handleReffedObject[dom.Element, jsTypeParam]
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
	prototype.Set("insertAdjacentHTML", wrapV8Callback(e.scriptHost, e.insertAdjacentHTML))
	prototype.SetAccessorProperty(
		"outerHTML",
		v8.NewFunctionTemplateWithError(iso, e.outerHTML),
		nil,
		v8.None,
	)
}

func (e *elementV8Wrapper) insertAdjacentHTML(cbCtx jsCallbackContext) (val jsValue, err error) {
	element, e0 := js.As[dom.Element](cbCtx.Instance())
	position, e1 := consumeArgument(cbCtx, "position", nil, decodeString)
	html, e2 := consumeArgument(cbCtx, "html", nil, decodeString)
	err = errors.Join(e0, e1, e2)
	if err == nil {
		err = element.InsertAdjacentHTML(position, html)
	}
	return nil, err
}

func (e *elementV8Wrapper) outerHTML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	if i, err := e.getInstance(info); err == nil {
		return v8.NewValue(e.scriptHost.iso, i.OuterHTML())
	} else {
		return nil, err
	}
}

func (e elementV8Wrapper) classList(cbCtx *v8CallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	cl := instance.ClassList()
	tokenList := cbCtx.ScriptCtx().getConstructor("DOMTokenList")
	return tokenList.NewInstance(cbCtx.ScriptCtx(), cl)
}

func (e *elementV8Wrapper) toNamedNodeMap(
	cbCtx jsCallbackContext,
	n dom.NamedNodeMap,
) (jsValue, error) {
	return e.toJSWrapper(cbCtx, n)
}
