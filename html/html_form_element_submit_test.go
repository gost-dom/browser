package html_test

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/fixtures"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/gost-dom/fixture"
)

type HTTPHandlerFixture struct{ *http.ServeMux }

func (f *HTTPHandlerFixture) Setup() {
	if f.ServeMux == nil {
		f.ServeMux = http.NewServeMux()
	}
}

type WindowFixture struct {
	fixtures.AssertFixture
	*HTTPHandlerFixture

	Window       htmltest.WindowHelper
	BaseLocation string
}

func (f *WindowFixture) Setup() {
	if f.Window.Window != nil {
		return
	}
	win := html.NewWindow(html.WindowOptions{
		BaseLocation: f.BaseLocation,
		HttpClient:   gosthttp.NewHttpClientFromHandler(f.HTTPHandlerFixture),
	})
	f.Window = htmltest.NewWindowHelper(f.TB, win)

	f.Helper()
	f.Assert().NoError(f.Window.LoadHTML(
		`<body>
			<form>
				<input name="foo" value="bar" />
			</form>
		</body>`,
	))
}

type DefaultWindowFixture struct {
	WindowFixture
	*HTTPHandlerFixture
	initialized   bool
	actualRequest *http.Request
	submittedForm url.Values
	requests      []*http.Request
}

func (f *DefaultWindowFixture) Setup() {
	if !f.initialized {
		f.HTTPHandlerFixture.ServeMux.HandleFunc(
			"/",
			func(res http.ResponseWriter, req *http.Request) {
				if req.ParseForm() != nil {
					panic("Error parsing form")
				}
				f.actualRequest = req
				f.submittedForm = req.Form
				f.requests = append(f.requests, req)
			},
		)
		f.initialized = true
	}
}

func AssertType[T any](t testing.TB, actual any) (res T) {
	t.Helper()
	var ok bool
	if res, ok = actual.(T); !ok {
		t.Errorf(
			"Value is not expected type: Value - %v. Type: %s",
			actual, reflect.TypeFor[T]().Name(),
		)
	}
	return
}

func (f *WindowFixture) Form() html.HTMLFormElement {
	f.Helper()
	e, err := f.Window.HTMLDocument().QuerySelector("form")
	f.Assert().NoError(err)
	return AssertType[html.HTMLFormElement](f, e)
}

type HTMLFormFixture struct {
	WindowFixture
}

func TestSubmitForm(t *testing.T) {
	w, setup := fixture.Init(t, &DefaultWindowFixture{})
	w.BaseLocation = "http://example.com/forms/example-form.html?original-query=original-value"
	setup.Setup()

	var submitEventDispatched bool
	w.Form().AddEventListener("submit", eventtest.NewTestHandler(func(e *event.Event) {
		submitEventDispatched = true
	}))
	w.Form().Submit()

	w.Expect(w.actualRequest.Method).To(Equal("GET"))
	w.Expect(w.actualRequest.URL.String()).
		To(Equal("http://example.com/forms/example-form.html?foo=bar"))

	w.Assert().Equal("/forms/example-form.html", w.actualRequest.URL.Path,
		"Form post path")
	w.Assert().Equal("foo=bar", w.actualRequest.URL.RawQuery, "Request query")

	w.Assert().False(submitEventDispatched,
		"'submit' event should not be dispatched on form.submit()")
}

func TestHTMLFormElementSubmitPost(t *testing.T) {
	w, setup := fixture.Init(t, &DefaultWindowFixture{})
	w.BaseLocation = "http://example.com/forms/example-form.html?original-query=original-value"
	setup.Setup()

	form := w.Form()
	form.SetMethod("post")
	form.Submit()

	w.Expect(w.actualRequest.Method).To(Equal("POST"))
	w.Expect(w.actualRequest.URL.Path).To(Equal("/forms/example-form.html"), "Requested path")
	w.Expect(w.actualRequest.URL.RawQuery).
		To(Equal("original-query=original-value"), "Requested query")
	w.Expect(w.submittedForm["foo"]).To(Equal([]string{"bar"}))
}

type HTMLFormSubmitButtonFixture struct {
	DefaultWindowFixture
	Submitter html.HTMLButtonElement
}

func (f *HTMLFormSubmitButtonFixture) Setup() {
	f.Submitter = htmltest.UnwrapHTMLElement[html.HTMLButtonElement](
		f.Window.HTMLDocument().CreateHTMLElement("button"),
	)
	f.Submitter.SetAttribute("type", "submit")
	f.Submitter.SetAttribute("name", "submitter")
	f.Submitter.SetAttribute("value", "submitbtn")
	f.Form().AppendChild(f.Submitter)
}

func TestHTMLFormElementSubmitWithClickButton(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Submitter.Click()
	w.Assert().NotNil(w.submittedForm, "A form was submitted")
	w.Assert().
		Equal([]string{"submitbtn"}, w.submittedForm["submitter"], "Submit button added to the form")
	w.Assert().Equal([]string{"bar"}, w.submittedForm["foo"], "Input value from the form")
}

