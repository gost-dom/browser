package v8host

import (
	"errors"
	"fmt"

	v8 "github.com/gost-dom/v8go"
)

type constructorBuilder[T any] struct {
	host           *V8ScriptHost
	constructor    *v8.FunctionTemplate
	instanceLookup func(*V8ScriptContext, *v8.Object) (T, error)
}

func createIllegalConstructor(host *V8ScriptHost) *v8.FunctionTemplate {
	result := v8.NewFunctionTemplateWithError(
		host.iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(host.iso, "Illegal Constructor")
		},
	)
	result.InstanceTemplate().SetInternalFieldCount(1)
	return result
}

func newConstructorBuilder[T any](
	host *V8ScriptHost,
	cb v8.FunctionCallbackWithError,
) constructorBuilder[T] {
	constructor := v8.NewFunctionTemplateWithError(
		host.iso,
		cb,
	)
	constructor.InstanceTemplate().SetInternalFieldCount(1)

	builder := constructorBuilder[T]{host: host,
		constructor: constructor,
	}
	return builder
}

func newIllegalConstructorBuilder[T any](host *V8ScriptHost) constructorBuilder[T] {
	constructor := createIllegalConstructor(host)

	builder := constructorBuilder[T]{host: host,
		constructor: constructor,
	}
	return builder
}

func getInstanceFromThis[T any](ctx *V8ScriptContext, this *v8.Object) (instance T, err error) {
	cachedEntity, ok := ctx.getCachedNode(this)
	if !ok {
		err = errors.New("No cached entity could be found for `this`")
		return
	}
	if i, e_ok := cachedEntity.(T); e_ok && ok {
		return i, nil
	} else {
		err = v8.NewTypeError(ctx.host.iso, "Not an object of the correct type")
		return
	}
}

func (c *constructorBuilder[T]) SetDefaultInstanceLookup() {
	c.instanceLookup = func(ctx *V8ScriptContext, this *v8.Object) (val T, err error) {
		instance, ok := ctx.getCachedNode(this)
		if instance, e_ok := instance.(T); e_ok && ok {
			return instance, nil
		} else {
			err = v8.NewTypeError(ctx.host.iso, "Not an instance of NamedNodeMap")
			return
		}
	}
}

func (c constructorBuilder[T]) NewPrototypeBuilder() prototypeBuilder[T] {
	if c.instanceLookup == nil {
		panic("Cannot build prototype builder if instance lookup not specified")
	}
	return prototypeBuilder[T]{
		host:   c.host,
		proto:  c.constructor.PrototypeTemplate(),
		lookup: c.instanceLookup,
	}
}

func (c constructorBuilder[T]) NewInstanceBuilder() prototypeBuilder[T] {
	if c.instanceLookup == nil {
		panic("Cannot build prototype builder if instance lookup not specified")
	}
	return prototypeBuilder[T]{
		host:   c.host,
		proto:  c.constructor.InstanceTemplate(),
		lookup: c.instanceLookup,
	}
}

type prototypeBuilder[T any] struct {
	host   *V8ScriptHost
	proto  *v8.ObjectTemplate
	lookup func(*V8ScriptContext, *v8.Object) (T, error)
}

func (b constructorBuilder[T]) GetInstance(info *v8.FunctionCallbackInfo) (T, error) {
	ctx := b.host.mustGetContext(info.Context())
	return b.instanceLookup(ctx, info.This())
}

func (b prototypeBuilder[T]) GetInstance(info *v8.FunctionCallbackInfo) (T, error) {
	ctx := b.host.mustGetContext(info.Context())
	return b.lookup(ctx, info.This())
}

func (h prototypeBuilder[T]) CreateReadonlyProp2(
	name string,
	fn func(T, *V8ScriptContext) (*v8.Value, error),
) {
	h.proto.SetAccessorProperty(
		name,
		v8.NewFunctionTemplateWithError(
			h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := h.host.mustGetContext(info.Context())
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				return fn(instance, ctx)
			},
		),
		nil,
		v8.ReadOnly,
	)
}

