// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateDOMTokenListPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESDOMTokenList(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.Set("item", v8.NewFunctionTemplateWithError(iso, wrapper.Item))
	prototype.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.Contains))
	prototype.Set("add", v8.NewFunctionTemplateWithError(iso, wrapper.Add))
	prototype.Set("remove", v8.NewFunctionTemplateWithError(iso, wrapper.Remove))
	prototype.Set("toggle", v8.NewFunctionTemplateWithError(iso, wrapper.Toggle))
	prototype.Set("replace", v8.NewFunctionTemplateWithError(iso, wrapper.Replace))
	prototype.Set("supports", v8.NewFunctionTemplateWithError(iso, wrapper.Supports))

	prototype.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.Length),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetValue),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetValue),
		v8.None)
	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (u ESDOMTokenList) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(u.host.iso, "Illegal Constructor")
}

func (u ESDOMTokenList) Item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	index, err := TryParseArg(args, 0, u.DecodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return u.ToNullableDOMString(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u ESDOMTokenList) Contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	token, err := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return u.ToBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u ESDOMTokenList) Add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	tokens, err := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		if err != nil {
			return nil, err
		}
		err = instance.Add(tokens)
		return nil, err
	}
	return nil, errors.New("Missing arguments")
}

func (u ESDOMTokenList) Remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	tokens, err := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (u ESDOMTokenList) Replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	token, err0 := TryParseArg(args, 0, u.DecodeDOMString)
	newToken, err1 := TryParseArg(args, 1, u.DecodeDOMString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return u.ToBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u ESDOMTokenList) Supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: DOMTokenList.supports")
}

func (u ESDOMTokenList) Length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return u.ToUnsignedLong(ctx, result)
}

func (u ESDOMTokenList) GetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetValue()
	return u.ToDOMString(ctx, result)
}

func (u ESDOMTokenList) SetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := newArgumentHelper(u.host, info)
	val, err := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		if err != nil {
			return nil, err
		}
		instance.SetValue(val)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}
func CreateHTMLTemplateElementPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESHTMLTemplateElement(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, wrapper.Content),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("shadowRootMode",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootMode),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootMode),
		v8.None)
	prototype.SetAccessorProperty("shadowRootDelegatesFocus",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootDelegatesFocus),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootDelegatesFocus),
		v8.None)
	prototype.SetAccessorProperty("shadowRootClonable",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootClonable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootClonable),
		v8.None)
	prototype.SetAccessorProperty("shadowRootSerializable",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetShadowRootSerializable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootSerializable),
		v8.None)
	return constructor
}

func (e ESHTMLTemplateElement) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	return e.CreateInstance(ctx, info.This())
}

func (e ESHTMLTemplateElement) Content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	instance, err := e.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return e.ToDocumentFragment(ctx, result)
}

func (e ESHTMLTemplateElement) GetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootMode")
}

func (e ESHTMLTemplateElement) SetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootMode")
}

func (e ESHTMLTemplateElement) GetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootDelegatesFocus")
}

func (e ESHTMLTemplateElement) SetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootDelegatesFocus")
}

func (e ESHTMLTemplateElement) GetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootClonable")
}

func (e ESHTMLTemplateElement) SetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootClonable")
}

func (e ESHTMLTemplateElement) GetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.GetShadowRootSerializable")
}

func (e ESHTMLTemplateElement) SetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootSerializable")
}
