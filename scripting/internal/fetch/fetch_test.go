package fetch_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
)

func TestFetch(t *testing.T) {
	g := gomega.NewWithT(t)
	b := browsertest.InitBrowser(t, gosttest.HttpHandlerMap{
		"/file.json": gosttest.StaticJSON(`{"foo": "Foo value", "bar": "Bar value"}`),
	})
	w := htmltest.NewWindowHelper(t, b.NewWindow())
	g.Expect(w.Eval("typeof fetch")).To(Equal("function"))
}
