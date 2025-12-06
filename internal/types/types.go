// Package types defines fundamental web types such as ByteString
package types

import (
	"strings"

	"github.com/gost-dom/browser/internal/gosterror"
)

// ByteString is a sequence of printable ASCII characters. This includes
// characters in the range 0x20-0x7E.
type ByteString string

func ToByteString(s string) (ByteString, error) {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < 0x20 || b > 0x7E {
			return "", gosterror.NewTypeError("cannot decode bytestring")
		}
	}
	return ByteString(s), nil
}

func ByteStringsToStrings(s []ByteString) []string {
	res := make([]string, len(s))
	for i, ss := range s {
		res[i] = string(ss)
	}
	return res
}

func (b ByteString) ToLower() ByteString { return ByteString(strings.ToLower(string(b))) }
