package key

import (
	"iter"
	"time"
)

// Key represents a single keyboard input.
//
// This is currently a simple abstraction, not taking the sequence of modifier
// keys into consideration. E.g., shift+A would result in keydown (shift),
// keydown (A with modifer shift: true), keyup (A), and keyup (shift).
//
// This general sequence of events is not yet properly representable in the
// types.
type Key struct {
	// Delay is not used yet. It merely shows an possible solution to simulating
	// delays betwen keystrokes, which can be valuable in order to verify
	// throttling/debounce behaviour.
	Delay  time.Duration
	Letter string
}

// RuneToKey returns a Key representing the keyboard key with the letter
// specified.
func RuneToKey(r rune) Key {
	return Key{Letter: string(r)}
}

// StringToKeys returns a sequence of [Key]
func StringToKeys(s string) iter.Seq[Key] {
	return func(yield func(Key) bool) {
		for _, r := range s {
			if !yield(RuneToKey(r)) {
				return
			}
		}
	}
}
