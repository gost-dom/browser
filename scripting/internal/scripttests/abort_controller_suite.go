package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type AbortControllerSuite struct {
	ScriptHostSuite
}

func NewAbortControllerSuite(h html.ScriptEngine) *AbortControllerSuite {
	return &AbortControllerSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *AbortControllerSuite) TestGlobals() {
	s.Expect(s.Eval("typeof AbortController")).To(Equal("function"))
	s.Expect(s.Eval("typeof AbortSignal")).To(Equal("function"))
}

func (s *AbortControllerSuite) TestAbortSignal() {
	s.Expect(s.Eval(`typeof (new AbortController())`)).To(Equal("object"))
	s.Expect(s.Eval(
		`Object.getPrototypeOf(new AbortController()) === globalThis.AbortController.prototype`,
	)).To(BeTrue())

	s.Expect(s.Eval(`
		const ctrl = new AbortController()
		const sig = ctrl.signal
		Object.getPrototypeOf(sig).constructor.name
	`)).To(Equal("AbortSignal"))
}
