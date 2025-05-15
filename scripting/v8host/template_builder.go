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
	h.proto.SetAccessorPropertyCallback(name,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.mustGetContext(info.Context())
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			return fn(instance, ctx)
		}, nil, v8.ReadOnly)
}

func (h prototypeBuilder[T]) CreateReadonlyProp(name string, fn func(T) string) {
	h.proto.SetAccessorPropertyCallback(name,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			value := fn(instance)
			return v8.NewValue(h.host.iso, value)
		}, nil, v8.ReadOnly)
}

func (h prototypeBuilder[T]) CreateReadWriteProp(
	name string,
	get func(T) string,
	set func(T, string),
) {
	h.proto.SetAccessorPropertyCallback(name,
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.mustGetContext(arg.Context())
			instance, err := h.lookup(ctx, arg.This())
			if err != nil {
				return nil, err
			}
			value := get(instance)
			return v8.NewValue(h.host.iso, value)
		},
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			newVal := info.Args()[0].String()
			set(instance, newVal)
			return nil, nil
		}, v8.None)
}

func (h prototypeBuilder[T]) CreateFunction(
	name string,
	fn func(T, argumentHelper) (*v8.Value, error),
) {
	h.proto.Set(
		name,
		v8.NewFunctionTemplateWithError(
			h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := h.host.mustGetContext(info.Context())
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				return fn(instance, argumentHelper{info, ctx, 0})
			},
		),
		v8.ReadOnly,
	)
}

func (h prototypeBuilder[T]) CreateFunctionStringToString(name string, fn func(T, string) string) {
	h.CreateFunction(name, func(instance T, info argumentHelper) (*v8.Value, error) {
		value := fn(instance, info.Args()[0].String())
		return v8.NewValue(h.host.iso, value)
	})
}

// parseSetterArg parses a single argument and is intended for attribute
// setters, where exactly one argument must be passed by v8.
func parseSetterArg[T any](
	ctx *V8ScriptContext,
	info *v8.FunctionCallbackInfo,
	parsers ...func(*V8ScriptContext, *v8.Value) (T, error),
) (result T, err error) {
	args := info.Args()
	if len(args) != 1 {
		err = fmt.Errorf("parseSetterArg: expected one argument. got: %d", len(args))
		return
	}

	value := args[0]
	errs := make([]error, len(parsers))
	for i, parser := range parsers {
		result, errs[i] = parser(ctx, value)
		if errs[i] == nil {
			return
		}
	}
	err = fmt.Errorf("parseSetterArg: conversion errors: %w", errors.Join(errs...))
	return
}

func tryParseArg[T any](
	args *argumentHelper,
	index int,
	parsers ...func(*V8ScriptContext, *v8.Value) (T, error),
) (result T, err error) {
	value := args.getArg(index)
	if value == nil {
		return
	}
	errs := make([]error, len(parsers))
	for i, parser := range parsers {
		result, errs[i] = parser(args.ctx, value)
		if errs[i] == nil {
			return
		}
	}
	err = fmt.Errorf("tryParseArg: argument at index %d: %w", index, errors.Join(errs...))
	return
}

func tryParseArgNullableType[T any](
	args *argumentHelper,
	index int,
	parsers ...func(*V8ScriptContext, *v8.Value) (res T, err error),
) (result T, err error) {
	value := args.getArg(index)
	if value == nil {
		args.acceptIndex(index)
		return
	}
	for _, parser := range parsers {
		result, err = parser(args.ctx, value)
		if err == nil {
			args.acceptIndex(index)
			return
		}
	}
	err = errors.New("TODO")
	return
}

func tryParseArgWithDefault[T any](
	args *argumentHelper,
	index int,
	defaultValue func() T,
	parsers ...func(*V8ScriptContext, *v8.Value) (T, error),
) (result T, err error) {
	if index >= len(args.Args()) {
		args.acceptIndex(index)
		return defaultValue(), nil
	} else {
		return tryParseArg(args, index, parsers...)
	}
}
