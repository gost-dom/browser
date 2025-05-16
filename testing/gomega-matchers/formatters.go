package matchers

import (
	"fmt"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/onsi/gomega/format"
)

func FormatElement(value any) (result string, ok bool) {
	var element dom.Element
	if element, ok = value.(dom.Element); ok {
		result = element.OuterHTML()
	}
	return
}

func formatNode(value any) (result string, ok bool) {
	var node dom.Node
	if node, ok = value.(dom.Node); ok {
		result = node.NodeName()
	}
	return
}

func formatNodeList(value any) (result string, ok bool) {
	var list dom.NodeList
	if list, ok = value.(dom.NodeList); ok {
		b := strings.Builder{}
		b.WriteString("NodeList{")
		for i, n := range list.All() {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("%v", n))
		}
		b.WriteString("}")

		result = b.String()
	}
	return
}

func init() {
	format.RegisterCustomFormatter(FormatElement)
	format.RegisterCustomFormatter(formatNodeList)
	format.RegisterCustomFormatter(formatNode)
}
