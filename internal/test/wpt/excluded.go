package main

var excludedSuites = []string{
	// uses document.write()
	"dom/nodes/Node-cloneNode-document-allow-declarative-shadow-roots.window.html",
}
