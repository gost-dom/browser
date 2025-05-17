package htmltest

import (
	"strings"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

type BrowserHelper struct {
	*browser.Browser
	t testing.TB
}

func NewBrowserHelper(t testing.TB, b *browser.Browser) BrowserHelper {
	return BrowserHelper{b, t}
}

func (b BrowserHelper) OpenWindow(url string) WindowHelper {
	b.t.Helper()
	win, err := b.Open(url)
	if err != nil {
		b.t.Errorf("Error opening path '%s': %s", url, err.Error())
	}
	return NewWindowHelper(b.t, win)
}

// Helper type on top of html.HTMLDocument to provide useful helper functions
// for testing.
type WindowHelper struct {
	html.Window
	t testing.TB
}

func NewWindowHelper(t testing.TB, win html.Window) WindowHelper {
	return WindowHelper{win, t}
}

func (win WindowHelper) MustEval(script string) any {
	win.t.Helper()
	res, err := win.Eval(script)
	assert.NoError(win.t, err)
	return res
}

func (win WindowHelper) HTMLDocument() HTMLDocumentHelper {
	return NewHTMLDocumentHelper(win.t, win.Document())
}

func (h WindowHelper) MustLoadHTML(html string) {
	h.t.Helper()

	if err := h.Window.LoadHTML(html); err != nil {
		h.t.Errorf("Error loading HTML: %s\nHTML: %s", err.Error(), html)
	}
}

// Helper type on top of html.HTMLDocument to provide useful helper functions
// for testing.
type HTMLDocumentHelper struct {
	dom.Document
	t testing.TB
}

func NewHTMLDocumentHelper(t testing.TB, doc dom.Document) HTMLDocumentHelper {
	if doc == nil {
		doc = html.NewHTMLDocument(nil)
	}
	return HTMLDocumentHelper{doc, t}
}

func ParseHTMLDocumentHelper(t testing.TB, s string) (res HTMLDocumentHelper) {
	t.Helper()
	win, err := html.NewWindowReader(strings.NewReader(s))
	if err != nil {
		t.Errorf("Error parsing HTML")
		return
	}
	return NewWindowHelper(t, win).HTMLDocument()
}

// GetHTMLElementById works as [html/HTMLDocument.GetElementById] but assumes the
// found element to be an [html/HTMLElement]
func (h HTMLDocumentHelper) GetHTMLElementById(id string) html.HTMLElement {
	h.t.Helper()
	el := h.GetElementById(id)
	if el == nil {
		return nil
	}
	e, ok := el.(html.HTMLElement)
	if !ok {
		h.t.Errorf("Element with ID '%s' was expected to be an HTMLElement", id)
	}
	return e
}

func (h HTMLDocumentHelper) CreateHTMLElement(tag string) html.HTMLElement {
	h.t.Helper()
	el := h.CreateElement(tag)
	if el == nil {
		return nil
	}
	e, ok := el.(html.HTMLElement)
	if !ok {
		// This really shouldn't happen, elements created by an HTMLDocument
		// should all be HTMLElements. Right?
		h.t.Errorf("Element with tag '%s' was expected to be an HTMLElement", tag)
	}
	return e
}

func (h HTMLDocumentHelper) QuerySelectorHTML(pattern string) (res html.HTMLElement) {
	h.t.Helper()
	e, err := h.QuerySelector(pattern)
	if err != nil {
		h.t.Errorf("QuerySelector error. Pattern: '%s'. Error: %s", pattern, err.Error())
	}
	if e != nil {
		var ok bool
		if res, ok = e.(html.HTMLElement); !ok {
			h.t.Errorf("Element found by query '%s' was expected to be an HTMLElement", pattern)
		}
	}
	return
}
