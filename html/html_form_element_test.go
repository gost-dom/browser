package html_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/suite"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type HTMLFormElementTestSuite struct {
	gosttest.GomegaSuite
	win  html.Window
	doc  html.HTMLDocument
	form html.HTMLFormElement
}

func TestHTMLFormElement(t *testing.T) {
	suite.Run(t, new(HTMLFormElementTestSuite))
}

func (s *HTMLFormElementTestSuite) SetupTest() {
	s.win = NewWindowHelper(s.T(), NewWindow(WindowOptions{
		BaseLocation: "http://example.com/forms/example-form.html?original-query=original-value",
	}))

	s.doc = NewHTMLDocument(s.win)
	s.form = s.doc.CreateElement("form").(HTMLFormElement)
}

func (s *HTMLFormElementTestSuite) createForm() HTMLFormElement {
	return s.doc.CreateElement("form").(HTMLFormElement)
}

func (s *HTMLFormElementTestSuite) TestMethodIDLAttribute() {
	form := s.createForm()

	s.Assert().Equal("get", form.Method(), "Default value for HTMLFormElement.method")

	form.SetMethod("new value")
	s.Expect(form).To(HaveAttribute("method", "new value"),
		"Content attribute when setting IDL to invalid value")
	s.Expect(form.Method()).To(Equal("get"),
		"IDL method attribute when set to an invalid value")

	for _, value := range []string{"post", "POST", "PoSt"} {
		form.SetMethod(value)
		s.Expect(form.Method()).To(Equal("post"),
			fmt.Sprintf("IDL method attribute, set to %s", value),
		)
	}
}

func (s *HTMLFormElementTestSuite) TestActionIDLAttribute() {
	form := s.createForm()

	s.Assert().Equal(s.win.LocationHREF(), form.Action(), "Default value for action IDL attribute")

	form.SetAction("/foo-bar")
	s.Expect(form).
		To(HaveAttribute("action", "/foo-bar"), "Action content attribute with absolute path")
	s.Expect(form.Action()).
		To(Equal("http://example.com/foo-bar"), "Action IDL attribute when set to an absolute path")

	form.SetAttribute("action", "submit-target")
	s.Expect(form).To(
		HaveAttribute("action", "submit-target"),
		"Action content attribute with relative path")

	s.Expect(form.Action()).To(
		Equal("http://example.com/forms/submit-target"),
		"Action IDL attribute when set to a relative path")
}

