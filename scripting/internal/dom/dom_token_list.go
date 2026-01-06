package dom

import (
	"errors"

	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func DOMTokenListCustomInitializer[T any](class js.Class[T]) {
	js.InstallIterator(class, codec.EncodeString[T])
}

func DOMTokenList_toggle[T any](args js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInstance := js.As[dom.DOMTokenList](args.Instance())
	token, err0 := js.ConsumeArgument(args, "toggle", nil, codec.DecodeString)
	force, found, err1 := js.ConsumeOptionalArg(args, "force", codec.DecodeBoolean)
	if err = errors.Join(errInstance, err0, err1); err != nil {
		return nil, err
	}

	var b bool
	if found {
		if force {
			err = instance.Add(token)
			b = true
		} else {
			err = instance.Remove(token)
			b = false
		}
	} else {
		b, err = instance.Toggle(token)
	}
	if err == nil {
		res = args.NewBoolean(b)
	}
	return
}
