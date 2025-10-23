package gosttest

import (
	"fmt"
	"net/http"
)

// HttpHandlerMap is a simple [http.Handler] implementing a mux (router) based
// on a simple map[string]http.Handler, mapping a local path to a handler.
//
// Having a map as an underlying type, you can create the entire http handler as
// a Go map literal, making it simpler to configure than creating a new mux.
// Combine with handlers like [StaticHTML], [StaticJS], or [StaticJSON]
// makes it easy build the necessary context for a many test cases.
type HttpHandlerMap map[string]http.Handler

// A simple [http.Handler] that serves static file content. This type is a pair
// of MIMEType and body content.
type StaticFile struct {
	MIMEType string
	Body     string
}

func (f StaticFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", f.MIMEType)
	fmt.Fprint(w, f.Body)
}

// StaticHTML creates an [http.Handler] that serves static content with the MIME
// type, "text/html".
func StaticHTML(html string) StaticFile { return StaticFile{"text/html", html} }

// StaticJS creates an [http.Handler] that serves static content with the MIME
// type, "text/javascript".
func StaticJS(js string) StaticFile { return StaticFile{"text/javascript", js} }

// StaticJSON creates an [http.Handler] that serves static content with the MIME
// type, "text/javascript".
func StaticJSON(json string) StaticFile { return StaticFile{"application/json", json} }

func (s HttpHandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if content, found := s[r.URL.Path]; found {
		content.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
