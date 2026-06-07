package key

import (
	"fmt"
	"iter"
	"time"
	"unicode"

	"github.com/gost-dom/browser/internal/uievents"
)

// Key represents a single keyboard input. WARNING: This is experimental.
//
// The fields Up and Down indicate if up/down events should be dispatched, which
// is used for simple modifier keys, e.g., the sequence Shift+A would generate
// the following sequence:
//
//   - Key: Shift, Down: true, Up: false
//   - Key: A, Down: true, Up: true
//   - Key: Shift, Down: false, Up: true
//
// Modifier states are not properly represented, the individual keyboard events
// do not contain modifier information, i.e., in the previous example, the event
// for the "A" key doesn't include the shiftKey yet.
//
// When simulating a string of keyboard events, KeydownDelay and KeyupDelay
// describes simulated delay after the keydown/keyup event before the next
// keyboard event in the sequence.
type Key struct {
	Key    string
	Letter string
	// KeydownDelay defines the simulated delay after the keydown event before
	// the next keyevent is dispatched.
	KeydownDelay time.Duration
	// KeyupDelay defines the simulated delay after the keyup event before the
	// next keyevent is dispatched.
	KeyupDelay time.Duration
	Down, Up   bool
}

type keyOption func(*Key)

func applyOptions(k Key, o []keyOption) Key {
	for _, oo := range o {
		oo(&k)
	}
	return k
}

func WithKeydownDelay(d time.Duration) keyOption {
	if d < 0 {
		panic(fmt.Sprintf("WithKeydownDelay: Negative timeout: %v", d))
	}
	return func(k *Key) {
		k.KeydownDelay = d
	}
}

func WithKeyupDelay(d time.Duration) keyOption {
	if d < 0 {
		panic(fmt.Sprintf("WithKeyupDelay: Negative timeout: %v", d))
	}
	return func(k *Key) {
		k.KeyupDelay = d
	}
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

// StringToKeys returns a sequence of [Key] given an input string s containing
// text. If the string contains upper-case letters, the sequence will include a
// keydown event for the shift key, the keydown/keyup events for the actual key,
// and finally the keyup event for the shoft key.
func StringToKeys(s string, o ...keyOption) iter.Seq[Key] {
	return func(yield func(Key) bool) {
		for _, r := range s {
			isUpper := unicode.IsUpper(r)
			if isUpper && !yield(applyOptions(Key{Key: "Shift", Down: true}, o)) {
				return
			}
			if !yield(applyOptions(RuneToKey(r), o)) {
				return
			}
			if isUpper && !yield(applyOptions(Key{Key: "Shift", Up: true}, o)) {
				return
			}
		}
	}
}
