package fetch_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	g := gomega.NewWithT(t)
	rec := gosttest.NewHTTPRequestRecorder(t, gosttest.HttpHandlerMap{
		"/file.json": gosttest.StaticJSON(`{"foo": "Foo value", "bar": "Bar value"}`),
	})
	b := browsertest.InitBrowser(t, rec)
	w := htmltest.NewWindowHelper(t, b.NewWindow())
	g.Expect(w.Eval("typeof fetch")).To(Equal("function"))

	w.MustRun(`
		let result;
		fetch("http://example.com/file.json", { body: null }).
			then(response => response.json()).
			then(js => { result = js });
	`)
	err := w.Clock().ProcessEvents(t.Context())
	assert.NoError(t, err)
	assert.Equal(t, "Foo value", w.MustEval("result.foo"))
	assert.Equal(t, "Bar value", w.MustEval("result.bar"))
}
