package app

import (
	"fmt"
	"net/http"
	"time"
)

func NewDatastarMux() http.Handler {
	mux := http.NewServeMux()

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

		fmt.Fprintf(w, "event: datastar-patch-elements\n")
		fmt.Fprintf(w, `data: elements <div id="click-target">Foobar</div>`)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()
	})

	mux.HandleFunc("POST /form-test", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "text/event-stream")
		w.Header().Set("cache-control", "no-store")
		w.Header().Set("Connection", "keep-alive")
		w.WriteHeader(200)

		fmt.Fprintf(w, "event: datastar-patch-elements\n")
		fmt.Fprintf(w, `data: fragments <div id="form-post-target">Processing .</div>`)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()

		time.Sleep(time.Second)

		fmt.Fprintf(w, "event: datastar-patch-elements\n")
		fmt.Fprintf(w, `data: elements <div id="form-post-target">Processing ..</div>`)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()

		time.Sleep(time.Second)

		fmt.Fprintf(w, "event: datastar-patch-elements\n")
		fmt.Fprintf(w, `data: elements <div id="form-post-target">Processing ...</div>`)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()

		fmt.Fprintf(w, "event: datastar-patch-elements\n")
		fmt.Fprintf(
			w,
			`data: elements <div id="form-post-target"><span>Input1: %s</span><br /><span>Input2: %s</span></div>`,
			r.PostForm.Get("input-1"),
			r.PostForm.Get("input-2"),
		)
		fmt.Fprintf(w, "\n\n")
		flusher.Flush()
	})

	return mux
}
