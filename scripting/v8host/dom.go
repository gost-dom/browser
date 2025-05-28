package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

func (l domTokenListV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	it := newIterator(l.scriptHost, l.toString_)
	it.installPrototype(constructor)
}

func (l domTokenListV8Wrapper) toggle(args *v8CallbackContext) (jsValue, error) {
	token, err0 := consumeArgument(args, "toggle", nil, l.decodeString)
	force, err1 := consumeArgument(args, "force", nil, l.decodeBoolean)
	instance, errInstance := js.As[dom.DOMTokenList](args.Instance())
	if args.noOfReadArguments >= 2 {
		if err := errors.Join(err0, err1, errInstance); err != nil {
			return args.ReturnWithError(err)
		}
		if force {
			instance.Add(token)
			v, err := v8.NewValue(l.scriptHost.iso, true)
			return args.ReturnWithValueErr(v, err)
		} else {
			instance.Remove(token)
			v, err := v8.NewValue(l.scriptHost.iso, false)
			return args.ReturnWithValueErr(v, err)
		}
	}
	if err := errors.Join(err0, errInstance); err != nil {
		return args.ReturnWithError(err)
	}
	v, err := v8.NewValue(l.scriptHost.iso, instance.Toggle(token))
	return args.ReturnWithValueErr(v, err)
}

func (e htmlTemplateElementV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}
