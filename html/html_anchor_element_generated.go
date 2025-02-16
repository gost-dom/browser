// This file is generated. Do not edit.

package html

import dom "github.com/gost-dom/browser/dom"

type HTMLAnchorElement interface {
	HTMLElement
	HTMLHyperlinkElementUtils
	Target() string
	SetTarget(string)
	Download() string
	SetDownload(string)
	Ping() string
	SetPing(string)
	Rel() string
	SetRel(string)
	RelList() dom.DOMTokenList
	Hreflang() string
	SetHreflang(string)
	Type() string
	SetType(string)
	Text() string
	SetText(string)
	ReferrerPolicy() string
	SetReferrerPolicy(string)
}
