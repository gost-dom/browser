package fetch_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/assert"
)

type TestBrowsingContext struct {
	logger   *slog.Logger
	Client   http.Client
	Location string
}

func NewBrowsingContext(t testing.TB, h http.Handler) TestBrowsingContext {
	return TestBrowsingContext{
		logger: gosttest.NewTestLogger(t),
		Client: gosthttp.NewHttpClientFromHandler(h),
	}
}

func (c TestBrowsingContext) HTTPClient() http.Client { return c.Client }
func (c TestBrowsingContext) LocationHREF() string    { return c.Location }
func (c TestBrowsingContext) Logger() *slog.Logger {
	if c.logger != nil {
		return c.logger
	}
	return log.Default()
}

func TestRequestURLUsesDocumentLocation(t *testing.T) {
	bc := TestBrowsingContext{Location: "https://example.com/users/joe"}
	f := fetch.New(bc)
	req := f.NewRequest("alice")
	assert.Equal(t, "https://example.com/users/alice", req.URL())
}

func TestFetchAborted(t *testing.T) {

	// Status code sent indicates the time when a Response object is returned by
	// the HTTP roundtripper, but the body is still not streamed.
	t.Run("Before status code has been sent", func(t *testing.T) {
		synctest.Run(func() {
			ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			defer cancel()

			handler := gosttest.NewPipeHandler(t)
			bc := NewBrowsingContext(t, handler)
			ac := dom.NewAbortController()
			f := fetch.New(bc)
			req := f.NewRequest("url")

			p := f.FetchAsync(ctx, req, fetch.WithSignal(ac.Signal()))

			synctest.Wait() // Doesn't affect the outcome, but the next assertion is useless without
			assert.False(t, handler.ClientDisconnected, "Client disconnected before cancel")

			ac.Abort("Dummy Reason")

			result := gosttest.ExpectReceive(t, p, gosttest.Context(t.Context()))
			assert.Error(t, result.Err, "Response should be an error")

			synctest.Wait()
			assert.True(t, handler.ClientDisconnected, "Client disconnected after cancel")
		})
	})

	t.Run("After status code has been sent", func(t *testing.T) {
		synctest.Run(func() {
			ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			defer cancel()

			handler := gosttest.NewPipeHandler(t)
			bc := NewBrowsingContext(t, handler)
			ac := dom.NewAbortController()
			f := fetch.New(bc)
			req := f.NewRequest("url")
			handler.WriteHeader(200)

			p := f.FetchAsync(ctx, req, fetch.WithSignal(ac.Signal()))

			synctest.Wait()
			res := <-p
			assert.Equal(t, 200, res.Value.Status)
			assert.NoError(t, res.Err, "response error")

			ac.Abort("Dummy reason")

			_, err := io.ReadAll(res.Value.Reader)
			assert.Error(t, err, "reading response body of cancelled response")

		})
	})
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
