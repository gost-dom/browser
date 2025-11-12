package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type DocumentFragmentSuite struct {
	ScriptHostSuite
}

func NewDocumentFragmentSuite(h html.ScriptEngine) *DocumentFragmentSuite {
	return &DocumentFragmentSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *DocumentFragmentSuite) TestInheritance() {
	s.Expect(
		s.Eval(`Object.getPrototypeOf(DocumentFragment.prototype) === Node.prototype`),
	).To(
		BeTrue(),
		"DocumentFragment extends Node",
	)

	s.Expect(s.Eval(`
		const fragment = document.createDocumentFragment();
		Object.getPrototypeOf(fragment) === DocumentFragment.prototype;
	`)).To(BeTrue(), "document.createDocumentFragment() returns a DocumentFragment")
}

func (s *DocumentFragmentSuite) TestMethods() {
	s.MustRunScript(`const fragment = document.createDocumentFragment()`)
	s.Expect(s.Eval(`typeof fragment.querySelector`)).To(Equal("function"))

}
