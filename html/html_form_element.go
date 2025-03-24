package html

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/url"
)

type FormEvent string

const (
	FormEventFormData FormEvent = "formdata"
	FormEventSubmit   FormEvent = "submit"
	FormEventReset    FormEvent = "reset"
)

type GetReader interface {
	GetReader() io.ReadCloser
}

type FormDataEventInit struct {
	FormData *FormData
}

type SubmitEventInit struct {
	Submitter dom.Element
}

// type FormDataEvent interface {
// 	dom.Event
// 	FormData() *FormData
// }

type FormSubmitEvent interface {
	event.Event
	Submitter() dom.Element
}

type formDataEvent struct {
	event.Event
	formData *FormData
}

func (e *formDataEvent) FormData() *FormData { return e.formData }

type formSubmitEvent struct {
	event.Event
	submitter dom.Element
}

func (e *formSubmitEvent) Submitter() dom.Element {
	return e.submitter
}

func newFormDataEvent(data *FormData) *event.Event {
	return &event.Event{
		Type:    string(FormEventFormData),
		Bubbles: true,
		Data: FormDataEventInit{
			FormData: data,
		}}
}

func newSubmitEvent(submitter dom.Element) *event.Event {
	eventInit := SubmitEventInit{submitter}
	return &event.Event{
		Type:       string(FormEventSubmit),
		Bubbles:    true,
		Cancelable: true,
		Data:       eventInit}
}

type HTMLFormElement interface {
	HTMLElement
	Action() string
	SetAction(val string)
	Method() string
	SetMethod(value string)
	Elements() dom.NodeList
	Submit() error
	RequestSubmit(submitter dom.Element) error
}

type htmlFormElement struct{ htmlElement }

func NewHtmlFormElement(ownerDocument HTMLDocument) HTMLFormElement {
	result := &htmlFormElement{
		newHTMLElement("form", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlFormElement) Submit() error {
	formData := NewFormDataForm(e)
	return e.submitFormData(formData)
}

func (e *htmlFormElement) Elements() dom.NodeList {
	inputs, err := e.QuerySelectorAll("input")
	if err == nil {
		return inputs
	}
	panic(err) // Should only be on invalid css pattern
}

func (e *htmlFormElement) submitFormData(formData *FormData) error {
	e.DispatchEvent(newFormDataEvent(formData))

	var (
		req *http.Request
		err error
	)
	if e.Method() == "get" {
		searchParams := formData.QueryString()
		targetURL := replaceSearchParams(e.action(), searchParams)
		req, err = http.NewRequest("GET", targetURL, nil)
	} else {
		req, err = http.NewRequest("POST", e.Action(), formData.GetReader())
		req.GetBody = func() (io.ReadCloser, error) { return formData.GetReader(), nil }
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if err != nil {
		return err
	}
	return e.htmlDocument.getWindow().fetchRequest(req)
}

func (e *htmlFormElement) RequestSubmit(submitter dom.Element) error {
	formData := NewFormDataForm(e)
	if submitter != nil {
		formData.AddElement(submitter)
	}
	if !e.DispatchEvent(newSubmitEvent(submitter)) {
		return nil
	}
	return e.submitFormData(formData)
}

func (e *htmlFormElement) Method() string {
	m, _ := e.GetAttribute("method")
	if strings.ToLower(m) == "post" {
		return "post"
	} else {
		return "get"
	}
}

func (e *htmlFormElement) SetAction(val string) { e.SetAttribute("action", val) }

func (e *htmlFormElement) action() *url.URL {
	window := e.window()
	action, found := e.GetAttribute("action")
	l := window.Location().(location)
	target := l.URL
	var err error
	if found {
		if target, err = url.NewUrlBase(action, window.Location().Href()); err != nil {
			// This _shouldn't_ happen. But let's refactor code, so err isn't a
			// possible return value
			panic(err)
		}
	}
	return target
}
func (e *htmlFormElement) Action() string {
	return e.action().Href()
}

func (e *htmlFormElement) SetMethod(value string) {
	e.SetAttribute("method", value)
}

func replaceSearchParams(location *url.URL, searchParams string) string {
	if searchParams == "" {
		return fmt.Sprintf("%s%s", location.Origin(), location.Pathname())
	} else {
		return fmt.Sprintf("%s%s?%s", location.Origin(), location.Pathname(), searchParams)
	}
}