func TestHTMLFormElementSubmitWithClickButtonAndWeirdCasing(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Submitter.SetType("SuBMit")
	w.Submitter.Click()
	w.Assert().NotNil(w.submittedForm, "A form was submitted")
	w.Assert().
		Equal([]string{"submitbtn"}, w.submittedForm["submitter"], "Submit button added to the form")
	w.Assert().Equal([]string{"bar"}, w.submittedForm["foo"], "Input value from the form")
}

func TestHTMLFormElementSubmitWithClickResetButton(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Submitter.SetType("reset")
	w.Submitter.Click()
	w.Assert().Nil(w.submittedForm, "A form was submitted")
}

func TestHTMLFormElementRequestSubmitWithoutSubmitter(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Assert().NoError(w.Form().RequestSubmit(nil))
	w.Assert().NotNil(w.submittedForm, "A form was submitted")
	w.Assert().Nil(w.submittedForm["submitter"], "Form submitter value")
	w.Assert().Equal([]string{"bar"}, w.submittedForm["foo"], "Input value from the form")
}

func TestHTMLFormElementRequestSubmitPreventDefault(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	var submitEventDispatched bool
	w.Form().AddEventListener("submit", eventtest.NewTestHandler(func(e *event.Event) {
		submitEventDispatched = true
		e.PreventDefault()
	}))
	w.Form().RequestSubmit(nil)
	w.Assert().True(submitEventDispatched, "'submit' event dispatched")
	w.Assert().Nil(w.actualRequest, "A request was sent")
}

func TestHTMLFormElementRequestSubmitWithSubmitter(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Assert().NoError(w.Form().RequestSubmit(w.Submitter))
	w.Assert().NotNil(w.submittedForm, "A form was submitted")
	w.Assert().NotNil(w.submittedForm["submitter"], "Form submitter value")
	w.Assert().Equal([]string{"bar"}, w.submittedForm["foo"], "Input value from the form")
}

func TestHTMLFormElementPreventDefaultOnSubmitButton(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	w.Submitter.AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) {
		e.PreventDefault()
	}))
	w.Submitter.Click()
	w.Assert().Nil(w.actualRequest)
}

func TestHTMLFormElementFormDataEvent(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitButtonFixture{})
	setup.Setup()

	var eventBubbles bool
	w.Form().ParentElement().
		AddEventListener("formdata", eventtest.NewTestHandler(func(e *event.Event) {
			eventBubbles = true
		}))
	var actualEvent *event.Event
	w.Form().AddEventListener(
		"formdata",
		event.NewEventHandlerFunc(func(e *event.Event) error {
			actualEvent = e
			return nil
		}),
	)
	w.Form().Submit()
	w.Expect(actualEvent).ToNot(BeNil())
	w.Expect(eventBubbles).To(BeTrue())
	formDataEventInit, ok := actualEvent.Data.(html.FormDataEventInit)
	w.Expect(ok).To(BeTrue())
	w.Expect(formDataEventInit.FormData).ToNot(BeNil())
	w.Expect(formDataEventInit.FormData).To(HaveFormDataValue("foo", "bar"))
}

type HTMLFormSubmitInputFixture struct {
	DefaultWindowFixture
	Submitter html.HTMLInputElement
}

func (f *HTMLFormSubmitInputFixture) Setup() {
	f.Submitter = htmltest.UnwrapHTMLElement[html.HTMLInputElement](
		f.Window.HTMLDocument().CreateHTMLElement("input"),
	)
	f.Submitter.SetAttribute("type", "submit")
	f.Submitter.SetAttribute("name", "submitter")
	f.Submitter.SetAttribute("value", "submitinput")
	f.Form().AppendChild(f.Submitter)
}

func TestHTMLFormElementSubmitButtonWithClickButton(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitInputFixture{})
	setup.Setup()

	w.Submitter.Click()
	w.Assert().NotNil(w.submittedForm, "A form was submitted")
	w.Assert().
		Equal([]string{"submitinput"}, w.submittedForm["submitter"], "Submit input value added to the form")
	w.Assert().Equal([]string{"bar"}, w.submittedForm["foo"], "Input value from the form")
}

func TestHTMLFormElementSubmitInputWithClickResetButton(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitInputFixture{})
	setup.Setup()

	w.Submitter.SetType("reset")
	w.Submitter.Click()
	w.Assert().Nil(w.submittedForm, "A form was submitted")
}

func TestResubmitFormOn307Redirects(t *testing.T) {
	w, setup := fixture.Init(t, &HTMLFormSubmitInputFixture{})
	w.BaseLocation = "http://example.com/forms"
	setup.Setup()

	rec := gosttest.NewHTTPRequestRecorder(t, ParseFormHandler)
	w.Handle("POST /form-destination", http.RedirectHandler("/form-redirected", 307))
	w.Handle("POST /form-redirected", rec)

	form := w.Form()
	form.SetMethod("post")
	form.SetAction("/form-destination")
	form.Submit()

	w.Assert().Equal(1, len(rec.Requests), "Request sent to the redirected location")
	w.Assert().Equal([]string{"bar"}, rec.Single().PostForm["foo"])
}
