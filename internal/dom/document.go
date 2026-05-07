package dom

import "github.com/gost-dom/browser/internal/entity"

type domKey string

const isHTMLDocumentKey domKey = "isHtmlDocument"

func SetIsHTMLDocument(doc entity.Components, v bool) {
	entity.SetComponent(doc, isHTMLDocumentKey, v)
}

func IsHTMLDocument(doc entity.Components) bool {
	v, _ := entity.Component[bool](doc, isHTMLDocumentKey)
	return v
}
