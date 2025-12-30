// This file is generated. Do not edit.

package uievents

import (
	"errors"
	uievents "github.com/gost-dom/browser/internal/uievents"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodeEventModifierInit[T any](scope js.Scope[T], options js.Object[T], init *uievents.EventModifierInit) error {
	return errors.Join(
		decodeUIEventInit(scope, options, &init.UIEventInit),
		js.DecodeInto(scope, &init.CtrlKey, options, "ctrlKey", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ShiftKey, options, "shiftKey", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.AltKey, options, "altKey", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.MetaKey, options, "metaKey", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierAltGraph, options, "modifierAltGraph", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierCapsLock, options, "modifierCapsLock", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierFn, options, "modifierFn", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierFnLock, options, "modifierFnLock", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierHyper, options, "modifierHyper", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierNumLock, options, "modifierNumLock", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierScrollLock, options, "modifierScrollLock", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierSuper, options, "modifierSuper", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierSymbol, options, "modifierSymbol", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.ModifierSymbolLock, options, "modifierSymbolLock", codec.DecodeBoolean),
	)
}

func decodeKeyboardEventInit[T any](scope js.Scope[T], options js.Object[T], init *uievents.KeyboardEventInit) error {
	return errors.Join(
		decodeEventModifierInit(scope, options, &init.EventModifierInit),
		js.DecodeInto(scope, &init.Key, options, "key", codec.DecodeString),
		js.DecodeInto(scope, &init.Code, options, "code", codec.DecodeString),
		js.DecodeInto(scope, &init.Location, options, "location", codec.DecodeInt),
		js.DecodeInto(scope, &init.Repeat, options, "repeat", codec.DecodeBoolean),
		js.DecodeInto(scope, &init.IsComposing, options, "isComposing", codec.DecodeBoolean),
	)
}
