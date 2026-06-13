package html

import (
	"encoding/base64"
	"strings"

	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// installBase64Globals adds the WindowOrWorkerGlobalScope base64 utility methods
// atob and btoa to the global scope. V8 does not provide these (they are Web
// APIs, not part of ECMAScript), so they are implemented here in Go.
func installBase64Globals[T any](global js.Class[T]) {
	global.CreateOperation("atob", atob)
	global.CreateOperation("btoa", btoa)
}

// atob decodes a base64-encoded ASCII string into a binary string, where each
// character of the result corresponds to one decoded byte.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Window/atob
func atob[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	data, err := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	decoded, err := forgivingBase64Decode(data)
	if err != nil {
		return nil, cbCtx.NewTypeError(
			"Failed to execute 'atob': The string to be decoded is not correctly encoded.")
	}
	// Produce a binary string: each byte becomes a UTF-16 code unit in 0..255.
	// Building the Go string from runes 0..255 yields UTF-8 that V8 decodes back
	// to the same code units.
	var sb strings.Builder
	sb.Grow(len(decoded))
	for _, b := range decoded {
		sb.WriteRune(rune(b))
	}
	return cbCtx.NewString(sb.String()), nil
}

// btoa encodes a binary string (one byte per character, each in 0..255) into a
// base64-encoded ASCII string.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Window/btoa
func btoa[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	data, err := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, 0, len(data))
	for _, r := range data {
		if r > 0xFF {
			return nil, cbCtx.NewTypeError(
				"Failed to execute 'btoa': The string to be encoded contains characters " +
					"outside of the Latin1 range.")
		}
		bytes = append(bytes, byte(r))
	}
	return cbCtx.NewString(base64.StdEncoding.EncodeToString(bytes)), nil
}

// forgivingBase64Decode implements the WHATWG "forgiving-base64 decode"
// algorithm closely enough for practical use: ASCII whitespace is removed and
// both standard and raw (unpadded) base64 are accepted.
func forgivingBase64Decode(s string) ([]byte, error) {
	s = strings.Map(func(r rune) rune {
		switch r {
		case ' ', '\t', '\n', '\r', '\f':
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
