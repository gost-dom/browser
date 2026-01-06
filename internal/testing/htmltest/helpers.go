package htmltest

import (
	"strings"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/url"
	"github.com/stretchr/testify/assert"
)

func SetTestingT(w html.Window, val testing.TB) {
	entity.SetComponentType(w, val)
}

func GetTestingT(w html.Window) (testing.TB, bool) {
	return entity.ComponentType[testing.TB](w)
}

type BrowserHelper struct {
	*browser.Browser
	t testing.TB
}

func NewBrowserHelper(t testing.TB, b *browser.Browser) BrowserHelper {
	return BrowserHelper{b, t}
}

func (b BrowserHelper) NewWindow() WindowHelper {
	b.t.Helper()
	return NewWindowHelper(b.t, b.Browser.NewWindow())
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

func NewWindowHTML(t testing.TB, s string) WindowHelper {
	t.Helper()
	win, err := html.NewWindowReader(strings.NewReader(s), url.ParseURL(s))
	assert.NoError(t, err, "htmltest: NewWindowHTML")
	return NewWindowHelper(t, win)
}

func NewWindowHelper(t testing.TB, win html.Window) WindowHelper {
	h := WindowHelper{win, t}
	SetTestingT(win, t)
	return h
}

func (win WindowHelper) MustRun(script string) {
	win.t.Helper()
	err := win.Run(script)
	assert.NoError(win.t, err)
}

func (win WindowHelper) MustEval(script string) any {
	win.t.Helper()
	res, err := win.Eval(script)
	assert.NoError(win.t, err)
	return res
}

// Return a [JsAssert] value that can be used to verify the state of JavaScript
// values.
func (h WindowHelper) Assert() JsAssert { return JsAssert{h.t, h.Window} }

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
	HTMLParentNodeHelper
	html.HTMLDocument
	t testing.TB
}

func NewHTMLDocumentHelper(t testing.TB, doc html.HTMLDocument) HTMLDocumentHelper {
	if doc == nil {
		doc = html.NewHTMLDocument(nil)
	}
	return HTMLDocumentHelper{NewHTMLParentNodeHelper(t, doc), doc, t}
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

func (h HTMLDocumentHelper) CreateHTMLElement(tag string) HTMLElementHelper {
	h.t.Helper()
	el := h.CreateElement(tag)
	if el == nil {
		h.t.Fatalf("htmltest: invalid html element tag: %s", tag)
	}
	e, ok := el.(html.HTMLElement)
	if !ok {
		// This really shouldn't happen, elements created by an HTMLDocument
		// should all be HTMLElements. Right?
		h.t.Errorf("Element with tag '%s' was expected to be an HTMLElement", tag)
	}
	return NewHTMLElementHelper(h.t, e)
}

type HTMLParentNodeHelper struct {
	t             testing.TB
	ElementParent dom.ElementParent
}

func NewHTMLParentNodeHelper(t testing.TB, n dom.ElementParent) HTMLParentNodeHelper {
	return HTMLParentNodeHelper{t, n}
}

func (h HTMLParentNodeHelper) QuerySelectorHTML(pattern string) (res HTMLElementHelper) {
	h.t.Helper()
	e, err := h.ElementParent.QuerySelector(pattern)
	h.t.Logf("Element: %T", e)
	if err != nil {
		h.t.Errorf("QuerySelector error. Pattern: '%s'. Error: %s", pattern, err.Error())
	}
	if e != nil {
		if e, ok := e.(html.HTMLElement); !ok {
			h.t.Errorf("Element found by query '%s' was expected to be an HTMLElement", pattern)
		} else {
			res = NewHTMLElementHelper(h.t, e)
		}
	}
	return
}

func (h HTMLParentNodeHelper) QuerySelectorHTMLOpt(pattern string) (res *HTMLElementHelper) {
	h.t.Helper()
	e, err := h.ElementParent.QuerySelector(pattern)
	h.t.Logf("Element: %T", e)
	if err != nil {
		h.t.Errorf("QuerySelector error. Pattern: '%s'. Error: %s", pattern, err.Error())
	}
	if e == nil {
		return nil
	}
	if e, ok := e.(html.HTMLElement); !ok {
		h.t.Errorf("Element found by query '%s' was expected to be an HTMLElement", pattern)
	} else {
		res = new(HTMLElementHelper)
		*res = NewHTMLElementHelper(h.t, e)
	}
	return
}

func ParseHTMLDocument(t testing.TB, s string) HTMLDocumentHelper {
	win, err := html.NewWindowReader(strings.NewReader(s), url.ParseURL(s))
	if err != nil {
		t.Fatalf("gost-dom/htmltest: ParseHTMLDocument: %v", err)
	}
	return NewHTMLDocumentHelper(t, win.Document())
}
