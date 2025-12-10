package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type WindowTestSuite struct {
	ScriptHostSuite
}

func (s *WindowTestSuite) TestGlobalInstance() {
	s.Expect(s.Eval("globalThis === window")).To(BeTrue())
}

func (s *WindowTestSuite) TestWindowInheritance() {
	s.Expect(s.Eval("window instanceof EventTarget")).To(BeTrue())
	s.Expect(s.Eval("window instanceof Window")).To(BeTrue(), "window instanceof Window")
	s.Expect(s.Eval("Object.getPrototypeOf(window).constructor === Window")).To(BeTrue())
}

func (s *WindowTestSuite) TestWindowConstructor() {
	s.Expect(s.Eval("Window && typeof Window")).To(Equal("function"))
	s.Expect(s.RunScript("Window()")).ToNot(Succeed())
	s.MustRunScript(`
		let error;
		try { new Window() } catch(err) { 
			error = err;
		}
	`)
	s.Expect(s.Eval("error instanceof TypeError")).To(BeTrue())
	s.Expect(s.Eval("error && Object.getPrototypeOf(error).constructor.name")).
		To(Equal("TypeError"))
}

func (s *WindowTestSuite) TestDocumentProperty() {
	s.Expect(s.Eval("document instanceof Document")).To(BeTrue())
	s.Expect(s.Eval(`
		const keys = []
		for (let key in window) {
			keys.push(key);
		}
		keys
	`)).To(ContainElement("document"), "document is an enumerable property")
	s.Expect(s.Eval(
		`const a = window.document;
		const b = window.document;
		a === b`,
	)).To(BeTrue())
}

func (s *WindowTestSuite) TestConstructorName() {
	s.Expect(s.Eval("window.constructor.name")).To(Equal("Window"))
}

func NewWindowTestSuite(h html.ScriptEngine) *WindowTestSuite {
	return &WindowTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *WindowTestSuite) TestDOMContentLoaded() {
	s.Expect(s.Window.LoadHTML(`<body><script>
  scripts = []
  function listener1() {
    scripts.push("DOMContentLoaded")
  }
  function listener2() {
    scripts.push("load")
  }
  window.document.addEventListener("DOMContentLoaded", listener1);
  window.document.addEventListener("load", listener2);
</script></body>`)).To(Succeed())
	s.Expect(s.Eval("scripts.join(',')")).To(Equal("DOMContentLoaded,load"))
}

func (s *WindowTestSuite) TestLocation() {
	s.OpenWindow("http://location.example.com/foo", nil)
	s.Expect(s.Eval("location.href")).To(Equal("http://location.example.com/foo"))
}
