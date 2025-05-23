package v8host

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	log "github.com/gost-dom/browser/internal/log"
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
	ctx *V8ScriptContext,
	this *v8go.Object,
	cb mutation.Callback,
) (*v8go.Value, error) {
	return w.store(mutation.NewObserver(ctx.clock, cb), ctx, this)
}

func (w mutationObserverV8Wrapper) decodeMutationCallback(
	ctx *V8ScriptContext,
	val *v8go.Value,
) (res mutation.Callback, err error) {
	var f *v8go.Function
	if f, err = val.AsFunction(); err == nil {
		res = MutationCallback{ctx, f}
	}
	return
}

func (w mutationObserverV8Wrapper) decodeMutationObserverInit(
	ctx *V8ScriptContext,
	val *v8go.Value,
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
	ctx *V8ScriptContext,
	records []mutation.Record,
) (*v8go.Value, error) {
	return toSequenceMutationRecord(ctx, records)
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
		log.Info(ctx.window.Logger(), "Record", "val", rec)
		res[i] = rec
	}
	return toArray(ctx.v8ctx, res...)
}
