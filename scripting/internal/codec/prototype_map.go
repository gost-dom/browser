package codec

import (
	"fmt"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/uievents"
)

func LookupJSPrototype(entity entity.ObjectIder) string {
	switch n := entity.(type) {
	case *event.Event:
		switch n.Data.(type) {
		case event.CustomEventInit:
			return "CustomEvent"
		case uievents.PointerEventInit:
			return "PointerEvent"
		case uievents.MouseEventInit:
			return "MouseEvent"
		case uievents.UIEventInit:
			return "UIEvent"
		default:
			return "Event"
		}
	case dom.Element:
		if constructor, ok := HtmlElements[strings.ToLower(n.TagName())]; ok {
			return constructor
		} else {
			return "Element"
		}
	case html.HTMLDocument:
		return "HTMLDocument"
	case dom.Document:
		return "Document"
	case dom.DocumentFragment:
		return "DocumentFragment"
	case dom.NamedNodeMap:
		return "NamedNodeMap"
	case dom.Attr:
		return "Attr"
	case dom.NodeList:
		return "NodeList"
	case dom.Text:
		return "Text"
	case dom.Node:
		return "Node"
	case *html.DOMStringMap:
		return "DOMStringMap"
	default:
		panic(fmt.Sprintf("Cannot lookup node: %v", n))
	}
}
