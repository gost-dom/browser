package v8host

import (
	"errors"
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/internal/constants"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	"github.com/gost-dom/browser/url"

	v8 "github.com/gost-dom/v8go"
)

type urlV8Wrapper struct {
	handleReffedObject[*url.URL]
}

func newURLV8Wrapper(host *V8ScriptHost) urlV8Wrapper {
	return urlV8Wrapper{newHandleReffedObject[*url.URL](host)}
}

type handleDisposable cgo.Handle

func (h handleDisposable) dispose() { cgo.Handle(h).Delete() }

func (w urlV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
	u string,
) (*v8.Value, error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return nil, err
	}
	w.store(value, ctx, this)
	return nil, nil
}

func (w urlV8Wrapper) CreateInstanceBase(
	ctx *V8ScriptContext,
	this *v8.Object,
	u string,
	base string,
) (*v8.Value, error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return nil, err
	}
	w.store(value, ctx, this)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := info.Args()
	var res url.URLSearchParams
	if len(args) > 0 {
		arg := args[0]
		if !arg.IsString() {
			return nil, fmt.Errorf(
				"URLSearchParams: Constructor only supports no arguments, or a string. If the argument is _valid_: %s",
				constants.BUG_USSUE_URL,
			)
		}
		var err error
		res, err = url.ParseURLSearchParams(arg.String())
		if err != nil {
			return nil, err
		}
	}
	ctx := w.mustGetContext(info)
	w.store(&res, ctx, info.This())
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) get(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return runV8FunctionCallback(w.scriptHost, info, v8URLSearchParamGet)
}

func v8URLSearchParamGet(ctx abstraction.CallbackContext) abstraction.CallbackRVal {
	instance, err0 := As[urlinterfaces.URLSearchParams](ctx.InternalInstance())
	name, err1 := ctx.ConsumeRequiredArg("name")
	if err := errors.Join(err0, err1); err != nil {
		return ctx.ReturnWithError(err)
	}
	rtnVal, hasValue := instance.Get(name.AsString())
	f := ctx.ValueFactory()
	if !hasValue {
		return ctx.ReturnWithValue(f.Null())
	}
	return ctx.ReturnWithValue(f.String(rtnVal))
}

func (w urlSearchParamsV8Wrapper) toSequenceUSVString(
	ctx *V8ScriptContext,
	values []string,
) (*v8.Value, error) {
	vs := make([]*v8.Value, len(values))
	for i, v := range values {
		vs[i], _ = v8.NewValue(ctx.host.iso, v)
	}
	return toArray(ctx.v8ctx, vs...)
}

func (w urlSearchParamsV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	iso := w.scriptHost.iso
	it := newPairIterator(
		w.scriptHost,
		func(k string, v string, ctx *V8ScriptContext) (*v8.Value, *v8.Value, error) {
			log.Info(w.scriptHost.logger, "Iterate", "key", k, "value", v)
			r1, e1 := v8.NewValue(iso, k)
			r2, e2 := v8.NewValue(iso, v)
			return r1, r2, errors.Join(e1, e2)
		},
	)
	fmt.Println("Install iterator")
	template := constructor.PrototypeTemplate()
	template.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				log.Info(w.logger(info), "ITERATOR")
				ctx := w.mustGetContext(info)
				instance, err := w.getInstance(info)
				if err != nil {
					return nil, err
				}
				return it.newIteratorInstanceOfIterable(ctx, instance)
			}),
	)
}
