package model

import (
	"slices"
	"strings"

	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/webref/idl"
)

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

func IsDefined(t idl.Type) bool { return t.Name != "undefined" }
