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
		decodeInto(scope, &init.CtrlKey, options, "ctrlKey", codec.DecodeBoolean),
		decodeInto(scope, &init.ShiftKey, options, "shiftKey", codec.DecodeBoolean),
		decodeInto(scope, &init.AltKey, options, "altKey", codec.DecodeBoolean),
		decodeInto(scope, &init.MetaKey, options, "metaKey", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierAltGraph, options, "modifierAltGraph", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierCapsLock, options, "modifierCapsLock", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierFn, options, "modifierFn", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierFnLock, options, "modifierFnLock", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierHyper, options, "modifierHyper", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierNumLock, options, "modifierNumLock", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierScrollLock, options, "modifierScrollLock", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierSuper, options, "modifierSuper", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierSymbol, options, "modifierSymbol", codec.DecodeBoolean),
		decodeInto(scope, &init.ModifierSymbolLock, options, "modifierSymbolLock", codec.DecodeBoolean),
	)
}

func decodeKeyboardEventInit[T any](scope js.Scope[T], options js.Object[T], init *uievents.KeyboardEventInit) error {
	return errors.Join(
		decodeEventModifierInit(scope, options, &init.EventModifierInit),
		decodeInto(scope, &init.Key, options, "key", codec.DecodeString),
		decodeInto(scope, &init.Code, options, "code", codec.DecodeString),
		decodeInto(scope, &init.Location, options, "location", codec.DecodeInt),
		decodeInto(scope, &init.Repeat, options, "repeat", codec.DecodeBoolean),
		decodeInto(scope, &init.IsComposing, options, "isComposing", codec.DecodeBoolean),
	)
}
