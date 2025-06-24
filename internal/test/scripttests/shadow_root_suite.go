package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ShadowRootSuite struct {
	ScriptHostSuite
}

func NewShadowRootSuite(h html.ScriptHost) *ShadowRootSuite {
	return &ShadowRootSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ShadowRootSuite) TestInheritance() {
	s.Expect(s.Eval(
		`Object.getPrototypeOf(ShadowRoot.prototype).constructor.name`)).To(
		Equal("DocumentFragment"),
		"ShadowRoot inheritd from DocumentFragment",
	)
}

func (s *ShadowRootSuite) TestNodeType() {
	s.T().Skip("ShadowRoot is defined as a type, but there isn't a way to construct one yet")
	s.Expect(s.Eval(`new ShadowRoot().nodeType`)).
		To(BeEquivalentTo(11))

}
