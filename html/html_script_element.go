package html

import (
	"bytes"
	"io"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/log"
)

type htmlScriptElement struct {
	htmlElement
	script string
	src    string
}

type HTMLScriptElement = HTMLElement

func NewHTMLScriptElement(ownerDocument HTMLDocument) HTMLElement {
	var result HTMLScriptElement = &htmlScriptElement{newHTMLElement("script", ownerDocument), "", ""}
	result.SetSelf(result)
	return result
}

func (e *htmlScriptElement) Connected() {
	var hasSrc bool
	e.src, hasSrc = e.GetAttribute("src")
	if !hasSrc {
		e.script = e.TextContent()
	} else {
		window, _ := e.htmlDocument.getWindow().(*window)
		e.src = window.resolveHref(e.src).Href()
		resp, err := window.httpClient.Get(e.src)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			log.Error(e.logger(), "Error from server", "body", string(body), "src", e.src)
			panic("Bad response")
		}

		buf := bytes.NewBuffer([]byte{})
		buf.ReadFrom(resp.Body)
		e.script = string(buf.Bytes())

		if _, deferScript := e.GetAttribute("defer"); deferScript {
			window.deferScript(e)
			return
		}
	}

	e.run()
}

func (e *htmlScriptElement) run() {
	if err := e.window().Run(e.script); err != nil {
		// TODO: Dispatch "error" event
		log.Error(e.Logger(), "Script error", "src", e.src, log.ErrAttr(err))
	}
}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