var _ = Describe("HTML Form", func() {
	Describe("Submit behaviour", func() {
		var window WindowHelper
		var requests []*http.Request
		var form HTMLFormElement
		var actualRequest *http.Request
		var submittedForm url.Values

		AfterEach(func() {
			// Make the values ready for garbage collection
			actualRequest = nil
			submittedForm = nil
			form = nil
		})

		BeforeEach(func() {
			DeferCleanup(func() { requests = nil; submittedForm = nil })

			window = NewWindowHelper(GinkgoTB(), NewWindow(WindowOptions{
				HttpClient: gosthttp.NewHttpClientFromHandler(
					http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
						if req.ParseForm() != nil {
							panic("Error parsing form")
						}
						actualRequest = req
						submittedForm = req.Form
						requests = append(requests, req)
					}),
				),
				BaseLocation: "http://example.com/forms/example-form.html?original-query=original-value",
			}))

			Expect(
				window.LoadHTML(
					`<body>
					<form>
						<input name="foo" value="bar" />
					</form>
				</body>`,
				),
			).To(Succeed())

			el, err := window.Document().QuerySelector("form")
			Expect(err).ToNot(HaveOccurred())
			f, ok := el.(HTMLFormElement)
			Expect(ok).To(BeTrue())
			form = f
		})

		Describe("Method and action behaviour", func() {
			Describe("No method of action specified", func() {
				It("Should make a GET request to the original location", func() {
					form.Submit()
					Expect(actualRequest.Method).To(Equal("GET"))
					Expect(
						actualRequest.URL.String(),
					).To(Equal("http://example.com/forms/example-form.html?foo=bar"))
				})

				It("Should handle path lookup for relative paths", func() {
					form.SetAttribute("action", "submit-target")
					form.Submit()
					Expect(actualRequest.Method).To(Equal("GET"))
					Expect(
						actualRequest.URL.String(),
					).To(Equal("http://example.com/forms/submit-target?foo=bar"))
				})
			})

			Describe("The form is a POST", func() {
				BeforeEach(func() {
					form.SetAttribute("method", "POST")
				})

				It("Should make a POST request", func() {
					form.Submit()
					Expect(actualRequest.Method).To(Equal("POST"))
					Expect(
						actualRequest.URL.String(),
					).To(Equal("http://example.com/forms/example-form.html?original-query=original-value"))
				})

				It("Should store the values in the form body", func() {
					form.Submit()
					Expect(submittedForm["foo"]).To(Equal([]string{"bar"}))
				})

				It("Should resolve a relative 'action' without a ./.. prefix", func() {
					form.SetAttribute("action", "example-form-post-target")
					form.Submit()
					Expect(
						actualRequest.URL.String(),
					).To(Equal("http://example.com/forms/example-form-post-target"))
				})
			})
		})

		Describe("ReqeustSubmit with a <input type='submit'>", func() {
			var submitter dom.Element
			BeforeEach(func() {
				submitter = window.Document().CreateElement("input")
				submitter.SetAttribute("type", "submit")
				submitter.SetAttribute("name", "submitter")
				form.AppendChild(submitter)
			})

			It("Should add the name of a submitter, if passed", func() {
				form.RequestSubmit(submitter)
				Expect(submittedForm).To(HaveKey("submitter"))
			})

			It("Should ignore the name of a submitter if not passed", func() {
				form.RequestSubmit(nil)
				Expect(submittedForm).ToNot(HaveKey("submitter"))
			})
		})

		Describe("React to <button> click", func() {
			var button html.HTMLElement

			BeforeEach(func() {
				button = window.HTMLDocument().CreateHTMLElement("button")
				form.Append(button)
			})

			Describe("The button is a type='submit'", func() {
				BeforeEach(func() {
					button.SetAttribute("type", "submit")
				})

				It("Should submit the form", func() {
					button.Click()
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should also submit the form if 'type' was weird casing", func() {
					button.SetAttribute("type", "sUBmiT")
					button.Click()
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should not submit if preventDefault is called", func() {
					button.AddEventListener(
						"click",
						event.NewEventHandlerFunc(func(e *event.Event) error {
							e.PreventDefault()
							return nil
						}),
					)
					button.Click()
					Expect(actualRequest).To(BeNil())
				})

				It("Should include the button name in the form data if set", func() {
					button.SetAttribute("name", "the-button")
					button.Click()
					Expect(submittedForm).To(HaveKey("the-button"))
				})
			})

			Describe("The button is not type='submit'", func() {
				It("should not submit the form", func() {
					button.SetAttribute("type", "reset")
					button.Click()
					Expect(actualRequest).To(BeNil())
				})
			})
		})

		Describe("Click a <input type='submit'>", func() {
			It("Should submit the form", func() {
				input := window.Document().CreateElement("input").(HTMLInputElement)
				input.SetType("submit")
				form.Append(input)
				input.Click()
				Expect(actualRequest).ToNot(BeNil())
			})
		})

		Describe("Click a <input type='reset'>", func() {
			It("Should submit the form", func() {
				input := window.Document().CreateElement("input").(HTMLInputElement)
				input.SetType("reset")
				form.Append(input)
				input.Click()
				Expect(actualRequest).To(BeNil())
			})
		})

		Describe("Dispatched events", func() {
			Describe("Submit event", func() {
				It("Should not be dispatched when form.submit is called", func() {
					var actualEvent *event.Event
					form.AddEventListener(
						"submit",
						event.NewEventHandlerFunc(func(e *event.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.Submit()
					Expect(actualEvent).To(BeNil())
				})

				It("Should be dispatched when form.requestSubmit is called", func() {
					var actualEvent *event.Event
					form.AddEventListener(
						"submit",
						event.NewEventHandlerFunc(func(e *event.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.RequestSubmit(nil)
					Expect(actualEvent).ToNot(BeNil())
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should be abort the request on preventDefault()", func() {
					form.AddEventListener(
						"submit",
						event.NewEventHandlerFunc(func(e *event.Event) error {
							e.PreventDefault()
							return nil
						}),
					)
					form.RequestSubmit(nil)
					Expect(actualRequest).To(BeNil())
				})
			})

			Describe("formdata event", func() {
				It("Should be dispatched when a form is submitted", func() {
					var eventBubbles bool
					form.ParentElement().
						AddEventListener("formdata", eventtest.NewTestHandler(func(e *event.Event) {
							eventBubbles = true
						}))
					var actualEvent *event.Event
					form.AddEventListener(
						"formdata",
						event.NewEventHandlerFunc(func(e *event.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.Submit()
					Expect(actualEvent).ToNot(BeNil())
					Expect(eventBubbles).To(BeTrue())
					formDataEventInit, ok := actualEvent.Data.(FormDataEventInit)
					Expect(ok).To(BeTrue())
					Expect(formDataEventInit.FormData).ToNot(BeNil())
					Expect(formDataEventInit.FormData).To(HaveFormDataValue("foo", "bar"))
				})
			})
		})
	})

	Describe("Elements", func() {
		It("Should be a live collection", func() {
			Skip("TODO")
		})
		It("Should include all relevant element types", func() {
			Skip("TODO")
			/*
				https://developer.mozilla.org/en-US/docs/Web/API/HTMLFormElement/elements

				Value

				An HTMLFormControlsCollection containing all non-image controls in the form. This is a live collection; if form controls are added to or removed from the form, this collection will update to reflect the change.

				The form controls in the returned collection are in the same order in which they appear in the form by following a preorder, depth-first traversal of the tree. This is called tree order.

				Only the following elements are returned:

				    <button>
				    <fieldset>
				    <input> (with the exception that any whose type is "image" are omitted for historical reasons)
				    <object>
				    <output>
				    <select>
				    <textarea>
				    form-associated custom elements

			*/
		})
	})
})
