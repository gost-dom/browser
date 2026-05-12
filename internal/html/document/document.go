package document

import "github.com/gost-dom/browser/html"

type NewDocumentOption = func(html.HTMLDocument)

func WithTitle(t string) NewDocumentOption {
	return func(d html.HTMLDocument) {
		e := d.CreateElement("title")
		e.SetTextContent(t)
		d.Head().Append(e)
	}
}
