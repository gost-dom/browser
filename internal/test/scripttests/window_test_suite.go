package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type WindowTestSuite struct {
	ScriptHostSuite
}

func (s *WindowTestSuite) TestGlobalInstance() {
	s.Expect(s.eval("globalThis === window")).To(BeTrue())
}

func (s *WindowTestSuite) TestWindowInheritance() {
	s.Expect(s.eval("window instanceof EventTarget")).To(BeTrue())
	s.Expect(s.eval("Object.getPrototypeOf(window).constructor === Window")).To(BeTrue())
}

func (s *WindowTestSuite) TestWindowConstructor() {
	s.Expect(s.eval("Window && typeof Window")).To(Equal("function"))
	s.Expect(s.run("Window()")).ToNot(Succeed())
	s.Expect(s.eval(
		`let error;
		try { new Window() } catch(err) { 
			error = err;
		}
		error && Object.getPrototypeOf(error).constructor.name
	`)).To(Equal("TypeError"))
}

func (s *WindowTestSuite) TestDocumentProperty() {
	s.Expect(s.eval("document instanceof Document")).To(BeTrue())
	s.Expect(s.eval(`
		const keys = []
		for (let key in window) {
			keys.push(key);
		}
		keys
	`)).To(ContainElement("document"), "document is an enumerable property")
	s.Expect(s.eval(
		`const a = window.document;
		const b = window.document;
		a === b`,
	)).To(BeTrue())
}

func (s *WindowTestSuite) TestConstructorName() {
	s.T().Skip("The Window function doesn't have a name. We might need to add setName to v8go")
	s.Expect(s.eval("window.constructor.name")).To(Equal("Window"))
}

func NewWindowTestSuite(h html.ScriptHost) *WindowTestSuite {
	return &WindowTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *WindowTestSuite) TestDOMContentLoaded() {
	s.Expect(s.window.LoadHTML(`<body><script>
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
	s.Expect(s.eval("scripts.join(',')")).To(Equal("DOMContentLoaded,load"))
}
