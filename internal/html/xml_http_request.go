package html

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/url"
)

// TODO: Events for async
// abort
// error
// load
// loadend
// loadstart
// progress
// readystatechange
// timeout

type XHREvent = string

const (
	XHREventLoad      XHREvent = "load"
	XHREventLoadstart XHREvent = "loadstart"
	XHREventLoadend   XHREvent = "loadend"
)

type XmlHttpRequest interface {
	event.EventTarget
	Abort() error
	Open(string, string, ...RequestOption)
	Send() error
	SendBody(body io.Reader) error
	Status() int
	StatusText() string
	ResponseText() string
	SetRequestHeader(name string, value string)
	GetAllResponseHeaders() (res string, err error)
	OverrideMimeType(mimeType string) error
	GetResponseHeader(headerName string) *string
	SetWithCredentials(val bool) error
	WithCredentials() bool
	ResponseURL() string
	Response() string
	SetTimeout(int) error
	Timeout() int
}

type xmlHttpRequest struct {
	event.EventTarget
	logSource log.LogSource
	location  string
	client    http.Client
	async     bool
	status    int
	method    string
	url       string
	response  []byte
	res       *http.Response
	headers   http.Header
	clock     *clock.Clock
}

func NewXmlHttpRequest(ctx html.BrowsingContext, clock *clock.Clock) XmlHttpRequest {
	location := ctx.LocationHREF()
	result := &xmlHttpRequest{
		EventTarget: event.NewEventTarget(),
		logSource:   ctx,
		location:    location,
		client:      ctx.HTTPClient(),
		headers:     make(map[string][]string),
		clock:       clock,
		async:       true,
	}
	event.SetEventTargetSelf(result)
	log.Info(result.logSource.Logger(), "NewXmlHttpRequest", "location", location)
	return result
}
func (r *xmlHttpRequest) logger() log.Logger {
	return r.logSource.Logger()
}

type RequestOption = func(req *xmlHttpRequest)

func (req *xmlHttpRequest) Open(
	method string,
	// TODO: Should this be a `string` or a stringer? The JS object should accept
	// stringable objects, e.g., a URL, but should we convert here; or on the JS
	// binding layer? Or different methods?
	url string,
	options ...RequestOption) {
	log.Info(req.logger(), "XmlHttpRequest.Open", "method", method, "url", url)

	req.method = method
	req.url = url
	for _, o := range options {
		o(req)
	}
	// TODO: if (req.open) { req.Abort() }
}

func (req *xmlHttpRequest) send(body io.Reader) error {
	reqUrl := req.url
	if u := url.ParseURLBase(req.url, req.location); u != nil {
		reqUrl = u.Href()
	}
	log.Info(req.logger(), "XmlHttpRequest.send", "url", reqUrl)
	httpRequest, err := http.NewRequest(req.method, reqUrl, body)
	if err != nil {
		return err
	}
	httpRequest.Header = req.headers
	res, err := req.client.Do(httpRequest)
	if err != nil {
		return err
	}
	req.status = res.StatusCode
	req.res = res
	b := new(bytes.Buffer) // TODO, branch out depending on content-type
	_, err = b.ReadFrom(res.Body)
	req.response = b.Bytes()
	log.Debug(req.logger(), "Response received", "Status", res.StatusCode)
	req.DispatchEvent(&event.Event{Type: XHREventLoad})
	return err
}

func (req *xmlHttpRequest) Send() error {
	return req.SendBody(nil)
}

func (req *xmlHttpRequest) SendBody(body io.Reader) error {
	if body != nil {
		// TODO: Set content type or not?
		req.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	}
	if req.async {
		req.DispatchEvent(&event.Event{Type: XHREventLoadstart})
		req.clock.AddSafeTask(func() {
			req.send(body)
		}, 0)
		return nil
	}
	return req.send(body)
}

func (req *xmlHttpRequest) Status() int { return req.status }

// GetStatusText implements the [statusText] property
// [statusText]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/statusText
// TODO: Should this exist? It's just a wrapper around [http.GetStatusText], could
// be in JS wrapper layer
func (req *xmlHttpRequest) StatusText() string { return http.StatusText(req.status) }

func (req *xmlHttpRequest) ResponseURL() string { return req.url }

func (req *xmlHttpRequest) Response() string { return req.ResponseText() }

func (req *xmlHttpRequest) ResponseText() string { return string(req.response) }

func (req *xmlHttpRequest) SetRequestHeader(name string, value string) {
	req.headers.Add(name, value)
}

func (req *xmlHttpRequest) Abort() error {
	return errors.New("XmlHttpRequest.Abort called - not implemented - ignoring call")
}

func (req *xmlHttpRequest) GetAllResponseHeaders() (res string, err error) {
	if req.res == nil {
		return
	}
	builder := strings.Builder{}
	for k, vs := range req.res.Header {
		key := strings.ToLower(k)
		if key != "set-cookie" {
			for _, v := range vs {
				_, err = builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, v))
				if err != nil {
					return
				}
			}
		}
	}
	return builder.String(), nil
}

func (req *xmlHttpRequest) OverrideMimeType(mimeType string) error {
	// This has no effect at the moment, but has an empty implementation to be
	// compatible with HTMX.
	return nil
}

func (req *xmlHttpRequest) GetResponseHeader(headerName string) *string {
	if req.res == nil {
		return nil
	}
	key := http.CanonicalHeaderKey(headerName)
	if val, ok := req.res.Header[key]; ok && len(val) > 0 {
		res := new(string)
		*res = strings.Join(val, ", ")
		return res

	}
	return nil
}

func (req *xmlHttpRequest) SetWithCredentials(val bool) error {
	return nil
}

func (req *xmlHttpRequest) WithCredentials() bool {
	return false
}

func (req *xmlHttpRequest) SetTimeout(val int) error {
	return nil
}

func (req *xmlHttpRequest) Timeout() int {
	return 0
}

/* -------- Options -------- */

func RequestOptionAsync(
	val bool,
) RequestOption {
	return func(req *xmlHttpRequest) { req.async = val }
}

func RequestOptionUserName(_ string) RequestOption {
	return func(req *xmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}

func RequestOptionPassword(_ string) RequestOption {
	return func(req *xmlHttpRequest) {
		// TODO
		panic("Not implemented")
	}
}
