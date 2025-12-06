package fetch

import (
	"context"
	"fmt"
	"io"
	"iter"
	"net/http"
	"slices"
	"strings"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/url"
)

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

type Header struct {
	key types.ByteString
	val types.ByteString
}

type Headers struct{ headers []Header }

func parseHeaders(h http.Header) *Headers {
	res := &Headers{}
	for k, v := range h {
		for _, val := range v {
			res.Append(types.ByteString(k), types.ByteString(val))
		}
	}
	return res
}

func compareHeaders(a, b Header) int {
	fmt.Println("Compare", a.key, b.key)
	return strings.Compare(string(a.key), string(b.key))
}

func (h *Headers) Append(name, val types.ByteString) {
	// In order to have correct iteration behaviour, it's imperative that the
	// list is kept sorted.
	h.headers = insertSorted(h.headers, Header{key: name.ToLower(), val: val}, compareHeaders)
}

func (h *Headers) Delete(name types.ByteString) {
	name = name.ToLower()
	h.headers = slices.DeleteFunc(h.headers, func(h Header) bool { return h.key == name })
}

func (h *Headers) Get(name types.ByteString) (string, bool) {
	name = name.ToLower()
	var res []types.ByteString
	header := Header{key: name}
	if i, found := slices.BinarySearchFunc(h.headers, header, compareHeaders); found {
		l := len(h.headers)
		for {
			curr := h.headers[i]
			res = append(res, curr.val)
			i++
			if i >= l {
				break
			}
			if compareHeaders(header, h.headers[i]) < 0 {
				break
			}
		}
	}
	return h.formatValue(res), len(res) > 0
}

func (h *Headers) formatValue(val []types.ByteString) string {
	return strings.Join(types.ByteStringsToStrings(val), ", ")
}

func (h *Headers) Has(name types.ByteString) bool {
	name = name.ToLower()
	_, ok := h.Get(name)
	return ok
}

func (h *Headers) Set(name, value types.ByteString) {
	name = name.ToLower()
	idx := slices.IndexFunc(h.headers, func(h Header) bool { return h.key == name })
	if idx != -1 {
		h.headers[idx].val = value
	} else {
		h.Append(name, value)
	}
}

func (h *Headers) All() iter.Seq2[types.ByteString, types.ByteString] {
	return func(yield func(types.ByteString, types.ByteString) bool) {
		var collect []types.ByteString
		var i = 0
		for i < len(h.headers) {
			v := h.headers[i]
			if i > 1000 {
				return
			}
			var nextI = i + 1
			collect = append(collect, v.val)
			if v.key != "set-cookie" {
				if len(h.headers) > nextI {
					if h.headers[nextI].key == v.key {
						i++
						continue
					}
				}
			}
			if !yield(v.key, types.ByteString(h.formatValue(collect))) {
				return
			}
			i++
			collect = nil
		}
	}
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
	method string
	signal *dom.AbortSignal
	body   io.Reader
}

func (r *Request) URL() string { return url.ParseURLBase(r.url, r.bc.LocationHREF()).Href() }

func (r *Request) do(ctx context.Context) (*http.Response, error) {
	method := r.method
	if method == "" {
		method = "GET"
	}
	url := r.URL()
	r.bc.Logger().Info("gost-dom/fetch: Request.do", "method", method, "url", url)
	req, err := http.NewRequestWithContext(ctx, method, url, r.body)
	if err != nil {
		return nil, err
	}
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

func (f Fetch) Fetch(req Request) (*Response, error) {
	res := <-f.FetchAsync(req)
	return res.Value, res.Err
}

func (f Fetch) FetchAsync(req Request) promise.Promise[*Response] {
	ctx := f.BrowsingContext.Context()
	if req.signal != nil {
		ctx = dom.AbortContext(ctx, req.signal)
	}

	return promise.New(func() (*Response, error) {
		resp, err := req.do(ctx)
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
}

type Response struct {
	io.Reader
	Status  int
	Headers *Headers

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
