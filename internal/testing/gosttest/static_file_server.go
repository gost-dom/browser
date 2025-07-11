package gosttest

import (
	"fmt"
	"net/http"
)

// StaticFileServer is a simple [http.Handler] implementing a mux (router) based
// on a simple map[string]http.Handler, mapping a local path to a handler.
//
// Having a map as an underlying type, you can create the entire http handler as
// a Go map literal, making it simpler to configure than creating a new mux.
// Combine with [StaticHTML], [StaticJS], or [StaticJSON] (or other variants
// added after this comment was written), makes it easy to serve static files
// for test content.
//
// TODO: Give this a new name. The first version was _only_ static files, but
// after extracting the StaticFile type as a dedicated and valid http.Handler
// implementation, this has been elevated to a more general use case.
type StaticFileServer map[string]http.Handler

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

func (s StaticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
