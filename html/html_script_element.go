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
	e.logger().Info("<script> connected", "element", e, "src", e.src())
	var (
		err         error
		deferScript bool
	)
	window := e.htmlDocument.window()
	e.script, deferScript, err = e.compile()
	if err != nil {
		e.logger().Error("HTMLScriptElement: script compile error", log.ErrAttr(err))
		return
	}
	if deferScript {
		window.deferScript(e)
	} else {
		e.run()
	}
}

func (e *htmlScriptElement) compile() (script Script, deferred bool, err error) {
	src, hasSrc := e.GetAttribute("src")
	if !hasSrc {
		script, err := e.compileInline()
		return script, false, err
	} else {
		return e.downloadAndCompile(src)
	}
}

func (e *htmlScriptElement) downloadAndCompile(
	src string,
) (script Script, deferred bool, err error) {
	window := e.htmlDocument.window()
	src = window.resolveHref(src).Href()
	scriptType, _ := e.GetAttribute("type")
	if scriptType == "module" {
		script, err = window.scriptContext.DownloadModule(src)
		deferred = true
	} else {
		script, err = window.scriptContext.DownloadScript(src)
		_, deferred = e.GetAttribute("defer")
	}
	return
}

func (e *htmlScriptElement) compileInline() (Script, error) {
	window := e.htmlDocument.window()
	script, err := window.scriptContext.Compile(e.TextContent())
	if err != nil {
		e.logger().Warn("HTMLScriptElement: compile error",
			slog.String("script", e.TextContent()),
			log.ErrAttr(err))
	}
	return script, err
}

func (e *htmlScriptElement) run() {
	e.logger().Info("Run script", "src", e.src())
	if err := e.script.Run(); err != nil {
		e.logger().Error("Script error", "src", e.src(), log.ErrAttr(err))
	}
	e.logger().Debug("Script execution completed", "src", e.src())
}

func (e *htmlScriptElement) src() string {
	res, _ := e.GetAttribute("src")
	return res
}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
