package fetch

import (
	"context"
	"errors"
	"fmt"
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
	var abortEvents <-chan *event.Event

	if opt.signal != nil {
		// // TODO: Get context from BrowsingContext
		// ctx, cancel := context.WithCancel(ctx)
		// defer cancel()
		abortEvents = event.NewEventSource(opt.signal).Listen(ctx, "abort", event.BufSize(1))
	}

	p2 := NewPromise[*Response]()
	go func() {
		p := NewPromise[*Response]()
		reqCtx, cancel := context.WithCancelCause(ctx)

		go func() {
			resp, err := req.do(reqCtx)
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
		if abortEvents == nil {
			p2 <- <-p
			return
		}

		select {
		case res := <-p:
			p2.Send(res.Value, res.Err)
		case e := <-abortEvents:
			err, ok := e.Data.(error)
			if !ok {
				err = ErrAbortSignal{Data: e.Data}
			}
			cancel(err)
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

type ErrAbortSignal struct{ Data any }

func (err ErrAbortSignal) Error() string {
	return fmt.Sprintf("aborted: readon: %v", err.Data)
}
