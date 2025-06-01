package v8host

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
)

type MutationCallback struct {
	ctx      jsCallbackContext
	function jsFunction
}

func (cb MutationCallback) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	v8Recs, _ := toSequenceMutationRecord(cb.ctx, recs)
	scope := cb.ctx.Scope()
	if _, err := cb.function.Call(scope.GlobalThis(), v8Recs); err != nil {
		UnhandledError(scope, err)
	}
}

func (w mutationObserverV8Wrapper) CreateInstance(
	cbCtx jsCallbackContext,
	cb mutation.Callback,
) (jsValue, error) {
	return w.store(mutation.NewObserver(cbCtx.Scope().Clock(), cb), cbCtx)
}

func (w mutationObserverV8Wrapper) decodeMutationCallback(
	cbCtx jsCallbackContext,
	val jsValue,
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback{cbCtx, f}, nil
	}
	return nil, cbCtx.ValueFactory().NewTypeError("Not a function")
}

func (w mutationObserverV8Wrapper) decodeObserveOption(
	cbCtx jsCallbackContext,
	val jsValue,
) ([]mutation.ObserveOption, error) {
	obj, ok := val.AsObject()
	if !ok {
		return nil, cbCtx.ValueFactory().NewTypeError("Obtions not an object")
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
	cbCtx jsCallbackContext,
	records []mutation.Record,
) (jsValue, error) {
	return toSequenceMutationRecord(cbCtx, records)
}

func toSequenceMutationRecord(
	cbCtx jsCallbackContext,
	records []mutation.Record,
) (jsValue, error) {
	res := make([]jsValue, len(records))
	prototype := cbCtx.Scope().Constructor("MutationRecord")
	for i, r := range records {
		rec, err := prototype.NewInstance(cbCtx.Scope(), &r)
		if err != nil {
			return nil, fmt.Errorf("v8host: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	return cbCtx.ValueFactory().NewArray(res...), nil
}
