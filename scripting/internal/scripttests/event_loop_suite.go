package scripttests

import (
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type EventLoopTestSuite struct {
	suite.Suite
	e html.ScriptEngine
	// ctx html.ScriptContext
	w htmltest.WindowHelper
}

func NewEventLoopTestSuite(e html.ScriptEngine) *EventLoopTestSuite {
	return &EventLoopTestSuite{e: e}
}

func (s *EventLoopTestSuite) SetupTest() {
	s.w = browsertest.InitWindow(s.T(), s.e, browsertest.WithIgnoreErrorLogs())
}

func (s *EventLoopTestSuite) TeardownTest() {
	s.w.Close()
}

func (s *EventLoopTestSuite) TestDeferExecution() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.w.Eval(`
				let val; 
				setTimeout(() => { val = 42 }, 1);
				val`,
		),
	).To(BeNil())
	s.w.Clock().Advance(time.Millisecond)
	Expect(s.w.Eval(`val`)).To(BeEquivalentTo(42))
}

func (s *EventLoopTestSuite) TestClearTimeout() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.w.Eval(`
				let val;
				let handle = setTimeout(() => { val = 42 }, 1);
				clearTimeout(handle);
				val`,
		),
	).To(BeNil())
	s.w.Clock().Advance(time.Millisecond)
	Expect(s.w.Eval(`val`)).To(BeNil())
}

func (s *EventLoopTestSuite) TestDispatchError() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.w.MustEval(`
				let val;
				window.addEventListener('error', () => {
					val = 42;
				});
				setTimeout(() => {
					throw new Error()
				}, 0);
				val`,
		),
	).To(BeNil())
	Expect(s.w.MustEval(`val`)).To(BeEquivalentTo(42))
}

func (s *EventLoopTestSuite) TestQueueMicrotask() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.w.Eval(`
				window.taskExecuted = false;
				window.queueMicrotask(() => {
					window.taskExecuted = true;
				});
				window.taskExecuted `,
		),
	).To(BeFalse())
	Expect(s.w.Eval(`window.taskExecuted`)).To(BeTrue())
}

func (s *EventLoopTestSuite) TestInterval() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(s.w.Eval(`
		let count = 0;
		const h = globalThis.setInterval(() => {
			count++;
		}, 100);
		count
	`)).To(BeEquivalentTo(0))
	s.w.Clock().Advance(100 * time.Millisecond)
	Expect(s.w.Eval("count")).To(BeEquivalentTo(1))
	s.w.Clock().Advance(300 * time.Millisecond)
	Expect(s.w.Eval("count")).To(BeEquivalentTo(4))
	s.w.Run("clearInterval(h)")
	s.w.Clock().Advance(300 * time.Millisecond)
	Expect(s.w.Eval("count")).To(BeEquivalentTo(4))
}

func (s *EventLoopTestSuite) TestGlobals() {
	// Methods should be in the window _instance_ not the Window prototype
	Expect := gomega.NewWithT(s.T()).Expect
	windowNames, err := s.w.Eval("Object.getOwnPropertyNames(window)")
	Expect(err).ToNot(HaveOccurred())
	Expect(windowNames).To(ContainElement("setTimeout"))
	Expect(windowNames).To(ContainElement("clearTimeout"))
	Expect(windowNames).To(ContainElement("setInterval"))
	Expect(windowNames).To(ContainElement("clearInterval"))
	Expect(windowNames).To(ContainElement("queueMicrotask"))

	prototypeNames, err := s.w.Eval("Object.getOwnPropertyNames(Window.prototype)")
	Expect(err).ToNot(HaveOccurred())
	Expect(prototypeNames).ToNot(ContainElement("setTimeout"))
	Expect(prototypeNames).ToNot(ContainElement("clearTimeout"))
	Expect(prototypeNames).ToNot(ContainElement("setInterval"))
	Expect(prototypeNames).ToNot(ContainElement("clearInterval"))
	Expect(prototypeNames).ToNot(ContainElement("queueMicrotask"))
}
