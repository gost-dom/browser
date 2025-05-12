package gosttest

import (
	"fmt"
	"net/http"
)

// StaticFileServer is a simple HTTP server serving GET requests based on a map
// of URLs to static file content.
//
// The content is represented by a 2-element string array containing the
// Content-Type response header, and response body respectively. The type is
// optimised for succinct test code. As a result, the type does not properly
// describe what element is for, but working test code is readable, as
// the actual values are easily recognised.
type StaticFileServer map[string]StaticFile

type StaticFile [2]string

func (f StaticFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", f[0])
	fmt.Fprint(w, f[1])
}

func StaticHTML(html string) StaticFile { return StaticFile{"text/html", html} }
func StaticJS(js string) StaticFile     { return StaticFile{"text/javascript", js} }

func (s StaticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, found := s[r.URL.String()]
	if !found {
		content, found = s[r.URL.Path]
	}
	if !found {
		w.WriteHeader(404)
		return
	}
	content.ServeHTTP(w, r)
}
