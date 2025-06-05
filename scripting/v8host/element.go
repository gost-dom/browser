package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
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

func (e *elementV8Wrapper) CustomInitializer(class js.Class[jsTypeParam]) {
	class.CreatePrototypeMethod("insertAdjacentHTML", e.insertAdjacentHTML)
	class.CreatePrototypeAttribute("outerHTML", e.outerHTML, nil)
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

func (e *elementV8Wrapper) outerHTML(cbCtx jsCallbackContext) (jsValue, error) {
	if element, err := js.As[dom.Element](cbCtx.Instance()); err == nil {
		return e.toString_(cbCtx, element.OuterHTML())
	} else {
		return nil, err
	}
}

func (e elementV8Wrapper) classList(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()
	tokenList := cbCtx.Scope().Constructor("DOMTokenList")
	return tokenList.NewInstance(cl)
}

func (e *elementV8Wrapper) toNamedNodeMap(
	cbCtx jsCallbackContext,
	n dom.NamedNodeMap,
) (jsValue, error) {
	return e.toJSWrapper(cbCtx, n)
}
