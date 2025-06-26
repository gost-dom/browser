package fetch

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
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

func (r *Request) do(ctx context.Context) (*http.Response, error) {
	r.bc.Logger().Info("Get", "url", r.URL())
	req, err := http.NewRequestWithContext(ctx, "GET", r.URL(), nil)
	if err != nil {
		return nil, err
	}
	c := r.bc.HTTPClient()
	return c.Do(req)
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
	var abortEvents <-chan *event.Event

	ctx := context.Background()
	f.BrowsingContext.Logger().Info("Signal?", "signal", opt.signal)
	if opt.signal != nil {
		// TODO: Get context from BrowsingContext
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		abortEvents = event.NewEventSource(opt.signal).Listen(ctx, "abort", event.BufSize(1))
	}

	p2 := NewPromise[*Response]()
	go func() {
		p := NewPromise[*Response]()
		reqCtx, cancel := context.WithCancel(ctx)

		go func() {
			f.BrowsingContext.Logger().Info("Send request")
			resp, err := req.do(reqCtx)
			f.BrowsingContext.Logger().Info("Got response")
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
		f.BrowsingContext.Logger().Info("Wait for response")
		if abortEvents == nil {
			p2 <- <-p
			return
		}
		select {
		case res := <-p:
			p2.Send(res.Value, res.Err)
		case <-abortEvents:
			f.BrowsingContext.Logger().Info("Abort event!")
			cancel()
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
