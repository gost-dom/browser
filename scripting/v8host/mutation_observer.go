package v8host

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

type MutationCallback struct {
	ctx      *V8ScriptContext
	function *v8go.Function
}

func (cb MutationCallback) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	v8Recs, _ := toSequenceMutationRecord(cb.ctx, recs)

	cb.function.Call(cb.ctx.v8ctx.Global(), v8Recs)
}

func (w mutationObserverV8Wrapper) CreateInstance(
	cbCtx *argumentHelper,
	cb mutation.Callback,
) js.CallbackRVal {
	ctx := cbCtx.ScriptCtx()
	return cbCtx.ReturnWithValueErr(w.store(mutation.NewObserver(ctx.clock, cb), ctx, cbCtx.This()))
}

func (w mutationObserverV8Wrapper) decodeMutationCallback(
	cbCtx jsCallbackContext,
	val jsValue,
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback{cbCtx.ScriptCtx(), f.v8fn}, nil
	}
	return nil, v8go.NewTypeError(cbCtx.iso(), "Not a function")
}

func (w mutationObserverV8Wrapper) decodeObserveOption(
	cbCtx jsCallbackContext,
	val jsValue,
) ([]mutation.ObserveOption, error) {
	obj, err := val.AsObject()
	if err != nil {
		return nil, err
	}
	var res []mutation.ObserveOption
	ap := func(key string, o mutation.ObserveOption) {
		if err == nil {
			var v *v8go.Value
			if v, err = obj.Get(key); err == nil && v.Boolean() {
				res = append(res, o)
			}
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
) js.CallbackRVal {
	return cbCtx.ReturnWithValueErr(toSequenceMutationRecord(cbCtx.ScriptCtx(), records))
}

func toSequenceMutationRecord(
	ctx *V8ScriptContext,
	records []mutation.Record,
) (*v8go.Value, error) {
	res := make([]*v8.Value, len(records))
	for i, r := range records {
		rec, err := ctx.createJSInstanceForObjectOfType("MutationRecord", &r)
		if err != nil {
			return nil, fmt.Errorf("v8host: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	return toArray(ctx.v8ctx, res...)
}
