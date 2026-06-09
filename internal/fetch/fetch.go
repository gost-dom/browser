package fetch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/url"
)

var defaultDelay time.Duration

func DefaultDelay() time.Duration     { return defaultDelay }
func SetDefaultDelay(d time.Duration) { defaultDelay = d }

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

type Header struct {
	key types.ByteString
	val types.ByteString
}

func New(bc html.BrowsingContext) Fetch { return Fetch{bc} }

func (f Fetch) NewRequest(url string, opts ...RequestOption) Request {
	req := Request{
		url:     url,
		bc:      f.BrowsingContext,
		Headers: Headers{invalidHeaders: invalidRequestHeaders},
	}
	for _, o := range opts {
		o(&req)
	}
	return req
}

type RequestOption func(*Request)

type Request struct {
	Headers Headers
	url     string
	bc      html.BrowsingContext
	method  string
	signal  *dom.AbortSignal
	body    io.Reader
}

func (r *Request) URL() string { return url.ParseURLBase(r.url, r.bc.LocationHREF()).Href() }

func (r *Request) createHttpReq(ctx context.Context) (*http.Request, error) {
	method := r.method
	if method == "" {
		method = "GET"
	}
	url := r.URL()
	r.bc.Logger().Info("gost-dom/fetch: Request.do", "method", method, "url", url)
	return http.NewRequestWithContext(ctx, method, url, r.body)
}

func (r *Request) do(req *http.Request) (*http.Response, error) {
	c := r.bc.HTTPClient()
	return c.Do(req)
}

func WithSignal(s *dom.AbortSignal) RequestOption {
	return func(opt *Request) { opt.signal = s }
}

func WithMethod(m string) RequestOption {
	return func(opt *Request) { opt.method = m }
}

func WithBody(b io.Reader) RequestOption {
	return func(opt *Request) { opt.body = b }
}

func WithHeaders(h [][2]types.ByteString) RequestOption {
	return func(opt *Request) {
		for _, kv := range h {
			opt.Headers.Append(kv[0], kv[1])
		}
	}
}

func (f Fetch) Fetch(req Request) (*Response, error) {
	res := <-f.FetchAsync(req).C
	return res.Value, res.Err
}

// RoundtripOptions describes properties for an individual fetch request.
type RoundtripOptions struct {
	// Delay controls the Simulated Time passing from issuing a request to
	// receiving a response.
	Delay time.Duration
}

func defaultRoundtripOptions() RoundtripOptions {
	return RoundtripOptions{Delay: defaultDelay}
}

type RequestOptionFunc func(*http.Request, *RoundtripOptions)

func (f Fetch) FetchAsync(req Request) promise.Promise[*Response] {
	ctx := f.BrowsingContext.Context()
	if req.signal != nil {
		ctx = dom.AbortContext(ctx, req.signal)
	}

	httpReq, err := req.createHttpReq(ctx)
	optsFn, _ := entity.ComponentType[RequestOptionFunc](f.BrowsingContext)
	opts := defaultRoundtripOptions()
	if optsFn != nil {
		optsFn(httpReq, &opts)
	}
	p := promise.New(func() (*Response, error) {
		if err != nil {
			return nil, err
		}
		resp, err := req.do(httpReq)
		if err != nil {
			return nil, err
		}
		headers := parseHeaders(resp.Header)
		return &Response{
			Reader:       resp.Body,
			Status:       resp.StatusCode,
			httpResponse: resp,
			Headers:      headers,
		}, nil
	})
	p.Delay = opts.Delay
	return p
}

type Response struct {
	io.Reader
	Status  int
	Headers Headers

	httpResponse *http.Response
}

type ReadableStream struct {
	Reader io.Reader
}

func (s ReadableStream) GetReader(opts ...streams.GetReaderOption) streams.Reader {
	return &Reader{s.Reader, false}
}

type Reader struct {
	Reader io.Reader
	Done   bool
}

func (r *Reader) Read() promise.Promise[streams.ReadResult] {
	return promise.New(
		func() (streams.ReadResult, error) {
			if r.Done {
				return streams.ReadResult{Done: true}, nil
			}
			buf := make([]byte, 1024)
			l, err := r.Reader.Read(buf)
			buf = buf[0:l]
			if err == nil {
				return streams.ReadResult{Value: buf}, nil
			}
			if err == io.EOF {
				r.Done = true
				if l == 0 {
					return streams.ReadResult{Done: true}, nil
				} else {
					return streams.ReadResult{Value: buf}, nil
				}
			}
			return streams.ReadResult{}, err
		},
	)
}

func (r Response) Body() streams.ReadableStream { return ReadableStream{r.Reader} }

func assertHeaderCountWithinLimit(count int) {
	if count > MAX_HEADER_COUNT {
		msg := fmt.Sprintf(
			"gost-dom/fetch: exceeded header count limit during iteration: %d",
			MAX_HEADER_COUNT,
		)
		panic(msg)
	}
}

func init() {
	defaultDelay = 5 * time.Millisecond
}
