package fetch_test

import (
	"log/slog"
	"net/http"
	"testing"

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

func TestFetch(t *testing.T) {
	handler := gosttest.StaticFileServer{
		"data.json": gosttest.StaticJSON(`{"foo": "foo"}`),
	}
	f := fetch.Fetch{NewBrowsingContext(t, handler)}
	f.Fetch(f.NewRequest("data.json"))
}

func TestRequestURLUsesDocumentLocation(t *testing.T) {
	bc := TestBrowsingContext{Location: "https://example.com/users/joe"}
	f := fetch.New(bc)
	req := f.NewRequest("alice")
	assert.Equal(t, "https://example.com/users/alice", req.URL())
}
