package html

import (
	"bytes"
	"io"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/log"
)

type htmlScriptElement struct{ htmlElement }

type HTMLScriptElement = HTMLElement

func NewHTMLScriptElement(ownerDocument HTMLDocument) HTMLElement {
	var result HTMLScriptElement = &htmlScriptElement{newHTMLElement("script", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlScriptElement) Connected() {
	var script = ""
	src, hasSrc := e.GetAttribute("src")
	if !hasSrc {
		script = e.TextContent()
	} else {
		window, _ := e.htmlDocument.getWindow().(*window)
		baseURL := window.baseLocation
		var absoluteURL string

		if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
			absoluteURL = src
		} else {
			if strings.HasPrefix(src, "/") {
				baseURLParsed, err := url.Parse(baseURL)
				if err != nil {
					panic(err)
				}
				schemeAndHost := fmt.Sprintf("%s://%s", baseURLParsed.Scheme, baseURLParsed.Host)
				absoluteURL = schemeAndHost + src
			} else {
				baseURLParsed, err := url.Parse(baseURL)
				if err != nil {
					panic(err)
				}
				relativeURLParsed, err := url.Parse(src)
				if err != nil {
					panic(err)
				}
				absoluteURL = baseURLParsed.ResolveReference(relativeURLParsed).String()
			}
		}

		resp, err := window.httpClient.Get(absoluteURL)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			body, _ := io.ReadAll(resp.Body)
			log.Error(e.logger(), "Error from server", "body", string(body), "src", absoluteURL)
			panic("Bad response")
		}

		buf := bytes.NewBuffer([]byte{})
		buf.ReadFrom(resp.Body)
		script = string(buf.Bytes())
	}
	if err := e.window().Run(script); err != nil {
		log.Error(e.Logger(), "Script error", "src", src)
	}

}

func (e *htmlScriptElement) AppendChild(n dom.Node) (dom.Node, error) {
	return e.htmlElement.AppendChild(n)
}
