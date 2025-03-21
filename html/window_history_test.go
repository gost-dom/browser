package html_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

type WindowHistoryTestSuite struct {
	gosttest.GomegaSuite
	win html.Window
	h   *gosttest.EchoHandler
}

func (s *WindowHistoryTestSuite) History() *html.History {
	return s.win.History()
}

func (s *WindowHistoryTestSuite) SetupTest() {
	s.h = new(gosttest.EchoHandler)
	s.win = html.NewWindow(windowOptionHandler(s.h))
}

func (s *WindowHistoryTestSuite) TestWithNewWindow() {
	s.Expect(s.History().Length()).To(Equal(1), "History entries")
}

func (s *WindowHistoryTestSuite) TestHistoryAfterNavigating() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(2), "History length after navigating")
}

func (s *WindowHistoryTestSuite) TestReloadWithGoMethod() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.History().Length()).To(Equal(2), "History length after navigate")
	// about:blank wasn't a request
	s.Expect(s.h.RequestCount()).To(Equal(1), "No of HTTP requests")
	s.Expect(s.History().Go(0)).To(Succeed())
	s.Expect(s.History().Length()).To(Equal(2), "History length after reload")
	s.Expect(s.h.RequestCount()).To(Equal(2), "No of HTTP requests") // about:blank wasn't a request
}

func (s *WindowHistoryTestSuite) TestNavigate() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(2))
	s.Expect(s.win.Document()).To(HaveH1("/page-2"), "Document contents reflect page")
}

func (s *WindowHistoryTestSuite) TestGoBackOnce() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.History().Go(-1)).To(Succeed())
	s.Expect(s.win.Document()).To(HaveH1("Gost-DOM"), "Page content after pack")
	s.Expect(s.win.Location().Href()).To(Equal("about:blank"), "Location after back")
}

func (s *WindowHistoryTestSuite) TestGoForward() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.Navigate("/page-3")).To(Succeed())
	s.Expect(s.win.Navigate("/page-4")).To(Succeed())
	s.Expect(s.win.Navigate("/page-5")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(5), "History length after 4 navigates")
	s.Expect(s.win.History().Go(-3)).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(5), "History length after Go(-3)")
	s.Expect(s.win.Navigate("/page-6")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(3), "History length after Navigate")
	s.Expect(s.win.Location().Pathname()).To(Equal("/page-6"), "Location after Navigate")
}

func (s *WindowHistoryTestSuite) TestPopStateEvent() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.History().ReplaceState("page-2 state", "")).To(Succeed())
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-3")).To(Succeed())

	var actualEvent *event.Event
	s.win.AddEventListener(
		"popstate",
		event.NewEventHandlerFunc(func(e *event.Event) error {
			actualEvent = e
			return nil
		}),
	)
	s.Expect(s.win.History().Go(-1)).To(Succeed())

	s.Expect(actualEvent).ToNot(BeNil(), "Event was dispatched")
	popEvent, ok := actualEvent.Data.(PopStateEventInit)
	s.Expect(ok).To(BeTrue(), "Event is a popstateevent")
	s.Expect(popEvent.State).To(BeEquivalentTo("page-2 state"), "Event state")
}

func (s *WindowHistoryTestSuite) TestReplaceState() {
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(2))
	s.Expect(s.win.History().ReplaceState(EMPTY_STATE, "/page-3"))
	s.Expect(s.h.RequestCount()).To(Equal(1), "No of HTTP requests")
	s.Expect(s.win.History().Length()).To(Equal(2), "History length after ReplaceState")
	s.Expect(s.win.Location().Pathname()).To(Equal("/page-3"), "Path after ReplaceState")
}

func TestWindowHistory(t *testing.T) {
	suite.Run(t, new(WindowHistoryTestSuite))
}

type WindowHistoryPushStateTestSuite struct {
	gosttest.GomegaSuite
	win html.Window
	h   *gosttest.EchoHandler
}

func (s *WindowHistoryPushStateTestSuite) SetupTest() {
	s.h = new(gosttest.EchoHandler)
	s.win = html.NewWindow(windowOptionHandler(s.h))
	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.Navigate("/page-3")).To(Succeed())

	s.Expect(s.win.History().Length()).To(Equal(3))
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-4"))
}

