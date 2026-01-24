package html

import (
	"log/slog"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/log"
)

type htmlScriptElement struct {
	htmlElement
	script Script
}

type HTMLScriptElement = HTMLElement

func NewHTMLScriptElement(ownerDocument HTMLDocument) HTMLElement {
	var result HTMLScriptElement = &htmlScriptElement{newHTMLElement("script", ownerDocument), nil}
	result.SetSelf(result)
	return result
}

func (e *htmlScriptElement) Connected() {
	l := e.logger()
	if src := e.src(); src != "" {
		l = l.With(slog.String("src", src))
	}
	l.Debug("<script> connected", "element", e)
	var (
		err         error
		deferScript bool
	)
	window := e.htmlDocument.window()
	e.script, deferScript, err = e.compile()
	if err != nil {
		l.Error("HTMLScriptElement: script error", log.ErrAttr(err))
		return
	}
	if deferScript {
		window.deferScript(e)
	} else {
		e.run(l)
	}
}

func (e *htmlScriptElement) compile() (script Script, deferred bool, err error) {
	src, hasSrc := e.GetAttribute("src")
	window := e.htmlDocument.window()
	if !hasSrc {
		script, err = window.scriptContext.Compile(e.TextContent())
		if err != nil {
			e.logger().Debug("HTMLScriptElement: compile error", "script", e.TextContent())
		}
	} else {
		src = window.resolveHref(src).Href()
		scriptType, _ := e.GetAttribute("type")
		if scriptType == "module" {
			script, err = window.scriptContext.DownloadModule(src)
			deferred = true
		} else {
			script, err = window.scriptContext.DownloadScript(src)
			_, deferred = e.GetAttribute("defer")
		}
	}
	return
}

func (e *htmlScriptElement) run(l *slog.Logger) {
	l.Info("Run")
	if err := e.script.Run(); err != nil {
		l.Error("Script error", "src", e.src, log.ErrAttr(err))
	}
	l.Debug("Script execution completed", "src", e.src)
}

func (e *htmlScriptElement) src() string {
	res, _ := e.GetAttribute("src")
	return res
}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
