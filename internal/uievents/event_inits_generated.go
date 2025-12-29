// This file is generated. Do not edit.

package uievents

type EventModifierInit struct {
	UIEventInit
	CtrlKey            bool
	ShiftKey           bool
	AltKey             bool
	MetaKey            bool
	ModifierAltGraph   bool
	ModifierCapsLock   bool
	ModifierFn         bool
	ModifierFnLock     bool
	ModifierHyper      bool
	ModifierNumLock    bool
	ModifierScrollLock bool
	ModifierSuper      bool
	ModifierSymbol     bool
	ModifierSymbolLock bool
}

type KeyboardEventInit struct {
	EventModifierInit
	Key         string
	Code        string
	Location    int
	Repeat      bool
	IsComposing bool
}
