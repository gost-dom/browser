package v8host

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

type MutationCallback struct {
	ctx      *V8ScriptContext
	function jsFunction
}

func (cb MutationCallback) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	v8Recs, _ := toSequenceMutationRecord(cb.ctx, recs)

	cb.function.Call(cb.ctx.global, v8Recs)
}

func (w mutationObserverV8Wrapper) CreateInstance(
	cbCtx *argumentHelper,
	cb mutation.Callback,
) (jsValue, error) {
	ctx := cbCtx.ScriptCtx()
	return cbCtx.ReturnWithJSValueErr(
		w.store(mutation.NewObserver(ctx.clock, cb), ctx, cbCtx.This()),
	)
}

func (w mutationObserverV8Wrapper) decodeMutationCallback(
	cbCtx jsCallbackContext,
	val jsValue,
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback{cbCtx.ScriptCtx(), f}, nil
	}
	return nil, v8go.NewTypeError(cbCtx.iso(), "Not a function")
}

func (w mutationObserverV8Wrapper) decodeObserveOption(
	cbCtx jsCallbackContext,
	val jsValue,
) ([]mutation.ObserveOption, error) {
	obj, ok := val.AsObject()
	if !ok {
		return nil, v8go.NewTypeError(cbCtx.iso(), "Obtions not an object")
	}
	var res []mutation.ObserveOption
	ap := func(key string, o mutation.ObserveOption) {
		if v, err := obj.Get(key); err == nil && v.Boolean() {
			res = append(res, o)
		}
	}
	ap("subtree", mutation.Subtree)
	ap("childList", mutation.ChildList)
	ap("attributes", mutation.Attributes)
	ap("attributeOldValue", mutation.AttributeOldValue)
	ap("characterData", mutation.CharacterData)
	ap("characterDataOldValue", mutation.CharacterDataOldValue)

	filter, err := obj.Get("attributeFilter")
	if filter.Boolean() {
		return nil, errors.New("mutationObject filter not yet implemented")
	}
	return res, err
}

func (w mutationObserverV8Wrapper) toSequenceMutationRecord(
	cbCtx *argumentHelper,
	records []mutation.Record,
) (jsValue, error) {
	return cbCtx.ReturnWithJSValueErr(toSequenceMutationRecord(cbCtx.ScriptCtx(), records))
}

func toSequenceMutationRecord(
	ctx *V8ScriptContext,
	records []mutation.Record,
) (jsValue, error) {
	res := make([]*v8.Value, len(records))
	for i, r := range records {
		rec, err := ctx.createJSInstanceForObjectOfType("MutationRecord", &r)
		if err != nil {
			return nil, fmt.Errorf("v8host: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	recs, err := toArray(ctx.v8ctx, res...)
	return newV8Value(ctx.host.iso, recs), err
}
