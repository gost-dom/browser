package fetch

import (
	"context"
	"io"
	"net/http"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/dom"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/url"
)

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

func New(bc html.BrowsingContext) Fetch { return Fetch{bc} }

func (f Fetch) NewRequest(url string, opts ...RequestOption) Request {
	req := Request{
		url: url,
		bc:  f.BrowsingContext,
	}
	for _, o := range opts {
		o(&req)
	}
	return req
}

type RequestOption func(*Request)

type Request struct {
	url    string
	bc     html.BrowsingContext
	signal dominterfaces.AbortSignal
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

func WithSignal(s dominterfaces.AbortSignal) RequestOption {
	return func(opt *Request) { opt.signal = s }
}

func (f Fetch) Fetch(req Request) (*Response, error) {
	// TODO: Get context from outside
	res := <-f.FetchAsync(context.Background(), req)
	return res.Value, res.Err
}

func (f Fetch) FetchAsync(
	ctx context.Context,
	req Request,
) promise.Promise[*Response] {
	// TODO: Get context from BrowsingContext
	if req.signal != nil {
		ctx = dom.AbortContext(ctx, req.signal)
	}

	return promise.New(func() (*Response, error) {
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
