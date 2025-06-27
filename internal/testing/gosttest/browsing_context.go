package gosttest

import (
	"context"
	"log/slog"
	"net/http"
	"testing"

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

func NewBrowsingContext(t testing.TB, h http.Handler) BrowsingContext {
	return BrowsingContext{
		logger: NewTestLogger(t),
		Client: gosthttp.NewHttpClientFromHandler(h),
	}
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
