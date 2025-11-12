package browser_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

// initBrowser creates a browser with the most useful options. This browser will
// by default fail a test if an error is logged, meaning an uncaught JavaScript
// error will result in a test error.
func initBrowser(t testing.TB, handler http.Handler) (b *browser.Browser) {
	defer t.Cleanup(func() { b.Close() })

	return browser.New(
		browser.WithHandler(handler),
		browser.WithLogger(gosttest.NewTestLogger(t)),
	)
}

func TestModule(t *testing.T) {
	const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head><script src="module.js" type="module"></script></head>
			<body>
				<h1>Module test</h1>
				<div id="tgt"></div>
			</body>
		</html>`
	const moduleJS = `
		// Verify that compilation doesn't fail on export
		export const foo = 43;
		document.addEventListener("DOMContentLoaded", () => {
			document.getElementById("tgt").textContent="CONTENT";
		})
	`
	server := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(indexHTML),
		"/module.js":  gosttest.StaticJS(moduleJS),
	}
	win, err := initBrowser(t, server).Open("https://example.com/index.html")
	assert.NoError(t, err)
	g := gomega.NewWithT(t)
	g.Expect(
		win.Document().GetElementById("tgt")).To(HaveTextContent("CONTENT"))
}

func TestModuleImportPaths(t *testing.T) {
	const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head><script src="module.js" type="module"></script></head>
			<body>
				<h1>Module test</h1>
				<div id="tgt"></div>
			</body>
		</html>`
	const moduleJS = `
		// Verify that compilation doesn't fail on export
		export const foo = 43;
		document.addEventListener("DOMContentLoaded", () => {
			document.getElementById("tgt").textContent="CONTENT";
		})
	`
	server := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(indexHTML),
		"/module.js": gosttest.StaticJS(`
				import * as importValues from "./script.js";
				import featureA from "./feature-a/index.js";
				document.addEventListener("DOMContentLoaded", () => {
					document.getElementById("tgt").textContent = JSON.stringify({
						importValues,
						featureA,
					})
				})
			`),
		"/script.js": gosttest.StaticJS(`
				export const self = "/script.js";
			`),
		"/feature-a/index.js": gosttest.StaticJS(`
				import * as importValues from "./script.js";
				export default {
					self: "/feature-a/index.js",
					importValues,
				}
			`),
		"/feature-a/script.js": gosttest.StaticJS(`
				export const self = "/feature-a/script.js";
			`),
	}
	win, err := initBrowser(t, server).Open("https://example.com/index.html")
	assert.NoError(t, err)
	data := win.Document().GetElementById("tgt").TextContent()
	var m map[string]any
	json.Unmarshal([]byte(data), &m)
	featureA := m["featureA"].(jsonObject)
	featureAImport := featureA["importValues"].(jsonObject)
	importValues := m["importValues"].(jsonObject)
	assert.Equal(t, "/script.js", importValues["self"])
	assert.Equal(t, "/feature-a/index.js", featureA["self"])
	assert.Equal(t, "/feature-a/script.js", featureAImport["self"])
}

func TestModuleCyclicDeps(t *testing.T) {
	const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head><script src="module.js" type="module"></script></head>
			<body>
				<h1>Module test</h1>
				<div id="tgt-1"></div>
				<div id="tgt-2"></div>
				<div id="tgt-3"></div>
				<div id="tgt-4"></div>
			</body>
		</html>`

	server := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(indexHTML),
		"/a.js": gosttest.StaticJS(`
				import * as b from "./b.js";
				let keys
				try {
					keys = Object.keys(b)
				} catch {
					keys = []
				}
				export const keysFromB = keys.join(",")
				export const valueFromA = () => "valueFromA";
				export const printA = () => "a-" + b.valueFromB();
			`),
		"/b.js": gosttest.StaticJS(`
				import * as a from "./a.js";
				let keys
				try {
					keys = Object.keys(a)
				} catch {
					keys = []
				}
				export const keysFromA = keys.join(",")
				export const valueFromB = () => "valueFromB";
				export const printB = () => "b-" + a.valueFromA();
			`),
		"/module.js": gosttest.StaticJS(`
				import * as a from "./a.js"
				import * as b from "./b.js"
				document.addEventListener("DOMContentLoaded", () => {
					document.getElementById("tgt-1").textContent = a.printA()
					document.getElementById("tgt-2").textContent = b.printB()
					document.getElementById("tgt-3").textContent = String(a.keysFromB)
					document.getElementById("tgt-4").textContent = String(b.keysFromA)
				})
			`),
	}

	win, err := initBrowser(t, server).Open("https://example.com/index.html")
	assert.NoError(t, err)
	g := gomega.NewWithT(t)
	g.Expect(win.Document().GetElementById("tgt-1")).To(HaveTextContent("a-valueFromB"))
	g.Expect(win.Document().GetElementById("tgt-2")).To(HaveTextContent("b-valueFromA"))

	// Test cyclic dependencies root scope access
	//
	// When ES Modules have a cyclic dependency, the first time a module is
	// _seen again_, the importer gets an empty namespace object when the module
	// is evaluated.
	//
	// After the module has been evaluated, the namespace object is filled out,
	// so later execution of exported functions, or event handlers being
	// triggered will now be able to access the properties of the namespace
	// object. (Namespace object is the the exotic object created by an import
	// statement)
	//
	// E.g., When evaluating `a` that has an import cycle of a->b->c->a , then
	// `a` can access the exports of `b` during module evaluation, and `b` can
	// access the exports of `c`. But `c` cannot access the exports of `a`
	// during evaluation. But an exported function in `c` can access the exports
	// of `a` when called.
	g.Expect(win.Document().GetElementById("tgt-3")).
		To(HaveTextContent(gomega.ContainSubstring("keysFromA")), "")
	g.Expect(win.Document().GetElementById("tgt-4")).
		To(HaveTextContent(""), "Empty namespace object when evaluating the cycle")
}

type jsonObject = map[string]any
