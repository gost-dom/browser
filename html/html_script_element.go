package html

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/log"
)

type htmlScriptElement struct {
	htmlElement
	script Script
	src    string
}

type HTMLScriptElement = HTMLElement

func NewHTMLScriptElement(ownerDocument HTMLDocument) HTMLElement {
	var result HTMLScriptElement = &htmlScriptElement{newHTMLElement("script", ownerDocument), nil, ""}
	result.SetSelf(result)
	return result
}

func (e *htmlScriptElement) Connected() {
	var err error
	src, hasSrc := e.GetAttribute("src")
	scriptType, _ := e.GetAttribute("type")
	window, _ := e.htmlDocument.getWindow().(*window)
	if !hasSrc {
		if e.script, err = window.scriptContext.Compile(e.TextContent()); err != nil {
			log.Error(e.Logger(), "HTMLScriptElement: compile error", "src", src, "err", err)
			return
		}
	} else {
		src = window.resolveHref(src).Href()
		if scriptType == "module" {
			if e.script, err = window.scriptContext.DownloadModule(src); err != nil {
				log.Error(e.Logger(), "HTMLScriptElement: download script error", "src", src, "err", err)
				return
			}
		} else {
			if e.script, err = window.scriptContext.DownloadScript(src); err != nil {
				log.Error(e.Logger(), "HTMLScriptElement: download script error", "src", src, "err", err)
				return
			}
		}
		if _, deferScript := e.GetAttribute("defer"); deferScript {
			window.deferScript(e)
			return
		}
	}
	e.run()
}

func (e *htmlScriptElement) run() {
	if err := e.script.Run(); err != nil {
		// TODO: Dispatch "error" event
		log.Error(e.Logger(), "Script error", "src", e.src, "err", err)
	}
}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
