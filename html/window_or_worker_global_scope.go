package html

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/gost-dom/browser/internal/clock"
)

type windowOrWorkerGlobalScope struct {
	clock *clock.Clock
}

// Btoa encodes the bytes of a binary string into a base64-encoded ASCII string.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/btoa
func (w windowOrWorkerGlobalScope) Btoa(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Atob decodes a base64-encoded ASCII string into its bytes, following the
// WHATWG "forgiving-base64 decode" algorithm.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/atob
func (w windowOrWorkerGlobalScope) Atob(data string) ([]byte, error) {
	return forgivingBase64Decode(data)
}

// forgivingBase64Decode implements the WHATWG "forgiving-base64 decode"
// algorithm closely enough for practical use: the four ASCII whitespace
// characters are removed and both padded and unpadded base64 are accepted.
//
// https://infra.spec.whatwg.org/#forgiving-base64-decode
func forgivingBase64Decode(s string) ([]byte, error) {
	s = strings.Map(func(r rune) rune {
		// "ASCII whitespace" per the Infra standard. unicode.IsSpace is
		// deliberately not used, as it would also strip non-ASCII whitespace
		// that the algorithm requires to be treated as invalid input.
		switch r {
		case '\t', '\n', '\f', '\r', ' ':
			return -1
		}
		return r
	}, s)
	if decoded, err := base64.StdEncoding.DecodeString(s); err == nil {
		return decoded, nil
	}
	// Fall back to unpadded decoding for inputs missing trailing '='.
	return base64.RawStdEncoding.DecodeString(strings.TrimRight(s, "="))
}

func (w windowOrWorkerGlobalScope) SetTimeout(
	task clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	return w.clock.SetTimeout(task, delay)
}

func (w windowOrWorkerGlobalScope) ClearTimeout(handle clock.TaskHandle) {
	w.clock.Cancel(handle)
}

func (w windowOrWorkerGlobalScope) SetInterval(
	task clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	return w.clock.SetInterval(task, delay)
}

func (w windowOrWorkerGlobalScope) ClearInterval(handle clock.TaskHandle) {
	w.clock.Cancel(handle)
}

func (w windowOrWorkerGlobalScope) QueueMicrotask(task clock.TaskCallback) {
	w.clock.QueueMicrotask(task)
}
