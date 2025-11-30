package main

import (
	"fmt"
	"os"

	xhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func element(tag string, opts ...any) *xhtml.Node {
	res := &xhtml.Node{
		Type:     xhtml.ElementNode,
		DataAtom: atom.Lookup([]byte(tag)),
		Data:     tag,
	}
	for _, opt := range opts {
		switch t := opt.(type) {
		case string:
			res.AppendChild(&xhtml.Node{
				Type: xhtml.TextNode,
				Data: t,
			})
		case xhtml.Attribute:
			res.Attr = append(res.Attr, t)
		case *xhtml.Node:
			res.AppendChild(t)
		default:
			panic(fmt.Sprintf("invalid element option: %v", opt))
		}
	}
	return res
}

func save(body *xhtml.Node) (err error) {
	f, err := os.Create("report.html")
	if err == nil {
		doc := &xhtml.Node{Type: xhtml.DocumentNode}
		html := element("html")
		doc.AppendChild(html)
		html.AppendChild(body)
		err = xhtml.Render(f, doc)
	}
	return err
}
