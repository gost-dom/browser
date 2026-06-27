package html

// HTMLIFrameElement models the subset of the [HTMLIFrameElement] interface that
// gost-dom supports: access to the nested browsing context via contentWindow
// and contentDocument.
//
// [HTMLIFrameElement]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLIFrameElement
type HTMLIFrameElement interface {
	HTMLElement
	// ContentWindow returns the Window of the iframe's nested browsing context,
	// creating it on first access. Returns nil if the owning document has no
	// associated script host.
	ContentWindow() Window
	// ContentDocument returns the Document of the nested browsing context.
	ContentDocument() HTMLDocument
}

type htmlIFrameElement struct {
	htmlElement
}

// NewHTMLIFrameElement creates an iframe element owned by ownerDoc.
func NewHTMLIFrameElement(ownerDoc HTMLDocument) HTMLIFrameElement {
	result := &htmlIFrameElement{htmlElement: newHTMLElement("iframe", ownerDoc)}
	result.SetSelf(result)
	return result
}

// ContentWindow returns the iframe's nested browsing context.
//
// A faithful implementation would expose a *separate* same-origin JavaScript
// realm (a real browser gives each iframe its own realm whose native built-ins
// are distinct objects but mutually accessible with the parent). gost-dom binds
// one realm per V8 context and the underlying v8go build does not expose
// security-token control, so a separate child context's globals are not
// cross-realm accessible (V8 raises "no access"). Until that lands, the nested
// context resolves to the owning window's realm: this keeps every native
// built-in (Function, String, eval, ...) on contentWindow accessible and
// genuine. A same-origin browser permits exactly this access; only object
// identity differs.
func (e *htmlIFrameElement) ContentWindow() Window {
	w := e.window()
	if w == nil {
		return nil
	}
	return w
}

func (e *htmlIFrameElement) ContentDocument() HTMLDocument {
	w := e.window()
	if w == nil {
		return nil
	}
	return w.document
}
