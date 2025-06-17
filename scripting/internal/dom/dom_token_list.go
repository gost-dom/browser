package dom

import (
	"errors"

	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (l DOMTokenList[T]) CustomInitializer(class js.Class[T]) {
	it := js.NewIterator(l.toString_)
	it.InstallPrototype(class)
}

func (w DOMTokenList[T]) toString_(
	cbCtx js.Scope[T],
	val string,
) (js.Value[T], error) {
	return cbCtx.NewString(val), nil
}

func (l DOMTokenList[T]) toggle(args js.CallbackContext[T]) (js.Value[T], error) {
	instance, errInstance := js.As[dom.DOMTokenList](args.Instance())
	token, err0 := js.ConsumeArgument(args, "toggle", nil, codec.DecodeString)
	if err := errors.Join(err0, errInstance); err != nil {
		return nil, err
	}

	force, found, err1 := js.ConsumeOptionalArg(args, "force", codec.DecodeBoolean)
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

func (w DOMTokenList[T]) remove(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.remove")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	tokens, errArg1 := js.ConsumeArgument(cbCtx, "tokens", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Remove(tokens)
	return nil, nil
}
