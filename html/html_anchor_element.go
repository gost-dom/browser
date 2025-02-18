package html

import (
	"github.com/gost-dom/browser/dom"
)

type htmlAnchorElement struct {
	HTMLElement
	URL dom.URL
}

func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{
		HTMLElement: NewHTMLElement("a", ownerDoc),
		URL:         nil,
	}
	result.SetSelf(result)
	return result
}

func (e *htmlAnchorElement) Click() bool {
	result := e.HTMLElement.Click()
	if href := e.Href(); result && href != "" {
		w := e.window()
		w.Navigate(w.resolveHref(href).Href())
	}
	return result
}

func (e *htmlAnchorElement) SetAttribute(name string, val string) {
	win := e.window().History().window
	e.HTMLElement.SetAttribute(name, val)
	if name == "href" {
		e.URL = dom.ParseURLBase(val, win.baseLocation)
	}
}

func (e *htmlAnchorElement) setUrl(f func(dom.URL, string), val string) {
	if e.URL == nil {
		return
	}
	f(e.URL, val)
	e.updateDataAttribute()
}

func (e htmlAnchorElement) updateDataAttribute() {
	e.HTMLElement.SetAttribute("href", e.URL.Href())

}

func (e *htmlAnchorElement) getUrl(f func(dom.URL) string) string {
	if e.URL == nil {
		return ""
	}
	return f(e.URL)
}

func (e *htmlAnchorElement) SetHref(val string) {
	window := e.getHTMLDocument().getWindow()
	newUrl := window.resolveHref(val)
	e.URL = newUrl
	e.updateDataAttribute()
}

func (e *htmlAnchorElement) String() string { return e.Href() }

func (e *htmlAnchorElement) SetProtocol(val string) { e.setUrl(dom.URL.SetProtocol, val) }
func (e *htmlAnchorElement) SetUsername(val string) { e.setUrl(dom.URL.SetUsername, val) }
func (e *htmlAnchorElement) SetPassword(val string) { e.setUrl(dom.URL.SetPassword, val) }
func (e *htmlAnchorElement) SetHost(val string)     { e.setUrl(dom.URL.SetHost, val) }
func (e *htmlAnchorElement) SetHostname(val string) { e.setUrl(dom.URL.SetHostname, val) }
func (e *htmlAnchorElement) SetPort(val string)     { e.setUrl(dom.URL.SetPort, val) }
func (e *htmlAnchorElement) SetPathname(val string) { e.setUrl(dom.URL.SetPathname, val) }
func (e *htmlAnchorElement) SetSearch(val string)   { e.setUrl(dom.URL.SetSearch, val) }
func (e *htmlAnchorElement) SetHash(val string)     { e.setUrl(dom.URL.SetHash, val) }

func (e *htmlAnchorElement) Href() string     { return e.getUrl(dom.URL.Href) }
func (e *htmlAnchorElement) Origin() string   { return e.getUrl(dom.URL.Origin) }
func (e *htmlAnchorElement) Protocol() string { return e.getUrl(dom.URL.Protocol) }
func (e *htmlAnchorElement) Username() string { return e.getUrl(dom.URL.Username) }
func (e *htmlAnchorElement) Password() string { return e.getUrl(dom.URL.Password) }
func (e *htmlAnchorElement) Host() string     { return e.getUrl(dom.URL.Host) }
func (e *htmlAnchorElement) Hostname() string { return e.getUrl(dom.URL.Hostname) }
func (e *htmlAnchorElement) Port() string     { return e.getUrl(dom.URL.Port) }
func (e *htmlAnchorElement) Pathname() string { return e.getUrl(dom.URL.Pathname) }
func (e *htmlAnchorElement) Search() string   { return e.getUrl(dom.URL.Search) }
func (e *htmlAnchorElement) Hash() string     { return e.getUrl(dom.URL.Hash) }

func (e *htmlAnchorElement) Target() string {
	result, _ := e.GetAttribute("target")
	return result
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.SetAttribute("target", val)
}
func (e *htmlAnchorElement) Download() string {
	result, _ := e.GetAttribute("download")
	return result
}

func (e *htmlAnchorElement) SetDownload(val string) {
	e.SetAttribute("download", val)
}
func (e *htmlAnchorElement) Ping() string {
	result, _ := e.GetAttribute("ping")
	return result
}

func (e *htmlAnchorElement) SetPing(val string) {
	e.SetAttribute("ping", val)
}
func (e *htmlAnchorElement) Rel() string {
	result, _ := e.GetAttribute("rel")
	return result
}

func (e *htmlAnchorElement) SetRel(val string) {
	e.SetAttribute("rel", val)
}
func (e *htmlAnchorElement) RelList() dom.DOMTokenList {
	return dom.NewDOMTokenList("rel", e)
}

func (e *htmlAnchorElement) Hreflang() string {
	result, _ := e.GetAttribute("hreflang")
	return result
}

func (e *htmlAnchorElement) SetHreflang(val string) {
	e.SetAttribute("hreflang", val)
}
func (e *htmlAnchorElement) Type() string {
	result, _ := e.GetAttribute("type")
	return result
}

func (e *htmlAnchorElement) SetType(val string) {
	e.SetAttribute("type", val)
}
func (e *htmlAnchorElement) Text() string {
	result, _ := e.GetAttribute("text")
	return result
}

func (e *htmlAnchorElement) SetText(val string) {
	e.SetAttribute("text", val)
}
func (e *htmlAnchorElement) ReferrerPolicy() string {
	result, _ := e.GetAttribute("referrerPolicy")
	return result
}

func (e *htmlAnchorElement) SetReferrerPolicy(val string) {
	e.SetAttribute("referrerPolicy", val)
}
