package html_test

import (
	"fmt"
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
type dummyScriptHost struct {
	client  http.Client
	scripts []dummyScript
}

func (h *dummyScriptHost) record(script dummyScript) (html.Script, error) {
	h.scripts = append(h.scripts, script)
	return script, nil
}

func (h *dummyScriptHost) Close() {}
func (h *dummyScriptHost) NewContext(win html.Window) html.ScriptContext {
	return dummyScriptContext{h, win, h.client}
}

type dummyScriptContext struct {
	host   *dummyScriptHost
	win    html.Window
	client http.Client
}

func (c dummyScriptContext) Close()                   {}
func (c dummyScriptContext) Clock() html.Clock        { return c.win.Clock() }
func (c dummyScriptContext) Eval(string) (any, error) { return nil, nil }
func (c dummyScriptContext) Run(string) error         { return nil }
func (c dummyScriptContext) Compile(string) (html.Script, error) {
	return c.host.record(dummyScript{scriptType: "inline"})
}

func (c dummyScriptContext) download(url string, scriptType string) (html.Script, error) {
	res, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad status code: %d, scr: %s", res.StatusCode, url)
	}
	return c.host.record(dummyScript{
		url:        url,
		scriptType: scriptType,
	})
}

func (c dummyScriptContext) DownloadScript(url string) (html.Script, error) {
	return c.download(url, "classic")
}

func (c dummyScriptContext) DownloadModule(url string) (html.Script, error) {
	return c.download(url, "module")
}

type dummyScript struct {
	url        string
	scriptType string
}

func (c dummyScript) Eval() (any, error) { return nil, nil }
func (c dummyScript) Run() error         { return nil }

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
	httpClient := gosthttp.NewHttpClientFromHandler(rec)
	scriptHost := &dummyScriptHost{httpClient, nil}
	options := []html.WindowOption{
		html.WindowOptionHTTPClient(httpClient),
		html.WindowOptionHost(scriptHost),
	}

	// Import script relative to current file in root
	_, err := html.OpenWindowFromLocation("http://example.com/index.html", options...)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"http://example.com/index.html",
		"http://example.com/script.js",
	}, rec.URLs())

	assert.Equal(t, 1, len(scriptHost.scripts))
	assert.Equal(t, "classic", scriptHost.scripts[0].scriptType)

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

func TestScriptElementModuleResolution(t *testing.T) {
	indexHTML := `
	<!DOCTYPE html>
	<html>
		<head><script src="script.js" type="module"></script></head>
		<body><h1>Script Test Page</h1></body>
	</html>`
	dummyScript := "// dummy script with no behaviour"
	srv := gosttest.StaticFileServer{
		"/index.html": gosttest.StaticHTML(indexHTML),
		"/script.js":  gosttest.StaticJS(dummyScript),
	}

	rec := gosttest.NewHTTPRequestRecorder(t, srv)
	httpClient := gosthttp.NewHttpClientFromHandler(rec)
	scriptHost := &dummyScriptHost{httpClient, nil}
	options := []html.WindowOption{
		html.WindowOptionHTTPClient(httpClient),
		html.WindowOptionHost(scriptHost),
	}

	// Import script relative to current file in root
	_, err := html.OpenWindowFromLocation("http://example.com/index.html", options...)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"http://example.com/index.html",
		"http://example.com/script.js",
	}, rec.URLs())

	assert.Equal(t, 1, len(scriptHost.scripts))
	assert.Equal(t, "module", scriptHost.scripts[0].scriptType, "Expect an ESM module")
}
