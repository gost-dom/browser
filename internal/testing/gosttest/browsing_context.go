package gosttest

import (
	"context"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/log"
)

// Assert the type conforms to the interface
var _ html.BrowsingContext = BrowsingContext{}

// BrowsingContext implements [html.BrowsingContext] to help test code depending
// on a context, e.g, fetch and XHR.
type BrowsingContext struct {
	logger   *slog.Logger
	Ctx      context.Context
	Client   http.Client
	Location string
}

type browsingContextOption struct {
	timeout time.Duration
}

type BrowsingContextOption func(*browsingContextOption)

func WithTimeoutMS(ms int) BrowsingContextOption {
	return func(o *browsingContextOption) {
		o.timeout = time.Duration(ms) * time.Millisecond
	}
}

// NewBrowsingContextFromT creates [html.BrowsingContext] with the logger
// connected to the current test case, as well as a default context deadline.
func NewBrowsingContextFromT(
	t testing.TB,
	h http.Handler,
	opts ...BrowsingContextOption,
) (BrowsingContext, context.CancelFunc) {
	var opt browsingContextOption
	for _, o := range opts {
		o(&opt)
	}
	timeout := opt.timeout
	if timeout == 0 {
		timeout = time.Second
	}

	ctx, cancel := context.WithTimeout(t.Context(), timeout)
	return BrowsingContext{
		logger: NewTestLogger(t),
		Client: gosthttp.NewHttpClientFromHandler(h),
		Ctx:    ctx,
	}, cancel
}

func (c BrowsingContext) Context() context.Context { return c.Ctx }
func (c BrowsingContext) HTTPClient() http.Client  { return c.Client }
func (c BrowsingContext) LocationHREF() string     { return c.Location }
func (c BrowsingContext) Logger() *slog.Logger {
	if c.logger != nil {
		return c.logger
	}
	return log.Default()
}
