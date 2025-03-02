package internal

import (
	"log/slog"
	"strings"
	"unicode"
)

// A list of possible acronyms that an interface can start with. Used as a very
// simple way to generate sensible unexported names for wrapper types. E.g.,
// HTMLTemplateElement -> htmlTemplateElement.
var KnownAcronyms = []string{"HTML", "URL", "DOM", "XML"}

func UpperCaseFirstLetter(s string) string {
	strLen := len(s)
	if strLen == 0 {
		slog.Warn("Passing empty string to UpperCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToUpper([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

func LowerCaseFirstLetter(s string) string {
	for _, acro := range KnownAcronyms {
		if strings.HasPrefix(s, acro) {
			return strings.ToLower(acro) + s[len(acro):]
		}
	}
	strLen := len(s)
	if strLen == 0 {
		slog.Warn("Passing empty string to lowerCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToLower([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}
