package gosthttp

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
)

// A TestRoundTripper is an implementation of the [http.RoundTripper] interface
// that communicates directly with an [http.Handler] instance.
type TestRoundTripper struct{ http.Handler }

func (h TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// You could possibly test on req.Host and apply different behaviour, e.g.
	// forwarding to external site, or have mocked external sites, such as IDPs
	body := req.Body
	if body == nil {
		body = nullReader{}
	}
	serverReq, err := http.NewRequestWithContext(req.Context(), req.Method, req.URL.String(), body)
	if err != nil {
		return nil, err
	}
	serverReq.Header = req.Header
	serverReq.Trailer = req.Trailer
	rw := newTestResponseWriter(req)
	go func() {
		h.ServeHTTP(&rw, serverReq)
		rw.CloseWriter()
	}()
	return rw.Response(req.Context())
}

// nullReader is just a reader with no content. When _sending_ an HTTP request,
// a _nil_ body is allowed, but when receiving; there _is_ a body. This fixes
// the request so the valid output request body is also a valid incoming request
// body.
type nullReader struct{}

func (_ nullReader) Read(b []byte) (int, error) { return 0, io.EOF }
func (_ nullReader) Close() error               { return nil }

type testResponseWriter struct {
	// readerWriter *testReaderWriter
	// writer       io.Writer
	req       *http.Request
	response  *http.Response
	BodyReady chan struct{}
	Reader    *io.PipeReader
	Writer    *io.PipeWriter
	bufW      *bufio.Writer
}

func newTestResponseWriter(req *http.Request) testResponseWriter {
	r, w := io.Pipe()
	return testResponseWriter{
		req:    req,
		Reader: r,
		Writer: w,
		bufW:   bufio.NewWriter(w),
		response: &http.Response{
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Body:       r,
			Header:     make(http.Header),
		},
		BodyReady: make(chan struct{}),
	}
}

func (w *testResponseWriter) Response(ctx context.Context) (*http.Response, error) {
	getErr := func() error {
		err := ctx.Err()
		cause := context.Cause(ctx)
		if cause == err { // No explicit cause was provided
			return fmt.Errorf("gosthttp: roundtripper: %w", ctx.Err())
		} else {
			return fmt.Errorf("gosthttp: roundtripper: %w (%w)", ctx.Err(), cause)
		}
	}
	select {
	case <-ctx.Done():
		return nil, getErr()
	case <-w.BodyReady:
	}
	w.response.Request = w.req
	context.AfterFunc(ctx, func() {
		// Close the pipe with an error when context is done. There is no need
		// to check if there was an error. If the writer was closed normally,
		// the context error will not overwrite the EOF marker already present.
		w.Writer.CloseWithError(getErr())
	})
	return w.response, nil
}

func (w *testResponseWriter) WriteHeader(status int) {
	if w.response.StatusCode == 0 {
		if status == 0 {
			status = 200
		}
		w.response.StatusCode = status
		close(w.BodyReady)
	} else {
		if status != 0 {
			panic("Setting HTTP status twice")
		}
	}
}

func (w *testResponseWriter) Header() http.Header {
	return w.response.Header
}

func (w *testResponseWriter) Write(b []byte) (int, error) {
	w.WriteHeader(0)
	return w.Writer.Write(b)
}

func (w *testResponseWriter) CloseWriter() {
	w.WriteHeader(0)
	w.bufW.Flush()
	w.Writer.Close()
}

func (w *testResponseWriter) Flush() {
	w.bufW.Flush()
}
