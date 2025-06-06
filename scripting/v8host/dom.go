package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

func (l domTokenListV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
	it := js.NewIterator(l.toString_)
	it.InstallPrototype(class)
}

func (w domTokenListV8Wrapper[T]) toString_(cbCtx js.Scope[T], val string) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewString(val), nil
}

func (l domTokenListV8Wrapper[T]) toggle(args js.CallbackContext[T]) (js.Value[T], error) {
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

func (e htmlTemplateElementV8Wrapper[T]) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return nil, errors.New("TODO")
}

func (w domTokenListV8Wrapper[T]) remove(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func (w parentNodeV8Wrapper[T]) decodeNodeOrText(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (dom.Node, error) {
	if val.IsString() {
		return cbCtx.Scope().Window().Document().CreateTextNode(val.String()), nil
	}
	return codec.DecodeNode(cbCtx, val)
}
