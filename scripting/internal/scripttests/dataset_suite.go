package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type DatasetSuite struct {
	ScriptHostSuite
}

func NewDatasetSuite(h html.ScriptEngine) *DatasetSuite {
	return &DatasetSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *DatasetSuite) TestIsAFunction() {
	s.Expect(s.Eval(`typeof DOMStringMap`)).To(Equal("function"))

	s.Expect(s.Eval(
		"Object.getPrototypeOf(document.body.dataset) === DOMStringMap.prototype",
	)).To(BeTrue(), "dataset is a DOMStringMap")
}

func (s *DatasetSuite) TestDataset() {
	win := s.Window

	win.LoadHTML(
		`<div id="target" data-foo="Foo value" data-bar="Bar value" data-foo-bar="Foo bar value"></div>`,
	)
	win.MustEval(`const target = document.getElementById("target")`)
	s.Expect(win.Eval("target.dataset.foo")).To(Equal("Foo value"))
	s.Expect(
		win.Eval(`Object.getPrototypeOf(target.dataset).constructor.name`)).
		To(Equal("DOMStringMap"), "Dataset type")

	// This test expect a specific ordering of keys, which Object.keys doesn't
	// guarantee. However, dataset keys _should_ be iterated in the order they
	// appear.
	s.Expect(win.Eval(`Object.keys(target.dataset)`)).
		To(Equal([]any{"foo", "bar", "fooBar"}), "dataset keys")

	s.Expect(win.Eval("'bar' in target.dataset")).To(BeTrue(), "bar in dataset")
	s.Expect(win.Eval("'notThere' in target.dataset")).To(BeFalse(), "notThere in dataset")

	win.MustRun(`
		target.dataset.bar = "new bar";
		target.dataset.barBaz = "bar baz";
	`)

	target := win.HTMLDocument().GetHTMLElementById("target")
	v1, _ := target.GetAttribute("data-bar")
	v2, _ := target.GetAttribute("data-bar-baz")
	s.Expect(v1).To(Equal("new bar"), "Setting a new value")
	s.Expect(v2).To(Equal("bar baz"), "Setting a value with camelcased name")

	win.MustRun(`delete target.dataset.bar`)
	_, hasBar := target.GetAttribute("data-bar")
	s.Expect(hasBar).To(BeFalse())
}
