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
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/url"
)

// When iterating headers, have a max header count, to prevent overflow if
// client code continuously add new headers while iterating.
const MAX_HEADER_COUNT = 10000

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

type Header struct {
	key types.ByteString
	val types.ByteString
}

type Headers struct {
	entity.Entity
	headers []Header
}

func parseHeaders(h http.Header) Headers {
	res := Headers{}
	for k, v := range h {
		for _, val := range v {
			// Cast the string values to ByteString without validation. This
			// assumes HTTP headers received from a server are valid headers.
			res.Append(types.ByteString(k), types.ByteString(val))
		}
	}
	return res
}

func compareHeaders(a, b Header) int {
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

func (h *Headers) getRange(name types.ByteString) (first int, last int, found bool) {
	first = -1
	last = -1
	header := Header{key: name.ToLower()}
	first, found = slices.BinarySearchFunc(h.headers, header, compareHeaders)
	if !found {
		first = -1
		last = first
		return
	}
	l := len(h.headers)
	last = first
	for {
		last++
		if last >= l {
			return
		}
		if compareHeaders(header, h.headers[last]) < 0 {
			return
		}
	}
}

func (h *Headers) Get(name types.ByteString) (string, bool) {
	if h.headers == nil {
		return "", false
	}
	first, last, found := h.getRange(name)
	return h.formatValue(h.headers[first:last]), found
}

func (h *Headers) formatValue(val []Header) string {
	vv := make([]types.ByteString, len(val))
	for i, v := range val {
		vv[i] = v.val
	}
	return strings.Join(types.ByteStringsToStrings(vv), ", ")
}

func (h *Headers) Has(name types.ByteString) bool {
	_, ok := h.Get(name)
	return ok
}

func (h *Headers) Set(name, value types.ByteString) {
	name = name.ToLower()
	idx := slices.IndexFunc(h.headers, func(h Header) bool { return h.key == name })
	first, last, found := h.getRange(name)
	if found {
		h.headers = slices.Delete(h.headers, first+1, last)
		h.headers[idx].val = value
	} else {
		h.Append(name, value)
	}
}

// All() returns an iterator of key/value pairs of header keys and values.
// The same key may appear multiple times in the output. The returned order is
// guaranteed to be sorted by name. The iterator operates on a live list, i.e.
// new headers can be inserted while iterating. Panics if the number of headers
// iterated exceed MAX_HEADER_COUNT.
func (h *Headers) All() iter.Seq2[types.ByteString, types.ByteString] {
	return func(yield func(types.ByteString, types.ByteString) bool) {
		var collect []Header
		i := 0
		for i < len(h.headers) {
			assertHeaderCountWithinLimit(i)
			curr := h.headers[i]
			collect = append(collect, curr)
			if curr.key != "set-cookie" {
				var nextI = i + 1
				if len(h.headers) > nextI {
					if h.headers[nextI].key == curr.key {
						i++
						continue
					}
				}
			}
			if !yield(curr.key, types.ByteString(h.formatValue(collect))) {
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
	Headers Headers
	url     string
	bc      html.BrowsingContext
	method  string
	signal  *dom.AbortSignal
	body    io.Reader
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

func WithHeaders(h [][2]types.ByteString) RequestOption {
	return func(opt *Request) {
		for _, kv := range h {
			opt.Headers.Append(kv[0], kv[1])
		}
	}
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
