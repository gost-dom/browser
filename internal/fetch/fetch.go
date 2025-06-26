package fetch

import (
	"context"
	"io"
	"net/http"

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

func (r *Request) do(ctx context.Context) (*http.Response, error) {
	method := "GET"
	url := r.URL()
	r.bc.Logger().Info("gost-dom/fetch: Request.do", "method", method, "url", url)
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
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
	// TODO: Get context from outside
	res := <-f.FetchAsync(context.Background(), req, opts...)
	return res.Value, res.Err
}

func (f Fetch) FetchAsync(
	ctx context.Context,
	req Request,
	opts ...FetchOption,
) Promise[*Response] {
	var opt fetchOption
	for _, o := range opts {
		o(&opt)
	}

	// TODO: Get context from BrowsingContext
	if opt.signal != nil {
		ctx = dom.AbortContext(ctx, opt.signal)
	}

	return NewPromiseFunc(func() (*Response, error) {
		resp, err := req.do(ctx)
		if err != nil {
			return nil, err
		}
		return &Response{
			Reader:       resp.Body,
			Status:       resp.StatusCode,
			httpResponse: resp,
		}, nil
	})
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

// NewPromiseFunc returns a promise that will settle with the outcome of running
// function f. It runs f in a separate gorouting, allowing early return.
func NewPromiseFunc[T any](f func() (T, error)) Promise[T] {
	p := NewPromise[T]()
	go func() { p.Send(f()) }()
	return p
}
