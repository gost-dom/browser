package mutation

import dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"

type Options = dominterfaces.MutationObserverInit

func Subtree(o *Options)               { o.Subtree = true }
func ChildList(o *Options)             { o.ChildList = true }
func Attributes(o *Options)            { o.Attributes = true }
func CharacterData(o *Options)         { o.CharacterData = true }
func AttributeOldValue(o *Options)     { o.AttributeOldValue = true }
func CharacterDataOldValue(o *Options) { o.CharacterDataOldValue = true }

func AttributeFilter(val []string) func(*Options) {
	return func(o *Options) { o.AttributeFilter = val }
}
