package dom

import (
	"errors"
	"fmt"

	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationCallback[T any] struct {
	s js.Scope[T]
	f js.Function[T]
}

func (cb MutationCallback[T]) HandleMutation(recs []mutation.Record, obs *mutation.Observer) {
	jsRecs, err := encodeSequenceMutationRecord(cb.s, recs)
	if err == nil {
		_, err = cb.f.Call(cb.s.GlobalThis(), jsRecs)
	}
	if err != nil {
		HandleJSCallbackError(cb.s, "HandleMutation", err)
	}
}

func CreateMutationObserver[T any](
	cbCtx js.CallbackContext[T],
	cb mutation.Callback,
) (js.Value[T], error) {
	return codec.EncodeConstructedValue(cbCtx, mutation.NewObserver(cbCtx.Clock(), cb))
}

func decodeMutationCallback[T any](s js.Scope[T], val js.Value[T],
) (mutation.Callback, error) {
	if f, ok := val.AsFunction(); ok {
		return MutationCallback[T]{s, f}, nil
	}
	return nil, s.NewTypeError("Not a function")
}

func decodeObserveOption[T any](
	s js.Scope[T],
	v js.Value[T],
) ([]mutation.ObserveOption, error) {
	obj, ok := v.AsObject()
	if !ok {
		return nil, s.NewTypeError("Options not an object")
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

func encodeSequenceMutationRecord[T any](
	s js.Scope[T], records []mutation.Record,
) (js.Value[T], error) {
	res := make([]js.Value[T], len(records))
	prototype := s.Constructor("MutationRecord")
	for i, r := range records {
		rec, err := prototype.NewInstance(&r)
		if err != nil {
			return nil, fmt.Errorf("gost-dom/scripting/dom: constructing mutation record: %w", err)
		}
		res[i] = rec
	}
	return s.NewArray(res...), nil
}
