package v8host_test

import (
	"net/http"

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
		window, err := OpenTestWindowFromHandler("/index.html", server)
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
		win, err := OpenTestWindowFromHandler("/index.html", server)
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Eval("window.scriptLoaded")).To(BeTrue())
	})
})
