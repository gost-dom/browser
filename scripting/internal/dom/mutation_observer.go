package dom

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationCallback[T any] struct {
	ctx      js.CallbackContext[T]
	function js.Function[T]
}

func (cb MutationCallback[T]) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	v8Recs, _ := toSequenceMutationRecord(cb.ctx, recs)
	if _, err := cb.function.Call(cb.ctx.GlobalThis(), v8Recs); err != nil {
		js.UnhandledError(cb.ctx, err)
	}
}

func (w MutationObserver[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	cb mutation.Callback,
) (js.Value[T], error) {
	return codec.EncodeConstrucedValue(cbCtx, mutation.NewObserver(cbCtx.Clock(), cb))
}

func (w MutationObserver[T]) decodeMutationCallback(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback[T]{cbCtx, f}, nil
	}
	return nil, cbCtx.NewTypeError("Not a function")
}

func (w MutationObserver[T]) decodeObserveOption(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) ([]mutation.ObserveOption, error) {
	obj, ok := val.AsObject()
	if !ok {
		return nil, cbCtx.NewTypeError("Obtions not an object")
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

func (w MutationObserver[T]) toSequenceMutationRecord(
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
	prototype := cbCtx.Constructor("MutationRecord")
	for i, r := range records {
		rec, err := prototype.NewInstance(&r)
		if err != nil {
			return nil, fmt.Errorf("v8host: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	return cbCtx.NewArray(res...), nil
}