func (h prototypeBuilder[T]) CreateReadonlyProp(name string, fn func(T) string) {
	h.proto.SetAccessorProperty(name,
		v8.NewFunctionTemplateWithError(h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				value := fn(instance)
				return v8.NewValue(h.host.iso, value)
			}), nil, v8.ReadOnly)
}

func (h prototypeBuilder[T]) CreateReadWriteProp(
	name string,
	get func(T) string,
	set func(T, string),
) {
	h.proto.SetAccessorProperty(name,
		v8.NewFunctionTemplateWithError(h.host.iso,
			func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := h.host.mustGetContext(arg.Context())
				instance, err := h.lookup(ctx, arg.This())
				if err != nil {
					return nil, err
				}
				value := get(instance)
				return v8.NewValue(h.host.iso, value)
			}),
		v8.NewFunctionTemplateWithError(h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				newVal := info.Args()[0].String()
				set(instance, newVal)
				return nil, nil
			}), v8.None)
}

func (h prototypeBuilder[T]) CreateFunction(
	name string,
	fn func(T, *v8CallbackContext) (*v8.Value, error),
) {
	h.proto.Set(
		name,
		v8.NewFunctionTemplateWithError(
			h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				return fn(instance, newCallbackContext(h.host, info))
			},
		),
		v8.ReadOnly,
	)
}

// parseSetterArg parses a single argument and is intended for attribute
// setters, where exactly one argument must be passed by v8.
func parseSetterArg[T any](
	ctx jsCallbackContext,
	parsers ...func(jsCallbackContext, jsValue) (T, error),
) (result T, err error) {
	arg := ctx.ConsumeArg()
	if arg == nil {
		err = fmt.Errorf("parseSetterArg: expected one argument. got: %d", len(ctx.v8Info.Args()))
	}

	errs := make([]error, len(parsers))
	for i, parser := range parsers {
		result, errs[i] = parser(ctx, arg)
		if errs[i] == nil {
			return
		}
	}
	err = fmt.Errorf("parseSetterArg: conversion errors: %w", errors.Join(errs...))
	return
}

func zeroValue[T any]() (res T) { return }

func ignoreArgument(args *v8CallbackContext) {
	args.ConsumeArg()
	args.acceptIndex(args.noOfReadArguments)
}

// consumeArgument pulls one of the passed arguments and tries to convert it to
// target type T using one of the passed decoders. The return value will be
// taken from the first decode that does not return an error. If no decoder is
// succeeds, an error is returned.
//
// If no more arguments are present, or the next argument is undefined, the
// defaultValue function will be used if not nil; otherwise an error is
// returned.
//
// If the function returns with an error, the name will be used in the error
// message. Otherwise, name has ho effect on the function.
func consumeArgument[T any](
	args *v8CallbackContext,
	name string,
	defaultValue func() T,
	decoders ...func(*v8CallbackContext, jsValue) (T, error),
) (result T, err error) {
	index := args.currentIndex
	value := args.ConsumeArg()
	if value == nil && defaultValue != nil {
		args.acceptIndex(index)
		return defaultValue(), nil
	} else {
		errs := make([]error, len(decoders))
		if value != nil {
			for i, parser := range decoders {
				result, errs[i] = parser(args, value)
				if errs[i] == nil {
					return
				}
			}
		}
		// TODO: This should eventually become a TypeError in JS
		err = fmt.Errorf("tryParseArg: %s: %w", name, errors.Join(errs...))
		return
	}
}

func consumeOptionalArg[T any](
	args *v8CallbackContext,
	name string,
	decoders ...func(*v8CallbackContext, jsValue) (T, error),
) (result T, found bool, err error) {
	value := args.ConsumeArg()
	if value == nil {
		return
	}
	found = true
	errs := make([]error, len(decoders))
	if value != nil {
		for i, parser := range decoders {
			result, errs[i] = parser(args, value)
			if errs[i] == nil {
				return
			}
		}
	}
	// TODO: This should eventually become a TypeError in JS
	err = fmt.Errorf("tryParseArg: %s: %w", name, errors.Join(errs...))
	return
}
