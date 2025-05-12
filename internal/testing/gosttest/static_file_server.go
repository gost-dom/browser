package gosttest

import (
	"fmt"
	"net/http"
)

// StaticFileServer is a simple [http.Handler] that can simplify test code that
// only needs to configure static files.
//
// As it has a map as an underlying type, you can create the entire http handler
// as a Go map literal, making it simpler to configure than creating a new mux.
type StaticFileServer map[string]StaticFile

// A simple [http.Handler] that serves static file content. Type type is a pair
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

func (s StaticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}
	if content, found := s[r.URL.Path]; found {
		content.ServeHTTP(w, r)
	} else {
		w.WriteHeader(404)
	}
}
