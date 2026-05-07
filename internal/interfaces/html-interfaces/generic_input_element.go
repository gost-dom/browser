package htmlinterfaces

import "github.com/gost-dom/browser/dom"

// HTMLInputtableElement represents an HTMLElement that has a string "value" attribute
type HTMLInputtableElement interface {
	dom.Element
	Value() string
	SetValue(string)
}
