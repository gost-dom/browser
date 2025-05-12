package app

import (
	"fmt"
	"net/http"
)

func NewDatastarMux() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /events/", func(w http.ResponseWriter, r *http.Request) {
	mux.HandleFunc("GET /events", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "text/event-stream")
		w.Header().Set("cache-control", "no-store")
		w.Header().Set("Connection", "keep-alive")
		w.WriteHeader(200)

		fmt.Fprintf(w, "event: datastar-merge-fragments\n")
		fmt.Fprintf(w, `data: fragments <div id="click-target">Foobar</div>`)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()
	})

	return mux
}
