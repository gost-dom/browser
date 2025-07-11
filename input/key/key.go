package key

import (
	"iter"
	"time"
	"unicode"

	"github.com/gost-dom/browser/internal/uievents"
)

// Key represents a single keyboard input. WARNING: This is experimental.
//
// This is currently a simple abstraction, not taking the sequence of modifier
// keys into consideration. E.g., shift+A would result in keydown (shift),
// keydown (A with modifer shift: true), keyup (A), and keyup (shift).
//
// This general sequence of events is not yet properly representable in the
// types.
type Key struct {
	Key      string
	Letter   string
	Down, Up bool
	// delay is not used yet, but express n possible solution to simulating
	// delays betwen keystrokes, which can be valuable in order to verify
	// throttling/debounce behaviour.
	delay time.Duration
}

// EventInit creates a KeyboardEventInit representing the key stroke.
func (k Key) EventInit() uievents.KeyboardEventInit {
	return uievents.KeyboardEventInit{
		Key: k.Key,
	}
}

// RuneToKey returns a Key representing the keyboard key with the letter
// specified.
func RuneToKey(r rune) Key {
	return Key{Key: string(r), Letter: string(r), Down: true, Up: true}
}

// StringToKeys returns a sequence of [Key]
func StringToKeys(s string) iter.Seq[Key] {
	return func(yield func(Key) bool) {
		for _, r := range s {
			isUpper := unicode.IsUpper(r)
			if isUpper && !yield(Key{Key: "Shift", Down: true}) {
				return
			}
			if !yield(RuneToKey(r)) {
				return
			}
			if isUpper && !yield(Key{Key: "Shift", Up: true}) {
				return
			}
		}
	}
}
