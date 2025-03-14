package html

import (
	"regexp"
	"strings"

	"github.com/gost-dom/browser/dom"
)

// DOMStringMap provides access to data-* attributes of an HTML or SVG element.
// In JavaScript, it is a dictionary-like object wrapping content with a "data-"
// prefix, converting kebab-case to camel-case, and stripping the prefix.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset
type DOMStringMap struct{ Element dom.Element }

var camelCaseDetector = regexp.MustCompile("[a-z][A-Z]")

// toKebab converts a camelCase string to a kebab-case string.
//
// This is intended for [HTMLElement.Dataset] that provides a camelCase API over
// the kebab-case data- content attribute names.
func toKebab(str string) string {
	return camelCaseDetector.ReplaceAllStringFunc(str, func(match string) string {
		lower := []rune(strings.ToLower(match))
		return string([]rune{lower[0], '-', lower[1]})
	})
}

func encodeDataAttrKey(key string) string {
	return "data-" + toKebab(key)
}

func (m DOMStringMap) Get(key string) (val string, exists bool) {
	return m.Element.GetAttribute(encodeDataAttrKey(key))
}

func (m DOMStringMap) Set(key string, val string) {
	m.Element.SetAttribute(encodeDataAttrKey(key), val)
}

func (m DOMStringMap) Delete(key string) {
	if attr := m.Element.GetAttributeNode(encodeDataAttrKey(key)); attr != nil {
		m.Element.RemoveAttributeNode(attr)
	}
}
