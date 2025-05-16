package v8host

import (
	"errors"

	"github.com/gost-dom/browser/html"

	v8 "github.com/gost-dom/v8go"
)

type formDataV8Wrapper struct {
	handleReffedObject[*html.FormData]
}

func newFormDataV8Wrapper(host *V8ScriptHost) formDataV8Wrapper {
	return formDataV8Wrapper{newHandleReffedObject[*html.FormData](host)}
}

func (w formDataV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	info *v8.FunctionCallbackInfo,
) (*v8.Value, error) {
	this := info.This()
	var form html.HTMLFormElement
	if len(info.Args()) > 0 {
		arg := info.Args()[0]
		n, err := w.decodeNode(ctx, arg)
		if err != nil {
			return nil, err
		}
		var ok bool
		form, ok = n.(html.HTMLFormElement)
		if !ok {
			return nil, v8.NewTypeError(w.scriptHost.iso, "form must be an HTMLFormElement")
		}
	}
	var value *html.FormData
	if form != nil {
		value = html.NewFormDataForm(form)
	} else {
		value = html.NewFormData()
	}
	w.store(value, ctx, this)
	return nil, nil
}

func createFormData(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newFormDataV8Wrapper(host)
	builder := newConstructorBuilder[*html.FormData](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			return wrapper.CreateInstance(ctx, info)
		},
	)
	stringIterator := newIterator(
		host,
		func(instance string, ctx *V8ScriptContext) (*v8.Value, error) {
			return v8.NewValue(ctx.host.iso, instance)
		},
	)

	entryIterator := newIterator(
		host,
		func(instance html.FormDataEntry, ctx *V8ScriptContext) (*v8.Value, error) {
			k, e1 := v8.NewValue(iso, instance.Name)
			v, e2 := v8.NewValue(iso, string(instance.Value))
			err := errors.Join(e1, e2)
			if err != nil {
				return nil, err
			}
			return toArray(ctx.v8ctx, k, v)
		})
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	prototype := protoBuilder.proto
	builder.constructor.InstanceTemplate().SetSymbol(
		v8.SymbolIterator(iso),
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			data, err := wrapper.getInstance(info)
			if err != nil {
				return nil, err
			}
			return stringIterator.newIteratorInstance(ctx, data.Keys())
		},
	)
	prototype.Set(
		"append",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				key, err1 := args.getStringArg(0)
				value, err2 := args.getStringArg(1)
				err := errors.Join(err0, err1, err2)
				if err != nil {
					return nil, err
				}
				instance.Append(key, html.FormDataValue(value))
				return nil, nil
			},
		),
	)

	prototype.Set(
		"get",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				if err0 != nil {
					return nil, err0
				}
				key, err := args.getStringArg(0)
				if err != nil {
					return nil, err
				}
				val := string(instance.Get(key))
				return v8.NewValue(iso, val)
			},
		),
	)
	prototype.Set(
		"delete",
		v8.NewFunctionTemplateWithError(host.iso,
			func(info *v8.FunctionCallbackInfo) (result *v8.Value, err error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				key, err1 := args.getStringArg(0)
				if err := errors.Join(err0, err1); err != nil {
					return nil, err
				}
				instance.Delete(key)
				return nil, nil
			}),
	)
	prototype.Set(
		"keys",
		v8.NewFunctionTemplateWithError(host.iso,
			func(info *v8.FunctionCallbackInfo) (result *v8.Value, err error) {
				args := newArgumentHelper(host, info)
				instance, err0 := wrapper.getInstance(info)
				if err0 != nil {
					return nil, err0
				}
				return stringIterator.newIteratorInstance(args.ctx, instance.Keys())
			}),
	)

	getEntries := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			instance, err := wrapper.getInstance(info)
			if err != nil {
				return nil, err
			}
			return entryIterator.newIteratorInstance(ctx, instance.Entries)
		},
	)
	protoBuilder.proto.Set("entries", getEntries)
	protoBuilder.proto.SetSymbol(v8.SymbolIterator(iso), getEntries)
	return builder.constructor
}
