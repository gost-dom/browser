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

	g.Expect(win.Eval("'bar' in target.dataset")).To(BeTrue(), "bar in dataset")
	g.Expect(win.Eval("'notThere' in target.dataset")).To(BeFalse(), "notThere in dataset")

	win.MustRun(`
		target.dataset.bar = "new bar";
		target.dataset.barBaz = "bar baz";
	`)

	target := win.HTMLDocument().GetHTMLElementById("target")
	v1, _ := target.GetAttribute("data-bar")
	v2, _ := target.GetAttribute("data-bar-baz")
	g.Expect(v1).To(Equal("new bar"), "Setting a new value")
	g.Expect(v2).To(Equal("bar baz"), "Setting a value with camelcased name")

	win.MustRun(`delete target.dataset.bar`)
	_, hasBar := target.GetAttribute("data-bar")
	g.Expect(hasBar).To(BeFalse())

}