func TestWindowHistoryPushState(t *testing.T) {
	suite.Run(t, new(WindowHistoryPushStateTestSuite))
}

func (s *WindowHistoryPushStateTestSuite) TestPushStateWithFragment() {
	eventDispatched := false
	s.win.AddEventListener(
		"hashchange",
		event.NewEventHandlerFunc(func(e *event.Event) error {
			eventDispatched = true
			return nil
		}),
	)
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-4#target"))
	s.Expect(eventDispatched).To(BeFalse(), "Fragment hashevent should not be emitted")
}

func (s *WindowHistoryPushStateTestSuite) TestPushStateWithRealPage() {
	s.Expect(s.win.History().Length()).To(Equal(4))
	s.Expect(s.win.History().Length()).To(Equal(4))
	s.Expect(s.win.Location().Pathname()).To(Equal("/page-4"))
	s.Expect(s.h.RequestCount()).To(Equal(2), "No of request _after_ pushState")

	s.Expect(s.win.History().Back()).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(4))
	s.Expect(s.win.Location().Pathname()).To(Equal("/page-3"))
	s.Expect(s.h.RequestCount()).To(Equal(2), "No of request _after_ back")
}

func (s *WindowHistoryPushStateTestSuite) TestPushStateWithEmptyURL() {
	s.T().Skip("TODO: Should add a new history entry")
}

type WindowHistoryPushStateWithMultipleEntriesTestSuite struct {
	gosttest.GomegaSuite
	win html.Window
	h   *gosttest.EchoHandler
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) SetupTest() {
	s.h = new(gosttest.EchoHandler)
	s.win = html.NewWindow(windowOptionHandler(s.h))

	s.Expect(s.win.Navigate("/page-2")).To(Succeed())
	s.Expect(s.win.Navigate("/page-3")).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(3))
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-4"))
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-5"))
	s.Expect(s.win.Navigate("/page-6")).To(Succeed())
	s.Expect(s.win.Navigate("/page-7")).To(Succeed())
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-8"))
	s.Expect(s.win.History().PushState(EMPTY_STATE, "/page-9"))
	s.Expect(s.win.History().Length()).To(Equal(9))
	s.Expect(s.h.RequestCount()).To(Equal(4))
}

func TestWindowHistoryPushStateWithMultipleEntries(t *testing.T) {
	suite.Run(t, new(WindowHistoryPushStateWithMultipleEntriesTestSuite))
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) TestGoBackTwo() {
	s.Expect(s.win.History().Go(-2)).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(9))
	s.Expect(s.h.RequestCount()).
		To(Equal(4), "No of HTTP requests, a new should _not_ have been made")
	s.Expect(s.win.Document().GetElementById("heading")).To(HaveTextContent("/page-7"))
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) TestGoBackThree() {
	s.Expect(s.win.History().Go(-3)).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(9))
	s.Expect(s.h.RequestCount()).To(Equal(5), "No of HTTP requests, a new _should_ have been made")
	s.Expect(s.win.Document().GetElementById("heading")).To(HaveTextContent("/page-6"))
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) TestGoBackFive() {
	s.Expect(s.win.History().Go(-5)).To(Succeed())
	s.Expect(s.win.Document().GetElementById("heading")).
		To(HaveTextContent("/page-4"), "Page for loaded")

	s.Expect(s.win.History().Length()).To(Equal(9), "History length - should be unchanged")
	s.Expect(s.h.RequestCount()).To(Equal(5), "No of HTTP requests - a new should have been made")
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) TestBackFiveAndForwardOne() {
	s.Expect(s.win.History().Go(-5)).To(Succeed())
	s.Expect(s.win.History().Go(1)).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(9))
	s.Expect(s.h.RequestCount()).
		To(Equal(5), "No of HTTP requests - a new should _not_ have been made")
}

func (s *WindowHistoryPushStateWithMultipleEntriesTestSuite) TestBackFiveAndForwardTwo() {
	s.Expect(s.win.History().Go(-5)).To(Succeed())
	s.Expect(s.win.History().Go(2)).To(Succeed())
	s.Expect(s.win.History().Length()).To(Equal(9))
	s.Expect(s.h.RequestCount()).To(Equal(6), "No of HTTP requests - a new should have been made")
	s.Expect(s.win.Document().GetElementById("heading")).To(HaveTextContent("/page-6"))
}
