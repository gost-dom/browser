package internal

import (
	"regexp"
	"strings"
)

var matchKnownWord = regexp.MustCompile("(HTML|URL|DOM|SVG)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func TypeNameToFileName(name string) string {
	snake := matchKnownWord.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
