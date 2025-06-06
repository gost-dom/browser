package v8host

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationCallback[T any] struct {
	ctx      js.CallbackContext[T]
	function js.Function[T]
}

func (cb MutationCallback[T]) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	v8Recs, _ := toSequenceMutationRecord(cb.ctx, recs)
	scope := cb.ctx.Scope()
	if _, err := cb.function.Call(scope.GlobalThis(), v8Recs); err != nil {
		js.UnhandledError(scope, err)
	}
}

func (w mutationObserverV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	cb mutation.Callback,
) (js.Value[T], error) {
	return storeNewValue(mutation.NewObserver(cbCtx.Scope().Clock(), cb), cbCtx)
}

func (w mutationObserverV8Wrapper[T]) decodeMutationCallback(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback[T]{cbCtx, f}, nil
	}
	return nil, cbCtx.ValueFactory().NewTypeError("Not a function")
}

func (w mutationObserverV8Wrapper[T]) decodeObserveOption(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
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

func (w mutationObserverV8Wrapper[T]) toSequenceMutationRecord(
	cbCtx js.CallbackContext[T],
	records []mutation.Record,
) (js.Value[T], error) {
	return toSequenceMutationRecord(cbCtx, records)
}

func toSequenceMutationRecord[T any](
	cbCtx js.CallbackContext[T],
	records []mutation.Record,
) (js.Value[T], error) {
	res := make([]js.Value[T], len(records))
	prototype := cbCtx.Scope().Constructor("MutationRecord")
	for i, r := range records {
		rec, err := prototype.NewInstance(&r)
		if err != nil {
			return nil, fmt.Errorf("v8host: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	return cbCtx.ValueFactory().NewArray(res...), nil
}
