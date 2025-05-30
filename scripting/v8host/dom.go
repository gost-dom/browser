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
	instance, errInstance := js.As[dom.DOMTokenList](args.Instance())
	token, err0 := consumeArgument(args, "toggle", nil, l.decodeString)
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}

	force, found, err1 := consumeOptionalArg(args, "force", l.decodeBoolean)
	if found {
		if err1 != nil {
			return nil, err1
		}
		if force {
			instance.Add(token)
			return args.ValueFactory().NewBoolean(true), nil
		} else {
			instance.Remove(token)
			return args.ValueFactory().NewBoolean(false), nil
		}
	}
	return args.ValueFactory().NewBoolean(instance.Toggle(token)), nil
}

func (e htmlTemplateElementV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}
