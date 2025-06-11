package v8host_test

import (
	"testing"

	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestElementDataset(t *testing.T) {
	g := gomega.NewWithT(t)
	win := initWindow(t)

	g.Expect(win.Eval(`typeof DOMStringMap`)).To(Equal("function"))

	win.LoadHTML(
		`<div id="target" data-foo="Foo value" data-bar="Bar value" data-foo-bar="Foo bar value"></div>`,
	)
	win.MustEval(`const target = document.getElementById("target")`)
	g.Expect(win.Eval("target.dataset.foo")).To(Equal("Foo value"))
	g.Expect(
		win.Eval(`Object.getPrototypeOf(target.dataset).constructor.name`)).
		To(Equal("DOMStringMap"), "Dataset type")
	g.Expect(win.Eval(`Object.keys(target.dataset)`)).
		To(Equal([]any{"foo", "bar", "fooBar"}), "dataset keys")

}
