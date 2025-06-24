package fetch

import (
	"errors"
	"io"
	"net/http"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/dom"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/gost-dom/browser/url"
)

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

func New(bc html.BrowsingContext) Fetch { return Fetch{bc} }

func (f Fetch) NewRequest(url string) Request {
	return Request{
		url: url,
		bc:  f.BrowsingContext,
	}
}

type RequestOption func(*Request)

type Request struct {
	url string
	bc  html.BrowsingContext
}

func (r *Request) URL() string { return url.ParseURLBase(r.url, r.bc.LocationHREF()).Href() }

func (r *Request) do() (*http.Response, error) {
	r.bc.Logger().Info("Get", "url", r.URL())
	c := r.bc.HTTPClient()
	return c.Get(r.URL())
}

type fetchOption struct {
	signal dominterfaces.AbortSignal
}

type FetchOption func(*fetchOption)

func WithSignal(s dominterfaces.AbortSignal) FetchOption {
	return func(opt *fetchOption) { opt.signal = s }
}

func (f Fetch) Fetch(req Request, opts ...FetchOption) (*Response, error) {
	res := <-f.FetchAsync(req, opts...)
	return res.Value, res.Err
}

func (f Fetch) FetchAsync(req Request, opts ...FetchOption) Promise[*Response] {
	var opt fetchOption
	for _, o := range opts {
		o(&opt)
	}
	signal := make(chan struct{})

	f.BrowsingContext.Logger().Info("Signal?", "signal", opt.signal)
	if opt.signal != nil {
		handler := event.NewEventHandlerFunc(
			func(e *event.Event) error {
				f.BrowsingContext.Logger().Info("AbortSignal")
				go func() { signal <- struct{}{} }()
				return nil
			},
		)
		f.BrowsingContext.Logger().Info("Add event listener")
		opt.signal.AddEventListener(dom.EventTypeAbort, handler)
	}
	p := NewPromise[*Response]()
	p2 := NewPromise[*Response]()
	go func() {
		resp, err := req.do()
		if err != nil {
			p.Reject(err)
		} else {
			p.Resolve(&Response{
				Reader:       resp.Body,
				Status:       resp.StatusCode,
				httpResponse: resp,
			})
		}
	}()
	go func() {
		select {
		case res := <-p:
			p2.Send(res.Value, res.Err)
		case <-signal:
			p2.Reject(errors.New("Aborted"))
		}
	}()

	return p2
}

type Response struct {
	io.Reader
	Status int

	httpResponse *http.Response
}

type Result[T any] struct {
	Value T
	Err   error
}

type Promise[T any] chan Result[T]

func NewPromise[T any]() Promise[T]      { return make(Promise[T], 1) }
func (p Promise[T]) Close()              { close(p) }
func (p Promise[T]) Resolve(v T)         { p <- Result[T]{Value: v} }
func (p Promise[T]) Reject(err error)    { p <- Result[T]{Err: err} }
func (p Promise[T]) Send(v T, err error) { p <- Result[T]{Value: v, Err: err} }
