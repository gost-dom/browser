package fetch_test

import (
	"context"
	"log/slog"
	"net/http"
	"testing"
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

func TestFetchWithAbortSignal(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	handler := gosttest.NewPipeHandler(t, ctx)
	bc := TestBrowsingContext{
		Location: "https://example.com/users/joe",
		Client:   gosthttp.NewHttpClientFromHandler(handler),
	}
	ac := dom.NewAbortController()
	f := fetch.New(bc)
	req := f.NewRequest("url")

	p := f.FetchAsync(req, fetch.WithSignal(ac.Signal()))

	ac.Abort("Dummy Reason")

	select {
	case <-ctx.Done():
		t.Error("timeout")
	case <-p:
	}
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
