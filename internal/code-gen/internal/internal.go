package internal

import (
	"log/slog"
	"slices"
	"strings"
	"unicode"

	g "github.com/gost-dom/generators"
)

// A list of possible acronyms that an interface can start with. Used as a very
// simple way to generate sensible unexported names for wrapper types. E.g.,
// HTMLTemplateElement -> htmlTemplateElement.
var KnownAcronyms = []string{"HTML", "URL", "DOM", "XML", "ID", "JSON"}

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
		slog.Warn("Passing empty string to LowerCaseFirstLetter")
		return ""
	}
	buffer := make([]rune, 0, strLen)
	buffer = append(buffer, unicode.ToLower([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

type BoundFunction struct {
	fn     g.Value
	values []g.Generator
}

func BindValues(val g.Value, values ...g.Generator) BoundFunction {
	return BoundFunction{val, values}
}

func (f BoundFunction) Call(values ...g.Generator) g.Value {
	return f.fn.Call(append(f.values, values...)...)
}

func IdlNameToGoName(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		uppered := strings.ToUpper(word)
		if slices.Contains(KnownAcronyms, uppered) {
			return uppered
		} else {
			words[i] = UpperCaseFirstLetter(word)
		}
	}
	return strings.Join(words, "")
}
