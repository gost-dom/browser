package dom

import "strings"

type attributeName struct {
	prefix    string
	localName string
}

func decodeAttrQualifiedName(name string) attributeName {
	if a, b, found := strings.Cut(name, ":"); found {
		return attributeName{a, b}
	} else {
		return attributeName{"", a}
	}
}
