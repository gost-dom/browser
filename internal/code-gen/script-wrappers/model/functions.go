package model

import (
	"strings"

	. "github.com/gost-dom/code-gen/internal"
)

func IdlNameToGoName(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = UpperCaseFirstLetter(word)
	}
	return strings.Join(words, "")
}
