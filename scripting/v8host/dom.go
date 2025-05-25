package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

func (l domTokenListV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	// constructor.InstanceTemplate().SetSymbol(
	// 	v8.SymbolIterator(l.scriptHost.iso),
	// 	v8.NewFunctionTemplateWithError(l.scriptHost.iso, l.GetIterator),
	// )
	it := newIterator(l.scriptHost, func(val string, ctx *V8ScriptContext) (*v8.Value, error) {
		return v8.NewValue(ctx.host.iso, val)
	})
	it.installPrototype(constructor)
}

func (l domTokenListV8Wrapper) toggle(args *argumentHelper) (*v8.Value, error) {
	token, err0 := consumeArgument(args, "toggle", nil, l.decodeString)
	force, err1 := consumeArgument(args, "force", nil, l.decodeBoolean)
	instance, errInstance := js.As[dom.DOMTokenList](args.Instance())
	if args.noOfReadArguments >= 2 {
		if err := errors.Join(err0, err1, errInstance); err != nil {
			return nil, err
		}
		if force {
			instance.Add(token)
			return v8.NewValue(l.scriptHost.iso, true)
		} else {
			instance.Remove(token)
			return v8.NewValue(l.scriptHost.iso, false)
		}
	}
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}
	return v8.NewValue(l.scriptHost.iso, instance.Toggle(token))
}

func (e htmlTemplateElementV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}

func (e htmlTemplateElementV8Wrapper) ToDocumentFragment(
	ctx *V8ScriptContext,
	fragment dom.DocumentFragment,
) (*v8.Value, error) {
	return ctx.getInstanceForNode(fragment)
}
