package html_test

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/assert"
)

// dummyScriptHost implements the ScriptHost interface but does not execute
// scripts. It is to help test behaviour of script downloading and execution
// behaviour separately from actual script engine.
type dummyScriptHost struct{}

func (h dummyScriptHost) Close() {}
func (h dummyScriptHost) NewContext(win html.Window) html.ScriptContext {
	return dummyScriptContext{win}
}

type dummyScriptContext struct{ win html.Window }

func (c dummyScriptContext) Close()                   {}
func (c dummyScriptContext) Clock() html.Clock        { return c.win.Clock() }
func (c dummyScriptContext) Eval(string) (any, error) { return nil, nil }
func (c dummyScriptContext) Run(string) error         { return nil }

func TestScriptElementSourceResolution(t *testing.T) {
	indexHTML := `
	<!DOCTYPE html>
	<html>
		<head><script src="script.js"></script></head>
		<body><h1>Script Test Page</h1></body>
	</html>`
	rootSrcHTML := `
	<!DOCTYPE html>
	<html>
		<head><script src="/script.js"></script></head>
		<body><h1>Script Test Page</h1></body>
	</html>`
	dummyScript := "// dummy script with no behaviour"
	srv := gosttest.StaticFileServer{
		"http://example.com/index.html":              [2]string{"text/html", indexHTML},
		"http://example.com/script.js":               [2]string{"text/javascript", dummyScript},
		"http://example.com/folder/index.html":       [2]string{"text/html", indexHTML},
		"http://example.com/folder/root-source.html": [2]string{"text/html", rootSrcHTML},
		"http://example.com/folder/script.js":        [2]string{"text/javascript", dummyScript},
	}

	rec := &HTTPRequestRecorder{Handler: srv}
	options := []html.WindowOption{
		html.WindowOptionHTTPClient(gosthttp.NewHttpClientFromHandler(rec)),
		html.WindowOptionHost(dummyScriptHost{}),
	}

	// Import script relative to current file in root
	_, err := html.OpenWindowFromLocation("http://example.com/index.html", options...)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"http://example.com/index.html",
		"http://example.com/script.js",
	}, rec.URLs())

	// Import script relative to current file in subfolder
	rec.Clear()
	_, err = html.OpenWindowFromLocation("http://example.com/folder/index.html", options...)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"http://example.com/folder/index.html",
		"http://example.com/folder/script.js",
	}, rec.URLs())

	// Import script relative with root path in root
	rec.Clear()
	_, err = html.OpenWindowFromLocation("http://example.com/folder/root-source.html", options...)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"http://example.com/folder/root-source.html",
		"http://example.com/script.js",
	}, rec.URLs())
}

// HTTPRequestRecorder is an HTTPHandler middleware that keeps a record of all
// incoming request objects.
type HTTPRequestRecorder struct {
	Handler  http.Handler
	Requests []*http.Request
}

func (rec *HTTPRequestRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec.Requests = append(rec.Requests, r)
	rec.Handler.ServeHTTP(w, r)
}

// URLs return all URL strings recorded
func (r HTTPRequestRecorder) URLs() []string {
	res := make([]string, len(r.Requests))
	for i, req := range r.Requests {
		res[i] = req.URL.String()
	}
	return res
}

// Clear deletes all recorded Requests.
func (r *HTTPRequestRecorder) Clear() { r.Requests = nil }
