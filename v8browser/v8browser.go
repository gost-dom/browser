package v8browser

import (
	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/scripting/v8host"
)

// New initializes a new [browser.Browser] using V8 for JavaScript execution.
// See documentation for [browser.New] for description of options.
func New(options ...browser.BrowserOption) *browser.Browser {
	return browser.New(append(options, browser.WithScriptEngine(v8host.DefaultEngine()))...)
}
