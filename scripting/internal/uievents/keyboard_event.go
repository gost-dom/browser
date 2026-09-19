package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func KeyboardEvent_getModifierState[T any](
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	key, err := js.ConsumeArgument(cbCtx, "key", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}

	init, err := codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}

	return codec.EncodeBoolean(cbCtx, modifierState(init.EventModifierInit, key))
}

func modifierState(init uievents.EventModifierInit, key string) bool {
	switch key {
	case "Alt":
		return init.AltKey
	case "AltGraph":
		return init.ModifierAltGraph
	case "CapsLock":
		return init.ModifierCapsLock
	case "Control":
		return init.CtrlKey
	case "Fn":
		return init.ModifierFn
	case "FnLock":
		return init.ModifierFnLock
	case "Hyper":
		return init.ModifierHyper
	case "Meta":
		return init.MetaKey
	case "NumLock":
		return init.ModifierNumLock
	case "ScrollLock":
		return init.ModifierScrollLock
	case "Shift":
		return init.ShiftKey
	case "Super":
		return init.ModifierSuper
	case "Symbol":
		return init.ModifierSymbol
	case "SymbolLock":
		return init.ModifierSymbolLock
	case "Accel":
		return init.CtrlKey || init.MetaKey
	default:
		return false
	}
}
