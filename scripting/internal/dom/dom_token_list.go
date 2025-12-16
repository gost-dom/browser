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

func (w DOMTokenList[T]) toString_(s js.Scope[T], val string) (js.Value[T], error) {
	return s.NewString(val), nil
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
			return args.NewBoolean(true), nil
		} else {
			instance.Remove(token)
			return args.NewBoolean(false), nil
		}
	}
	res, err := instance.Toggle(token)
	if err != nil {
		return nil, err
	}
	return args.NewBoolean(res), nil
}
