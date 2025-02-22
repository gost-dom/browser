package v8host_test

import (
	"testing"
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type EventLoopTestSuite struct {
	suite.Suite
	host html.ScriptHost
	ctx  html.ScriptContext
}

func (s *EventLoopTestSuite) SetupTest() {
	w := html.NewWindow(html.WindowOptionHost(s.host))
	s.ctx = w.ScriptContext()
}

func (s *EventLoopTestSuite) TeardownTest() {
	s.ctx.Close()
}

func (s *EventLoopTestSuite) TestDeferExecution() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.ctx.Eval(`
				let val; 
				setTimeout(() => { val = 42 }, 1);
				val`,
		),
	).To(BeNil())
	s.ctx.Clock().Advance(time.Millisecond)
	Expect(s.ctx.Eval(`val`)).To(BeEquivalentTo(42))
}

func (s *EventLoopTestSuite) TestClearTimeout() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.ctx.Eval(`
				let val; 
				let handle =setTimeout(() => { val = 42 }, 1);
				clearTimeout(handle);
				val`,
		),
	).To(BeNil())
	s.ctx.Clock().Advance(time.Millisecond)
	Expect(s.ctx.Eval(`val`)).To(BeNil())
}

func (s *EventLoopTestSuite) TestDispatchError() {
	Expect := gomega.NewWithT(s.T()).Expect
	Expect(
		s.ctx.Eval(`
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
	Expect(s.ctx.Eval(`val`)).To(BeEquivalentTo(42))
}

func TestEventLoop(t *testing.T) {
	suite.Run(t, &EventLoopTestSuite{host: host})
}
