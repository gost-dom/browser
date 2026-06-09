package worker

import (
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/clock"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
)

var _ htmlinterfaces.WorkerGlobalScope = &workerGlobalScope{}

type workerGlobalScope struct{ worker *Worker }

func newGlobal(w *Worker) *workerGlobalScope { return &workerGlobalScope{worker: w} }

func (g *workerGlobalScope) SetTimeout(
	cb clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	handle := g.worker.clock.SetTimeout(
		func() error {
			g.worker.Enqueue(func(GlobalScope) error {
				return cb()
			})
			return nil
		}, delay)
	return handle
}

// PostMessage causes a MessageEvent to be emitted in the main event loop.
//
// In JavaScript, postMessage sends a message to the Window object. This
// implementation just needs to know the main event loop, and the event target
// where the message should be emitted.
func (s *workerGlobalScope) PostMessage(data any) {
	e := s.worker.winClock.BeginEvent()
	e.AddEvent(func() error {
		s.worker.winEventTarget.DispatchEvent(&event.Event{
			Type: "message",
			Data: htmlinterfaces.MessageEventInit{
				Data: data,
			},
		})
		return nil
	})
}

func (g *workerGlobalScope) ClearTimeout(clock.TaskHandle) {
	panic("Not implemented")
}

func (g *workerGlobalScope) SetInterval(clock.TaskCallback, time.Duration) clock.TaskHandle {
	panic("Not implemented")
}

func (g *workerGlobalScope) ClearInterval(clock.TaskHandle) {
	panic("Not implemented")
}

func (g *workerGlobalScope) QueueMicrotask(clock.TaskCallback) {
	panic("Not implemented")
}
