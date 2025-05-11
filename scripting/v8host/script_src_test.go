package v8host_test

import (
	"net/http"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/scripting/v8host"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load from server", func() {
	It("Loads from an HTTP server", func() {
		server := http.NewServeMux()
		server.Handle(
			"GET /index.html",
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("<body>Hello, World!</body>"))
			}),
		)
		client := gosthttp.NewHttpClientFromHandler(server)
		window, err := html.OpenWindowFromLocation("/index.html", html.WindowOptions{
			ScriptHost: v8host.New(v8host.WithHTTPClient(&client)),
			HttpClient: client,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(window.Document().Body().OuterHTML()).To(Equal("<body>Hello, World!</body>"))
	})

	It("Should download and execute script from script tags", func() {
		// Create a simple server, serving an HTML file and JS
		server := http.NewServeMux()
		server.HandleFunc(
			"GET /index.html",
			func(res http.ResponseWriter, req *http.Request) {
				res.Write(
					[]byte(
						`<html><head><script src="/js/script.js"></script></head><body>Hello, World!</body>`,
					),
				)
			},
		)
		// The script is pretty basic. In order to verify it has been executed, it
		// produces an observable side effect; setting a variable in global scope
		server.HandleFunc(
			"GET /js/script.js",
			func(res http.ResponseWriter, req *http.Request) {
				res.Header().Add("Content-Type", "text/javascript")
				res.Write([]byte(`var scriptLoaded = true`))
			},
		)
		// Verify, create a browser communicating with this. Open the HTML file, and
		// verify the side effect by inspecting global JS scope.
		client := gosthttp.NewHttpClientFromHandler(server)
		win, err := html.OpenWindowFromLocation("/index.html", html.WindowOptions{
			ScriptHost: v8host.New(v8host.WithHTTPClient(&client)),
			HttpClient: client,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Eval("window.scriptLoaded")).To(BeTrue())
	})
})
