package htmltest

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

type HTMLElementHelper struct {
	html.HTMLElement
	t testing.TB
}

func NewHTMLElementHelper(t testing.TB, e html.HTMLElement) HTMLElementHelper {
	return HTMLElementHelper{e, t}
}

func (h *HTMLElementHelper) AttributeValue(key string) string {
	h.t.Helper()
	v, found := h.GetAttribute(key)
	if !found {
		h.t.Errorf("Attribute '%s' not found on element '%v'", key, h.HTMLElement)
	}
	return v
}
