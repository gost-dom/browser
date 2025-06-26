package gosttest

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// PipeHandler is reminescent of an io pipe, connecting an output stream to an
// input stream. This handler works as an HTTP handler that can be connected to,
// but client code can control the response.
//
// This is intended to be used in a test context, where it is designed to
// process exactly one request.
//
// All "write" methods return immediately, and will eventually perform the
// requested operation on the actual [http.ResponseWriter]. Errors returned from
// the ResponseWriter will result in a call to Error on the [testing.TB]
// instance.
//
// "Write" methods are sent to a buffered channel. BufSize controls the size of
// the buffer. If zero, a default value of 16 is used.
//
// If no TB instance is passed, errors are silently ignored.
//
// The zero value is safe to use but will not communicate errors, nor will it
// cancel on timeout.
type PipeHandler struct {
	T       testing.TB
	BufSize uint
	Req     *http.Request
	// ClientDisconnected tells whther the HTTP client disconnects before the
	// handler has completed.
	ClientDisconnected bool

	mu     sync.RWMutex
	init   sync.Once
	cmds   chan func(http.ResponseWriter)
	served bool
	closed bool
}

// Creates a new [PipeHandler] with a default [testing.TB] and [context.Context]
// instance.
//
// While both are optional, they are recommended options.
func NewPipeHandler(t testing.TB) *PipeHandler {
	return &PipeHandler{T: t}
}

func (h *PipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.served {
		panic("gosttest: PipeHandler: ServeHTTP: received multiple requests")
	}
	h.Req = r
	h.served = true
	h.ensureChannel()
	for {
		select {
		case f, ok := <-h.cmds:
			if !ok {
				return
			}
			f(w)
		case <-r.Context().Done():
			h.ClientDisconnected = true
			h.close() // probably not necessary, as the channel is subject to gc
			return
		}
	}
}

func (h *PipeHandler) ensureChannel() {
	h.init.Do(func() {
		if h.cmds == nil {
			buf := h.BufSize
			if buf == 0 {
				buf = 16
			}
			h.cmds = make(chan func(http.ResponseWriter), buf)
		}
	})
}

func (h *PipeHandler) WriteHeader(s int) {
	h.addF("WriteHeader", func(w http.ResponseWriter) { w.WriteHeader(s) })
}

func (h *PipeHandler) errorF(format string, a ...any) {
	if h.T == nil {
		return
	}
	h.T.Errorf(format, a...)
}

func (h *PipeHandler) Print(a ...any) {
	h.addF("Print", func(w http.ResponseWriter) {
		if _, err := fmt.Fprint(w, a...); err != nil {
			h.errorF("PipeHandler.Print: %v", err)
		}
	})
}

func (h *PipeHandler) Printf(format string, a ...any) {
	h.addF("Printf", func(w http.ResponseWriter) {
		if _, err := fmt.Fprintf(w, format, a...); err != nil {
			h.errorF("PipeHandler.Printf: %v", err)
		}
	})
}

func (h *PipeHandler) Do(f func(http.ResponseWriter)) {
	h.addF("Do", f)
}

func (h *PipeHandler) addF(n string, f func(http.ResponseWriter)) {
	h.ensureChannel()
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.closed {
		h.errorF("delayed http response: %s: context closed", n)
		return
	}
	h.cmds <- f
}

// Close closes the "pipe", completing the HTTP response. Panics if already
// closed.
func (h *PipeHandler) Close() {
	h.ensureChannel()
	h.close()
}

func (h *PipeHandler) close() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if !h.closed {
		h.closed = true
		close(h.cmds)
	}
}

func (h *PipeHandler) Flush() {
	h.addF("Flush", func(w http.ResponseWriter) {
		// While Flush() may return an error, we can't really use it for
		// anything. If the buffer isn't flushed, it will cause a bug.
		http.NewResponseController(w).Flush()
	})
}
