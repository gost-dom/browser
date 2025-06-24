package htmltest

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

type HTMLElementHelper struct {
	html.HTMLElement
	HTMLParentNodeHelper
	t testing.TB
}

func NewHTMLElementHelper(t testing.TB, e html.HTMLElement) HTMLElementHelper {
	return HTMLElementHelper{e, NewHTMLParentNodeHelper(t, e), t}
}

func (h *HTMLElementHelper) AttributeValue(key string) string {
	h.t.Helper()
	v, found := h.GetAttribute(key)
	if !found {
		h.t.Errorf("Attribute '%s' not found on element '%v'", key, h.HTMLElement)
	}
	return v
}

func UnwrapHTMLElement[T html.HTMLElement](h HTMLElementHelper) (res T) {
	h.t.Helper()

	var ok bool
	if res, ok = h.HTMLElement.(T); !ok {
		h.t.Fatalf(
			"htmltest: UnwrapHTMLElement: element with tag name %s cannot be converted to type %T",
			h.HTMLElement.TagName(),
			res,
		)
	}
	return
}
