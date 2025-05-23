package html_test

import (
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
		"/index.html":              gosttest.StaticHTML(indexHTML),
		"/script.js":               gosttest.StaticJS(dummyScript),
		"/folder/index.html":       gosttest.StaticHTML(indexHTML),
		"/folder/root-source.html": gosttest.StaticHTML(rootSrcHTML),
		"/folder/script.js":        gosttest.StaticJS(dummyScript),
	}

	rec := gosttest.NewHTTPRequestRecorder(t, srv)
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
