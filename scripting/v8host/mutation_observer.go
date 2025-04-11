package v8host

import (
	mutation "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/v8go"
)

func (w mutationObserverV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8go.Object,
	cb mutation.Callback,
) (*v8go.Value, error) {
	return nil, nil
}

func (w mutationObserverV8Wrapper) decodeMutationCallback(
	ctx *V8ScriptContext,
	val *v8go.Value,
) (mutation.Callback, error) {
	panic("TODO")
}

func (w mutationObserverV8Wrapper) decodeMutationObserverInit(
	ctx *V8ScriptContext,
	val *v8go.Value,
) ([]mutation.ObserveOption, error) {
	panic("TODO")
}

func (w mutationObserverV8Wrapper) toSequenceMutationRecord(
	ctx *V8ScriptContext,
	records []mutation.Record,
) (*v8go.Value, error) {
	panic("TODO")
}
